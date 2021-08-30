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

// BSCSwapAgentABI is the input ABI used to generate the binding from.
const BSCSwapAgentABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bep20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ethTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SwapFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ethRegisterTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bep20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"SwapPairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bep20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"SwapStarted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bep20Implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bep20ProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ethTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createSwapPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ethTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"erc20Addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fillETH2BSCSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filledETHTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bep20Impl\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"ownerAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bep20ProxyAdminAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bep20Addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swapBSC2ETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapMappingBSC2ETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapMappingETH2BSC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BSCSwapAgentBin is the compiled bytecode used for deploying new contracts.
var BSCSwapAgentBin = "0x608060405234801561001057600080fd5b50612f3f806100206000396000f3fe608060405260043610620000f25760003560e01c806360b810f1116200008b5780638da5cb5b11620000615780638da5cb5b1462000582578063be0ace6914620005dc578063e307b9311462000671578063f2fde38b146200071257620000f2565b806360b810f1146200047957806366fec65c146200050e578063715018a6146200056857620000f2565b806334e1990711620000cd57806334e199071462000316578063358394d814620003555780634e2dc7f114620003f457806354cf2aeb146200044b57620000f2565b80630344165a14620000f75780631ba3b150146200015157806332bd6e3114620001ba575b600080fd5b3480156200010457600080fd5b506200010f62000767565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b620001a0600480360360408110156200016957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506200078d565b604051808215151515815260200191505060405180910390f35b348015620001c757600080fd5b50620002d4600480360360a0811015620001e057600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156200022857600080fd5b8201836020820111156200023b57600080fd5b803590602001918460018302840111640100000000831117156200025e57600080fd5b9091929391929390803590602001906401000000008111156200028057600080fd5b8201836020820111156200029357600080fd5b80359060200191846001830284011164010000000083111715620002b657600080fd5b9091929391929390803560ff16906020019092919050505062000bf4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156200032357600080fd5b5062000353600480360360208110156200033c57600080fd5b8101908080359060200190929190505050620011c2565b005b3480156200036257600080fd5b50620003f2600480360360808110156200037b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505062001299565b005b3480156200040157600080fd5b5062000431600480360360208110156200041a57600080fd5b81019080803590602001909291905050506200146e565b604051808215151515815260200191505060405180910390f35b3480156200045857600080fd5b50620004636200148e565b6040518082815260200191505060405180910390f35b3480156200048657600080fd5b50620004cc600480360360208110156200049f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505062001494565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156200051b57600080fd5b5062000526620014c7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156200057557600080fd5b5062000580620014ed565b005b3480156200058f57600080fd5b506200059a6200167b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b348015620005e957600080fd5b506200062f600480360360208110156200060257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050620016a1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156200067e57600080fd5b50620006f8600480360360808110156200069757600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050620016d4565b604051808215151515815260200191505060405180910390f35b3480156200071f57600080fd5b5062000765600480360360208110156200073857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505062001aa5565b005b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006200079a3362001cba565b156200080e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f636f6e7472616374206973206e6f7420616c6c6f77656420746f20737761700081525060200191505060405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620008b0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6e6f2070726f787920636f6e747261637420697320616c6c6f7765640000000081525060200191505060405180910390fd5b6000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620009b8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f6e6f2073776170207061697220666f72207468697320746f6b656e000000000081525060200191505060405180910390fd5b600754341462000a30576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f7377617020666565206e6f7420657175616c000000000000000000000000000081525060200191505060405180910390fd5b62000a5f3330858773ffffffffffffffffffffffffffffffffffffffff1662001ccd909392919063ffffffff16565b8373ffffffffffffffffffffffffffffffffffffffff166342966c68846040518263ffffffff1660e01b815260040180828152602001915050602060405180830381600087803b15801562000ab357600080fd5b505af115801562000ac8573d6000803e3d6000fd5b505050506040513d602081101562000adf57600080fd5b8101908080519060200190929190505050506000341462000b6557600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f1935050505015801562000b63573d6000803e3d6000fd5b505b3373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f49c08ff11118922c1e8298915531eff9ef6f8b39b44b3e9952b75d47e1d0cdd08634604051808381526020018281526020019250505060405180910390a4600191505092915050565b600062000c0062001dbc565b73ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161462000cc3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161462000dc5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6475706c6963617465642073776170207061697200000000000000000000000081525060200191505060405180910390fd5b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405162000e1b9062002183565b808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018281038252600081526020016020019350505050604051809103906000f08015801562000eb5573d6000803e3d6000fd5b50905060008190508073ffffffffffffffffffffffffffffffffffffffff1663ef3ebcb8898989898960006001306040518963ffffffff1660e01b81526004018080602001806020018760ff1660ff168152602001868152602001851515151581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200183810383528b8b82818152602001925080828437600081840152601f19601f8201169050808301925050508381038252898982818152602001925080828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b15801562000fc957600080fd5b505af115801562000fde573d6000803e3d6000fd5b5050505080600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555088600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508873ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff168b7fcc0314763eabceb74cd3d30ae785c09bfe4e204af2088b3bfcdbbe5082133db589898d8d8b6040518080602001806020018460ff1660ff1681526020018381038352888882818152602001925080828437600081840152601f19601f8201169050808301925050508381038252868682818152602001925080828437600081840152601f19601f82011690508083019250505097505050505050505060405180910390a48092505050979650505050505050565b620011cc62001dbc565b73ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146200128f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b8060078190555050565b600060019054906101000a900460ff1680620012bb5750620012ba62001dc4565b5b80620012d357506000809054906101000a900460ff16155b6200132a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e81526020018062002eb2602e913960400191505060405180910390fd5b60008060019054906101000a900460ff1615905080156200137b576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b84600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360078190555082600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015620014675760008060016101000a81548160ff0219169083151502179055505b5050505050565b60036020528060005260406000206000915054906101000a900460ff1681565b60075481565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b620014f762001dbc565b73ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614620015ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000620016e062001dbc565b73ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614620017a3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6003600086815260200190815260200160002060009054906101000a900460ff161562001838576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f6574682074782066696c6c656420616c7265616479000000000000000000000081525060200191505060405180910390fd5b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141562001940576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f6e6f2073776170207061697220666f72207468697320746f6b656e000000000081525060200191505060405180910390fd5b60016003600088815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff1663b723b34e84866040518363ffffffff1660e01b8152600401808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200192505050602060405180830381600087803b158015620019f457600080fd5b505af115801562001a09573d6000803e3d6000fd5b505050506040513d602081101562001a2057600080fd5b8101908080519060200190929190505050508373ffffffffffffffffffffffffffffffffffffffff16868273ffffffffffffffffffffffffffffffffffffffff167f3bebd9a738291e69898b5dbfadb6329b4b09fc648bdef68762928e521463abd9866040518082815260200191505060405180910390a46001915050949350505050565b62001aaf62001dbc565b73ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161462001b72576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141562001bfa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602681526020018062002e666026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080823b905060008111915050919050565b62001db6846323b872dd60e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505062001ddb565b50505050565b600033905090565b6000803090506000813b9050600081149250505090565b606062001e3f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1662001ed09092919063ffffffff16565b905060008151111562001ecb5780806020019051602081101562001e6257600080fd5b810190808051906020019092919050505062001eca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a81526020018062002ee0602a913960400191505060405180910390fd5b5b505050565b606062001ee1848460008562001eea565b90509392505050565b60608247101562001f47576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602681526020018062002e8c6026913960400191505060405180910390fd5b62001f52856200209e565b62001fc5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000081525060200191505060405180910390fd5b600060608673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b6020831062002017578051825260208201915060208101905060208303925062001ff2565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146200207b576040519150601f19603f3d011682016040523d82523d6000602084013e62002080565b606091505b509150915062002092828286620020b1565b92505050949350505050565b600080823b905060008111915050919050565b60608315620020c3578290506200217c565b600083511115620020d75782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156200214057808201518184015260208101905062002123565b50505050905090810190601f1680156200216e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b9392505050565b610cd480620021928339019056fe608060405234801561001057600080fd5b50604051610cd4380380610cd48339818101604052606081101561003357600080fd5b8101908080519060200190929190805190602001909291908051604051939291908464010000000082111561006757600080fd5b8382019150602082018581111561007d57600080fd5b825186600182028301116401000000008211171561009a57600080fd5b8083526020830192505050908051906020019080838360005b838110156100ce5780820151818401526020810190506100b3565b50505050905090810190601f1680156100fb5780820380516001836020036101000a031916815260200191505b506040525050508282828281600160405180807f656970313936372e70726f78792e696d706c656d656e746174696f6e00000000815250601c019050604051809103902060001c0360001b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b1461017157fe5b610180826102d260201b60201c565b60008151111561024c5760008273ffffffffffffffffffffffffffffffffffffffff16826040518082805190602001908083835b602083106101d757805182526020820191506020810190506020830392506101b4565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114610237576040519150601f19603f3d011682016040523d82523d6000602084013e61023c565b606091505b505090508061024a57600080fd5b505b5050600160405180807f656970313936372e70726f78792e61646d696e000000000000000000000000008152506013019050604051809103902060001c0360001b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b146102b857fe5b6102c78261036960201b60201c565b5050505050506103ab565b6102e58161039860201b6107e91760201c565b61033a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526036815260200180610c9e6036913960400191505060405180910390fd5b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b90508181555050565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b90508181555050565b600080823b905060008111915050919050565b6108e4806103ba6000396000f3fe60806040526004361061004e5760003560e01c80633659cfe6146100675780634f1ef286146100b85780635c60da1b146101515780638f283970146101a8578063f851a440146101f95761005d565b3661005d5761005b610250565b005b610065610250565b005b34801561007357600080fd5b506100b66004803603602081101561008a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061026a565b005b61014f600480360360408110156100ce57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561010b57600080fd5b82018360208201111561011d57600080fd5b8035906020019184600183028401116401000000008311171561013f57600080fd5b90919293919293905050506102bf565b005b34801561015d57600080fd5b50610166610395565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101b457600080fd5b506101f7600480360360208110156101cb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103ed565b005b34801561020557600080fd5b5061020e610566565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102586105be565b610268610263610654565b610685565b565b6102726106ab565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102b3576102ae816106dc565b6102bc565b6102bb610250565b5b50565b6102c76106ab565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561038757610303836106dc565b60008373ffffffffffffffffffffffffffffffffffffffff168383604051808383808284378083019250505092505050600060405180830381855af49150503d806000811461036e576040519150601f19603f3d011682016040523d82523d6000602084013e610373565b606091505b505090508061038157600080fd5b50610390565b61038f610250565b5b505050565b600061039f6106ab565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103e1576103da610654565b90506103ea565b6103e9610250565b5b90565b6103f56106ab565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561055a57600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156104ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603a8152602001806107fd603a913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6104d76106ab565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a16105558161072b565b610563565b610562610250565b5b50565b60006105706106ab565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156105b2576105ab6106ab565b90506105bb565b6105ba610250565b5b90565b6105c66106ab565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561064a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604281526020018061086d6042913960600191505060405180910390fd5b61065261075a565b565b6000807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b9050805491505090565b3660008037600080366000845af43d6000803e80600081146106a6573d6000f35b3d6000fd5b6000807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b9050805491505090565b6106e58161075c565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610360001b90508181555050565b565b610765816107e9565b6107ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001806108376036913960400191505060405180910390fd5b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b90508181555050565b600080823b90506000811191505091905056fe5472616e73706172656e745570677261646561626c6550726f78793a206e65772061646d696e20697320746865207a65726f20616464726573735570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a2646970667358221220954570f32ef9b171d2aed7c17e3bf75619dff899e18f390df6f4330ef01fb6f364736f6c634300060400335570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e74726163744f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a65645361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a26469706673582212201bc729f19df3271290300ce6b72e4cbd5b2984af23e85b0ae89e8d9cad8ee4b364736f6c63430006040033"

