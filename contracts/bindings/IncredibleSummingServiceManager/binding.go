// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIncredibleSummingServiceManager

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

// IPaymentCoordinatorRangePayment is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorRangePayment struct {
	StrategiesAndMultipliers []IPaymentCoordinatorStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IPaymentCoordinatorStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IPaymentCoordinatorStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractIncredibleSummingServiceManagerMetaData contains all meta data concerning the ContractIncredibleSummingServiceManager contract.
var ContractIncredibleSummingServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avsDirectory\",\"type\":\"address\",\"internalType\":\"contractIAVSDirectory\"},{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractIRegistryCoordinator\"},{\"name\":\"_stakeRegistry\",\"type\":\"address\",\"internalType\":\"contractIStakeRegistry\"},{\"name\":\"_incredibleSummingTaskManager\",\"type\":\"address\",\"internalType\":\"contractIIncredibleSummingTaskManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avsDirectory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"freezeOperator\",\"inputs\":[{\"name\":\"operatorAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorRestakedStrategies\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRestakeableStrategies\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"incredibleSummingTaskManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIIncredibleSummingTaskManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payForRange\",\"inputs\":[{\"name\":\"rangePayments\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.RangePayment[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIPaymentCoordinator.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"_metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162001b6338038062001b6383398101604081905262000035916200015a565b6001600160a01b03808516608052600060a081905281851660c05290831660e05284908484620000646200007f565b505050506001600160a01b03166101005250620001c2915050565b600054610100900460ff1615620000ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811610156200013f576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200015757600080fd5b50565b600080600080608085870312156200017157600080fd5b84516200017e8162000141565b6020860151909450620001918162000141565b6040860151909350620001a48162000141565b6060860151909250620001b78162000141565b939692955090935050565b60805160a05160c05160e051610100516118ec620002776000396000818161019b015261093b01526000818161065a015281816107b60152818161084d01528181610c7a01528181610dfe0152610e9d015260008181610485015281816105140152818161059401528181610a0c01528181610aa201528181610bb80152610d5901526000818161031501526103f301526000818161010c01528181610a6001528181610afe0152610b7d01526118ec6000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80639926ee7d116100715780639926ee7d1461015d578063a364f4da14610170578063a98fb35514610183578063b664a9b514610196578063e481af9d146101bd578063f2fde38b146101c557600080fd5b80631b445516146100b957806333cfb7b7146100ce57806338c8ee64146100f75780636b3aa72e1461010a578063715018a6146101445780638da5cb5b1461014c575b600080fd5b6100cc6100c7366004611192565b6101d8565b005b6100e16100dc36600461121c565b610460565b6040516100ee9190611240565b60405180910390f35b6100cc61010536600461121c565b610930565b7f00000000000000000000000000000000000000000000000000000000000000005b6040516001600160a01b0390911681526020016100ee565b6100cc6109ed565b6033546001600160a01b031661012c565b6100cc61016b366004611342565b610a01565b6100cc61017e36600461121c565b610a97565b6100cc6101913660046113ed565b610b5e565b61012c7f000000000000000000000000000000000000000000000000000000000000000081565b6100e1610bb2565b6100cc6101d336600461121c565b610f7c565b6101e0610ff2565b60005b818110156103db578282828181106101fd576101fd61143e565b905060200281019061020f9190611454565b61022090604081019060200161121c565b6001600160a01b03166323b872dd33308686868181106102425761024261143e565b90506020028101906102549190611454565b604080516001600160e01b031960e087901b1681526001600160a01b039485166004820152939092166024840152013560448201526064016020604051808303816000875af11580156102ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102cf9190611484565b508282828181106102e2576102e261143e565b90506020028101906102f49190611454565b61030590604081019060200161121c565b6001600160a01b031663095ea7b37f00000000000000000000000000000000000000000000000000000000000000008585858181106103465761034661143e565b90506020028101906103589190611454565b604080516001600160e01b031960e086901b1681526001600160a01b039093166004840152013560248201526044016020604051808303816000875af11580156103a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103ca9190611484565b506103d4816114bc565b90506101e3565b50604051630da22a8b60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690631b4455169061042a9085908590600401611570565b600060405180830381600087803b15801561044457600080fd5b505af1158015610458573d6000803e3d6000fd5b505050505050565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa1580156104cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104f0919061167e565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa15801561055b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057f9190611697565b90506001600160c01b038116158061061957507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105f0573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061491906116c0565b60ff16155b1561063557505060408051600081526020810190915292915050565b6000610649826001600160c01b031661104c565b90506000805b825181101561071f577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106106995761069961143e565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa1580156106dd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610701919061167e565b61070b90836116e3565b915080610717816114bc565b91505061064f565b5060008167ffffffffffffffff81111561073b5761073b61128d565b604051908082528060200260200182016040528015610764578160200160208202803683370190505b5090506000805b84518110156109235760008582815181106107885761078861143e565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa1580156107fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610821919061167e565b905060005b8181101561090d576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa15801561089b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bf91906116fb565b600001518686815181106108d5576108d561143e565b6001600160a01b0390921660209283029190910190910152846108f7816114bc565b9550508080610905906114bc565b915050610826565b505050808061091b906114bc565b91505061076b565b5090979650505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109ea5760405162461bcd60e51b815260206004820152604e60248201527f6f6e6c79496e6372656469626c6553756d6d696e675461736b4d616e6167657260448201527f3a206e6f742066726f6d206372656469626c652061727261792073756d6d696e60648201526d33903a30b9b59036b0b730b3b2b960911b608482015260a4015b60405180910390fd5b50565b6109f5610ff2565b6109ff600061110f565b565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a495760405162461bcd60e51b81526004016109e19061175a565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061042a908590859060040161181f565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610adf5760405162461bcd60e51b81526004016109e19061175a565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da906024015b600060405180830381600087803b158015610b4357600080fd5b505af1158015610b57573d6000803e3d6000fd5b5050505050565b610b66610ff2565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb35590610b2990849060040161186a565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3891906116c0565b60ff16905080610c5657505060408051600081526020810190915290565b6000805b82811015610d0b57604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610cc9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ced919061167e565b610cf790836116e3565b915080610d03816114bc565b915050610c5a565b5060008167ffffffffffffffff811115610d2757610d2761128d565b604051908082528060200260200182016040528015610d50578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610db5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dd991906116c0565b60ff16811015610f7257604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610e4d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e71919061167e565b905060005b81811015610f5d576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610eeb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0f91906116fb565b60000151858581518110610f2557610f2561143e565b6001600160a01b039092166020928302919091019091015283610f47816114bc565b9450508080610f55906114bc565b915050610e76565b50508080610f6a906114bc565b915050610d57565b5090949350505050565b610f84610ff2565b6001600160a01b038116610fe95760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109e1565b6109ea8161110f565b6033546001600160a01b031633146109ff5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109e1565b606060008061105a84611161565b61ffff1667ffffffffffffffff8111156110765761107661128d565b6040519080825280601f01601f1916602001820160405280156110a0576020820181803683370190505b5090506000805b8251821080156110b8575061010081105b15610f72576001811b9350858416156110ff578060f81b8383815181106110e1576110e161143e565b60200101906001600160f81b031916908160001a9053508160010191505b611108816114bc565b90506110a7565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000805b821561118c5761117660018461187d565b909216918061118481611894565b915050611165565b92915050565b600080602083850312156111a557600080fd5b823567ffffffffffffffff808211156111bd57600080fd5b818501915085601f8301126111d157600080fd5b8135818111156111e057600080fd5b8660208260051b85010111156111f557600080fd5b60209290920196919550909350505050565b6001600160a01b03811681146109ea57600080fd5b60006020828403121561122e57600080fd5b813561123981611207565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156112815783516001600160a01b03168352928401929184019160010161125c565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff811182821017156112c6576112c661128d565b60405290565b600067ffffffffffffffff808411156112e7576112e761128d565b604051601f8501601f19908116603f0116810190828211818310171561130f5761130f61128d565b8160405280935085815286868601111561132857600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561135557600080fd5b823561136081611207565b9150602083013567ffffffffffffffff8082111561137d57600080fd5b908401906060828703121561139157600080fd5b6113996112a3565b8235828111156113a857600080fd5b83019150601f820187136113bb57600080fd5b6113ca878335602085016112cc565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156113ff57600080fd5b813567ffffffffffffffff81111561141657600080fd5b8201601f8101841361142757600080fd5b611436848235602084016112cc565b949350505050565b634e487b7160e01b600052603260045260246000fd5b60008235609e1983360301811261146a57600080fd5b9190910192915050565b803561147f81611207565b919050565b60006020828403121561149657600080fd5b8151801515811461123957600080fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156114d0576114d06114a6565b5060010190565b6bffffffffffffffffffffffff811681146109ea57600080fd5b8183526000602080850194508260005b8581101561155157813561151481611207565b6001600160a01b031687528183013561152c816114d7565b6bffffffffffffffffffffffff16878401526040968701969190910190600101611501565b509495945050505050565b803563ffffffff8116811461147f57600080fd5b60208082528181018390526000906040808401600586901b8501820187855b8881101561167057878303603f190184528135368b9003609e190181126115b557600080fd5b8a0160a0813536839003601e190181126115ce57600080fd5b8201803567ffffffffffffffff8111156115e757600080fd5b8060061b36038413156115f957600080fd5b82875261160b838801828c85016114f1565b9250505061161a888301611474565b6001600160a01b0316888601528187013587860152606061163c81840161155c565b63ffffffff1690860152608061165383820161155c565b63ffffffff1695019490945250928501929085019060010161158f565b509098975050505050505050565b60006020828403121561169057600080fd5b5051919050565b6000602082840312156116a957600080fd5b81516001600160c01b038116811461123957600080fd5b6000602082840312156116d257600080fd5b815160ff8116811461123957600080fd5b600082198211156116f6576116f66114a6565b500190565b60006040828403121561170d57600080fd5b6040516040810181811067ffffffffffffffff821117156117305761173061128d565b604052825161173e81611207565b8152602083015161174e816114d7565b60208201529392505050565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b6000815180845260005b818110156117f8576020818501810151868301820152016117dc565b8181111561180a576000602083870101525b50601f01601f19169290920160200192915050565b60018060a01b038316815260406020820152600082516060604084015261184960a08401826117d2565b90506020840151606084015260408401516080840152809150509392505050565b60208152600061123960208301846117d2565b60008282101561188f5761188f6114a6565b500390565b600061ffff808316818114156118ac576118ac6114a6565b600101939250505056fea2646970667358221220cf012b8a6379d6df939a4f41f45bdc23b54459a0424eba0bcf6e362caa03143764736f6c634300080c0033",
}

// ContractIncredibleSummingServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIncredibleSummingServiceManagerMetaData.ABI instead.
var ContractIncredibleSummingServiceManagerABI = ContractIncredibleSummingServiceManagerMetaData.ABI

// ContractIncredibleSummingServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIncredibleSummingServiceManagerMetaData.Bin instead.
var ContractIncredibleSummingServiceManagerBin = ContractIncredibleSummingServiceManagerMetaData.Bin

// DeployContractIncredibleSummingServiceManager deploys a new Ethereum contract, binding an instance of ContractIncredibleSummingServiceManager to it.
func DeployContractIncredibleSummingServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _avsDirectory common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _incredibleSummingTaskManager common.Address) (common.Address, *types.Transaction, *ContractIncredibleSummingServiceManager, error) {
	parsed, err := ContractIncredibleSummingServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIncredibleSummingServiceManagerBin), backend, _avsDirectory, _registryCoordinator, _stakeRegistry, _incredibleSummingTaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIncredibleSummingServiceManager{ContractIncredibleSummingServiceManagerCaller: ContractIncredibleSummingServiceManagerCaller{contract: contract}, ContractIncredibleSummingServiceManagerTransactor: ContractIncredibleSummingServiceManagerTransactor{contract: contract}, ContractIncredibleSummingServiceManagerFilterer: ContractIncredibleSummingServiceManagerFilterer{contract: contract}}, nil
}

// ContractIncredibleSummingServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractIncredibleSummingServiceManager struct {
	ContractIncredibleSummingServiceManagerCaller     // Read-only binding to the contract
	ContractIncredibleSummingServiceManagerTransactor // Write-only binding to the contract
	ContractIncredibleSummingServiceManagerFilterer   // Log filterer for contract events
}

// ContractIncredibleSummingServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIncredibleSummingServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSummingServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIncredibleSummingServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSummingServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIncredibleSummingServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIncredibleSummingServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIncredibleSummingServiceManagerSession struct {
	Contract     *ContractIncredibleSummingServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                            // Call options to use throughout this session
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ContractIncredibleSummingServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIncredibleSummingServiceManagerCallerSession struct {
	Contract *ContractIncredibleSummingServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                  // Call options to use throughout this session
}

// ContractIncredibleSummingServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIncredibleSummingServiceManagerTransactorSession struct {
	Contract     *ContractIncredibleSummingServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                  // Transaction auth options to use throughout this session
}

// ContractIncredibleSummingServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIncredibleSummingServiceManagerRaw struct {
	Contract *ContractIncredibleSummingServiceManager // Generic contract binding to access the raw methods on
}

// ContractIncredibleSummingServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIncredibleSummingServiceManagerCallerRaw struct {
	Contract *ContractIncredibleSummingServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIncredibleSummingServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIncredibleSummingServiceManagerTransactorRaw struct {
	Contract *ContractIncredibleSummingServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIncredibleSummingServiceManager creates a new instance of ContractIncredibleSummingServiceManager, bound to a specific deployed contract.
func NewContractIncredibleSummingServiceManager(address common.Address, backend bind.ContractBackend) (*ContractIncredibleSummingServiceManager, error) {
	contract, err := bindContractIncredibleSummingServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingServiceManager{ContractIncredibleSummingServiceManagerCaller: ContractIncredibleSummingServiceManagerCaller{contract: contract}, ContractIncredibleSummingServiceManagerTransactor: ContractIncredibleSummingServiceManagerTransactor{contract: contract}, ContractIncredibleSummingServiceManagerFilterer: ContractIncredibleSummingServiceManagerFilterer{contract: contract}}, nil
}

// NewContractIncredibleSummingServiceManagerCaller creates a new read-only instance of ContractIncredibleSummingServiceManager, bound to a specific deployed contract.
func NewContractIncredibleSummingServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIncredibleSummingServiceManagerCaller, error) {
	contract, err := bindContractIncredibleSummingServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingServiceManagerCaller{contract: contract}, nil
}

