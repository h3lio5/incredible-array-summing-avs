package types

import (
	"errors"

	cstaskmanager "github.com/h3lio5/incredible-array-summing-avs/contracts/bindings/IncredibleSummingTaskManager"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.IIncredibleSummingTaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.IIncredibleSummingTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
