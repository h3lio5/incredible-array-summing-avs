// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIncredibleSummingTaskManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IIncredibleSummingTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSummingTaskManagerTask struct {
	ArrayToBeSummed           [3]uint64
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

// IIncredibleSummingTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSummingTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	ArraySummed        uint64
}

// IIncredibleSummingTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IIncredibleSummingTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractIncredibleSummingTaskManagerMetaData contains all meta data concerning the ContractIncredibleSummingTaskManager contract.
var ContractIncredibleSummingTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"arrayToBeSummed\",\"type\":\"uint64[3]\",\"internalType\":\"uint64[3]\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSummingTaskManager.Task\",\"components\":[{\"name\":\"arrayToBeSummed\",\"type\":\"uint64[3]\",\"internalType\":\"uint64[3]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSummingTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"arraySummed\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSummingTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"respondToTask\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSummingTaskManager.Task\",\"components\":[{\"name\":\"arrayToBeSummed\",\"type\":\"uint64[3]\",\"internalType\":\"uint64[3]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSummingTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"arraySummed\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSummingTaskManager.Task\",\"components\":[{\"name\":\"arrayToBeSummed\",\"type\":\"uint64[3]\",\"internalType\":\"uint64[3]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSummingTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"arraySummed\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSummingTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b506040516200601a3803806200601a8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615d23620002f76000396000818161027d0152818161059c0152613203015260008181610552015261201b01526000818161040b015281816121fd0152612d47015260008181610432015281816123d3015261259501526000818161045901528181610e0b01528181611ce601528181611e7e01526120b80152615d236000f3fe608060405234801561001057600080fd5b506004361061021c5760003560e01c80636efb463611610125578063cefdc1d4116100ad578063f2fde38b1161007c578063f2fde38b14610587578063f5c9899d1461059a578063f63c5bab146105c0578063f8c8765e146105c8578063fabc1cbc146105db57600080fd5b8063cefdc1d414610519578063d1c2a3871461053a578063df5cf7231461054d578063e3f6ec281461057457600080fd5b8063886f1195116100f4578063886f1195146104c55780638b00ce7c146104d85780638d1e4cb2146104e85780638da5cb5b146104fb578063b98d09081461050c57600080fd5b80636efb46361461047b578063715018a61461049c57806372d18e8d146104a45780637afa1eed146104b257600080fd5b80634f739f74116101a85780635c975abb116101775780635c975abb146103db5780635decc3f5146103e35780635df4594614610406578063683048351461042d5780636d14a9871461045457600080fd5b80634f739f7414610360578063595c6a67146103805780635ac86ab7146103885780635c155662146103bb57600080fd5b8063245a7bfc116101ef578063245a7bfc146102b45780632cb223d5146102df5780632d89f6fc1461030d5780633563b0d11461032d578063416c7e5e1461034d57600080fd5b806310d67a2f14610221578063136439dd14610236578063171f1d5b146102495780631ad4318914610278575b600080fd5b61023461022f366004614753565b6105ee565b005b610234610244366004614770565b6106aa565b61025c6102573660046148ee565b6107e9565b6040805192151583529015156020830152015b60405180910390f35b61029f7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161026f565b60cd546102c7906001600160a01b031681565b6040516001600160a01b03909116815260200161026f565b6102ff6102ed36600461495c565b60cb6020526000908152604090205481565b60405190815260200161026f565b6102ff61031b36600461495c565b60ca6020526000908152604090205481565b61034061033b366004614979565b610973565b60405161026f9190614ad4565b61023461035b366004614afc565b610e09565b61037361036e366004614b61565b610f7e565b60405161026f9190614c65565b6102346116a4565b6103ab610396366004614d2f565b606654600160ff9092169190911b9081161490565b604051901515815260200161026f565b6103ce6103c9366004614d6f565b61176b565b60405161026f9190614e20565b6066546102ff565b6103ab6103f136600461495c565b60cc6020526000908152604090205460ff1681565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b61048e6104893660046150e0565b611933565b60405161026f9291906151a0565b61023461284a565b60c95463ffffffff1661029f565b60ce546102c7906001600160a01b031681565b6065546102c7906001600160a01b031681565b60c95461029f9063ffffffff1681565b6102346104f6366004615213565b61285e565b6033546001600160a01b03166102c7565b6097546103ab9060ff1681565b61052c610527366004615299565b612e80565b60405161026f9291906152db565b6102346105483660046152fc565b613012565b6102c77f000000000000000000000000000000000000000000000000000000000000000081565b610234610582366004615370565b613491565b610234610595366004614753565b613627565b7f000000000000000000000000000000000000000000000000000000000000000061029f565b61029f606481565b6102346105d63660046153d7565b61369d565b6102346105e9366004614770565b6137ee565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610641573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106659190615433565b6001600160a01b0316336001600160a01b03161461069e5760405162461bcd60e51b815260040161069590615450565b60405180910390fd5b6106a78161394a565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106f2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610716919061549a565b6107325760405162461bcd60e51b8152600401610695906154b7565b606654818116146107ab5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610695565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000187876000015188602001518860000151600060028110610831576108316154ff565b60200201518951600160200201518a60200151600060028110610856576108566154ff565b60200201518b60200151600160028110610872576108726154ff565b602090810291909101518c518d8301516040516108cf9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108f29190615515565b905061096561090b6109048884613a41565b8690613ad8565b610913613b6c565b61095b61094c85610946604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613a41565b6109558c613c2c565b90613ad8565b886201d4c0613cbc565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d99190615433565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3f9190615433565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa59190615433565b9050600086516001600160401b03811115610ac257610ac2614789565b604051908082528060200260200182016040528015610af557816020015b6060815260200190600190039081610ae05790505b50905060005b8751811015610dfd576000888281518110610b1857610b186154ff565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610b79573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ba19190810190615537565b905080516001600160401b03811115610bbc57610bbc614789565b604051908082528060200260200182016040528015610c0757816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610bda5790505b50848481518110610c1a57610c1a6154ff565b602002602001018190525060005b8151811015610de7576040518060600160405280876001600160a01b03166347b314e8858581518110610c5d57610c5d6154ff565b60200260200101516040518263ffffffff1660e01b8152600401610c8391815260200190565b602060405180830381865afa158015610ca0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc49190615433565b6001600160a01b03168152602001838381518110610ce457610ce46154ff565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610d1257610d126154ff565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610d6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d9291906155c7565b6001600160601b0316815250858581518110610db057610db06154ff565b60200260200101518281518110610dc957610dc96154ff565b60200260200101819052508080610ddf90615606565b915050610c28565b5050508080610df590615606565b915050610afb565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8b9190615433565b6001600160a01b0316336001600160a01b031614610f375760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610695565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b610fa96040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fe9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061100d9190615433565b905061103a6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061106a908b9089908990600401615621565b600060405180830381865afa158015611087573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110af919081019061566b565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906110e1908b908b908b90600401615722565b600060405180830381865afa1580156110fe573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611126919081019061566b565b6040820152856001600160401b0381111561114357611143614789565b60405190808252806020026020018201604052801561117657816020015b60608152602001906001900390816111615790505b50606082015260005b60ff81168711156115b5576000856001600160401b038111156111a4576111a4614789565b6040519080825280602002602001820160405280156111cd578160200160208202803683370190505b5083606001518360ff16815181106111e7576111e76154ff565b602002602001018190525060005b868110156114b55760008c6001600160a01b03166304ec63518a8a85818110611220576112206154ff565b905060200201358e8860000151868151811061123e5761123e6154ff565b60200260200101516040518463ffffffff1660e01b815260040161127b9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611298573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112bc919061574b565b90506001600160c01b0381166113605760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610695565b8a8a8560ff16818110611375576113756154ff565b6001600160c01b03841692013560f81c9190911c6001908116141590506114a257856001600160a01b031663dd9846b98a8a858181106113b7576113b76154ff565b905060200201358d8d8860ff168181106113d3576113d36154ff565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611429573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061144d9190615774565b85606001518560ff1681518110611466576114666154ff565b6020026020010151848151811061147f5761147f6154ff565b63ffffffff909216602092830291909101909101528261149e81615606565b9350505b50806114ad81615606565b9150506111f5565b506000816001600160401b038111156114d0576114d0614789565b6040519080825280602002602001820160405280156114f9578160200160208202803683370190505b50905060005b8281101561157a5784606001518460ff1681518110611520576115206154ff565b60200260200101518181518110611539576115396154ff565b6020026020010151828281518110611553576115536154ff565b63ffffffff909216602092830291909101909101528061157281615606565b9150506114ff565b508084606001518460ff1681518110611595576115956154ff565b6020026020010181905250505080806115ad90615791565b91505061117f565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061161a9190615433565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c9061164d908b908b908e906004016157b1565b600060405180830381865afa15801561166a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611692919081019061566b565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156116ec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611710919061549a565b61172c5760405162461bcd60e51b8152600401610695906154b7565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b815260040161179d9291906157db565b600060405180830381865afa1580156117ba573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526117e2919081019061566b565b9050600084516001600160401b038111156117ff576117ff614789565b604051908082528060200260200182016040528015611828578160200160208202803683370190505b50905060005b855181101561192957866001600160a01b03166304ec6351878381518110611858576118586154ff565b602002602001015187868581518110611873576118736154ff565b60200260200101516040518463ffffffff1660e01b81526004016118b09392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156118cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f1919061574b565b6001600160c01b031682828151811061190c5761190c6154ff565b60209081029190910101528061192181615606565b91505061182e565b5095945050505050565b60408051808201909152606080825260208201526000846119aa5760405162461bcd60e51b81526020600482015260376024820152600080516020615cce83398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610695565b604083015151851480156119c2575060a08301515185145b80156119d2575060c08301515185145b80156119e2575060e08301515185145b611a4c5760405162461bcd60e51b81526020600482015260416024820152600080516020615cce83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610695565b82515160208401515114611ac45760405162461bcd60e51b815260206004820152604460248201819052600080516020615cce833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610695565b4363ffffffff168463ffffffff1610611b335760405162461bcd60e51b815260206004820152603c6024820152600080516020615cce83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610695565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611b7457611b74614789565b604051908082528060200260200182016040528015611b9d578160200160208202803683370190505b506020820152866001600160401b03811115611bbb57611bbb614789565b604051908082528060200260200182016040528015611be4578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611c1857611c18614789565b604051908082528060200260200182016040528015611c41578160200160208202803683370190505b5081526020860151516001600160401b03811115611c6157611c61614789565b604051908082528060200260200182016040528015611c8a578160200160208202803683370190505b5081602001819052506000611d5c8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611d33573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d57919061582f565b613ee0565b905060005b876020015151811015611ff757611da688602001518281518110611d8757611d876154ff565b6020026020010151805160009081526020918201519091526040902090565b83602001518281518110611dbc57611dbc6154ff565b60209081029190910101528015611e7c576020830151611ddd60018361584c565b81518110611ded57611ded6154ff565b602002602001015160001c83602001518281518110611e0e57611e0e6154ff565b602002602001015160001c11611e7c576040805162461bcd60e51b8152602060048201526024810191909152600080516020615cce83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610695565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110611ec157611ec16154ff565b60200260200101518b8b600001518581518110611ee057611ee06154ff565b60200260200101516040518463ffffffff1660e01b8152600401611f1d9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611f3a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f5e919061574b565b6001600160c01b031683600001518281518110611f7d57611f7d6154ff565b602002602001018181525050611fe3610904611fb78486600001518581518110611fa957611fa96154ff565b602002602001015116613f73565b8a602001518481518110611fcd57611fcd6154ff565b6020026020010151613f9e90919063ffffffff16565b945080611fef81615606565b915050611d61565b505061200283614082565b60975490935060ff1660008161201957600061209b565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612077573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061209b9190615863565b905060005b8a8110156127195782156121fb578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106120f7576120f76154ff565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612137573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061215b9190615863565b612165919061587c565b116121fb5760405162461bcd60e51b81526020600482015260666024820152600080516020615cce83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610695565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061223c5761223c6154ff565b9050013560f81c60f81b60f81c8c8c60a001518581518110612260576122606154ff565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa1580156122bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122e09190615894565b6001600160401b0319166123038a604001518381518110611d8757611d876154ff565b67ffffffffffffffff19161461239f5760405162461bcd60e51b81526020600482015260616024820152600080516020615cce83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610695565b6123cf896040015182815181106123b8576123b86154ff565b602002602001015187613ad890919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110612412576124126154ff565b9050013560f81c60f81b60f81c8c8c60c001518581518110612436576124366154ff565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612492573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124b691906155c7565b856020015182815181106124cc576124cc6154ff565b6001600160601b039092166020928302919091018201528501518051829081106124f8576124f86154ff565b602002602001015185600001518281518110612516576125166154ff565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156127045761258e86600001518281518110612560576125606154ff565b60200260200101518f8f8681811061257a5761257a6154ff565b600192013560f81c9290921c811614919050565b156126f2577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106125d4576125d46154ff565b9050013560f81c60f81b60f81c8e896020015185815181106125f8576125f86154ff565b60200260200101518f60e001518881518110612616576126166154ff565b6020026020010151878151811061262f5761262f6154ff565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612693573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126b791906155c7565b87518051859081106126cb576126cb6154ff565b602002602001018181516126df91906158bf565b6001600160601b03169052506001909101905b806126fc81615606565b91505061253a565b5050808061271190615606565b9150506120a0565b5050506000806127338c868a606001518b608001516107e9565b91509150816127a45760405162461bcd60e51b81526020600482015260436024820152600080516020615cce83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610695565b806128055760405162461bcd60e51b81526020600482015260396024820152600080516020615cce83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610695565b505060008782602001516040516020016128209291906158e7565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b61285261411d565b61285c6000614177565b565b600061286d602085018561495c565b63ffffffff8116600090815260cb602052604090205490915085906128de5760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610695565b84846040516020016128f1929190615977565b60408051601f19818403018152918152815160209283012063ffffffff8516600090815260cb909352912054146129905760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610695565b63ffffffff8216600090815260cc602052604090205460ff1615612a285760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610695565b6064612a37602086018661495c565b612a4191906159ad565b63ffffffff164363ffffffff161115612ac25760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610695565b6000805b6003816001600160401b03161015612b245782816001600160401b031660038110612af357612af36154ff565b602002016020810190612b0691906159d5565b612b1090836159f0565b915080612b1c81615a12565b915050612ac6565b506000612b3760408801602089016159d5565b6001600160401b0383811691161490506001811415612b8c57604051339063ffffffff8616907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a350505050612e7a565b600085516001600160401b03811115612ba757612ba7614789565b604051908082528060200260200182016040528015612bd0578160200160208202803683370190505b50905060005b8651811015612c2357612bf4878281518110611d8757611d876154ff565b828281518110612c0657612c066154ff565b602090810291909101015280612c1b81615606565b915050612bd6565b506000612c3660808b0160608c0161495c565b82604051602001612c489291906158e7565b60405160208183030381529060405280519060200120905087602001358114612cf25760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610695565b600087516001600160401b03811115612d0d57612d0d614789565b604051908082528060200260200182016040528015612d36578160200160208202803683370190505b50905060005b8851811015612e29577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae6858381518110612d8657612d866154ff565b60200260200101516040518263ffffffff1660e01b8152600401612dac91815260200190565b602060405180830381865afa158015612dc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ded9190615433565b828281518110612dff57612dff6154ff565b6001600160a01b039092166020928302919091019091015280612e2181615606565b915050612d3c565b5063ffffffff8716600081815260cc6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a3505050505050505b50505050565b6040805160018082528183019092526000916060918391602080830190803683370190505090508481600081518110612ebb57612ebb6154ff565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e90612ef790889086906004016157db565b600060405180830381865afa158015612f14573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612f3c919081019061566b565b600081518110612f4e57612f4e6154ff565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa158015612fba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fde919061574b565b6001600160c01b031690506000612ff4826141c9565b9050816130028a838a610973565b9550955050505050935093915050565b60cd546001600160a01b0316331461306c5760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610695565b600061307e608085016060860161495c565b90503660006130906080870187615a39565b909250905060006130a760c0880160a0890161495c565b905060ca60006130ba602089018961495c565b63ffffffff1663ffffffff16815260200190815260200160002054876040516020016130e69190615a7f565b604051602081830303815290604052805190602001201461316f5760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610695565b600060cb8161318160208a018a61495c565b63ffffffff1663ffffffff16815260200190815260200160002054146131fe5760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610695565b6132287f0000000000000000000000000000000000000000000000000000000000000000856159ad565b63ffffffff164363ffffffff1611156132995760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610695565b6000866040516020016132ac9190615b51565b6040516020818303038152906040528051906020012090506000806132d48387878a8c611933565b9150915060005b858110156133d3578460ff16836020015182815181106132fd576132fd6154ff565b602002602001015161330f9190615b5f565b6001600160601b0316606484600001518381518110613330576133306154ff565b60200260200101516001600160601b031661334b9190615b8e565b10156133c1576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610695565b806133cb81615606565b9150506132db565b5060408051808201825263ffffffff43168152602080820184905291519091613400918c91849101615bad565b6040516020818303038152906040528051906020012060cb60008c600001602081019061342d919061495c565b63ffffffff1663ffffffff168152602001908152602001600020819055507fd182f9a99b685ad96deeeffe824b62b2e00b48247636e91fc1cd479177fca5c18a8260405161347c929190615bad565b60405180910390a15050505050505050505050565b60ce546001600160a01b031633146134f55760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610695565b6134fd614636565b60408051606081810190925290869060039083908390808284376000920191909152505050815263ffffffff438116602080840191909152908516606083015260408051601f850183900483028101830190915283815290849084908190840183828082843760009201919091525050505060408083019190915251613587908290602001615bd9565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907f9e2f71ee6820f1e61a65e453b006caef0f911d17e5fdf06d567711f5578525da906135ea908490615bd9565b60405180910390a260c9546136069063ffffffff1660016159ad565b60c9805463ffffffff191663ffffffff929092169190911790555050505050565b61362f61411d565b6001600160a01b0381166136945760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610695565b6106a781614177565b600054610100900460ff16158080156136bd5750600054600160ff909116105b806136d75750303b1580156136d7575060005460ff166001145b61373a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610695565b6000805460ff19166001179055801561375d576000805461ff0019166101001790555b613768856000614295565b61377184614177565b60cd80546001600160a01b038086166001600160a01b03199283161790925560ce80549285169290911691909117905580156137e7576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613841573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138659190615433565b6001600160a01b0316336001600160a01b0316146138955760405162461bcd60e51b815260040161069590615450565b6066541981196066541916146139135760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610695565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107de565b6001600160a01b0381166139d85760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610695565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152613a5d614664565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa9050808015613a9057613a92565bfe5b5080613ad05760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610695565b505092915050565b6040805180820190915260008082526020820152613af4614682565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015613a90575080613ad05760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610695565b613b746146a0565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082019091526000808252602082015260008080613c5c600080516020615cae83398151915286615515565b90505b613c688161437f565b9093509150600080516020615cae833981519152828309831415613ca2576040805180820190915290815260208101919091529392505050565b600080516020615cae833981519152600182089050613c5f565b604080518082018252868152602080820186905282518084019093528683528201849052600091829190613cee6146c5565b60005b6002811015613eb3576000613d07826006615b8e565b9050848260028110613d1b57613d1b6154ff565b60200201515183613d2d83600061587c565b600c8110613d3d57613d3d6154ff565b6020020152848260028110613d5457613d546154ff565b60200201516020015183826001613d6b919061587c565b600c8110613d7b57613d7b6154ff565b6020020152838260028110613d9257613d926154ff565b6020020151515183613da583600261587c565b600c8110613db557613db56154ff565b6020020152838260028110613dcc57613dcc6154ff565b6020020151516001602002015183613de583600361587c565b600c8110613df557613df56154ff565b6020020152838260028110613e0c57613e0c6154ff565b602002015160200151600060028110613e2757613e276154ff565b602002015183613e3883600461587c565b600c8110613e4857613e486154ff565b6020020152838260028110613e5f57613e5f6154ff565b602002015160200151600160028110613e7a57613e7a6154ff565b602002015183613e8b83600561587c565b600c8110613e9b57613e9b6154ff565b60200201525080613eab81615606565b915050613cf1565b50613ebc6146e4565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613eec84614401565b9050808360ff166001901b11613f6a5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610695565b90505b92915050565b6000805b8215613f6d57613f8860018461584c565b9092169180613f9681615c95565b915050613f77565b60408051808201909152600080825260208201526102008261ffff1610613ffa5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610695565b8161ffff166001141561400e575081613f6d565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff161061407757600161ffff871660ff83161c8116141561405a576140578484613ad8565b93505b6140648384613ad8565b92506201fffe600192831b16910161402a565b509195945050505050565b604080518082019091526000808252602082015281511580156140a757506020820151155b156140c5575050604080518082019091526000808252602082015290565b604051806040016040528083600001518152602001600080516020615cae83398151915284602001516140f89190615515565b61411090600080516020615cae83398151915261584c565b905292915050565b919050565b6033546001600160a01b0316331461285c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610695565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000806141d784613f73565b61ffff166001600160401b038111156141f2576141f2614789565b6040519080825280601f01601f19166020018201604052801561421c576020820181803683370190505b5090506000805b825182108015614234575061010081105b1561428b576001811b93508584161561427b578060f81b83838151811061425d5761425d6154ff565b60200101906001600160f81b031916908160001a9053508160010191505b61428481615606565b9050614223565b5090949350505050565b6065546001600160a01b03161580156142b657506001600160a01b03821615155b6143385760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610695565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a261437b8261394a565b5050565b60008080600080516020615cae8339815191526003600080516020615cae83398151915286600080516020615cae8339815191528889090908905060006143f5827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020615cae83398151915261458e565b91959194509092505050565b60006101008251111561448a5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610695565b815161449857506000919050565b600080836000815181106144ae576144ae6154ff565b0160200151600160f89190911c81901b92505b8451811015614585578481815181106144dc576144dc6154ff565b0160200151600160f89190911c1b91508282116145715760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610695565b9181179161457e81615606565b90506144c1565b50909392505050565b6000806145996146e4565b6145a1614702565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa9250828015613a9057508261462b5760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610695565b505195945050505050565b6040518060800160405280614649614664565b81526000602082018190526060604083018190529091015290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806146b3614720565b81526020016146c0614720565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146106a757600080fd5b60006020828403121561476557600080fd5b8135613f6a8161473e565b60006020828403121561478257600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156147c1576147c1614789565b60405290565b60405161010081016001600160401b03811182821017156147c1576147c1614789565b604051601f8201601f191681016001600160401b038111828210171561481257614812614789565b604052919050565b60006040828403121561482c57600080fd5b61483461479f565b9050813581526020820135602082015292915050565b600082601f83011261485b57600080fd5b604051604081018181106001600160401b038211171561487d5761487d614789565b806040525080604084018581111561489457600080fd5b845b81811015614077578035835260209283019201614896565b6000608082840312156148c057600080fd5b6148c861479f565b90506148d4838361484a565b81526148e3836040840161484a565b602082015292915050565b600080600080610120858703121561490557600080fd5b84359350614916866020870161481a565b925061492586606087016148ae565b91506149348660e0870161481a565b905092959194509250565b63ffffffff811681146106a757600080fd5b80356141188161493f565b60006020828403121561496e57600080fd5b8135613f6a8161493f565b60008060006060848603121561498e57600080fd5b83356149998161473e565b92506020848101356001600160401b03808211156149b657600080fd5b818701915087601f8301126149ca57600080fd5b8135818111156149dc576149dc614789565b6149ee601f8201601f191685016147ea565b91508082528884828501011115614a0457600080fd5b8084840185840137600084828401015250809450505050614a2760408501614951565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b86811015614ac6578385038a52825180518087529087019087870190845b81811015614ab157835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614a6d565b50509a87019a95505091850191600101614a4f565b509298975050505050505050565b602081526000614ae76020830184614a30565b9392505050565b80151581146106a757600080fd5b600060208284031215614b0e57600080fd5b8135613f6a81614aee565b60008083601f840112614b2b57600080fd5b5081356001600160401b03811115614b4257600080fd5b602083019150836020828501011115614b5a57600080fd5b9250929050565b60008060008060008060808789031215614b7a57600080fd5b8635614b858161473e565b95506020870135614b958161493f565b945060408701356001600160401b0380821115614bb157600080fd5b614bbd8a838b01614b19565b90965094506060890135915080821115614bd657600080fd5b818901915089601f830112614bea57600080fd5b813581811115614bf957600080fd5b8a60208260051b8501011115614c0e57600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b83811015614c5a57815163ffffffff1687529582019590820190600101614c38565b509495945050505050565b600060208083528351608082850152614c8160a0850182614c24565b905081850151601f1980868403016040870152614c9e8383614c24565b92506040870151915080868403016060870152614cbb8383614c24565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614d125784878303018452614d00828751614c24565b95880195938801939150600101614ce6565b509998505050505050505050565b60ff811681146106a757600080fd5b600060208284031215614d4157600080fd5b8135613f6a81614d20565b60006001600160401b03821115614d6557614d65614789565b5060051b60200190565b600080600060608486031215614d8457600080fd5b8335614d8f8161473e565b92506020848101356001600160401b03811115614dab57600080fd5b8501601f81018713614dbc57600080fd5b8035614dcf614dca82614d4c565b6147ea565b81815260059190911b82018301908381019089831115614dee57600080fd5b928401925b82841015614e0c57833582529284019290840190614df3565b8096505050505050614a2760408501614951565b6020808252825182820181905260009190848201906040850190845b81811015614e5857835183529284019291840191600101614e3c565b50909695505050505050565b600082601f830112614e7557600080fd5b81356020614e85614dca83614d4c565b82815260059290921b84018101918181019086841115614ea457600080fd5b8286015b84811015614ec8578035614ebb8161493f565b8352918301918301614ea8565b509695505050505050565b600082601f830112614ee457600080fd5b81356020614ef4614dca83614d4c565b82815260069290921b84018101918181019086841115614f1357600080fd5b8286015b84811015614ec857614f29888261481a565b835291830191604001614f17565b600082601f830112614f4857600080fd5b81356020614f58614dca83614d4c565b82815260059290921b84018101918181019086841115614f7757600080fd5b8286015b84811015614ec85780356001600160401b03811115614f9a5760008081fd5b614fa88986838b0101614e64565b845250918301918301614f7b565b60006101808284031215614fc957600080fd5b614fd16147c7565b905081356001600160401b0380821115614fea57600080fd5b614ff685838601614e64565b8352602084013591508082111561500c57600080fd5b61501885838601614ed3565b6020840152604084013591508082111561503157600080fd5b61503d85838601614ed3565b604084015261504f85606086016148ae565b60608401526150618560e0860161481a565b608084015261012084013591508082111561507b57600080fd5b61508785838601614e64565b60a08401526101408401359150808211156150a157600080fd5b6150ad85838601614e64565b60c08401526101608401359150808211156150c757600080fd5b506150d484828501614f37565b60e08301525092915050565b6000806000806000608086880312156150f857600080fd5b8535945060208601356001600160401b038082111561511657600080fd5b61512289838a01614b19565b9096509450604088013591506151378261493f565b9092506060870135908082111561514d57600080fd5b5061515a88828901614fb6565b9150509295509295909350565b600081518084526020808501945080840160005b83811015614c5a5781516001600160601b03168752958201959082019060010161517b565b60408152600083516040808401526151bb6080840182615167565b90506020850151603f198483030160608501526151d88282615167565b925050508260208301529392505050565b600060c082840312156151fb57600080fd5b50919050565b6000604082840312156151fb57600080fd5b60008060008060c0858703121561522957600080fd5b84356001600160401b038082111561524057600080fd5b61524c888389016151e9565b955061525b8860208901615201565b945061526a8860608901615201565b935060a087013591508082111561528057600080fd5b5061528d87828801614ed3565b91505092959194509250565b6000806000606084860312156152ae57600080fd5b83356152b98161473e565b92506020840135915060408401356152d08161493f565b809150509250925092565b8281526040602082015260006152f46040830184614a30565b949350505050565b60008060006080848603121561531157600080fd5b83356001600160401b038082111561532857600080fd5b615334878388016151e9565b94506153438760208801615201565b9350606086013591508082111561535957600080fd5b5061536686828701614fb6565b9150509250925092565b60008060008060a0858703121561538657600080fd5b606085018681111561539757600080fd5b859450356153a48161493f565b925060808501356001600160401b038111156153bf57600080fd5b6153cb87828801614b19565b95989497509550505050565b600080600080608085870312156153ed57600080fd5b84356153f88161473e565b935060208501356154088161473e565b925060408501356154188161473e565b915060608501356154288161473e565b939692955090935050565b60006020828403121561544557600080fd5b8151613f6a8161473e565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b6000602082840312156154ac57600080fd5b8151613f6a81614aee565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261553257634e487b7160e01b600052601260045260246000fd5b500690565b6000602080838503121561554a57600080fd5b82516001600160401b0381111561556057600080fd5b8301601f8101851361557157600080fd5b805161557f614dca82614d4c565b81815260059190911b8201830190838101908783111561559e57600080fd5b928401925b828410156155bc578351825292840192908401906155a3565b979650505050505050565b6000602082840312156155d957600080fd5b81516001600160601b0381168114613f6a57600080fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561561a5761561a6155f0565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561564e57600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561567e57600080fd5b82516001600160401b0381111561569457600080fd5b8301601f810185136156a557600080fd5b80516156b3614dca82614d4c565b81815260059190911b820183019083810190878311156156d257600080fd5b928401925b828410156155bc5783516156ea8161493f565b825292840192908401906156d7565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006157426040830184866156f9565b95945050505050565b60006020828403121561575d57600080fd5b81516001600160c01b0381168114613f6a57600080fd5b60006020828403121561578657600080fd5b8151613f6a8161493f565b600060ff821660ff8114156157a8576157a86155f0565b60010192915050565b6040815260006157c56040830185876156f9565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b8181101561582257845183529383019391830191600101615806565b5090979650505050505050565b60006020828403121561584157600080fd5b8151613f6a81614d20565b60008282101561585e5761585e6155f0565b500390565b60006020828403121561587557600080fd5b5051919050565b6000821982111561588f5761588f6155f0565b500190565b6000602082840312156158a657600080fd5b815167ffffffffffffffff1981168114613f6a57600080fd5b60006001600160601b03838116908316818110156158df576158df6155f0565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b8381101561592257815185529382019390820190600101615906565b5092979650505050505050565b80356001600160401b038116811461411857600080fd5b80356159518161493f565b63ffffffff1682526001600160401b0361596d6020830161592f565b1660208301525050565b608081016159858285615946565b82356159908161493f565b63ffffffff16604083015260209290920135606090910152919050565b600063ffffffff8083168185168083038211156159cc576159cc6155f0565b01949350505050565b6000602082840312156159e757600080fd5b614ae78261592f565b60006001600160401b038083168185168083038211156159cc576159cc6155f0565b60006001600160401b0380831681811415615a2f57615a2f6155f0565b6001019392505050565b6000808335601e19843603018112615a5057600080fd5b8301803591506001600160401b03821115615a6a57600080fd5b602001915036819003821315614b5a57600080fd5b602080825260009083838201835b6003811015615aba576001600160401b03615aa78461592f565b1682529183019190830190600101615a8d565b5050506060840135615acb8161493f565b63ffffffff81166080850152506080840135601e19853603018112615aef57600080fd5b840180356001600160401b03811115615b0757600080fd5b803603861315615b1657600080fd5b60c060a0860152615b2d60e08601828585016156f9565b92505050615b3d60a08501614951565b63ffffffff811660c0850152509392505050565b60408101613f6d8284615946565b60006001600160601b0380831681851681830481118215151615615b8557615b856155f0565b02949350505050565b6000816000190483118215151615615ba857615ba86155f0565b500290565b60808101615bbb8285615946565b63ffffffff8351166040830152602083015160608301529392505050565b6020808252825160009190828483015b6003821015615c115782516001600160401b0316815291830191600191909101908301615be9565b50505063ffffffff81850151166080840152604084015160c060a085015280518060e086015260005b81811015615c575782810184015186820161010001528301615c3a565b81811115615c6a57600061010083880101525b50606086015163ffffffff811660c08701529250601f01601f19169390930161010001949350505050565b600061ffff80831681811415615a2f57615a2f6155f056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122080fdf64eeb35941b3e4502660b91dc53c949252d2409b8e03f17f82f4d7def6964736f6c634300080c0033",
}

// ContractIncredibleSummingTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIncredibleSummingTaskManagerMetaData.ABI instead.
var ContractIncredibleSummingTaskManagerABI = ContractIncredibleSummingTaskManagerMetaData.ABI

// ContractIncredibleSummingTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIncredibleSummingTaskManagerMetaData.Bin instead.
var ContractIncredibleSummingTaskManagerBin = ContractIncredibleSummingTaskManagerMetaData.Bin

// DeployContractIncredibleSummingTaskManager deploys a new Ethereum contract, binding an instance of ContractIncredibleSummingTaskManager to it.
func DeployContractIncredibleSummingTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractIncredibleSummingTaskManager, error) {
	parsed, err := ContractIncredibleSummingTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIncredibleSummingTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIncredibleSummingTaskManager{ContractIncredibleSummingTaskManagerCaller: ContractIncredibleSummingTaskManagerCaller{contract: contract}, ContractIncredibleSummingTaskManagerTransactor: ContractIncredibleSummingTaskManagerTransactor{contract: contract}, ContractIncredibleSummingTaskManagerFilterer: ContractIncredibleSummingTaskManagerFilterer{contract: contract}}, nil
}

// ContractIncredibleSummingTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractIncredibleSummingTaskManager struct {
	ContractIncredibleSummingTaskManagerCaller     // Read-only binding to the contract
	ContractIncredibleSummingTaskManagerTransactor // Write-only binding to the contract
	ContractIncredibleSummingTaskManagerFilterer   // Log filterer for contract events
}

// ContractIncredibleSummingTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIncredibleSummingTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSummingTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIncredibleSummingTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSummingTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIncredibleSummingTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSummingTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIncredibleSummingTaskManagerSession struct {
	Contract     *ContractIncredibleSummingTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                         // Call options to use throughout this session
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractIncredibleSummingTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIncredibleSummingTaskManagerCallerSession struct {
	Contract *ContractIncredibleSummingTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                               // Call options to use throughout this session
}

// ContractIncredibleSummingTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIncredibleSummingTaskManagerTransactorSession struct {
	Contract     *ContractIncredibleSummingTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                               // Transaction auth options to use throughout this session
}

// ContractIncredibleSummingTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIncredibleSummingTaskManagerRaw struct {
	Contract *ContractIncredibleSummingTaskManager // Generic contract binding to access the raw methods on
}

// ContractIncredibleSummingTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIncredibleSummingTaskManagerCallerRaw struct {
	Contract *ContractIncredibleSummingTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIncredibleSummingTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIncredibleSummingTaskManagerTransactorRaw struct {
	Contract *ContractIncredibleSummingTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIncredibleSummingTaskManager creates a new instance of ContractIncredibleSummingTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSummingTaskManager(address common.Address, backend bind.ContractBackend) (*ContractIncredibleSummingTaskManager, error) {
	contract, err := bindContractIncredibleSummingTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManager{ContractIncredibleSummingTaskManagerCaller: ContractIncredibleSummingTaskManagerCaller{contract: contract}, ContractIncredibleSummingTaskManagerTransactor: ContractIncredibleSummingTaskManagerTransactor{contract: contract}, ContractIncredibleSummingTaskManagerFilterer: ContractIncredibleSummingTaskManagerFilterer{contract: contract}}, nil
}

// NewContractIncredibleSummingTaskManagerCaller creates a new read-only instance of ContractIncredibleSummingTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSummingTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIncredibleSummingTaskManagerCaller, error) {
	contract, err := bindContractIncredibleSummingTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerCaller{contract: contract}, nil
}

