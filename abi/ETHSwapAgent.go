// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ETHSwapAgentABI is the input ABI used to generate the binding from.
const ETHSwapAgentABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bscTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SwapFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"SwapPairRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"SwapStarted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bscTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fillBSC2ETHSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filledBSCTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"ownerAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"}],\"name\":\"registerSwapPairToBSC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swapETH2BSC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ETHSwapAgentBin is the compiled bytecode used for deploying new contracts.
var ETHSwapAgentBin = "0x608060405234801561001057600080fd5b50611d63806100206000396000f3fe60806040526004361061009c5760003560e01c806389b156041161006457806389b15604146101da5780638da5cb5b146102435780639867df111461029a578063b9927a9c14610337578063da35a26f1461039d578063f2fde38b146103f85761009c565b806334e19907146100a157806350877c77146100dc57806354cf2aeb1461012f5780635c13c1511461015a578063715018a6146101c3575b600080fd5b3480156100ad57600080fd5b506100da600480360360208110156100c457600080fd5b8101908080359060200190929190505050610449565b005b3480156100e857600080fd5b50610115600480360360208110156100ff57600080fd5b810190808035906020019092919050505061051d565b604051808215151515815260200191505060405180910390f35b34801561013b57600080fd5b5061014461053d565b6040518082815260200191505060405180910390f35b34801561016657600080fd5b506101a96004803603602081101561017d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610543565b604051808215151515815260200191505060405180910390f35b3480156101cf57600080fd5b506101d8610ba6565b005b3480156101e657600080fd5b50610229600480360360208110156101fd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d31565b604051808215151515815260200191505060405180910390f35b34801561024f57600080fd5b50610258610d51565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102a657600080fd5b5061031d600480360360808110156102bd57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d77565b604051808215151515815260200191505060405180910390f35b6103836004803603604081101561034d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061105f565b604051808215151515815260200191505060405180910390f35b3480156103a957600080fd5b506103f6600480360360408110156103c057600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113ca565b005b34801561040457600080fd5b506104476004803603602081101561041b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611513565b005b610451611723565b73ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610513576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b8060048190555050565b60026020528060005260406000206000915054906101000a900460ff1681565b60045481565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610605576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f616c72656164792072656769737465726564000000000000000000000000000081525060200191505060405180910390fd5b60608273ffffffffffffffffffffffffffffffffffffffff166306fdde036040518163ffffffff1660e01b815260040160006040518083038186803b15801561064d57600080fd5b505afa158015610661573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250602081101561068b57600080fd5b81019080805160405193929190846401000000008211156106ab57600080fd5b838201915060208201858111156106c157600080fd5b82518660018202830111640100000000821117156106de57600080fd5b8083526020830192505050908051906020019080838360005b838110156107125780820151818401526020810190506106f7565b50505050905090810190601f16801561073f5780820380516001836020036101000a031916815260200191505b50604052505050905060608373ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b15801561079057600080fd5b505afa1580156107a4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156107ce57600080fd5b81019080805160405193929190846401000000008211156107ee57600080fd5b8382019150602082018581111561080457600080fd5b825186600182028301116401000000008211171561082157600080fd5b8083526020830192505050908051906020019080838360005b8381101561085557808201518184015260208101905061083a565b50505050905090810190601f1680156108825780820380516001836020036101000a031916815260200191505b50604052505050905060008473ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156108d357600080fd5b505afa1580156108e7573d6000803e3d6000fd5b505050506040513d60208110156108fd57600080fd5b810190808051906020019092919050505090506000835111610987576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f656d707479206e616d650000000000000000000000000000000000000000000081525060200191505060405180910390fd5b60008251116109fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f656d7074792073796d626f6c000000000000000000000000000000000000000081525060200191505060405180910390fd5b60018060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167ffe3bd005e346323fa452df8cafc28c55b99e3766ba8750571d139c6cf5bc08a08585856040518080602001806020018460ff1660ff168152602001838103835286818151815260200191508051906020019080838360005b83811015610af7578082015181840152602081019050610adc565b50505050905090810190601f168015610b245780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b83811015610b5d578082015181840152602081019050610b42565b50505050905090810190601f168015610b8a5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a360019350505050919050565b610bae611723565b73ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610c70576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60016020528060005260406000206000915054906101000a900460ff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000610d81611723565b73ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610e43576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6002600086815260200190815260200160002060009054906101000a900460ff1615610ed7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f6273632074782066696c6c656420616c7265616479000000000000000000000081525060200191505060405180910390fd5b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610f96576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6e6f74207265676973746572656420746f6b656e00000000000000000000000081525060200191505060405180910390fd5b60016002600087815260200190815260200160002060006101000a81548160ff021916908315150217905550610fed83838673ffffffffffffffffffffffffffffffffffffffff1661172b9092919063ffffffff16565b8273ffffffffffffffffffffffffffffffffffffffff16858573ffffffffffffffffffffffffffffffffffffffff167f3bebd9a738291e69898b5dbfadb6329b4b09fc648bdef68762928e521463abd9856040518082815260200191505060405180910390a460019050949350505050565b600061106a336117e3565b156110dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f636f6e7472616374206973206e6f7420616c6c6f77656420746f20737761700081525060200191505060405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461117e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6e6f2070726f787920636f6e747261637420697320616c6c6f7765640000000081525060200191505060405180910390fd5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661123d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6e6f74207265676973746572656420746f6b656e00000000000000000000000081525060200191505060405180910390fd5b60045434146112b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f7377617020666565206e6f7420657175616c000000000000000000000000000081525060200191505060405180910390fd5b6112e13330848673ffffffffffffffffffffffffffffffffffffffff166117f6909392919063ffffffff16565b6000341461135357600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015611351573d6000803e3d6000fd5b505b3373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167ff60309f865a6aa297da5fac6188136a02e5acfdf6e8f6d35257a9f4e9653170f8434604051808381526020018281526020019250505060405180910390a36001905092915050565b600060019054906101000a900460ff16806113e957506113e86118e3565b5b8061140057506000809054906101000a900460ff16155b611455576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611cd6602e913960400191505060405180910390fd5b60008060019054906101000a900460ff1615905080156114a5576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b8260048190555081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550801561150e5760008060016101000a81548160ff0219169083151502179055505b505050565b61151b611723565b73ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146115dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611663576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180611c8a6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b6117de8363a9059cbb60e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506118fa565b505050565b600080823b905060008111915050919050565b6118dd846323b872dd60e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506118fa565b50505050565b6000803090506000813b9050600081149250505090565b606061195c826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166119e99092919063ffffffff16565b90506000815111156119e45780806020019051602081101561197d57600080fd5b81019080805190602001909291905050506119e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180611d04602a913960400191505060405180910390fd5b5b505050565b60606119f88484600085611a01565b90509392505050565b606082471015611a5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180611cb06026913960400191505060405180910390fd5b611a6585611baa565b611ad7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b600060608673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b60208310611b275780518252602082019150602081019050602083039250611b04565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611b89576040519150601f19603f3d011682016040523d82523d6000602084013e611b8e565b606091505b5091509150611b9e828286611bbd565b92505050949350505050565b600080823b905060008111915050919050565b60608315611bcd57829050611c82565b600083511115611be05782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611c47578082015181840152602081019050611c2c565b50505050905090810190601f168015611c745780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b939250505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a65645361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a26469706673582212208355cc2dfdcbac18e2dcc47a9219045c35954bcf6039cf9f858516f6067dff8f64736f6c63430006040033"

