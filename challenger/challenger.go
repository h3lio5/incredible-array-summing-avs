package challenger

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"nhooyr.io/websocket"

	ethclient "github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/h3lio5/incredible-array-summing-avs/core/config"

	"github.com/h3lio5/incredible-array-summing-avs/challenger/types"
	cstaskmanager "github.com/h3lio5/incredible-array-summing-avs/contracts/bindings/IncredibleSummingTaskManager"
	"github.com/h3lio5/incredible-array-summing-avs/core/chainio"
)

type Challenger struct {
	logger             logging.Logger
	ethClient          ethclient.Client
	avsReader          chainio.AvsReaderer
	avsWriter          chainio.AvsWriterer
	avsSubscriber      chainio.AvsSubscriberer
	tasks              map[uint32]cstaskmanager.IIncredibleSummingTaskManagerTask
	taskResponses      map[uint32]ResultToAvailDa
	taskResponseChan   chan *cstaskmanager.ContractIncredibleSummingTaskManagerTaskResponded
	newTaskCreatedChan chan *cstaskmanager.ContractIncredibleSummingTaskManagerNewTaskCreated
	webSocketChan      chan WebSocketMessage
}

type ResultToAvailDa struct {
	TaskResponse                cstaskmanager.IIncredibleSummingTaskManagerTaskResponse
	TaskResponseMetadata        cstaskmanager.IIncredibleSummingTaskManagerTaskResponseMetadata
	NonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature
}

type WebSocketMessage struct {
	MessageType websocket.MessageType
	Data        []byte
}

type SubscriptionRequest struct {
	Topics     []string `json:"topics"`
	DataFields []string `json:"data_fields"`
}
type AvailDataBlob struct {
	Topic   string  `json:"topic"`
	Message Message `json:"message"`
}
type Message struct {
	BlockNumber      uint64            `json:"block_number"`
	DataTransactions []DataTransaction `json:"data_transactions"`
}
type DataTransaction struct {
	Data string `json:"data"`
}

func NewChallenger(c *config.Config) (*Challenger, error) {

	avsWriter, err := chainio.BuildAvsWriterFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create EthWriter", "err", err)
		return nil, err
	}
	avsReader, err := chainio.BuildAvsReaderFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create EthReader", "err", err)
		return nil, err
	}
	avsSubscriber, err := chainio.BuildAvsSubscriberFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create EthSubscriber", "err", err)
		return nil, err
	}

	webSocketChan := make(chan WebSocketMessage)

	challenger := &Challenger{
		ethClient:          c.EthHttpClient,
		logger:             c.Logger,
		avsWriter:          avsWriter,
		avsReader:          avsReader,
		avsSubscriber:      avsSubscriber,
		tasks:              make(map[uint32]cstaskmanager.IIncredibleSummingTaskManagerTask),
		taskResponses:      make(map[uint32]ResultToAvailDa),
		taskResponseChan:   make(chan *cstaskmanager.ContractIncredibleSummingTaskManagerTaskResponded),
		newTaskCreatedChan: make(chan *cstaskmanager.ContractIncredibleSummingTaskManagerNewTaskCreated),
		webSocketChan:      webSocketChan,
	}

	return challenger, nil
}

func (c *Challenger) Start(ctx context.Context) error {
	c.logger.Infof("Starting Challenger.")

	newTaskSub := c.avsSubscriber.SubscribeToNewTasks(c.newTaskCreatedChan)
	c.logger.Infof("Subscribed to new tasks")

	taskResponseSub := c.avsSubscriber.SubscribeToTaskResponses(c.taskResponseChan)
	c.logger.Infof("Subscribed to task responses")

	// Generate a subscription ID
	subscriptionID, err := generateSubscriptionID()
	if err != nil {
		c.logger.Errorf("Error generating subscription ID: %v", err)
		return err
	}

	c.logger.Infof("Subscription ID: %v", subscriptionID)

	// Subscribe to the WebSocket API
	err = c.subscribeToWebSocket(ctx, subscriptionID)
	if err != nil {
		c.logger.Errorf("Error subscribing to WebSocket: %v", err)
		return err
	}

	for {
		select {
		case err := <-newTaskSub.Err():
			// TODO(samlaf): Copied from operator. There was a comment about this on when should exactly do these errors occur? do we need to restart the websocket
			c.logger.Error("Error in websocket subscription for new Task", "err", err)
			newTaskSub.Unsubscribe()
			newTaskSub = c.avsSubscriber.SubscribeToNewTasks(c.newTaskCreatedChan)

		case err := <-taskResponseSub.Err():
			// TODO(samlaf): Copied from operator. There was a comment about this on when should exactly do these errors occur? do we need to restart the websocket
			c.logger.Error("Error in websocket subscription for task response", "err", err)
			taskResponseSub.Unsubscribe()
			taskResponseSub = c.avsSubscriber.SubscribeToTaskResponses(c.taskResponseChan)

		case newTaskCreatedLog := <-c.newTaskCreatedChan:
			c.logger.Info("New task created log received", "newTaskCreatedLog", newTaskCreatedLog)
			taskIndex := c.processNewTaskCreatedLog(newTaskCreatedLog)

			if _, found := c.taskResponses[taskIndex]; found {
				err := c.callChallengeModule(taskIndex)
				if err != nil {
					c.logger.Error("Error calling challenge module", "err", err)
				}
				continue
			}

		case availDataBlobMsg := <-c.webSocketChan:
			if availDataBlobMsg.MessageType == websocket.MessageBinary {
				// Handle binary data (blob)
				taskIndex, err := c.handleBinaryData(availDataBlobMsg.Data)
				if err != nil {
					c.logger.Info("Error decoding the avail data blob", "err", err)
					continue
				}

				if _, found := c.tasks[taskIndex]; found {
					err := c.callChallengeModule(taskIndex)
					if err != nil {
						c.logger.Info("Info:", "err", err)
					}
					continue
				}

			}
		}
	}

}