// NewContractIncredibleSummingServiceManagerTransactor creates a new write-only instance of ContractIncredibleSummingServiceManager, bound to a specific deployed contract.
func NewContractIncredibleSummingServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIncredibleSummingServiceManagerTransactor, error) {
	contract, err := bindContractIncredibleSummingServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingServiceManagerTransactor{contract: contract}, nil
}

// NewContractIncredibleSummingServiceManagerFilterer creates a new log filterer instance of ContractIncredibleSummingServiceManager, bound to a specific deployed contract.
func NewContractIncredibleSummingServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIncredibleSummingServiceManagerFilterer, error) {
	contract, err := bindContractIncredibleSummingServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingServiceManagerFilterer{contract: contract}, nil
}

// bindContractIncredibleSummingServiceManager binds a generic wrapper to an already deployed contract.
func bindContractIncredibleSummingServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIncredibleSummingServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSummingServiceManager.Contract.ContractIncredibleSummingServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.ContractIncredibleSummingServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.ContractIncredibleSummingServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIncredibleSummingServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.contract.Transact(opts, method, params...)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerCaller) AvsDirectory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingServiceManager.contract.Call(opts, &out, "avsDirectory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerSession) AvsDirectory() (common.Address, error) {
	return _ContractIncredibleSummingServiceManager.Contract.AvsDirectory(&_ContractIncredibleSummingServiceManager.CallOpts)
}

// AvsDirectory is a free data retrieval call binding the contract method 0x6b3aa72e.
//
// Solidity: function avsDirectory() view returns(address)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerCallerSession) AvsDirectory() (common.Address, error) {
	return _ContractIncredibleSummingServiceManager.Contract.AvsDirectory(&_ContractIncredibleSummingServiceManager.CallOpts)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractIncredibleSummingServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractIncredibleSummingServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractIncredibleSummingServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractIncredibleSummingServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractIncredibleSummingServiceManager.Contract.GetRestakeableStrategies(&_ContractIncredibleSummingServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractIncredibleSummingServiceManager.Contract.GetRestakeableStrategies(&_ContractIncredibleSummingServiceManager.CallOpts)
}

// IncredibleSummingTaskManager is a free data retrieval call binding the contract method 0xb664a9b5.
//
// Solidity: function incredibleSummingTaskManager() view returns(address)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerCaller) IncredibleSummingTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingServiceManager.contract.Call(opts, &out, "incredibleSummingTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IncredibleSummingTaskManager is a free data retrieval call binding the contract method 0xb664a9b5.
//
// Solidity: function incredibleSummingTaskManager() view returns(address)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerSession) IncredibleSummingTaskManager() (common.Address, error) {
	return _ContractIncredibleSummingServiceManager.Contract.IncredibleSummingTaskManager(&_ContractIncredibleSummingServiceManager.CallOpts)
}