// DeployBSCSwapAgent deploys a new Ethereum contract, binding an instance of BSCSwapAgent to it.
func DeployBSCSwapAgent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BSCSwapAgent, error) {
	parsed, err := abi.JSON(strings.NewReader(BSCSwapAgentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BSCSwapAgentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BSCSwapAgent{BSCSwapAgentCaller: BSCSwapAgentCaller{contract: contract}, BSCSwapAgentTransactor: BSCSwapAgentTransactor{contract: contract}, BSCSwapAgentFilterer: BSCSwapAgentFilterer{contract: contract}}, nil
}

// BSCSwapAgent is an auto generated Go binding around an Ethereum contract.
type BSCSwapAgent struct {
	BSCSwapAgentCaller     // Read-only binding to the contract
	BSCSwapAgentTransactor // Write-only binding to the contract
	BSCSwapAgentFilterer   // Log filterer for contract events
}

// BSCSwapAgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type BSCSwapAgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCSwapAgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BSCSwapAgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCSwapAgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BSCSwapAgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BSCSwapAgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BSCSwapAgentSession struct {
	Contract     *BSCSwapAgent     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BSCSwapAgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BSCSwapAgentCallerSession struct {
	Contract *BSCSwapAgentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BSCSwapAgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BSCSwapAgentTransactorSession struct {
	Contract     *BSCSwapAgentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BSCSwapAgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type BSCSwapAgentRaw struct {
	Contract *BSCSwapAgent // Generic contract binding to access the raw methods on
}

// BSCSwapAgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BSCSwapAgentCallerRaw struct {
	Contract *BSCSwapAgentCaller // Generic read-only contract binding to access the raw methods on
}

// BSCSwapAgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BSCSwapAgentTransactorRaw struct {
	Contract *BSCSwapAgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBSCSwapAgent creates a new instance of BSCSwapAgent, bound to a specific deployed contract.
func NewBSCSwapAgent(address common.Address, backend bind.ContractBackend) (*BSCSwapAgent, error) {
	contract, err := bindBSCSwapAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BSCSwapAgent{BSCSwapAgentCaller: BSCSwapAgentCaller{contract: contract}, BSCSwapAgentTransactor: BSCSwapAgentTransactor{contract: contract}, BSCSwapAgentFilterer: BSCSwapAgentFilterer{contract: contract}}, nil
}

// NewBSCSwapAgentCaller creates a new read-only instance of BSCSwapAgent, bound to a specific deployed contract.
func NewBSCSwapAgentCaller(address common.Address, caller bind.ContractCaller) (*BSCSwapAgentCaller, error) {
	contract, err := bindBSCSwapAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BSCSwapAgentCaller{contract: contract}, nil
}

// NewBSCSwapAgentTransactor creates a new write-only instance of BSCSwapAgent, bound to a specific deployed contract.
func NewBSCSwapAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*BSCSwapAgentTransactor, error) {
	contract, err := bindBSCSwapAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BSCSwapAgentTransactor{contract: contract}, nil
}

// NewBSCSwapAgentFilterer creates a new log filterer instance of BSCSwapAgent, bound to a specific deployed contract.
func NewBSCSwapAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*BSCSwapAgentFilterer, error) {
	contract, err := bindBSCSwapAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BSCSwapAgentFilterer{contract: contract}, nil
}

// bindBSCSwapAgent binds a generic wrapper to an already deployed contract.
func bindBSCSwapAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BSCSwapAgentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCSwapAgent *BSCSwapAgentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCSwapAgent.Contract.BSCSwapAgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCSwapAgent *BSCSwapAgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.BSCSwapAgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCSwapAgent *BSCSwapAgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.BSCSwapAgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BSCSwapAgent *BSCSwapAgentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BSCSwapAgent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BSCSwapAgent *BSCSwapAgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BSCSwapAgent *BSCSwapAgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.contract.Transact(opts, method, params...)
}