func generateSubscriptionID() (string, error) {
	url := "http://127.0.0.1:7000/v2/subscriptions"
	reqBody := SubscriptionRequest{
		Topics:     []string{"data-verified"},
		DataFields: []string{"data"},
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	subscriptionID, ok := result["subscription_id"].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse subscription ID")
	}

	return subscriptionID, nil
}

func (c *Challenger) subscribeToWebSocket(ctx context.Context, subscriptionID string) error {
	url := fmt.Sprintf("ws://127.0.0.1:7000/v2/ws/%s", subscriptionID)
	conn, _, err := websocket.Dial(ctx, url, &websocket.DialOptions{
		HTTPClient: &http.Client{},
	})

	if err != nil {
		c.logger.Error("Failed to establish the websocket connection", err)
		return err
	}
	defer conn.Close(websocket.StatusInternalError, "")

	c.logger.Info("Successfully established the websockets connection")

	// Read messages from the WebSocket connection
	go func() {
		for {
			messageType, message, err := conn.Read(ctx)
			if err != nil {
				c.logger.Errorf("Error reading from WebSocket: %v", err)
				return
			}

			// Send the message to the webSocketChan
			c.webSocketChan <- WebSocketMessage{
				MessageType: messageType,
				Data:        message,
			}
		}
	}()

	return nil
}

func (c *Challenger) handleBinaryData(data []byte) (uint32, error) {
	// Unmarshal the binary data into your desired struct
	var blob AvailDataBlob
	var taskResponseReferenceIndex uint32

	err := json.Unmarshal(data, &blob)
	if err != nil {
		c.logger.Errorf("Error unmarshaling JSON data: %v", err)
		return 0, err
	}

	for _, tx := range blob.Message.DataTransactions {
		decodedData, err := base64.StdEncoding.DecodeString(tx.Data)
		if err != nil {
			c.logger.Errorf("Error decoding base64 data: %v", err)
			continue
		}

		var taskResponse ResultToAvailDa
		err = json.Unmarshal(decodedData, &taskResponse)
		if err != nil {
			c.logger.Errorf("Error unmarshaling task response: %v", err)
			continue
		}
		// Use the unmarshaled struct for your other tasks
		taskResponseReferenceIndex = c.processTaskResponseData(taskResponse)
	}
	return taskResponseReferenceIndex, nil
}

func (c *Challenger) processNewTaskCreatedLog(newTaskCreatedLog *cstaskmanager.ContractIncredibleSummingTaskManagerNewTaskCreated) uint32 {
	c.tasks[newTaskCreatedLog.TaskIndex] = newTaskCreatedLog.Task
	return newTaskCreatedLog.TaskIndex
}

func (c *Challenger) processTaskResponseData(taskResponseData ResultToAvailDa) uint32 {
	c.taskResponses[taskResponseData.TaskResponse.ReferenceTaskIndex] = taskResponseData
	return taskResponseData.TaskResponse.ReferenceTaskIndex
}

func (c *Challenger) callChallengeModule(taskIndex uint32) error {
	arrayToBeSummed := c.tasks[taskIndex].ArrayToBeSummed
	answerInResponse := c.taskResponses[taskIndex].TaskResponse.ArraySummed

	trueAnswer := uint64(0)
	for _, num := range arrayToBeSummed {
		trueAnswer += num
	}

	// checking if the answer in the response submitted by aggregator is correct
	if trueAnswer != answerInResponse {
		c.logger.Infof("The array summed is not correct")

		// raise challenge
		c.raiseChallenge(taskIndex)

		return nil
	}
	return types.NoErrorInTaskResponse
}

func (c *Challenger) raiseChallenge(taskIndex uint32) error {
	c.logger.Info("Challenger raising challenge.", "taskIndex", taskIndex)
	c.logger.Info("Task", "Task", c.tasks[taskIndex])
	c.logger.Info("TaskResponse", "TaskResponse", c.taskResponses[taskIndex].TaskResponse)
	c.logger.Info("TaskResponseMetadata", "TaskResponseMetadata", c.taskResponses[taskIndex].TaskResponseMetadata)
	c.logger.Info("NonSignerStakesAndSignature", "NonSignerStakesAndSignature", c.taskResponses[taskIndex].NonSignerStakesAndSignature)

	receipt, err := c.avsWriter.RaiseChallenge(
		context.Background(),
		c.tasks[taskIndex],
		c.taskResponses[taskIndex].TaskResponse,
		c.taskResponses[taskIndex].TaskResponseMetadata,
		c.taskResponses[taskIndex].NonSignerStakesAndSignature,
	)

	if err != nil {
		c.logger.Error("Challenger failed to raise challenge:", "err", err)
		return err
	}
	c.logger.Infof("Tx hash of the challenge tx: %v", receipt.TxHash.Hex())
	return nil
}