// IncredibleSummingTaskManager is a free data retrieval call binding the contract method 0xb664a9b5.
//
// Solidity: function incredibleSummingTaskManager() view returns(address)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerCallerSession) IncredibleSummingTaskManager() (common.Address, error) {
	return _ContractIncredibleSummingServiceManager.Contract.IncredibleSummingTaskManager(&_ContractIncredibleSummingServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIncredibleSummingServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSummingServiceManager.Contract.Owner(&_ContractIncredibleSummingServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractIncredibleSummingServiceManager.Contract.Owner(&_ContractIncredibleSummingServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractIncredibleSummingServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractIncredibleSummingServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.FreezeOperator(&_ContractIncredibleSummingServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.FreezeOperator(&_ContractIncredibleSummingServiceManager.TransactOpts, operatorAddr)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactor) PayForRange(opts *bind.TransactOpts, rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.contract.Transact(opts, "payForRange", rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.PayForRange(&_ContractIncredibleSummingServiceManager.TransactOpts, rangePayments)
}

// PayForRange is a paid mutator transaction binding the contract method 0x1b445516.
//
// Solidity: function payForRange(((address,uint96)[],address,uint256,uint32,uint32)[] rangePayments) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactorSession) PayForRange(rangePayments []IPaymentCoordinatorRangePayment) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.PayForRange(&_ContractIncredibleSummingServiceManager.TransactOpts, rangePayments)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.RegisterOperatorToAVS(&_ContractIncredibleSummingServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.RegisterOperatorToAVS(&_ContractIncredibleSummingServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.RenounceOwnership(&_ContractIncredibleSummingServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.RenounceOwnership(&_ContractIncredibleSummingServiceManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.TransferOwnership(&_ContractIncredibleSummingServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.TransferOwnership(&_ContractIncredibleSummingServiceManager.TransactOpts, newOwner)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.contract.Transact(opts, "updateAVSMetadataURI", _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.UpdateAVSMetadataURI(&_ContractIncredibleSummingServiceManager.TransactOpts, _metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string _metadataURI) returns()
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerTransactorSession) UpdateAVSMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractIncredibleSummingServiceManager.Contract.UpdateAVSMetadataURI(&_ContractIncredibleSummingServiceManager.TransactOpts, _metadataURI)
}

// ContractIncredibleSummingServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractIncredibleSummingServiceManager contract.
type ContractIncredibleSummingServiceManagerInitializedIterator struct {
	Event *ContractIncredibleSummingServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSummingServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingServiceManagerInitialized)
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
		it.Event = new(ContractIncredibleSummingServiceManagerInitialized)
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
func (it *ContractIncredibleSummingServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingServiceManagerInitialized represents a Initialized event raised by the ContractIncredibleSummingServiceManager contract.
type ContractIncredibleSummingServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractIncredibleSummingServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractIncredibleSummingServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingServiceManagerInitializedIterator{contract: _ContractIncredibleSummingServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractIncredibleSummingServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingServiceManagerInitialized)
				if err := _ContractIncredibleSummingServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractIncredibleSummingServiceManagerInitialized, error) {
	event := new(ContractIncredibleSummingServiceManagerInitialized)
	if err := _ContractIncredibleSummingServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIncredibleSummingServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractIncredibleSummingServiceManager contract.
type ContractIncredibleSummingServiceManagerOwnershipTransferredIterator struct {
	Event *ContractIncredibleSummingServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractIncredibleSummingServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIncredibleSummingServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractIncredibleSummingServiceManagerOwnershipTransferred)
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
func (it *ContractIncredibleSummingServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIncredibleSummingServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIncredibleSummingServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractIncredibleSummingServiceManager contract.
type ContractIncredibleSummingServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractIncredibleSummingServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSummingServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIncredibleSummingServiceManagerOwnershipTransferredIterator{contract: _ContractIncredibleSummingServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractIncredibleSummingServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractIncredibleSummingServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIncredibleSummingServiceManagerOwnershipTransferred)
				if err := _ContractIncredibleSummingServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractIncredibleSummingServiceManager *ContractIncredibleSummingServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractIncredibleSummingServiceManagerOwnershipTransferred, error) {
	event := new(ContractIncredibleSummingServiceManagerOwnershipTransferred)
	if err := _ContractIncredibleSummingServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