// NewContractIncredibleSummingTaskManagerTransactor creates a new write-only instance of ContractIncredibleSummingTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSummingTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIncredibleSummingTaskManagerTransactor, error) {
	contract, err := bindContractIncredibleSummingTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerTransactor{contract: contract}, nil
}

// NewContractIncredibleSummingTaskManagerFilterer creates a new log filterer instance of ContractIncredibleSummingTaskManager, bound to a specific deployed contract.
func NewContractIncredibleSummingTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIncredibleSummingTaskManagerFilterer, error) {
	contract, err := bindContractIncredibleSummingTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerFilterer{contract: contract}, nil
}

// bindContractIncredibleSummingTaskManager binds a generic wrapper to an already deployed contract.
func bindContractIncredibleSummingTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIncredibleSummingTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSummingTaskManager.Contract.ContractIncredibleSummingTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.ContractIncredibleSummingTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.ContractIncredibleSummingTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSummingTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSummingTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSummingTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSummingTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractIncredibleSummingTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Aggregator(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Aggregator(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSummingTaskManager.Contract.AllTaskHashes(&_ContractIncredibleSummingTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSummingTaskManager.Contract.AllTaskHashes(&_ContractIncredibleSummingTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSummingTaskManager.Contract.AllTaskResponses(&_ContractIncredibleSummingTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractIncredibleSummingTaskManager.Contract.AllTaskResponses(&_ContractIncredibleSummingTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.BlsApkRegistry(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.BlsApkRegistry(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleSummingTaskManager.Contract.CheckSignatures(&_ContractIncredibleSummingTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractIncredibleSummingTaskManager.Contract.CheckSignatures(&_ContractIncredibleSummingTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Delegation(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Delegation(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) Generator() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Generator(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Generator(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleSummingTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleSummingTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractIncredibleSummingTaskManager.Contract.GetCheckSignaturesIndices(&_ContractIncredibleSummingTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSummingTaskManager.Contract.GetOperatorState(&_ContractIncredibleSummingTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSummingTaskManager.Contract.GetOperatorState(&_ContractIncredibleSummingTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSummingTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSummingTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractIncredibleSummingTaskManager.Contract.GetOperatorState0(&_ContractIncredibleSummingTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractIncredibleSummingTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractIncredibleSummingTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractIncredibleSummingTaskManager.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractIncredibleSummingTaskManager.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleSummingTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractIncredibleSummingTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleSummingTaskManager.Contract.LatestTaskNum(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractIncredibleSummingTaskManager.Contract.LatestTaskNum(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Owner(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Owner(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Paused(&_ContractIncredibleSummingTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Paused(&_ContractIncredibleSummingTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Paused0(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Paused0(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.PauserRegistry(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.PauserRegistry(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.RegistryCoordinator(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.StakeRegistry(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractIncredibleSummingTaskManager.Contract.StakeRegistry(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractIncredibleSummingTaskManager.Contract.StaleStakesForbidden(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractIncredibleSummingTaskManager.Contract.StaleStakesForbidden(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleSummingTaskManager.Contract.TaskNumber(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractIncredibleSummingTaskManager.Contract.TaskNumber(&_ContractIncredibleSummingTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleSummingTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractIncredibleSummingTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractIncredibleSummingTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractIncredibleSummingTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractIncredibleSummingTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractIncredibleSummingTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleSummingTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractIncredibleSummingTaskManager.Contract.TrySignatureAndApkVerification(&_ContractIncredibleSummingTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xe3f6ec28.
//
// Solidity: function createNewTask(uint64[3] arrayToBeSummed, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, arrayToBeSummed [3]uint64, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.contract.Transact(opts, "createNewTask", arrayToBeSummed, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xe3f6ec28.
//
// Solidity: function createNewTask(uint64[3] arrayToBeSummed, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) CreateNewTask(arrayToBeSummed [3]uint64, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.CreateNewTask(&_ContractIncredibleSummingTaskManager.TransactOpts, arrayToBeSummed, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xe3f6ec28.
//
// Solidity: function createNewTask(uint64[3] arrayToBeSummed, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorSession) CreateNewTask(arrayToBeSummed [3]uint64, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.CreateNewTask(&_ContractIncredibleSummingTaskManager.TransactOpts, arrayToBeSummed, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Initialize(&_ContractIncredibleSummingTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Initialize(&_ContractIncredibleSummingTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Pause(&_ContractIncredibleSummingTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Pause(&_ContractIncredibleSummingTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.PauseAll(&_ContractIncredibleSummingTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.PauseAll(&_ContractIncredibleSummingTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x8d1e4cb2.
//
// Solidity: function raiseAndResolveChallenge((uint64[3],uint32,bytes,uint32) task, (uint32,uint64) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IIncredibleSummingTaskManagerTask, taskResponse IIncredibleSummingTaskManagerTaskResponse, taskResponseMetadata IIncredibleSummingTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x8d1e4cb2.
//
// Solidity: function raiseAndResolveChallenge((uint64[3],uint32,bytes,uint32) task, (uint32,uint64) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) RaiseAndResolveChallenge(task IIncredibleSummingTaskManagerTask, taskResponse IIncredibleSummingTaskManagerTaskResponse, taskResponseMetadata IIncredibleSummingTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleSummingTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0x8d1e4cb2.
//
// Solidity: function raiseAndResolveChallenge((uint64[3],uint32,bytes,uint32) task, (uint32,uint64) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorSession) RaiseAndResolveChallenge(task IIncredibleSummingTaskManagerTask, taskResponse IIncredibleSummingTaskManagerTaskResponse, taskResponseMetadata IIncredibleSummingTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleSummingTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.RenounceOwnership(&_ContractIncredibleSummingTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.RenounceOwnership(&_ContractIncredibleSummingTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xd1c2a387.
//
// Solidity: function respondToTask((uint64[3],uint32,bytes,uint32) task, (uint32,uint64) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IIncredibleSummingTaskManagerTask, taskResponse IIncredibleSummingTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xd1c2a387.
//
// Solidity: function respondToTask((uint64[3],uint32,bytes,uint32) task, (uint32,uint64) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) RespondToTask(task IIncredibleSummingTaskManagerTask, taskResponse IIncredibleSummingTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.RespondToTask(&_ContractIncredibleSummingTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xd1c2a387.
//
// Solidity: function respondToTask((uint64[3],uint32,bytes,uint32) task, (uint32,uint64) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorSession) RespondToTask(task IIncredibleSummingTaskManagerTask, taskResponse IIncredibleSummingTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.RespondToTask(&_ContractIncredibleSummingTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.SetPauserRegistry(&_ContractIncredibleSummingTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.SetPauserRegistry(&_ContractIncredibleSummingTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.SetStaleStakesForbidden(&_ContractIncredibleSummingTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.SetStaleStakesForbidden(&_ContractIncredibleSummingTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.TransferOwnership(&_ContractIncredibleSummingTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.TransferOwnership(&_ContractIncredibleSummingTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Unpause(&_ContractIncredibleSummingTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.Unpause(&_ContractIncredibleSummingTaskManager.TransactOpts, newPausedStatus)
}

// ContractIncredibleSummingTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerInitializedIterator struct {
	Event *ContractIncredibleSummingTaskManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSummingTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingTaskManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSummingTaskManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSummingTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingTaskManagerInitialized represents a Initialized event raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractIncredibleSummingTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerInitializedIterator{contract: _ContractIncredibleSummingTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingTaskManagerInitialized)
				if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractIncredibleSummingTaskManagerInitialized, error) {
	event := new(ContractIncredibleSummingTaskManagerInitialized)
	if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSummingTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerNewTaskCreatedIterator struct {
	Event *ContractIncredibleSummingTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSummingTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingTaskManagerNewTaskCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSummingTaskManagerNewTaskCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSummingTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IIncredibleSummingTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x9e2f71ee6820f1e61a65e453b006caef0f911d17e5fdf06d567711f5578525da.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint64[3],uint32,bytes,uint32) task)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSummingTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerNewTaskCreatedIterator{contract: _ContractIncredibleSummingTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x9e2f71ee6820f1e61a65e453b006caef0f911d17e5fdf06d567711f5578525da.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint64[3],uint32,bytes,uint32) task)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingTaskManagerNewTaskCreated)
				if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTaskCreated is a log parse operation binding the contract event 0x9e2f71ee6820f1e61a65e453b006caef0f911d17e5fdf06d567711f5578525da.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint64[3],uint32,bytes,uint32) task)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractIncredibleSummingTaskManagerNewTaskCreated, error) {
	event := new(ContractIncredibleSummingTaskManagerNewTaskCreated)
	if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSummingTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerOwnershipTransferredIterator struct {
	Event *ContractIncredibleSummingTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSummingTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingTaskManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSummingTaskManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSummingTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractIncredibleSummingTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerOwnershipTransferredIterator{contract: _ContractIncredibleSummingTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingTaskManagerOwnershipTransferred)
				if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractIncredibleSummingTaskManagerOwnershipTransferred, error) {
	event := new(ContractIncredibleSummingTaskManagerOwnershipTransferred)
	if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSummingTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerPausedIterator struct {
	Event *ContractIncredibleSummingTaskManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSummingTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingTaskManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSummingTaskManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSummingTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingTaskManagerPaused represents a Paused event raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleSummingTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerPausedIterator{contract: _ContractIncredibleSummingTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingTaskManagerPaused)
				if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) ParsePaused(log types.Log) (*ContractIncredibleSummingTaskManagerPaused, error) {
	event := new(ContractIncredibleSummingTaskManagerPaused)
	if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSummingTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerPauserRegistrySetIterator struct {
	Event *ContractIncredibleSummingTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSummingTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingTaskManagerPauserRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSummingTaskManagerPauserRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSummingTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractIncredibleSummingTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerPauserRegistrySetIterator{contract: _ContractIncredibleSummingTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingTaskManagerPauserRegistrySet)
				if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractIncredibleSummingTaskManagerPauserRegistrySet, error) {
	event := new(ContractIncredibleSummingTaskManagerPauserRegistrySet)
	if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractIncredibleSummingTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractIncredibleSummingTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSummingTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractIncredibleSummingTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSummingTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingTaskManagerTaskChallengedSuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSummingTaskManagerTaskChallengedSuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSummingTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleSummingTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractIncredibleSummingTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingTaskManagerTaskChallengedSuccessfully)
				if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractIncredibleSummingTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractIncredibleSummingTaskManagerTaskChallengedSuccessfully)
	if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfully)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfully)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractIncredibleSummingTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractIncredibleSummingTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSummingTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerTaskCompletedIterator struct {
	Event *ContractIncredibleSummingTaskManagerTaskCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSummingTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingTaskManagerTaskCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSummingTaskManagerTaskCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSummingTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractIncredibleSummingTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerTaskCompletedIterator{contract: _ContractIncredibleSummingTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingTaskManagerTaskCompleted)
				if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractIncredibleSummingTaskManagerTaskCompleted, error) {
	event := new(ContractIncredibleSummingTaskManagerTaskCompleted)
	if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSummingTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerTaskRespondedIterator struct {
	Event *ContractIncredibleSummingTaskManagerTaskResponded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSummingTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingTaskManagerTaskResponded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSummingTaskManagerTaskResponded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSummingTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingTaskManagerTaskResponded represents a TaskResponded event raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerTaskResponded struct {
	TaskResponse         IIncredibleSummingTaskManagerTaskResponse
	TaskResponseMetadata IIncredibleSummingTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xd182f9a99b685ad96deeeffe824b62b2e00b48247636e91fc1cd479177fca5c1.
//
// Solidity: event TaskResponded((uint32,uint64) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractIncredibleSummingTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerTaskRespondedIterator{contract: _ContractIncredibleSummingTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xd182f9a99b685ad96deeeffe824b62b2e00b48247636e91fc1cd479177fca5c1.
//
// Solidity: event TaskResponded((uint32,uint64) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingTaskManagerTaskResponded)
				if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTaskResponded is a log parse operation binding the contract event 0xd182f9a99b685ad96deeeffe824b62b2e00b48247636e91fc1cd479177fca5c1.
//
// Solidity: event TaskResponded((uint32,uint64) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractIncredibleSummingTaskManagerTaskResponded, error) {
	event := new(ContractIncredibleSummingTaskManagerTaskResponded)
	if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSummingTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerUnpausedIterator struct {
	Event *ContractIncredibleSummingTaskManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractIncredibleSummingTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingTaskManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractIncredibleSummingTaskManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractIncredibleSummingTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingTaskManagerUnpaused represents a Unpaused event raised by the ContractIncredibleSummingTaskManager contract.
type ContractIncredibleSummingTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractIncredibleSummingTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerUnpausedIterator{contract: _ContractIncredibleSummingTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingTaskManagerUnpaused)
				if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractIncredibleSummingTaskManagerUnpaused, error) {
	event := new(ContractIncredibleSummingTaskManagerUnpaused)
	if err := _ContractIncredibleSummingTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
