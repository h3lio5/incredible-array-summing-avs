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
	TaskResponsedBlocktime uint64
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allTaskResponses\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createNewTask\",\"inputs\":[{\"name\":\"arrayToBeSummed\",\"type\":\"uint64[3]\",\"internalType\":\"uint64[3]\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResponseWindowBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_aggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_generator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestTaskNum\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"raiseAndResolveChallenge\",\"inputs\":[{\"name\":\"task\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSummingTaskManager.Task\",\"components\":[{\"name\":\"arrayToBeSummed\",\"type\":\"uint64[3]\",\"internalType\":\"uint64[3]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"taskResponse\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSummingTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"arraySummed\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"internalType\":\"structIIncredibleSummingTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlocktime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApks\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taskSuccesfullyChallenged\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trySignatureAndApkVerification\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"apk\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"pairingSuccessful\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"siganatureIsValid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTaskCreated\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"task\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSummingTaskManager.Task\",\"components\":[{\"name\":\"arrayToBeSummed\",\"type\":\"uint64[3]\",\"internalType\":\"uint64[3]\"},{\"name\":\"taskCreatedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedSuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskChallengedUnsuccessfully\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskCompleted\",\"inputs\":[{\"name\":\"taskIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskResponded\",\"inputs\":[{\"name\":\"taskResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSummingTaskManager.TaskResponse\",\"components\":[{\"name\":\"referenceTaskIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"arraySummed\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"taskResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIIncredibleSummingTaskManager.TaskResponseMetadata\",\"components\":[{\"name\":\"taskResponsedBlocktime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b506040516200596a3803806200596a8339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e0516101005161567a620002f060003960008181610272015261057e0152600081816105210152611ffd015260008181610400015281816121df0152612ffd015260008181610427015281816123b5015261257701526000818161044e01528181610ded01528181611cc801528181611e60015261209a015261567a6000f3fe608060405234801561001057600080fd5b50600436106102115760003560e01c80636d14a98711610125578063cefdc1d4116100ad578063f2fde38b1161007c578063f2fde38b14610569578063f5c9899d1461057c578063f63c5bab146105a2578063f8c8765e146105aa578063fabc1cbc146105bd57600080fd5b8063cefdc1d4146104fb578063df5cf7231461051c578063e3f6ec2814610543578063ee1c87c51461055657600080fd5b80637afa1eed116100f45780637afa1eed146104a7578063886f1195146104ba5780638b00ce7c146104cd5780638da5cb5b146104dd578063b98d0908146104ee57600080fd5b80636d14a987146104495780636efb463614610470578063715018a61461049157806372d18e8d1461049957600080fd5b8063416c7e5e116101a85780635c155662116101775780635c155662146103b05780635c975abb146103d05780635decc3f5146103d85780635df45946146103fb578063683048351461042257600080fd5b8063416c7e5e146103425780634f739f7414610355578063595c6a67146103755780635ac86ab71461037d57600080fd5b8063245a7bfc116101e4578063245a7bfc146102a95780632cb223d5146102d45780632d89f6fc146103025780633563b0d11461032257600080fd5b806310d67a2f14610216578063136439dd1461022b578063171f1d5b1461023e5780631ad431891461026d575b600080fd5b610229610224366004614268565b6105d0565b005b610229610239366004614285565b61068c565b61025161024c366004614403565b6107cb565b6040805192151583529015156020830152015b60405180910390f35b6102947f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610264565b60cd546102bc906001600160a01b031681565b6040516001600160a01b039091168152602001610264565b6102f46102e2366004614471565b60cb6020526000908152604090205481565b604051908152602001610264565b6102f4610310366004614471565b60ca6020526000908152604090205481565b61033561033036600461448e565b610955565b60405161026491906145e9565b610229610350366004614611565b610deb565b610368610363366004614676565b610f60565b604051610264919061477a565b610229611686565b6103a061038b366004614844565b606654600160ff9092169190911b9081161490565b6040519015158152602001610264565b6103c36103be366004614884565b61174d565b6040516102649190614935565b6066546102f4565b6103a06103e6366004614471565b60cc6020526000908152604090205460ff1681565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b61048361047e366004614bf5565b611915565b604051610264929190614cb5565b61022961282c565b60c95463ffffffff16610294565b60ce546102bc906001600160a01b031681565b6065546102bc906001600160a01b031681565b60c9546102949063ffffffff1681565b6033546001600160a01b03166102bc565b6097546103a09060ff1681565b61050e610509366004614cfe565b612840565b604051610264929190614d40565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b610229610551366004614d61565b6129d2565b610229610564366004614dc8565b612b68565b610229610577366004614268565b61313c565b7f0000000000000000000000000000000000000000000000000000000000000000610294565b610294606481565b6102296105b8366004614e6a565b6131b2565b6102296105cb366004614285565b613303565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610623573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106479190614ec6565b6001600160a01b0316336001600160a01b0316146106805760405162461bcd60e51b815260040161067790614ee3565b60405180910390fd5b6106898161345f565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f89190614f2d565b6107145760405162461bcd60e51b815260040161067790614f4a565b6066548181161461078d5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610677565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061081357610813614f92565b60200201518951600160200201518a6020015160006002811061083857610838614f92565b60200201518b6020015160016002811061085457610854614f92565b602090810291909101518c518d8301516040516108b19a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108d49190614fa8565b90506109476108ed6108e68884613556565b86906135ed565b6108f5613681565b61093d61092e85610928604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613556565b6109378c613741565b906135ed565b886201d4c06137d1565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610997573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109bb9190614ec6565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a219190614ec6565b90506000866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a879190614ec6565b9050600086516001600160401b03811115610aa457610aa461429e565b604051908082528060200260200182016040528015610ad757816020015b6060815260200190600190039081610ac25790505b50905060005b8751811015610ddf576000888281518110610afa57610afa614f92565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291506000906001600160a01b03871690638902624590604401600060405180830381865afa158015610b5b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b839190810190614fca565b905080516001600160401b03811115610b9e57610b9e61429e565b604051908082528060200260200182016040528015610be957816020015b6040805160608101825260008082526020808301829052928201528252600019909201910181610bbc5790505b50848481518110610bfc57610bfc614f92565b602002602001018190525060005b8151811015610dc9576040518060600160405280876001600160a01b03166347b314e8858581518110610c3f57610c3f614f92565b60200260200101516040518263ffffffff1660e01b8152600401610c6591815260200190565b602060405180830381865afa158015610c82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca69190614ec6565b6001600160a01b03168152602001838381518110610cc657610cc6614f92565b60200260200101518152602001896001600160a01b031663fa28c627858581518110610cf457610cf4614f92565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610d50573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d74919061505a565b6001600160601b0316815250858581518110610d9257610d92614f92565b60200260200101518281518110610dab57610dab614f92565b60200260200101819052508080610dc190615099565b915050610c0a565b5050508080610dd790615099565b915050610add565b50979650505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6d9190614ec6565b6001600160a01b0316336001600160a01b031614610f195760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610677565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b610f8b6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fcb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fef9190614ec6565b905061101c6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061104c908b90899089906004016150b4565b600060405180830381865afa158015611069573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261109191908101906150fe565b81526040516340e03a8160e11b81526001600160a01b038316906381c07502906110c3908b908b908b906004016151b5565b600060405180830381865afa1580156110e0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261110891908101906150fe565b6040820152856001600160401b038111156111255761112561429e565b60405190808252806020026020018201604052801561115857816020015b60608152602001906001900390816111435790505b50606082015260005b60ff8116871115611597576000856001600160401b038111156111865761118661429e565b6040519080825280602002602001820160405280156111af578160200160208202803683370190505b5083606001518360ff16815181106111c9576111c9614f92565b602002602001018190525060005b868110156114975760008c6001600160a01b03166304ec63518a8a8581811061120257611202614f92565b905060200201358e8860000151868151811061122057611220614f92565b60200260200101516040518463ffffffff1660e01b815260040161125d9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561127a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061129e91906151de565b90506001600160c01b0381166113425760405162461bcd60e51b815260206004820152605c60248201527f4f70657261746f7253746174655265747269657665722e676574436865636b5360448201527f69676e617475726573496e64696365733a206f70657261746f72206d7573742060648201527f6265207265676973746572656420617420626c6f636b6e756d62657200000000608482015260a401610677565b8a8a8560ff1681811061135757611357614f92565b6001600160c01b03841692013560f81c9190911c60019081161415905061148457856001600160a01b031663dd9846b98a8a8581811061139957611399614f92565b905060200201358d8d8860ff168181106113b5576113b5614f92565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa15801561140b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061142f9190615207565b85606001518560ff168151811061144857611448614f92565b6020026020010151848151811061146157611461614f92565b63ffffffff909216602092830291909101909101528261148081615099565b9350505b508061148f81615099565b9150506111d7565b506000816001600160401b038111156114b2576114b261429e565b6040519080825280602002602001820160405280156114db578160200160208202803683370190505b50905060005b8281101561155c5784606001518460ff168151811061150257611502614f92565b6020026020010151818151811061151b5761151b614f92565b602002602001015182828151811061153557611535614f92565b63ffffffff909216602092830291909101909101528061155481615099565b9150506114e1565b508084606001518460ff168151811061157757611577614f92565b60200260200101819052505050808061158f90615224565b915050611161565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115fc9190614ec6565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c9061162f908b908b908e90600401615244565b600060405180830381865afa15801561164c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261167491908101906150fe565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156116ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f29190614f2d565b61170e5760405162461bcd60e51b815260040161067790614f4a565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b60606000846001600160a01b031663c391425e84866040518363ffffffff1660e01b815260040161177f92919061526e565b600060405180830381865afa15801561179c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526117c491908101906150fe565b9050600084516001600160401b038111156117e1576117e161429e565b60405190808252806020026020018201604052801561180a578160200160208202803683370190505b50905060005b855181101561190b57866001600160a01b03166304ec635187838151811061183a5761183a614f92565b60200260200101518786858151811061185557611855614f92565b60200260200101516040518463ffffffff1660e01b81526004016118929392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa1580156118af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118d391906151de565b6001600160c01b03168282815181106118ee576118ee614f92565b60209081029190910101528061190381615099565b915050611810565b5095945050505050565b604080518082019091526060808252602082015260008461198c5760405162461bcd60e51b8152602060048201526037602482015260008051602061562583398151915260448201527f7265733a20656d7074792071756f72756d20696e7075740000000000000000006064820152608401610677565b604083015151851480156119a4575060a08301515185145b80156119b4575060c08301515185145b80156119c4575060e08301515185145b611a2e5760405162461bcd60e51b8152602060048201526041602482015260008051602061562583398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610677565b82515160208401515114611aa65760405162461bcd60e51b815260206004820152604460248201819052600080516020615625833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610677565b4363ffffffff168463ffffffff1610611b155760405162461bcd60e51b815260206004820152603c602482015260008051602061562583398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610677565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b03811115611b5657611b5661429e565b604051908082528060200260200182016040528015611b7f578160200160208202803683370190505b506020820152866001600160401b03811115611b9d57611b9d61429e565b604051908082528060200260200182016040528015611bc6578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b03811115611bfa57611bfa61429e565b604051908082528060200260200182016040528015611c23578160200160208202803683370190505b5081526020860151516001600160401b03811115611c4357611c4361429e565b604051908082528060200260200182016040528015611c6c578160200160208202803683370190505b5081602001819052506000611d3e8a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa158015611d15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d3991906152c2565b6139f5565b905060005b876020015151811015611fd957611d8888602001518281518110611d6957611d69614f92565b6020026020010151805160009081526020918201519091526040902090565b83602001518281518110611d9e57611d9e614f92565b60209081029190910101528015611e5e576020830151611dbf6001836152df565b81518110611dcf57611dcf614f92565b602002602001015160001c83602001518281518110611df057611df0614f92565b602002602001015160001c11611e5e576040805162461bcd60e51b815260206004820152602481019190915260008051602061562583398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610677565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110611ea357611ea3614f92565b60200260200101518b8b600001518581518110611ec257611ec2614f92565b60200260200101516040518463ffffffff1660e01b8152600401611eff9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611f1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f4091906151de565b6001600160c01b031683600001518281518110611f5f57611f5f614f92565b602002602001018181525050611fc56108e6611f998486600001518581518110611f8b57611f8b614f92565b602002602001015116613a88565b8a602001518481518110611faf57611faf614f92565b6020026020010151613ab390919063ffffffff16565b945080611fd181615099565b915050611d43565b5050611fe483613b97565b60975490935060ff16600081611ffb57600061207d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c448feb86040518163ffffffff1660e01b8152600401602060405180830381865afa158015612059573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061207d91906152f6565b905060005b8a8110156126fb5782156121dd578963ffffffff16827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f868181106120d9576120d9614f92565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015612119573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061213d91906152f6565b612147919061530f565b116121dd5760405162461bcd60e51b8152602060048201526066602482015260008051602061562583398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610677565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d8481811061221e5761221e614f92565b9050013560f81c60f81b60f81c8c8c60a00151858151811061224257612242614f92565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561229e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122c29190615327565b6001600160401b0319166122e58a604001518381518110611d6957611d69614f92565b67ffffffffffffffff1916146123815760405162461bcd60e51b8152602060048201526061602482015260008051602061562583398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610677565b6123b18960400151828151811061239a5761239a614f92565b6020026020010151876135ed90919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d848181106123f4576123f4614f92565b9050013560f81c60f81b60f81c8c8c60c00151858151811061241857612418614f92565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015612474573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612498919061505a565b856020015182815181106124ae576124ae614f92565b6001600160601b039092166020928302919091018201528501518051829081106124da576124da614f92565b6020026020010151856000015182815181106124f8576124f8614f92565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156126e6576125708660000151828151811061254257612542614f92565b60200260200101518f8f8681811061255c5761255c614f92565b600192013560f81c9290921c811614919050565b156126d4577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106125b6576125b6614f92565b9050013560f81c60f81b60f81c8e896020015185815181106125da576125da614f92565b60200260200101518f60e0015188815181106125f8576125f8614f92565b6020026020010151878151811061261157612611614f92565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa158015612675573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612699919061505a565b87518051859081106126ad576126ad614f92565b602002602001018181516126c19190615352565b6001600160601b03169052506001909101905b806126de81615099565b91505061251c565b505080806126f390615099565b915050612082565b5050506000806127158c868a606001518b608001516107cb565b91509150816127865760405162461bcd60e51b8152602060048201526043602482015260008051602061562583398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610677565b806127e75760405162461bcd60e51b8152602060048201526039602482015260008051602061562583398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610677565b5050600087826020015160405160200161280292919061537a565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b612834613c32565b61283e6000613c8c565b565b604080516001808252818301909252600091606091839160208083019080368337019050509050848160008151811061287b5761287b614f92565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906128b7908890869060040161526e565b600060405180830381865afa1580156128d4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526128fc91908101906150fe565b60008151811061290e5761290e614f92565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa15801561297a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061299e91906151de565b6001600160c01b0316905060006129b482613cde565b9050816129c28a838a610955565b9550955050505050935093915050565b60ce546001600160a01b03163314612a365760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610677565b612a3e61414b565b60408051606081810190925290869060039083908390808284376000920191909152505050815263ffffffff438116602080840191909152908516606083015260408051601f850183900483028101830190915283815290849084908190840183828082843760009201919091525050505060408083019190915251612ac89082906020016153c2565b60408051601f19818403018152828252805160209182012060c9805463ffffffff908116600090815260ca90945293909220555416907f9e2f71ee6820f1e61a65e453b006caef0f911d17e5fdf06d567711f5578525da90612b2b9084906153c2565b60405180910390a260c954612b479063ffffffff16600161547e565b60c9805463ffffffff191663ffffffff929092169190911790555050505050565b6000612b776020850185614471565b9050846000612b8c6080830160608401614471565b9050366000612b9e60808a018a6154a6565b90925090506000612bb560c08b0160a08c01614471565b63ffffffff8716600090815260cc602052604090205490915060ff1615612c505760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610677565b600089604051602001612c639190615503565b604051602081830303815290604052805190602001209050600080612c8b8387878a8e611915565b9150915060005b85811015612d8a578460ff1683602001518281518110612cb457612cb4614f92565b6020026020010151612cc6919061553a565b6001600160601b0316606484600001518381518110612ce757612ce7614f92565b60200260200101516001600160601b0316612d029190615569565b1015612d78576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610677565b80612d8281615099565b915050612c92565b506064612d9a60208d018d615588565b612da491906155a3565b6001600160401b03164363ffffffff161115612e285760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610677565b6000805b6003816001600160401b03161015612e8a5789816001600160401b031660038110612e5957612e59614f92565b602002016020810190612e6c9190615588565b612e7690836155a3565b915080612e82816155c5565b915050612e2c565b5060008d6020016020810190612ea09190615588565b6001600160401b0383811691161490506001811415612efc57604051339063ffffffff8d16907ffd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb0590600090a35050505050505050505050613136565b60008c60200151516001600160401b03811115612f1b57612f1b61429e565b604051908082528060200260200182016040528015612f44578160200160208202803683370190505b50905060005b8d6020015151811015612f9f57612f708e602001518281518110611d6957611d69614f92565b828281518110612f8257612f82614f92565b602090810291909101015280612f9781615099565b915050612f4a565b5060008d60200151516001600160401b03811115612fbf57612fbf61429e565b604051908082528060200260200182016040528015612fe8578160200160208202803683370190505b50905060005b8e60200151518110156130df577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae684838151811061303c5761303c614f92565b60200260200101516040518263ffffffff1660e01b815260040161306291815260200190565b602060405180830381865afa15801561307f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130a39190614ec6565b8282815181106130b5576130b5614f92565b6001600160a01b0390921660209283029190910190910152806130d781615099565b915050612fee565b5063ffffffff8d16600081815260cc6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a3505050505050505050505050505b50505050565b613144613c32565b6001600160a01b0381166131a95760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610677565b61068981613c8c565b600054610100900460ff16158080156131d25750600054600160ff909116105b806131ec5750303b1580156131ec575060005460ff166001145b61324f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610677565b6000805460ff191660011790558015613272576000805461ff0019166101001790555b61327d856000613daa565b61328684613c8c565b60cd80546001600160a01b038086166001600160a01b03199283161790925560ce80549285169290911691909117905580156132fc576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613356573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061337a9190614ec6565b6001600160a01b0316336001600160a01b0316146133aa5760405162461bcd60e51b815260040161067790614ee3565b6066541981196066541916146134285760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610677565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107c0565b6001600160a01b0381166134ed5760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610677565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b6040805180820190915260008082526020820152613572614179565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa90508080156135a5576135a7565bfe5b50806135e55760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610677565b505092915050565b6040805180820190915260008082526020820152613609614197565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa90508080156135a55750806135e55760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610677565b6136896141b5565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061377160008051602061560583398151915286614fa8565b90505b61377d81613e94565b90935091506000805160206156058339815191528283098314156137b7576040805180820190915290815260208101919091529392505050565b600080516020615605833981519152600182089050613774565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906138036141da565b60005b60028110156139c857600061381c826006615569565b905084826002811061383057613830614f92565b6020020151518361384283600061530f565b600c811061385257613852614f92565b602002015284826002811061386957613869614f92565b60200201516020015183826001613880919061530f565b600c811061389057613890614f92565b60200201528382600281106138a7576138a7614f92565b60200201515151836138ba83600261530f565b600c81106138ca576138ca614f92565b60200201528382600281106138e1576138e1614f92565b60200201515160016020020151836138fa83600361530f565b600c811061390a5761390a614f92565b602002015283826002811061392157613921614f92565b60200201516020015160006002811061393c5761393c614f92565b60200201518361394d83600461530f565b600c811061395d5761395d614f92565b602002015283826002811061397457613974614f92565b60200201516020015160016002811061398f5761398f614f92565b6020020151836139a083600561530f565b600c81106139b0576139b0614f92565b602002015250806139c081615099565b915050613806565b506139d16141f9565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613a0184613f16565b9050808360ff166001901b11613a7f5760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610677565b90505b92915050565b6000805b8215613a8257613a9d6001846152df565b9092169180613aab816155ec565b915050613a8c565b60408051808201909152600080825260208201526102008261ffff1610613b0f5760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610677565b8161ffff1660011415613b23575081613a82565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613b8c57600161ffff871660ff83161c81161415613b6f57613b6c84846135ed565b93505b613b7983846135ed565b92506201fffe600192831b169101613b3f565b509195945050505050565b60408051808201909152600080825260208201528151158015613bbc57506020820151155b15613bda575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020016000805160206156058339815191528460200151613c0d9190614fa8565b613c25906000805160206156058339815191526152df565b905292915050565b919050565b6033546001600160a01b0316331461283e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610677565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060600080613cec84613a88565b61ffff166001600160401b03811115613d0757613d0761429e565b6040519080825280601f01601f191660200182016040528015613d31576020820181803683370190505b5090506000805b825182108015613d49575061010081105b15613da0576001811b935085841615613d90578060f81b838381518110613d7257613d72614f92565b60200101906001600160f81b031916908160001a9053508160010191505b613d9981615099565b9050613d38565b5090949350505050565b6065546001600160a01b0316158015613dcb57506001600160a01b03821615155b613e4d5760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610677565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613e908261345f565b5050565b60008080600080516020615605833981519152600360008051602061560583398151915286600080516020615605833981519152888909090890506000613f0a827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526000805160206156058339815191526140a3565b91959194509092505050565b600061010082511115613f9f5760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610677565b8151613fad57506000919050565b60008083600081518110613fc357613fc3614f92565b0160200151600160f89190911c81901b92505b845181101561409a57848181518110613ff157613ff1614f92565b0160200151600160f89190911c1b91508282116140865760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610677565b9181179161409381615099565b9050613fd6565b50909392505050565b6000806140ae6141f9565b6140b6614217565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa92508280156135a55750826141405760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610677565b505195945050505050565b604051806080016040528061415e614179565b81526000602082018190526060604083018190529091015290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b60405180604001604052806141c8614235565b81526020016141d5614235565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461068957600080fd5b60006020828403121561427a57600080fd5b8135613a7f81614253565b60006020828403121561429757600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156142d6576142d661429e565b60405290565b60405161010081016001600160401b03811182821017156142d6576142d661429e565b604051601f8201601f191681016001600160401b03811182821017156143275761432761429e565b604052919050565b60006040828403121561434157600080fd5b6143496142b4565b9050813581526020820135602082015292915050565b600082601f83011261437057600080fd5b604051604081018181106001600160401b03821117156143925761439261429e565b80604052508060408401858111156143a957600080fd5b845b81811015613b8c5780358352602092830192016143ab565b6000608082840312156143d557600080fd5b6143dd6142b4565b90506143e9838361435f565b81526143f8836040840161435f565b602082015292915050565b600080600080610120858703121561441a57600080fd5b8435935061442b866020870161432f565b925061443a86606087016143c3565b91506144498660e0870161432f565b905092959194509250565b63ffffffff8116811461068957600080fd5b8035613c2d81614454565b60006020828403121561448357600080fd5b8135613a7f81614454565b6000806000606084860312156144a357600080fd5b83356144ae81614253565b92506020848101356001600160401b03808211156144cb57600080fd5b818701915087601f8301126144df57600080fd5b8135818111156144f1576144f161429e565b614503601f8201601f191685016142ff565b9150808252888482850101111561451957600080fd5b808484018584013760008482840101525080945050505061453c60408501614466565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b868110156145db578385038a52825180518087529087019087870190845b818110156145c657835180516001600160a01b031684528a8101518b8501526040908101516001600160601b03169084015292890192606090920191600101614582565b50509a87019a95505091850191600101614564565b509298975050505050505050565b6020815260006145fc6020830184614545565b9392505050565b801515811461068957600080fd5b60006020828403121561462357600080fd5b8135613a7f81614603565b60008083601f84011261464057600080fd5b5081356001600160401b0381111561465757600080fd5b60208301915083602082850101111561466f57600080fd5b9250929050565b6000806000806000806080878903121561468f57600080fd5b863561469a81614253565b955060208701356146aa81614454565b945060408701356001600160401b03808211156146c657600080fd5b6146d28a838b0161462e565b909650945060608901359150808211156146eb57600080fd5b818901915089601f8301126146ff57600080fd5b81358181111561470e57600080fd5b8a60208260051b850101111561472357600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b8381101561476f57815163ffffffff168752958201959082019060010161474d565b509495945050505050565b60006020808352835160808285015261479660a0850182614739565b905081850151601f19808684030160408701526147b38383614739565b925060408701519150808684030160608701526147d08383614739565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b828110156148275784878303018452614815828751614739565b958801959388019391506001016147fb565b509998505050505050505050565b60ff8116811461068957600080fd5b60006020828403121561485657600080fd5b8135613a7f81614835565b60006001600160401b0382111561487a5761487a61429e565b5060051b60200190565b60008060006060848603121561489957600080fd5b83356148a481614253565b92506020848101356001600160401b038111156148c057600080fd5b8501601f810187136148d157600080fd5b80356148e46148df82614861565b6142ff565b81815260059190911b8201830190838101908983111561490357600080fd5b928401925b8284101561492157833582529284019290840190614908565b809650505050505061453c60408501614466565b6020808252825182820181905260009190848201906040850190845b8181101561496d57835183529284019291840191600101614951565b50909695505050505050565b600082601f83011261498a57600080fd5b8135602061499a6148df83614861565b82815260059290921b840181019181810190868411156149b957600080fd5b8286015b848110156149dd5780356149d081614454565b83529183019183016149bd565b509695505050505050565b600082601f8301126149f957600080fd5b81356020614a096148df83614861565b82815260069290921b84018101918181019086841115614a2857600080fd5b8286015b848110156149dd57614a3e888261432f565b835291830191604001614a2c565b600082601f830112614a5d57600080fd5b81356020614a6d6148df83614861565b82815260059290921b84018101918181019086841115614a8c57600080fd5b8286015b848110156149dd5780356001600160401b03811115614aaf5760008081fd5b614abd8986838b0101614979565b845250918301918301614a90565b60006101808284031215614ade57600080fd5b614ae66142dc565b905081356001600160401b0380821115614aff57600080fd5b614b0b85838601614979565b83526020840135915080821115614b2157600080fd5b614b2d858386016149e8565b60208401526040840135915080821115614b4657600080fd5b614b52858386016149e8565b6040840152614b6485606086016143c3565b6060840152614b768560e0860161432f565b6080840152610120840135915080821115614b9057600080fd5b614b9c85838601614979565b60a0840152610140840135915080821115614bb657600080fd5b614bc285838601614979565b60c0840152610160840135915080821115614bdc57600080fd5b50614be984828501614a4c565b60e08301525092915050565b600080600080600060808688031215614c0d57600080fd5b8535945060208601356001600160401b0380821115614c2b57600080fd5b614c3789838a0161462e565b909650945060408801359150614c4c82614454565b90925060608701359080821115614c6257600080fd5b50614c6f88828901614acb565b9150509295509295909350565b600081518084526020808501945080840160005b8381101561476f5781516001600160601b031687529582019590820190600101614c90565b6040815260008351604080840152614cd06080840182614c7c565b90506020850151603f19848303016060850152614ced8282614c7c565b925050508260208301529392505050565b600080600060608486031215614d1357600080fd5b8335614d1e81614253565b9250602084013591506040840135614d3581614454565b809150509250925092565b828152604060208201526000614d596040830184614545565b949350505050565b60008060008060a08587031215614d7757600080fd5b6060850186811115614d8857600080fd5b85945035614d9581614454565b925060808501356001600160401b03811115614db057600080fd5b614dbc8782880161462e565b95989497509550505050565b60008060008084860360a0811215614ddf57600080fd5b85356001600160401b0380821115614df657600080fd5b9087019060c0828a031215614e0a57600080fd5b8196506040601f1984011215614e1f57600080fd5b60208881019650605f1984011215614e3657600080fd5b6060880194506080880135925080831115614e5057600080fd5b5050614e5e87828801614acb565b91505092959194509250565b60008060008060808587031215614e8057600080fd5b8435614e8b81614253565b93506020850135614e9b81614253565b92506040850135614eab81614253565b91506060850135614ebb81614253565b939692955090935050565b600060208284031215614ed857600080fd5b8151613a7f81614253565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614f3f57600080fd5b8151613a7f81614603565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b600082614fc557634e487b7160e01b600052601260045260246000fd5b500690565b60006020808385031215614fdd57600080fd5b82516001600160401b03811115614ff357600080fd5b8301601f8101851361500457600080fd5b80516150126148df82614861565b81815260059190911b8201830190838101908783111561503157600080fd5b928401925b8284101561504f57835182529284019290840190615036565b979650505050505050565b60006020828403121561506c57600080fd5b81516001600160601b0381168114613a7f57600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156150ad576150ad615083565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b038311156150e157600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561511157600080fd5b82516001600160401b0381111561512757600080fd5b8301601f8101851361513857600080fd5b80516151466148df82614861565b81815260059190911b8201830190838101908783111561516557600080fd5b928401925b8284101561504f57835161517d81614454565b8252928401929084019061516a565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006151d560408301848661518c565b95945050505050565b6000602082840312156151f057600080fd5b81516001600160c01b0381168114613a7f57600080fd5b60006020828403121561521957600080fd5b8151613a7f81614454565b600060ff821660ff81141561523b5761523b615083565b60010192915050565b60408152600061525860408301858761518c565b905063ffffffff83166020830152949350505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156152b557845183529383019391830191600101615299565b5090979650505050505050565b6000602082840312156152d457600080fd5b8151613a7f81614835565b6000828210156152f1576152f1615083565b500390565b60006020828403121561530857600080fd5b5051919050565b6000821982111561532257615322615083565b500190565b60006020828403121561533957600080fd5b815167ffffffffffffffff1981168114613a7f57600080fd5b60006001600160601b038381169083168181101561537257615372615083565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156153b557815185529382019390820190600101615399565b5092979650505050505050565b6020808252825160009190828483015b60038210156153fa5782516001600160401b03168152918301916001919091019083016153d2565b50505063ffffffff81850151166080840152604084015160c060a085015280518060e086015260005b818110156154405782810184015186820161010001528301615423565b8181111561545357600061010083880101525b50606086015163ffffffff811660c08701529250601f01601f19169390930161010001949350505050565b600063ffffffff80831681851680830382111561549d5761549d615083565b01949350505050565b6000808335601e198436030181126154bd57600080fd5b8301803591506001600160401b038211156154d757600080fd5b60200191503681900382131561466f57600080fd5b80356001600160401b0381168114613c2d57600080fd5b60408101823561551281614454565b63ffffffff1682526001600160401b0361552e602085016154ec565b16602083015292915050565b60006001600160601b038083168185168183048111821515161561556057615560615083565b02949350505050565b600081600019048311821515161561558357615583615083565b500290565b60006020828403121561559a57600080fd5b6145fc826154ec565b60006001600160401b0380831681851680830382111561549d5761549d615083565b60006001600160401b03808316818114156155e2576155e2615083565b6001019392505050565b600061ffff808316818114156155e2576155e261508356fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a2646970667358221220bf77d494aa719834ceed875af80bf8afb86481ecd40cc3674e799ec4165952ea64736f6c634300080c0033",
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

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xee1c87c5.
//
// Solidity: function raiseAndResolveChallenge((uint64[3],uint32,bytes,uint32) task, (uint32,uint64) taskResponse, (uint64) taskResponseMetadata, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IIncredibleSummingTaskManagerTask, taskResponse IIncredibleSummingTaskManagerTaskResponse, taskResponseMetadata IIncredibleSummingTaskManagerTaskResponseMetadata, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, nonSignerStakesAndSignature)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xee1c87c5.
//
// Solidity: function raiseAndResolveChallenge((uint64[3],uint32,bytes,uint32) task, (uint32,uint64) taskResponse, (uint64) taskResponseMetadata, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerSession) RaiseAndResolveChallenge(task IIncredibleSummingTaskManagerTask, taskResponse IIncredibleSummingTaskManagerTaskResponse, taskResponseMetadata IIncredibleSummingTaskManagerTaskResponseMetadata, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleSummingTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, nonSignerStakesAndSignature)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xee1c87c5.
//
// Solidity: function raiseAndResolveChallenge((uint64[3],uint32,bytes,uint32) task, (uint32,uint64) taskResponse, (uint64) taskResponseMetadata, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerTransactorSession) RaiseAndResolveChallenge(task IIncredibleSummingTaskManagerTask, taskResponse IIncredibleSummingTaskManagerTaskResponse, taskResponseMetadata IIncredibleSummingTaskManagerTaskResponseMetadata, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractIncredibleSummingTaskManager.Contract.RaiseAndResolveChallenge(&_ContractIncredibleSummingTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, nonSignerStakesAndSignature)
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

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xdec2b2c86c7f21aefb374e79a431e6bf194d79f59580cc5bd403bb694f40c63c.
//
// Solidity: event TaskResponded((uint32,uint64) taskResponse, (uint64) taskResponseMetadata)
func (_ContractIncredibleSummingTaskManager *ContractIncredibleSummingTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractIncredibleSummingTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractIncredibleSummingTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingTaskManagerTaskRespondedIterator{contract: _ContractIncredibleSummingTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xdec2b2c86c7f21aefb374e79a431e6bf194d79f59580cc5bd403bb694f40c63c.
//
// Solidity: event TaskResponded((uint32,uint64) taskResponse, (uint64) taskResponseMetadata)
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

// ParseTaskResponded is a log parse operation binding the contract event 0xdec2b2c86c7f21aefb374e79a431e6bf194d79f59580cc5bd403bb694f40c63c.
//
// Solidity: event TaskResponded((uint32,uint64) taskResponse, (uint64) taskResponseMetadata)
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
