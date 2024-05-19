package chainio

import (
	"context"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	logging "github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	cstaskmanager "github.com/h3lio5/incredible-array-summing-avs/contracts/bindings/IncredibleSummingTaskManager"
	"github.com/h3lio5/incredible-array-summing-avs/core/config"
)

type AvsWriterer interface {
	avsregistry.AvsRegistryWriter

	SendNewTaskArrayToSum(
		ctx context.Context,
		arrayToSum [3]uint64,
		quorumThresholdPercentage sdktypes.QuorumThresholdPercentage,
		quorumNumbers sdktypes.QuorumNums,
	) (cstaskmanager.IIncredibleSummingTaskManagerTask, uint32, error)
	RaiseChallenge(
		ctx context.Context,
		task cstaskmanager.IIncredibleSummingTaskManagerTask,
		taskResponse cstaskmanager.IIncredibleSummingTaskManagerTaskResponse,
		taskResponseMetadata cstaskmanager.IIncredibleSummingTaskManagerTaskResponseMetadata,
		nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
	) (*types.Receipt, error)
}

type AvsWriter struct {
	avsregistry.AvsRegistryWriter
	AvsContractBindings *AvsManagersBindings
	logger              logging.Logger
	TxMgr               txmgr.TxManager
	client              eth.Client
}

var _ AvsWriterer = (*AvsWriter)(nil)

func BuildAvsWriterFromConfig(c *config.Config) (*AvsWriter, error) {
	return BuildAvsWriter(c.TxMgr, c.IncredibleSummingRegistryCoordinatorAddr, c.OperatorStateRetrieverAddr, c.EthHttpClient, c.Logger)
}

func BuildAvsWriter(txMgr txmgr.TxManager, registryCoordinatorAddr, operatorStateRetrieverAddr gethcommon.Address, ethHttpClient eth.Client, logger logging.Logger) (*AvsWriter, error) {
	avsServiceBindings, err := NewAvsManagersBindings(registryCoordinatorAddr, operatorStateRetrieverAddr, ethHttpClient, logger)
	if err != nil {
		logger.Error("Failed to create contract bindings", "err", err)
		return nil, err
	}
	avsRegistryWriter, err := avsregistry.BuildAvsRegistryChainWriter(registryCoordinatorAddr, operatorStateRetrieverAddr, logger, ethHttpClient, txMgr)
	if err != nil {
		return nil, err
	}
	return NewAvsWriter(avsRegistryWriter, avsServiceBindings, logger, txMgr), nil
}
func NewAvsWriter(avsRegistryWriter avsregistry.AvsRegistryWriter, avsServiceBindings *AvsManagersBindings, logger logging.Logger, txMgr txmgr.TxManager) *AvsWriter {
	return &AvsWriter{
		AvsRegistryWriter:   avsRegistryWriter,
		AvsContractBindings: avsServiceBindings,
		logger:              logger,
		TxMgr:               txMgr,
	}
}

// returns the tx receipt, as well as the task index (which it gets from parsing the tx receipt logs)
func (w *AvsWriter) SendNewTaskArrayToSum(ctx context.Context, arrayToSum [3]uint64, quorumThresholdPercentage sdktypes.QuorumThresholdPercentage, quorumNumbers sdktypes.QuorumNums) (cstaskmanager.IIncredibleSummingTaskManagerTask, uint32, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return cstaskmanager.IIncredibleSummingTaskManagerTask{}, 0, err
	}
	tx, err := w.AvsContractBindings.TaskManager.CreateNewTask(txOpts, arrayToSum, uint32(quorumThresholdPercentage), quorumNumbers.UnderlyingType())
	if err != nil {
		w.logger.Errorf("Error assembling CreateNewTask tx")
		return cstaskmanager.IIncredibleSummingTaskManagerTask{}, 0, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting CreateNewTask tx")
		return cstaskmanager.IIncredibleSummingTaskManagerTask{}, 0, err
	}
	newTaskCreatedEvent, err := w.AvsContractBindings.TaskManager.ContractIncredibleSummingTaskManagerFilterer.ParseNewTaskCreated(*receipt.Logs[0])
	if err != nil {
		w.logger.Error("Aggregator failed to parse new task created event", "err", err)
		return cstaskmanager.IIncredibleSummingTaskManagerTask{}, 0, err
	}
	return newTaskCreatedEvent.Task, newTaskCreatedEvent.TaskIndex, nil
}

func (w *AvsWriter) RaiseChallenge(
	ctx context.Context,
	task cstaskmanager.IIncredibleSummingTaskManagerTask,
	taskResponse cstaskmanager.IIncredibleSummingTaskManagerTaskResponse,
	taskResponseMetadata cstaskmanager.IIncredibleSummingTaskManagerTaskResponseMetadata,
	nonSignerStakesAndSignature cstaskmanager.IBLSSignatureCheckerNonSignerStakesAndSignature,
) (*types.Receipt, error) {
	txOpts, err := w.TxMgr.GetNoSendTxOpts()
	if err != nil {
		w.logger.Errorf("Error getting tx opts")
		return nil, err
	}
	tx, err := w.AvsContractBindings.TaskManager.RaiseAndResolveChallenge(txOpts, task, taskResponse, taskResponseMetadata, nonSignerStakesAndSignature)
	if err != nil {
		w.logger.Errorf("Error assembling RaiseChallenge tx")
		return nil, err
	}
	receipt, err := w.TxMgr.Send(ctx, tx)
	if err != nil {
		w.logger.Errorf("Error submitting RaiseChallenge tx")
		return nil, err
	}
	return receipt, nil
}