// Bep20Implementation is a free data retrieval call binding the contract method 0x66fec65c.
//
// Solidity: function bep20Implementation() view returns(address)
func (_BSCSwapAgent *BSCSwapAgentCaller) Bep20Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCSwapAgent.contract.Call(opts, &out, "bep20Implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bep20Implementation is a free data retrieval call binding the contract method 0x66fec65c.
//
// Solidity: function bep20Implementation() view returns(address)
func (_BSCSwapAgent *BSCSwapAgentSession) Bep20Implementation() (common.Address, error) {
	return _BSCSwapAgent.Contract.Bep20Implementation(&_BSCSwapAgent.CallOpts)
}

// Bep20Implementation is a free data retrieval call binding the contract method 0x66fec65c.
//
// Solidity: function bep20Implementation() view returns(address)
func (_BSCSwapAgent *BSCSwapAgentCallerSession) Bep20Implementation() (common.Address, error) {
	return _BSCSwapAgent.Contract.Bep20Implementation(&_BSCSwapAgent.CallOpts)
}

// Bep20ProxyAdmin is a free data retrieval call binding the contract method 0x0344165a.
//
// Solidity: function bep20ProxyAdmin() view returns(address)
func (_BSCSwapAgent *BSCSwapAgentCaller) Bep20ProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCSwapAgent.contract.Call(opts, &out, "bep20ProxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bep20ProxyAdmin is a free data retrieval call binding the contract method 0x0344165a.
//
// Solidity: function bep20ProxyAdmin() view returns(address)
func (_BSCSwapAgent *BSCSwapAgentSession) Bep20ProxyAdmin() (common.Address, error) {
	return _BSCSwapAgent.Contract.Bep20ProxyAdmin(&_BSCSwapAgent.CallOpts)
}

// Bep20ProxyAdmin is a free data retrieval call binding the contract method 0x0344165a.
//
// Solidity: function bep20ProxyAdmin() view returns(address)
func (_BSCSwapAgent *BSCSwapAgentCallerSession) Bep20ProxyAdmin() (common.Address, error) {
	return _BSCSwapAgent.Contract.Bep20ProxyAdmin(&_BSCSwapAgent.CallOpts)
}

// FilledETHTx is a free data retrieval call binding the contract method 0x4e2dc7f1.
//
// Solidity: function filledETHTx(bytes32 ) view returns(bool)
func (_BSCSwapAgent *BSCSwapAgentCaller) FilledETHTx(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _BSCSwapAgent.contract.Call(opts, &out, "filledETHTx", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FilledETHTx is a free data retrieval call binding the contract method 0x4e2dc7f1.
//
// Solidity: function filledETHTx(bytes32 ) view returns(bool)
func (_BSCSwapAgent *BSCSwapAgentSession) FilledETHTx(arg0 [32]byte) (bool, error) {
	return _BSCSwapAgent.Contract.FilledETHTx(&_BSCSwapAgent.CallOpts, arg0)
}

// FilledETHTx is a free data retrieval call binding the contract method 0x4e2dc7f1.
//
// Solidity: function filledETHTx(bytes32 ) view returns(bool)
func (_BSCSwapAgent *BSCSwapAgentCallerSession) FilledETHTx(arg0 [32]byte) (bool, error) {
	return _BSCSwapAgent.Contract.FilledETHTx(&_BSCSwapAgent.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BSCSwapAgent *BSCSwapAgentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BSCSwapAgent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BSCSwapAgent *BSCSwapAgentSession) Owner() (common.Address, error) {
	return _BSCSwapAgent.Contract.Owner(&_BSCSwapAgent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BSCSwapAgent *BSCSwapAgentCallerSession) Owner() (common.Address, error) {
	return _BSCSwapAgent.Contract.Owner(&_BSCSwapAgent.CallOpts)
}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_BSCSwapAgent *BSCSwapAgentCaller) SwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BSCSwapAgent.contract.Call(opts, &out, "swapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_BSCSwapAgent *BSCSwapAgentSession) SwapFee() (*big.Int, error) {
	return _BSCSwapAgent.Contract.SwapFee(&_BSCSwapAgent.CallOpts)
}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_BSCSwapAgent *BSCSwapAgentCallerSession) SwapFee() (*big.Int, error) {
	return _BSCSwapAgent.Contract.SwapFee(&_BSCSwapAgent.CallOpts)
}

// SwapMappingBSC2ETH is a free data retrieval call binding the contract method 0xbe0ace69.
//
// Solidity: function swapMappingBSC2ETH(address ) view returns(address)
func (_BSCSwapAgent *BSCSwapAgentCaller) SwapMappingBSC2ETH(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _BSCSwapAgent.contract.Call(opts, &out, "swapMappingBSC2ETH", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingBSC2ETH is a free data retrieval call binding the contract method 0xbe0ace69.
//
// Solidity: function swapMappingBSC2ETH(address ) view returns(address)
func (_BSCSwapAgent *BSCSwapAgentSession) SwapMappingBSC2ETH(arg0 common.Address) (common.Address, error) {
	return _BSCSwapAgent.Contract.SwapMappingBSC2ETH(&_BSCSwapAgent.CallOpts, arg0)
}

// SwapMappingBSC2ETH is a free data retrieval call binding the contract method 0xbe0ace69.
//
// Solidity: function swapMappingBSC2ETH(address ) view returns(address)
func (_BSCSwapAgent *BSCSwapAgentCallerSession) SwapMappingBSC2ETH(arg0 common.Address) (common.Address, error) {
	return _BSCSwapAgent.Contract.SwapMappingBSC2ETH(&_BSCSwapAgent.CallOpts, arg0)
}

// SwapMappingETH2BSC is a free data retrieval call binding the contract method 0x60b810f1.
//
// Solidity: function swapMappingETH2BSC(address ) view returns(address)
func (_BSCSwapAgent *BSCSwapAgentCaller) SwapMappingETH2BSC(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _BSCSwapAgent.contract.Call(opts, &out, "swapMappingETH2BSC", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingETH2BSC is a free data retrieval call binding the contract method 0x60b810f1.
//
// Solidity: function swapMappingETH2BSC(address ) view returns(address)
func (_BSCSwapAgent *BSCSwapAgentSession) SwapMappingETH2BSC(arg0 common.Address) (common.Address, error) {
	return _BSCSwapAgent.Contract.SwapMappingETH2BSC(&_BSCSwapAgent.CallOpts, arg0)
}

// SwapMappingETH2BSC is a free data retrieval call binding the contract method 0x60b810f1.
//
// Solidity: function swapMappingETH2BSC(address ) view returns(address)
func (_BSCSwapAgent *BSCSwapAgentCallerSession) SwapMappingETH2BSC(arg0 common.Address) (common.Address, error) {
	return _BSCSwapAgent.Contract.SwapMappingETH2BSC(&_BSCSwapAgent.CallOpts, arg0)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x32bd6e31.
//
// Solidity: function createSwapPair(bytes32 ethTxHash, address erc20Addr, string name, string symbol, uint8 decimals) returns(address)
func (_BSCSwapAgent *BSCSwapAgentTransactor) CreateSwapPair(opts *bind.TransactOpts, ethTxHash [32]byte, erc20Addr common.Address, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BSCSwapAgent.contract.Transact(opts, "createSwapPair", ethTxHash, erc20Addr, name, symbol, decimals)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x32bd6e31.
//
// Solidity: function createSwapPair(bytes32 ethTxHash, address erc20Addr, string name, string symbol, uint8 decimals) returns(address)
func (_BSCSwapAgent *BSCSwapAgentSession) CreateSwapPair(ethTxHash [32]byte, erc20Addr common.Address, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.CreateSwapPair(&_BSCSwapAgent.TransactOpts, ethTxHash, erc20Addr, name, symbol, decimals)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x32bd6e31.
//
// Solidity: function createSwapPair(bytes32 ethTxHash, address erc20Addr, string name, string symbol, uint8 decimals) returns(address)
func (_BSCSwapAgent *BSCSwapAgentTransactorSession) CreateSwapPair(ethTxHash [32]byte, erc20Addr common.Address, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.CreateSwapPair(&_BSCSwapAgent.TransactOpts, ethTxHash, erc20Addr, name, symbol, decimals)
}

// FillETH2BSCSwap is a paid mutator transaction binding the contract method 0xe307b931.
//
// Solidity: function fillETH2BSCSwap(bytes32 ethTxHash, address erc20Addr, address toAddress, uint256 amount) returns(bool)
func (_BSCSwapAgent *BSCSwapAgentTransactor) FillETH2BSCSwap(opts *bind.TransactOpts, ethTxHash [32]byte, erc20Addr common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BSCSwapAgent.contract.Transact(opts, "fillETH2BSCSwap", ethTxHash, erc20Addr, toAddress, amount)
}

// FillETH2BSCSwap is a paid mutator transaction binding the contract method 0xe307b931.
//
// Solidity: function fillETH2BSCSwap(bytes32 ethTxHash, address erc20Addr, address toAddress, uint256 amount) returns(bool)
func (_BSCSwapAgent *BSCSwapAgentSession) FillETH2BSCSwap(ethTxHash [32]byte, erc20Addr common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.FillETH2BSCSwap(&_BSCSwapAgent.TransactOpts, ethTxHash, erc20Addr, toAddress, amount)
}

// FillETH2BSCSwap is a paid mutator transaction binding the contract method 0xe307b931.
//
// Solidity: function fillETH2BSCSwap(bytes32 ethTxHash, address erc20Addr, address toAddress, uint256 amount) returns(bool)
func (_BSCSwapAgent *BSCSwapAgentTransactorSession) FillETH2BSCSwap(ethTxHash [32]byte, erc20Addr common.Address, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.FillETH2BSCSwap(&_BSCSwapAgent.TransactOpts, ethTxHash, erc20Addr, toAddress, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x358394d8.
//
// Solidity: function initialize(address bep20Impl, uint256 fee, address ownerAddr, address bep20ProxyAdminAddr) returns()
func (_BSCSwapAgent *BSCSwapAgentTransactor) Initialize(opts *bind.TransactOpts, bep20Impl common.Address, fee *big.Int, ownerAddr common.Address, bep20ProxyAdminAddr common.Address) (*types.Transaction, error) {
	return _BSCSwapAgent.contract.Transact(opts, "initialize", bep20Impl, fee, ownerAddr, bep20ProxyAdminAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x358394d8.
//
// Solidity: function initialize(address bep20Impl, uint256 fee, address ownerAddr, address bep20ProxyAdminAddr) returns()
func (_BSCSwapAgent *BSCSwapAgentSession) Initialize(bep20Impl common.Address, fee *big.Int, ownerAddr common.Address, bep20ProxyAdminAddr common.Address) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.Initialize(&_BSCSwapAgent.TransactOpts, bep20Impl, fee, ownerAddr, bep20ProxyAdminAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x358394d8.
//
// Solidity: function initialize(address bep20Impl, uint256 fee, address ownerAddr, address bep20ProxyAdminAddr) returns()
func (_BSCSwapAgent *BSCSwapAgentTransactorSession) Initialize(bep20Impl common.Address, fee *big.Int, ownerAddr common.Address, bep20ProxyAdminAddr common.Address) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.Initialize(&_BSCSwapAgent.TransactOpts, bep20Impl, fee, ownerAddr, bep20ProxyAdminAddr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BSCSwapAgent *BSCSwapAgentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BSCSwapAgent.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BSCSwapAgent *BSCSwapAgentSession) RenounceOwnership() (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.RenounceOwnership(&_BSCSwapAgent.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BSCSwapAgent *BSCSwapAgentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.RenounceOwnership(&_BSCSwapAgent.TransactOpts)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 fee) returns()
func (_BSCSwapAgent *BSCSwapAgentTransactor) SetSwapFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _BSCSwapAgent.contract.Transact(opts, "setSwapFee", fee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 fee) returns()
func (_BSCSwapAgent *BSCSwapAgentSession) SetSwapFee(fee *big.Int) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.SetSwapFee(&_BSCSwapAgent.TransactOpts, fee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 fee) returns()
func (_BSCSwapAgent *BSCSwapAgentTransactorSession) SetSwapFee(fee *big.Int) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.SetSwapFee(&_BSCSwapAgent.TransactOpts, fee)
}

// SwapBSC2ETH is a paid mutator transaction binding the contract method 0x1ba3b150.
//
// Solidity: function swapBSC2ETH(address bep20Addr, uint256 amount) payable returns(bool)
func (_BSCSwapAgent *BSCSwapAgentTransactor) SwapBSC2ETH(opts *bind.TransactOpts, bep20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BSCSwapAgent.contract.Transact(opts, "swapBSC2ETH", bep20Addr, amount)
}

// SwapBSC2ETH is a paid mutator transaction binding the contract method 0x1ba3b150.
//
// Solidity: function swapBSC2ETH(address bep20Addr, uint256 amount) payable returns(bool)
func (_BSCSwapAgent *BSCSwapAgentSession) SwapBSC2ETH(bep20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.SwapBSC2ETH(&_BSCSwapAgent.TransactOpts, bep20Addr, amount)
}

// SwapBSC2ETH is a paid mutator transaction binding the contract method 0x1ba3b150.
//
// Solidity: function swapBSC2ETH(address bep20Addr, uint256 amount) payable returns(bool)
func (_BSCSwapAgent *BSCSwapAgentTransactorSession) SwapBSC2ETH(bep20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.SwapBSC2ETH(&_BSCSwapAgent.TransactOpts, bep20Addr, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BSCSwapAgent *BSCSwapAgentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BSCSwapAgent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BSCSwapAgent *BSCSwapAgentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.TransferOwnership(&_BSCSwapAgent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BSCSwapAgent *BSCSwapAgentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BSCSwapAgent.Contract.TransferOwnership(&_BSCSwapAgent.TransactOpts, newOwner)
}

// BSCSwapAgentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BSCSwapAgent contract.
type BSCSwapAgentOwnershipTransferredIterator struct {
	Event *BSCSwapAgentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BSCSwapAgentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCSwapAgentOwnershipTransferred)
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
		it.Event = new(BSCSwapAgentOwnershipTransferred)
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
func (it *BSCSwapAgentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCSwapAgentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCSwapAgentOwnershipTransferred represents a OwnershipTransferred event raised by the BSCSwapAgent contract.
type BSCSwapAgentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BSCSwapAgent *BSCSwapAgentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BSCSwapAgentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BSCSwapAgent.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BSCSwapAgentOwnershipTransferredIterator{contract: _BSCSwapAgent.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BSCSwapAgent *BSCSwapAgentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BSCSwapAgentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BSCSwapAgent.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCSwapAgentOwnershipTransferred)
				if err := _BSCSwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BSCSwapAgent *BSCSwapAgentFilterer) ParseOwnershipTransferred(log types.Log) (*BSCSwapAgentOwnershipTransferred, error) {
	event := new(BSCSwapAgentOwnershipTransferred)
	if err := _BSCSwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCSwapAgentSwapFilledIterator is returned from FilterSwapFilled and is used to iterate over the raw logs and unpacked data for SwapFilled events raised by the BSCSwapAgent contract.
type BSCSwapAgentSwapFilledIterator struct {
	Event *BSCSwapAgentSwapFilled // Event containing the contract specifics and raw log

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
func (it *BSCSwapAgentSwapFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCSwapAgentSwapFilled)
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
		it.Event = new(BSCSwapAgentSwapFilled)
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
func (it *BSCSwapAgentSwapFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCSwapAgentSwapFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCSwapAgentSwapFilled represents a SwapFilled event raised by the BSCSwapAgent contract.
type BSCSwapAgentSwapFilled struct {
	Bep20Addr common.Address
	EthTxHash [32]byte
	ToAddress common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapFilled is a free log retrieval operation binding the contract event 0x3bebd9a738291e69898b5dbfadb6329b4b09fc648bdef68762928e521463abd9.
//
// Solidity: event SwapFilled(address indexed bep20Addr, bytes32 indexed ethTxHash, address indexed toAddress, uint256 amount)
func (_BSCSwapAgent *BSCSwapAgentFilterer) FilterSwapFilled(opts *bind.FilterOpts, bep20Addr []common.Address, ethTxHash [][32]byte, toAddress []common.Address) (*BSCSwapAgentSwapFilledIterator, error) {

	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}
	var ethTxHashRule []interface{}
	for _, ethTxHashItem := range ethTxHash {
		ethTxHashRule = append(ethTxHashRule, ethTxHashItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _BSCSwapAgent.contract.FilterLogs(opts, "SwapFilled", bep20AddrRule, ethTxHashRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &BSCSwapAgentSwapFilledIterator{contract: _BSCSwapAgent.contract, event: "SwapFilled", logs: logs, sub: sub}, nil
}

// WatchSwapFilled is a free log subscription operation binding the contract event 0x3bebd9a738291e69898b5dbfadb6329b4b09fc648bdef68762928e521463abd9.
//
// Solidity: event SwapFilled(address indexed bep20Addr, bytes32 indexed ethTxHash, address indexed toAddress, uint256 amount)
func (_BSCSwapAgent *BSCSwapAgentFilterer) WatchSwapFilled(opts *bind.WatchOpts, sink chan<- *BSCSwapAgentSwapFilled, bep20Addr []common.Address, ethTxHash [][32]byte, toAddress []common.Address) (event.Subscription, error) {

	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}
	var ethTxHashRule []interface{}
	for _, ethTxHashItem := range ethTxHash {
		ethTxHashRule = append(ethTxHashRule, ethTxHashItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _BSCSwapAgent.contract.WatchLogs(opts, "SwapFilled", bep20AddrRule, ethTxHashRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCSwapAgentSwapFilled)
				if err := _BSCSwapAgent.contract.UnpackLog(event, "SwapFilled", log); err != nil {
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
// Solidity: event SwapFilled(address indexed bep20Addr, bytes32 indexed ethTxHash, address indexed toAddress, uint256 amount)
func (_BSCSwapAgent *BSCSwapAgentFilterer) ParseSwapFilled(log types.Log) (*BSCSwapAgentSwapFilled, error) {
	event := new(BSCSwapAgentSwapFilled)
	if err := _BSCSwapAgent.contract.UnpackLog(event, "SwapFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCSwapAgentSwapPairCreatedIterator is returned from FilterSwapPairCreated and is used to iterate over the raw logs and unpacked data for SwapPairCreated events raised by the BSCSwapAgent contract.
type BSCSwapAgentSwapPairCreatedIterator struct {
	Event *BSCSwapAgentSwapPairCreated // Event containing the contract specifics and raw log

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
func (it *BSCSwapAgentSwapPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCSwapAgentSwapPairCreated)
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
		it.Event = new(BSCSwapAgentSwapPairCreated)
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
func (it *BSCSwapAgentSwapPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCSwapAgentSwapPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCSwapAgentSwapPairCreated represents a SwapPairCreated event raised by the BSCSwapAgent contract.
type BSCSwapAgentSwapPairCreated struct {
	EthRegisterTxHash [32]byte
	Bep20Addr         common.Address
	Erc20Addr         common.Address
	Symbol            string
	Name              string
	Decimals          uint8
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapPairCreated is a free log retrieval operation binding the contract event 0xcc0314763eabceb74cd3d30ae785c09bfe4e204af2088b3bfcdbbe5082133db5.
//
// Solidity: event SwapPairCreated(bytes32 indexed ethRegisterTxHash, address indexed bep20Addr, address indexed erc20Addr, string symbol, string name, uint8 decimals)
func (_BSCSwapAgent *BSCSwapAgentFilterer) FilterSwapPairCreated(opts *bind.FilterOpts, ethRegisterTxHash [][32]byte, bep20Addr []common.Address, erc20Addr []common.Address) (*BSCSwapAgentSwapPairCreatedIterator, error) {

	var ethRegisterTxHashRule []interface{}
	for _, ethRegisterTxHashItem := range ethRegisterTxHash {
		ethRegisterTxHashRule = append(ethRegisterTxHashRule, ethRegisterTxHashItem)
	}
	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}
	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}

	logs, sub, err := _BSCSwapAgent.contract.FilterLogs(opts, "SwapPairCreated", ethRegisterTxHashRule, bep20AddrRule, erc20AddrRule)
	if err != nil {
		return nil, err
	}
	return &BSCSwapAgentSwapPairCreatedIterator{contract: _BSCSwapAgent.contract, event: "SwapPairCreated", logs: logs, sub: sub}, nil
}

// WatchSwapPairCreated is a free log subscription operation binding the contract event 0xcc0314763eabceb74cd3d30ae785c09bfe4e204af2088b3bfcdbbe5082133db5.
//
// Solidity: event SwapPairCreated(bytes32 indexed ethRegisterTxHash, address indexed bep20Addr, address indexed erc20Addr, string symbol, string name, uint8 decimals)
func (_BSCSwapAgent *BSCSwapAgentFilterer) WatchSwapPairCreated(opts *bind.WatchOpts, sink chan<- *BSCSwapAgentSwapPairCreated, ethRegisterTxHash [][32]byte, bep20Addr []common.Address, erc20Addr []common.Address) (event.Subscription, error) {

	var ethRegisterTxHashRule []interface{}
	for _, ethRegisterTxHashItem := range ethRegisterTxHash {
		ethRegisterTxHashRule = append(ethRegisterTxHashRule, ethRegisterTxHashItem)
	}
	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}
	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}

	logs, sub, err := _BSCSwapAgent.contract.WatchLogs(opts, "SwapPairCreated", ethRegisterTxHashRule, bep20AddrRule, erc20AddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCSwapAgentSwapPairCreated)
				if err := _BSCSwapAgent.contract.UnpackLog(event, "SwapPairCreated", log); err != nil {
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

// ParseSwapPairCreated is a log parse operation binding the contract event 0xcc0314763eabceb74cd3d30ae785c09bfe4e204af2088b3bfcdbbe5082133db5.
//
// Solidity: event SwapPairCreated(bytes32 indexed ethRegisterTxHash, address indexed bep20Addr, address indexed erc20Addr, string symbol, string name, uint8 decimals)
func (_BSCSwapAgent *BSCSwapAgentFilterer) ParseSwapPairCreated(log types.Log) (*BSCSwapAgentSwapPairCreated, error) {
	event := new(BSCSwapAgentSwapPairCreated)
	if err := _BSCSwapAgent.contract.UnpackLog(event, "SwapPairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BSCSwapAgentSwapStartedIterator is returned from FilterSwapStarted and is used to iterate over the raw logs and unpacked data for SwapStarted events raised by the BSCSwapAgent contract.
type BSCSwapAgentSwapStartedIterator struct {
	Event *BSCSwapAgentSwapStarted // Event containing the contract specifics and raw log

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
func (it *BSCSwapAgentSwapStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BSCSwapAgentSwapStarted)
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
		it.Event = new(BSCSwapAgentSwapStarted)
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
func (it *BSCSwapAgentSwapStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BSCSwapAgentSwapStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BSCSwapAgentSwapStarted represents a SwapStarted event raised by the BSCSwapAgent contract.
type BSCSwapAgentSwapStarted struct {
	Bep20Addr common.Address
	Erc20Addr common.Address
	FromAddr  common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapStarted is a free log retrieval operation binding the contract event 0x49c08ff11118922c1e8298915531eff9ef6f8b39b44b3e9952b75d47e1d0cdd0.
//
// Solidity: event SwapStarted(address indexed bep20Addr, address indexed erc20Addr, address indexed fromAddr, uint256 amount, uint256 feeAmount)
func (_BSCSwapAgent *BSCSwapAgentFilterer) FilterSwapStarted(opts *bind.FilterOpts, bep20Addr []common.Address, erc20Addr []common.Address, fromAddr []common.Address) (*BSCSwapAgentSwapStartedIterator, error) {

	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}
	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}
	var fromAddrRule []interface{}
	for _, fromAddrItem := range fromAddr {
		fromAddrRule = append(fromAddrRule, fromAddrItem)
	}

	logs, sub, err := _BSCSwapAgent.contract.FilterLogs(opts, "SwapStarted", bep20AddrRule, erc20AddrRule, fromAddrRule)
	if err != nil {
		return nil, err
	}
	return &BSCSwapAgentSwapStartedIterator{contract: _BSCSwapAgent.contract, event: "SwapStarted", logs: logs, sub: sub}, nil
}

// WatchSwapStarted is a free log subscription operation binding the contract event 0x49c08ff11118922c1e8298915531eff9ef6f8b39b44b3e9952b75d47e1d0cdd0.
//
// Solidity: event SwapStarted(address indexed bep20Addr, address indexed erc20Addr, address indexed fromAddr, uint256 amount, uint256 feeAmount)
func (_BSCSwapAgent *BSCSwapAgentFilterer) WatchSwapStarted(opts *bind.WatchOpts, sink chan<- *BSCSwapAgentSwapStarted, bep20Addr []common.Address, erc20Addr []common.Address, fromAddr []common.Address) (event.Subscription, error) {

	var bep20AddrRule []interface{}
	for _, bep20AddrItem := range bep20Addr {
		bep20AddrRule = append(bep20AddrRule, bep20AddrItem)
	}
	var erc20AddrRule []interface{}
	for _, erc20AddrItem := range erc20Addr {
		erc20AddrRule = append(erc20AddrRule, erc20AddrItem)
	}
	var fromAddrRule []interface{}
	for _, fromAddrItem := range fromAddr {
		fromAddrRule = append(fromAddrRule, fromAddrItem)
	}

	logs, sub, err := _BSCSwapAgent.contract.WatchLogs(opts, "SwapStarted", bep20AddrRule, erc20AddrRule, fromAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BSCSwapAgentSwapStarted)
				if err := _BSCSwapAgent.contract.UnpackLog(event, "SwapStarted", log); err != nil {
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

// ParseSwapStarted is a log parse operation binding the contract event 0x49c08ff11118922c1e8298915531eff9ef6f8b39b44b3e9952b75d47e1d0cdd0.
//
// Solidity: event SwapStarted(address indexed bep20Addr, address indexed erc20Addr, address indexed fromAddr, uint256 amount, uint256 feeAmount)
func (_BSCSwapAgent *BSCSwapAgentFilterer) ParseSwapStarted(log types.Log) (*BSCSwapAgentSwapStarted, error) {
	event := new(BSCSwapAgentSwapStarted)
	if err := _BSCSwapAgent.contract.UnpackLog(event, "SwapStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