// DeployETHSwapAgent deploys a new Ethereum contract, binding an instance of ETHSwapAgent to it.
func DeployETHSwapAgent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ETHSwapAgent, error) {
	parsed, err := abi.JSON(strings.NewReader(ETHSwapAgentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ETHSwapAgentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ETHSwapAgent{ETHSwapAgentCaller: ETHSwapAgentCaller{contract: contract}, ETHSwapAgentTransactor: ETHSwapAgentTransactor{contract: contract}, ETHSwapAgentFilterer: ETHSwapAgentFilterer{contract: contract}}, nil
}

// ETHSwapAgent is an auto generated Go binding around an Ethereum contract.
type ETHSwapAgent struct {
	ETHSwapAgentCaller     // Read-only binding to the contract
	ETHSwapAgentTransactor // Write-only binding to the contract
	ETHSwapAgentFilterer   // Log filterer for contract events
}

// ETHSwapAgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ETHSwapAgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHSwapAgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ETHSwapAgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHSwapAgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ETHSwapAgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHSwapAgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ETHSwapAgentSession struct {
	Contract     *ETHSwapAgent     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ETHSwapAgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ETHSwapAgentCallerSession struct {
	Contract *ETHSwapAgentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ETHSwapAgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ETHSwapAgentTransactorSession struct {
	Contract     *ETHSwapAgentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ETHSwapAgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ETHSwapAgentRaw struct {
	Contract *ETHSwapAgent // Generic contract binding to access the raw methods on
}

// ETHSwapAgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ETHSwapAgentCallerRaw struct {
	Contract *ETHSwapAgentCaller // Generic read-only contract binding to access the raw methods on
}

// ETHSwapAgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ETHSwapAgentTransactorRaw struct {
	Contract *ETHSwapAgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewETHSwapAgent creates a new instance of ETHSwapAgent, bound to a specific deployed contract.
func NewETHSwapAgent(address common.Address, backend bind.ContractBackend) (*ETHSwapAgent, error) {
	contract, err := bindETHSwapAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ETHSwapAgent{ETHSwapAgentCaller: ETHSwapAgentCaller{contract: contract}, ETHSwapAgentTransactor: ETHSwapAgentTransactor{contract: contract}, ETHSwapAgentFilterer: ETHSwapAgentFilterer{contract: contract}}, nil
}

// NewETHSwapAgentCaller creates a new read-only instance of ETHSwapAgent, bound to a specific deployed contract.
func NewETHSwapAgentCaller(address common.Address, caller bind.ContractCaller) (*ETHSwapAgentCaller, error) {
	contract, err := bindETHSwapAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ETHSwapAgentCaller{contract: contract}, nil
}

// NewETHSwapAgentTransactor creates a new write-only instance of ETHSwapAgent, bound to a specific deployed contract.
func NewETHSwapAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*ETHSwapAgentTransactor, error) {
	contract, err := bindETHSwapAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ETHSwapAgentTransactor{contract: contract}, nil
}

// NewETHSwapAgentFilterer creates a new log filterer instance of ETHSwapAgent, bound to a specific deployed contract.
func NewETHSwapAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*ETHSwapAgentFilterer, error) {
	contract, err := bindETHSwapAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ETHSwapAgentFilterer{contract: contract}, nil
}

// bindETHSwapAgent binds a generic wrapper to an already deployed contract.
func bindETHSwapAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ETHSwapAgentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHSwapAgent *ETHSwapAgentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHSwapAgent.Contract.ETHSwapAgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHSwapAgent *ETHSwapAgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.ETHSwapAgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHSwapAgent *ETHSwapAgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.ETHSwapAgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHSwapAgent *ETHSwapAgentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHSwapAgent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHSwapAgent *ETHSwapAgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHSwapAgent *ETHSwapAgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.contract.Transact(opts, method, params...)
}

// FilledBSCTx is a free data retrieval call binding the contract method 0x50877c77.
//
// Solidity: function filledBSCTx(bytes32 ) view returns(bool)
func (_ETHSwapAgent *ETHSwapAgentCaller) FilledBSCTx(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ETHSwapAgent.contract.Call(opts, &out, "filledBSCTx", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FilledBSCTx is a free data retrieval call binding the contract method 0x50877c77.
//
// Solidity: function filledBSCTx(bytes32 ) view returns(bool)
func (_ETHSwapAgent *ETHSwapAgentSession) FilledBSCTx(arg0 [32]byte) (bool, error) {
	return _ETHSwapAgent.Contract.FilledBSCTx(&_ETHSwapAgent.CallOpts, arg0)
}

// FilledBSCTx is a free data retrieval call binding the contract method 0x50877c77.
//
// Solidity: function filledBSCTx(bytes32 ) view returns(bool)
func (_ETHSwapAgent *ETHSwapAgentCallerSession) FilledBSCTx(arg0 [32]byte) (bool, error) {
	return _ETHSwapAgent.Contract.FilledBSCTx(&_ETHSwapAgent.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHSwapAgent *ETHSwapAgentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHSwapAgent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHSwapAgent *ETHSwapAgentSession) Owner() (common.Address, error) {
	return _ETHSwapAgent.Contract.Owner(&_ETHSwapAgent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHSwapAgent *ETHSwapAgentCallerSession) Owner() (common.Address, error) {
	return _ETHSwapAgent.Contract.Owner(&_ETHSwapAgent.CallOpts)
}

// RegisteredERC20 is a free data retrieval call binding the contract method 0x89b15604.
//
// Solidity: function registeredERC20(address ) view returns(bool)
func (_ETHSwapAgent *ETHSwapAgentCaller) RegisteredERC20(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ETHSwapAgent.contract.Call(opts, &out, "registeredERC20", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredERC20 is a free data retrieval call binding the contract method 0x89b15604.
//
// Solidity: function registeredERC20(address ) view returns(bool)
func (_ETHSwapAgent *ETHSwapAgentSession) RegisteredERC20(arg0 common.Address) (bool, error) {
	return _ETHSwapAgent.Contract.RegisteredERC20(&_ETHSwapAgent.CallOpts, arg0)
}

// RegisteredERC20 is a free data retrieval call binding the contract method 0x89b15604.
//
// Solidity: function registeredERC20(address ) view returns(bool)
func (_ETHSwapAgent *ETHSwapAgentCallerSession) RegisteredERC20(arg0 common.Address) (bool, error) {
	return _ETHSwapAgent.Contract.RegisteredERC20(&_ETHSwapAgent.CallOpts, arg0)
}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_ETHSwapAgent *ETHSwapAgentCaller) SwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHSwapAgent.contract.Call(opts, &out, "swapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_ETHSwapAgent *ETHSwapAgentSession) SwapFee() (*big.Int, error) {
	return _ETHSwapAgent.Contract.SwapFee(&_ETHSwapAgent.CallOpts)
}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_ETHSwapAgent *ETHSwapAgentCallerSession) SwapFee() (*big.Int, error) {
	return _ETHSwapAgent.Contract.SwapFee(&_ETHSwapAgent.CallOpts)
}

// FillBSC2ETHSwap is a paid mutator transaction binding the contract method 0x9867df11.
//
// Solidity: function fillBSC2ETHSwap(bytes32 bscTxHash, address erc20Addr, address toAddress, uint256 amount) returns(bool)
func (_ETHSwapAgent *ETHSwapAgentTransactor) FillBSC2ETHSwap(opts *bind.TransactOpts, bscTxHash [32]byte, erc20Addr common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHSwapAgent.contract.Transact(opts, "fillBSC2ETHSwap", bscTxHash, erc20Addr, toAddress, amount)
}

// FillBSC2ETHSwap is a paid mutator transaction binding the contract method 0x9867df11.
//
// Solidity: function fillBSC2ETHSwap(bytes32 bscTxHash, address erc20Addr, address toAddress, uint256 amount) returns(bool)
func (_ETHSwapAgent *ETHSwapAgentSession) FillBSC2ETHSwap(bscTxHash [32]byte, erc20Addr common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.FillBSC2ETHSwap(&_ETHSwapAgent.TransactOpts, bscTxHash, erc20Addr, toAddress, amount)
}

// FillBSC2ETHSwap is a paid mutator transaction binding the contract method 0x9867df11.
//
// Solidity: function fillBSC2ETHSwap(bytes32 bscTxHash, address erc20Addr, address toAddress, uint256 amount) returns(bool)
func (_ETHSwapAgent *ETHSwapAgentTransactorSession) FillBSC2ETHSwap(bscTxHash [32]byte, erc20Addr common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.FillBSC2ETHSwap(&_ETHSwapAgent.TransactOpts, bscTxHash, erc20Addr, toAddress, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 fee, address ownerAddr) returns()
func (_ETHSwapAgent *ETHSwapAgentTransactor) Initialize(opts *bind.TransactOpts, fee *big.Int, ownerAddr common.Address) (*types.Transaction, error) {
	return _ETHSwapAgent.contract.Transact(opts, "initialize", fee, ownerAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 fee, address ownerAddr) returns()
func (_ETHSwapAgent *ETHSwapAgentSession) Initialize(fee *big.Int, ownerAddr common.Address) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.Initialize(&_ETHSwapAgent.TransactOpts, fee, ownerAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 fee, address ownerAddr) returns()
func (_ETHSwapAgent *ETHSwapAgentTransactorSession) Initialize(fee *big.Int, ownerAddr common.Address) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.Initialize(&_ETHSwapAgent.TransactOpts, fee, ownerAddr)
}

// RegisterSwapPairToBSC is a paid mutator transaction binding the contract method 0x5c13c151.
//
// Solidity: function registerSwapPairToBSC(address erc20Addr) returns(bool)
func (_ETHSwapAgent *ETHSwapAgentTransactor) RegisterSwapPairToBSC(opts *bind.TransactOpts, erc20Addr common.Address) (*types.Transaction, error) {
	return _ETHSwapAgent.contract.Transact(opts, "registerSwapPairToBSC", erc20Addr)
}

// RegisterSwapPairToBSC is a paid mutator transaction binding the contract method 0x5c13c151.
//
// Solidity: function registerSwapPairToBSC(address erc20Addr) returns(bool)
func (_ETHSwapAgent *ETHSwapAgentSession) RegisterSwapPairToBSC(erc20Addr common.Address) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.RegisterSwapPairToBSC(&_ETHSwapAgent.TransactOpts, erc20Addr)
}

// RegisterSwapPairToBSC is a paid mutator transaction binding the contract method 0x5c13c151.
//
// Solidity: function registerSwapPairToBSC(address erc20Addr) returns(bool)
func (_ETHSwapAgent *ETHSwapAgentTransactorSession) RegisterSwapPairToBSC(erc20Addr common.Address) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.RegisterSwapPairToBSC(&_ETHSwapAgent.TransactOpts, erc20Addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETHSwapAgent *ETHSwapAgentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHSwapAgent.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETHSwapAgent *ETHSwapAgentSession) RenounceOwnership() (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.RenounceOwnership(&_ETHSwapAgent.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETHSwapAgent *ETHSwapAgentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.RenounceOwnership(&_ETHSwapAgent.TransactOpts)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 fee) returns()
func (_ETHSwapAgent *ETHSwapAgentTransactor) SetSwapFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _ETHSwapAgent.contract.Transact(opts, "setSwapFee", fee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 fee) returns()
func (_ETHSwapAgent *ETHSwapAgentSession) SetSwapFee(fee *big.Int) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.SetSwapFee(&_ETHSwapAgent.TransactOpts, fee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 fee) returns()
func (_ETHSwapAgent *ETHSwapAgentTransactorSession) SetSwapFee(fee *big.Int) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.SetSwapFee(&_ETHSwapAgent.TransactOpts, fee)
}

// SwapETH2BSC is a paid mutator transaction binding the contract method 0xb9927a9c.
//
// Solidity: function swapETH2BSC(address erc20Addr, uint256 amount) payable returns(bool)
func (_ETHSwapAgent *ETHSwapAgentTransactor) SwapETH2BSC(opts *bind.TransactOpts, erc20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHSwapAgent.contract.Transact(opts, "swapETH2BSC", erc20Addr, amount)
}

// SwapETH2BSC is a paid mutator transaction binding the contract method 0xb9927a9c.
//
// Solidity: function swapETH2BSC(address erc20Addr, uint256 amount) payable returns(bool)
func (_ETHSwapAgent *ETHSwapAgentSession) SwapETH2BSC(erc20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.SwapETH2BSC(&_ETHSwapAgent.TransactOpts, erc20Addr, amount)
}

// SwapETH2BSC is a paid mutator transaction binding the contract method 0xb9927a9c.
//
// Solidity: function swapETH2BSC(address erc20Addr, uint256 amount) payable returns(bool)
func (_ETHSwapAgent *ETHSwapAgentTransactorSession) SwapETH2BSC(erc20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.SwapETH2BSC(&_ETHSwapAgent.TransactOpts, erc20Addr, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETHSwapAgent *ETHSwapAgentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ETHSwapAgent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETHSwapAgent *ETHSwapAgentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.TransferOwnership(&_ETHSwapAgent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETHSwapAgent *ETHSwapAgentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ETHSwapAgent.Contract.TransferOwnership(&_ETHSwapAgent.TransactOpts, newOwner)
}

// ETHSwapAgentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ETHSwapAgent contract.
type ETHSwapAgentOwnershipTransferredIterator struct {
	Event *ETHSwapAgentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ETHSwapAgentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHSwapAgentOwnershipTransferred)
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
		it.Event = new(ETHSwapAgentOwnershipTransferred)
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
func (it *ETHSwapAgentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHSwapAgentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHSwapAgentOwnershipTransferred represents a OwnershipTransferred event raised by the ETHSwapAgent contract.
type ETHSwapAgentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ETHSwapAgent *ETHSwapAgentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ETHSwapAgentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ETHSwapAgent.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ETHSwapAgentOwnershipTransferredIterator{contract: _ETHSwapAgent.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ETHSwapAgent *ETHSwapAgentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ETHSwapAgentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ETHSwapAgent.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHSwapAgentOwnershipTransferred)
				if err := _ETHSwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ETHSwapAgent *ETHSwapAgentFilterer) ParseOwnershipTransferred(log types.Log) (*ETHSwapAgentOwnershipTransferred, error) {
	event := new(ETHSwapAgentOwnershipTransferred)
	if err := _ETHSwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHSwapAgentSwapFilledIterator is returned from FilterSwapFilled and is used to iterate over the raw logs and unpacked data for SwapFilled events raised by the ETHSwapAgent contract.
type ETHSwapAgentSwapFilledIterator struct {
	Event *ETHSwapAgentSwapFilled // Event containing the contract specifics and raw log

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
func (it *ETHSwapAgentSwapFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHSwapAgentSwapFilled)
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
		it.Event = new(ETHSwapAgentSwapFilled)
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
func (it *ETHSwapAgentSwapFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHSwapAgentSwapFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHSwapAgentSwapFilled represents a SwapFilled event raised by the ETHSwapAgent contract.
type ETHSwapAgentSwapFilled struct {
	Erc20Addr common.Address
	BscTxHash [32]byte
	ToAddress common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapFilled is a free log retrieval operation binding the contract event 0x3bebd9a738291e69898b5dbfadb6329b4b09fc648bdef68762928e521463abd9.
//
// Solidity: event SwapFilled(address indexed erc20Addr, bytes32 indexed bscTxHash, address indexed toAddress, uint256 amount)
func (_ETHSwapAgent *ETHSwapAgentFilterer) FilterSwapFilled(opts *bind.FilterOpts, erc20Addr []common.Address, bscTxHash [][32]byte, toAddress []common.Address) (*ETHSwapAgentSwapFilledIterator, error) {

	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}
	var bscTxHashRule []interface{}
	for _, bscTxHashItem := range bscTxHash {
		bscTxHashRule = append(bscTxHashRule, bscTxHashItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _ETHSwapAgent.contract.FilterLogs(opts, "SwapFilled", erc20AddrRule, bscTxHashRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &ETHSwapAgentSwapFilledIterator{contract: _ETHSwapAgent.contract, event: "SwapFilled", logs: logs, sub: sub}, nil
}

// WatchSwapFilled is a free log subscription operation binding the contract event 0x3bebd9a738291e69898b5dbfadb6329b4b09fc648bdef68762928e521463abd9.
//
// Solidity: event SwapFilled(address indexed erc20Addr, bytes32 indexed bscTxHash, address indexed toAddress, uint256 amount)
func (_ETHSwapAgent *ETHSwapAgentFilterer) WatchSwapFilled(opts *bind.WatchOpts, sink chan<- *ETHSwapAgentSwapFilled, erc20Addr []common.Address, bscTxHash [][32]byte, toAddress []common.Address) (event.Subscription, error) {

	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}
	var bscTxHashRule []interface{}
	for _, bscTxHashItem := range bscTxHash {
		bscTxHashRule = append(bscTxHashRule, bscTxHashItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _ETHSwapAgent.contract.WatchLogs(opts, "SwapFilled", erc20AddrRule, bscTxHashRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHSwapAgentSwapFilled)
				if err := _ETHSwapAgent.contract.UnpackLog(event, "SwapFilled", log); err != nil {
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

// ParseSwapFilled is a log parse operation binding the contract event 0x3bebd9a738291e69898b5dbfadb6329b4b09fc648bdef68762928e521463abd9.
//
// Solidity: event SwapFilled(address indexed erc20Addr, bytes32 indexed bscTxHash, address indexed toAddress, uint256 amount)
func (_ETHSwapAgent *ETHSwapAgentFilterer) ParseSwapFilled(log types.Log) (*ETHSwapAgentSwapFilled, error) {
	event := new(ETHSwapAgentSwapFilled)
	if err := _ETHSwapAgent.contract.UnpackLog(event, "SwapFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHSwapAgentSwapPairRegisterIterator is returned from FilterSwapPairRegister and is used to iterate over the raw logs and unpacked data for SwapPairRegister events raised by the ETHSwapAgent contract.
type ETHSwapAgentSwapPairRegisterIterator struct {
	Event *ETHSwapAgentSwapPairRegister // Event containing the contract specifics and raw log

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
func (it *ETHSwapAgentSwapPairRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHSwapAgentSwapPairRegister)
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
		it.Event = new(ETHSwapAgentSwapPairRegister)
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
func (it *ETHSwapAgentSwapPairRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHSwapAgentSwapPairRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHSwapAgentSwapPairRegister represents a SwapPairRegister event raised by the ETHSwapAgent contract.
type ETHSwapAgentSwapPairRegister struct {
	Sponsor   common.Address
	Erc20Addr common.Address
	Name      string
	Symbol    string
	Decimals  uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapPairRegister is a free log retrieval operation binding the contract event 0xfe3bd005e346323fa452df8cafc28c55b99e3766ba8750571d139c6cf5bc08a0.
//
// Solidity: event SwapPairRegister(address indexed sponsor, address indexed erc20Addr, string name, string symbol, uint8 decimals)
func (_ETHSwapAgent *ETHSwapAgentFilterer) FilterSwapPairRegister(opts *bind.FilterOpts, sponsor []common.Address, erc20Addr []common.Address) (*ETHSwapAgentSwapPairRegisterIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}

	logs, sub, err := _ETHSwapAgent.contract.FilterLogs(opts, "SwapPairRegister", sponsorRule, erc20AddrRule)
	if err != nil {
		return nil, err
	}
	return &ETHSwapAgentSwapPairRegisterIterator{contract: _ETHSwapAgent.contract, event: "SwapPairRegister", logs: logs, sub: sub}, nil
}

// WatchSwapPairRegister is a free log subscription operation binding the contract event 0xfe3bd005e346323fa452df8cafc28c55b99e3766ba8750571d139c6cf5bc08a0.
//
// Solidity: event SwapPairRegister(address indexed sponsor, address indexed erc20Addr, string name, string symbol, uint8 decimals)
func (_ETHSwapAgent *ETHSwapAgentFilterer) WatchSwapPairRegister(opts *bind.WatchOpts, sink chan<- *ETHSwapAgentSwapPairRegister, sponsor []common.Address, erc20Addr []common.Address) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}

	logs, sub, err := _ETHSwapAgent.contract.WatchLogs(opts, "SwapPairRegister", sponsorRule, erc20AddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHSwapAgentSwapPairRegister)
				if err := _ETHSwapAgent.contract.UnpackLog(event, "SwapPairRegister", log); err != nil {
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

// ParseSwapPairRegister is a log parse operation binding the contract event 0xfe3bd005e346323fa452df8cafc28c55b99e3766ba8750571d139c6cf5bc08a0.
//
// Solidity: event SwapPairRegister(address indexed sponsor, address indexed erc20Addr, string name, string symbol, uint8 decimals)
func (_ETHSwapAgent *ETHSwapAgentFilterer) ParseSwapPairRegister(log types.Log) (*ETHSwapAgentSwapPairRegister, error) {
	event := new(ETHSwapAgentSwapPairRegister)
	if err := _ETHSwapAgent.contract.UnpackLog(event, "SwapPairRegister", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHSwapAgentSwapStartedIterator is returned from FilterSwapStarted and is used to iterate over the raw logs and unpacked data for SwapStarted events raised by the ETHSwapAgent contract.
type ETHSwapAgentSwapStartedIterator struct {
	Event *ETHSwapAgentSwapStarted // Event containing the contract specifics and raw log

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
func (it *ETHSwapAgentSwapStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHSwapAgentSwapStarted)
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
		it.Event = new(ETHSwapAgentSwapStarted)
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
func (it *ETHSwapAgentSwapStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHSwapAgentSwapStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHSwapAgentSwapStarted represents a SwapStarted event raised by the ETHSwapAgent contract.
type ETHSwapAgentSwapStarted struct {
	Erc20Addr common.Address
	FromAddr  common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapStarted is a free log retrieval operation binding the contract event 0xf60309f865a6aa297da5fac6188136a02e5acfdf6e8f6d35257a9f4e9653170f.
//
// Solidity: event SwapStarted(address indexed erc20Addr, address indexed fromAddr, uint256 amount, uint256 feeAmount)
func (_ETHSwapAgent *ETHSwapAgentFilterer) FilterSwapStarted(opts *bind.FilterOpts, erc20Addr []common.Address, fromAddr []common.Address) (*ETHSwapAgentSwapStartedIterator, error) {

	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}
	var fromAddrRule []interface{}
	for _, fromAddrItem := range fromAddr {
		fromAddrRule = append(fromAddrRule, fromAddrItem)
	}

	logs, sub, err := _ETHSwapAgent.contract.FilterLogs(opts, "SwapStarted", erc20AddrRule, fromAddrRule)
	if err != nil {
		return nil, err
	}
	return &ETHSwapAgentSwapStartedIterator{contract: _ETHSwapAgent.contract, event: "SwapStarted", logs: logs, sub: sub}, nil
}

// WatchSwapStarted is a free log subscription operation binding the contract event 0xf60309f865a6aa297da5fac6188136a02e5acfdf6e8f6d35257a9f4e9653170f.
//
// Solidity: event SwapStarted(address indexed erc20Addr, address indexed fromAddr, uint256 amount, uint256 feeAmount)
func (_ETHSwapAgent *ETHSwapAgentFilterer) WatchSwapStarted(opts *bind.WatchOpts, sink chan<- *ETHSwapAgentSwapStarted, erc20Addr []common.Address, fromAddr []common.Address) (event.Subscription, error) {

	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}
	var fromAddrRule []interface{}
	for _, fromAddrItem := range fromAddr {
		fromAddrRule = append(fromAddrRule, fromAddrItem)
	}

	logs, sub, err := _ETHSwapAgent.contract.WatchLogs(opts, "SwapStarted", erc20AddrRule, fromAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHSwapAgentSwapStarted)
				if err := _ETHSwapAgent.contract.UnpackLog(event, "SwapStarted", log); err != nil {
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

// ParseSwapStarted is a log parse operation binding the contract event 0xf60309f865a6aa297da5fac6188136a02e5acfdf6e8f6d35257a9f4e9653170f.
//
// Solidity: event SwapStarted(address indexed erc20Addr, address indexed fromAddr, uint256 amount, uint256 feeAmount)
func (_ETHSwapAgent *ETHSwapAgentFilterer) ParseSwapStarted(log types.Log) (*ETHSwapAgentSwapStarted, error) {
	event := new(ETHSwapAgentSwapStarted)
	if err := _ETHSwapAgent.contract.UnpackLog(event, "SwapStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
