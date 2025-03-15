// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// TreasureManagerMetaData contains all meta data concerning the TreasureManager contract.
var TreasureManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimAllToken\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositERC20\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositETH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ethAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenWhiteList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantReward\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"granter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"granterRewardAmount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_treasureManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_withdrawManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryReward\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenWhiteList\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWithdrawManager\",\"inputs\":[{\"name\":\"_withdrawManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenWhiteList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"treasureManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawERC20\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawETH\",\"inputs\":[{\"name\":\"withdrawAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GrantRewardTokenAmount\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"granter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawManagerUpdate\",\"inputs\":[{\"name\":\"withdrawManger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f80fd5b506130408061001c5f395ff3fe6080604052600436106101ba575f3560e01c8063519acebf116100eb578063a2672ad711610089578063d547741f11610063578063d547741f14610681578063ec3e9da5146106a9578063f2fde38b146106d3578063f6326fb3146106fb5761022f565b8063a2672ad7146105f5578063c0c53b8b1461061d578063c37220ea146106455761022f565b80638da5cb5b116100c55780638da5cb5b1461053d57806391d148541461056757806397feb926146105a3578063a217fddf146105cb5761022f565b8063519acebf146104af578063523fba7f146104eb578063715018a6146105275761022f565b8063248a9ca31161015857806336568abe1161013257806336568abe146103f157806341398b151461041957806344004cc1146104435780634782f7791461047f5761022f565b8063248a9ca3146103655780632f2ff15d146103a157806332f289cf146103c95761022f565b8063084cfb2711610194578063084cfb27146102c15780630ec08b00146102d757806317e3e2e81461031357806323ecdbee1461033b5761022f565b806301ffc9a71461023357806302fef78b1461026f57806303186d22146102975761022f565b3661022f573073ffffffffffffffffffffffffffffffffffffffff1663f6326fb36040518163ffffffff1660e01b81526004016020604051808303815f875af1158015610209573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061022d919061248e565b005b5f80fd5b34801561023e575f80fd5b506102596004803603810190610254919061250e565b610719565b6040516102669190612548565b60405180910390f35b34801561027a575f80fd5b50610295600480360381019061029091906125ee565b610792565b005b3480156102a2575f80fd5b506102ab610998565b6040516102b8919061264d565b60405180910390f35b3480156102cc575f80fd5b506102d56109bb565b005b3480156102e2575f80fd5b506102fd60048036038101906102f89190612666565b610d60565b60405161030a919061264d565b60405180910390f35b34801561031e575f80fd5b5061033960048036038101906103349190612691565b610d9b565b005b348015610346575f80fd5b5061034f610e21565b60405161035c9190612773565b60405180910390f35b348015610370575f80fd5b5061038b600480360381019061038691906127c6565b610eac565b6040516103989190612800565b60405180910390f35b3480156103ac575f80fd5b506103c760048036038101906103c29190612819565b610ed6565b005b3480156103d4575f80fd5b506103ef60048036038101906103ea9190612691565b610ef8565b005b3480156103fc575f80fd5b5061041760048036038101906104129190612819565b6112fe565b005b348015610424575f80fd5b5061042d611379565b60405161043a919061264d565b60405180910390f35b34801561044e575f80fd5b5061046960048036038101906104649190612892565b611391565b6040516104769190612548565b60405180910390f35b6104996004803603810190610494919061291d565b6115e2565b6040516104a69190612548565b60405180910390f35b3480156104ba575f80fd5b506104d560048036038101906104d0919061295b565b61186a565b6040516104e291906129a8565b60405180910390f35b3480156104f6575f80fd5b50610511600480360381019061050c9190612691565b61188a565b60405161051e91906129a8565b60405180910390f35b348015610532575f80fd5b5061053b61189f565b005b348015610548575f80fd5b506105516118b2565b60405161055e919061264d565b60405180910390f35b348015610572575f80fd5b5061058d60048036038101906105889190612819565b6118e7565b60405161059a9190612548565b60405180910390f35b3480156105ae575f80fd5b506105c960048036038101906105c491906129c1565b611958565b005b3480156105d6575f80fd5b506105df611a92565b6040516105ec9190612800565b60405180910390f35b348015610600575f80fd5b5061061b60048036038101906106169190612691565b611a98565b005b348015610628575f80fd5b50610643600480360381019061063e91906129ff565b611afb565b005b348015610650575f80fd5b5061066b60048036038101906106669190612691565b611cfd565b60405161067891906129a8565b60405180910390f35b34801561068c575f80fd5b506106a760048036038101906106a29190612819565b611d7e565b005b3480156106b4575f80fd5b506106bd611da0565b6040516106ca919061264d565b60405180910390f35b3480156106de575f80fd5b506106f960048036038101906106f49190612691565b611dc5565b005b610703611e49565b6040516107109190612548565b60405180910390f35b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061078b575061078a82611f31565b5b9050919050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461081f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081690612aa9565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561088757505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b6108c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108bd90612b11565b60405180910390fd5b8060045f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff167f10621458f3ad2a9cfcb87c216122570629e44079d6af4d717035eb55d2c60424838360405161098b929190612b2f565b60405180910390a2505050565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6109c3611f9a565b5f5b600280549050811015610d55575f600282815481106109e7576109e6612b56565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7690612b11565b60405180910390fd5b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f811115610d465773eeeeeeeeeeeeeeeceeeeeeeeeeceeeeeeceeeece73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610bf8575f3373ffffffffffffffffffffffffffffffffffffffff1682604051610b6f90612bb0565b5f6040518083038185875af1925050503d805f8114610ba9576040519150601f19603f3d011682016040523d82523d5f602084013e610bae565b606091505b5050905080610bf2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be990612c0e565b60405180910390fd5b50610c75565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401610c33929190612b2f565b6020604051808303815f875af1158015610c4f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c73919061248e565b505b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610d3e9190612c59565b925050819055505b505080806001019150506109c5565b50610d5e611fee565b565b60028181548110610d6f575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f799e16a314d482c87bc47fd219011aaf33f4f9c7e302be5d7a0af286a52998b960405160405180910390a250565b60606002805480602002602001604051908101604052809291908181526020018280548015610ea257602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610e59575b5050505050905090565b5f80610eb6612005565b9050805f015f8481526020019081526020015f2060010154915050919050565b610edf82610eac565b610ee88161202c565b610ef28383612040565b50505050565b610f00611f9a565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610f6e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6590612b11565b60405180910390fd5b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f811161102c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161102390612cd6565b60405180910390fd5b5f60035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f81116110af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110a690612d3e565b60405180910390fd5b73eeeeeeeeeeeeeeeceeeeeeeeeeceeeeeeceeeece73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036111a4575f3373ffffffffffffffffffffffffffffffffffffffff168360405161111b90612bb0565b5f6040518083038185875af1925050503d805f8114611155576040519150601f19603f3d011682016040523d82523d5f602084013e61115a565b606091505b505090508061119e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161119590612c0e565b60405180910390fd5b50611221565b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b81526004016111df929190612b2f565b6020604051808303815f875af11580156111fb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061121f919061248e565b505b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508160035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546112ea9190612c59565b9250508190555050506112fb611fee565b50565b611306612138565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461136a576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611374828261213f565b505050565b73eeeeeeeeeeeeeeeceeeeeeeeeeceeeeeeceeeece81565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611421576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141890612da6565b60405180910390fd5b8160035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156114a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161149890612e0e565b60405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b81526004016114dc929190612b2f565b6020604051808303815f875af11580156114f8573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061151c919061248e565b508160035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546115699190612c59565b925050819055503373ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b74685856040516115cf929190612b2f565b60405180910390a3600190509392505050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611672576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161166990612da6565b60405180910390fd5b8160035f73eeeeeeeeeeeeeeeceeeeeeeeeeceeeeeeceeeece73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611706576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116fd90612e76565b60405180910390fd5b5f8373ffffffffffffffffffffffffffffffffffffffff168360405161172b90612bb0565b5f6040518083038185875af1925050503d805f8114611765576040519150601f19603f3d011682016040523d82523d5f602084013e61176a565b606091505b505090508061177c575f915050611864565b8260035f73eeeeeeeeeeeeeeeceeeeeeeeeeceeeeeeceeeece73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546117dc9190612c59565b925050819055503373ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeceeeeeeeeeeceeeeeeceeeece73ffffffffffffffffffffffffffffffffffffffff167f9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b7468686604051611856929190612eef565b60405180910390a360019150505b92915050565b6004602052815f5260405f20602052805f5260405f205f91509150505481565b6003602052805f5260405f205f915090505481565b6118a7612237565b6118b05f6122be565b565b5f806118bc61238f565b9050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691505090565b5f806118f1612005565b9050805f015f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1691505092915050565b8173ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b815260040161199593929190612f16565b6020604051808303815f875af11580156119b1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119d5919061248e565b508060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254611a229190612f4b565b925050819055503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd5783604051611a8691906129a8565b60405180910390a35050565b5f801b81565b600281908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f611b046123b6565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f808267ffffffffffffffff16148015611b4c5750825b90505f60018367ffffffffffffffff16148015611b7f57505f3073ffffffffffffffffffffffffffffffffffffffff163b145b905081158015611b8d575080155b15611bc4576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508315611c11576001855f0160086101000a81548160ff0219169083151502179055505b865f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508560015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611c99886122be565b8315611cf3575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d26001604051611cea9190612fca565b60405180910390a15b5050505050505050565b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b611d8782610eac565b611d908161202c565b611d9a838361213f565b50505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611dcd612237565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611e3d575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401611e34919061264d565b60405180910390fd5b611e46816122be565b50565b5f3460035f73eeeeeeeeeeeeeeeceeeeeeeeeeceeeeeeceeeece73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254611eaa9190612f4b565b925050819055503373ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeceeeeeeeeeeceeeeeeceeeece73ffffffffffffffffffffffffffffffffffffffff167f4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd5734604051611f2291906129a8565b60405180910390a36001905090565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f611fa36123dd565b90506002815f015403611fe2576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002815f018190555050565b5f611ff76123dd565b90506001815f018190555050565b5f7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800905090565b61203d81612038612138565b612404565b50565b5f8061204a612005565b905061205684846118e7565b61212d576001815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506120c9612138565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050612132565b5f9150505b92915050565b5f33905090565b5f80612149612005565b905061215584846118e7565b1561222c575f815f015f8681526020019081526020015f205f015f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506121c8612138565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16857ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a46001915050612231565b5f9150505b92915050565b61223f612138565b73ffffffffffffffffffffffffffffffffffffffff1661225d6118b2565b73ffffffffffffffffffffffffffffffffffffffff16146122bc57612280612138565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016122b3919061264d565b60405180910390fd5b565b5f6122c761238f565b90505f815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082825f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505050565b5f7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300905090565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00905090565b61240e82826118e7565b6124515780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401612448929190612fe3565b60405180910390fd5b5050565b5f80fd5b5f8115159050919050565b61246d81612459565b8114612477575f80fd5b50565b5f8151905061248881612464565b92915050565b5f602082840312156124a3576124a2612455565b5b5f6124b08482850161247a565b91505092915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6124ed816124b9565b81146124f7575f80fd5b50565b5f81359050612508816124e4565b92915050565b5f6020828403121561252357612522612455565b5b5f612530848285016124fa565b91505092915050565b61254281612459565b82525050565b5f60208201905061255b5f830184612539565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61258a82612561565b9050919050565b61259a81612580565b81146125a4575f80fd5b50565b5f813590506125b581612591565b92915050565b5f819050919050565b6125cd816125bb565b81146125d7575f80fd5b50565b5f813590506125e8816125c4565b92915050565b5f805f6060848603121561260557612604612455565b5b5f612612868287016125a7565b9350506020612623868287016125a7565b9250506040612634868287016125da565b9150509250925092565b61264781612580565b82525050565b5f6020820190506126605f83018461263e565b92915050565b5f6020828403121561267b5761267a612455565b5b5f612688848285016125da565b91505092915050565b5f602082840312156126a6576126a5612455565b5b5f6126b3848285016125a7565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6126ee81612580565b82525050565b5f6126ff83836126e5565b60208301905092915050565b5f602082019050919050565b5f612721826126bc565b61272b81856126c6565b9350612736836126d6565b805f5b8381101561276657815161274d88826126f4565b97506127588361270b565b925050600181019050612739565b5085935050505092915050565b5f6020820190508181035f83015261278b8184612717565b905092915050565b5f819050919050565b6127a581612793565b81146127af575f80fd5b50565b5f813590506127c08161279c565b92915050565b5f602082840312156127db576127da612455565b5b5f6127e8848285016127b2565b91505092915050565b6127fa81612793565b82525050565b5f6020820190506128135f8301846127f1565b92915050565b5f806040838503121561282f5761282e612455565b5b5f61283c858286016127b2565b925050602061284d858286016125a7565b9150509250929050565b5f61286182612580565b9050919050565b61287181612857565b811461287b575f80fd5b50565b5f8135905061288c81612868565b92915050565b5f805f606084860312156128a9576128a8612455565b5b5f6128b68682870161287e565b93505060206128c7868287016125a7565b92505060406128d8868287016125da565b9150509250925092565b5f6128ec82612561565b9050919050565b6128fc816128e2565b8114612906575f80fd5b50565b5f81359050612917816128f3565b92915050565b5f806040838503121561293357612932612455565b5b5f61294085828601612909565b9250506020612951858286016125da565b9150509250929050565b5f806040838503121561297157612970612455565b5b5f61297e858286016125a7565b925050602061298f858286016125a7565b9150509250929050565b6129a2816125bb565b82525050565b5f6020820190506129bb5f830184612999565b92915050565b5f80604083850312156129d7576129d6612455565b5b5f6129e48582860161287e565b92505060206129f5858286016125da565b9150509250929050565b5f805f60608486031215612a1657612a15612455565b5b5f612a23868287016125a7565b9350506020612a34868287016125a7565b9250506040612a45868287016125a7565b9150509250925092565b5f82825260208201905092915050565b7f6f6e6c792074726561737572654d616e616765722063616e2063616c6c0000005f82015250565b5f612a93601d83612a4f565b9150612a9e82612a5f565b602082019050919050565b5f6020820190508181035f830152612ac081612a87565b9050919050565b7f696e76616c6964206164647265737300000000000000000000000000000000005f82015250565b5f612afb600f83612a4f565b9150612b0682612ac7565b602082019050919050565b5f6020820190508181035f830152612b2881612aef565b9050919050565b5f604082019050612b425f83018561263e565b612b4f6020830184612999565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81905092915050565b50565b5f612b9b5f83612b83565b9150612ba682612b8d565b5f82019050919050565b5f612bba82612b90565b9150819050919050565b7f657468207472616e73666572206661696c6564000000000000000000000000005f82015250565b5f612bf8601383612a4f565b9150612c0382612bc4565b602082019050919050565b5f6020820190508181035f830152612c2581612bec565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f612c63826125bb565b9150612c6e836125bb565b9250828203905081811115612c8657612c85612c2c565b5b92915050565b7f6e6f2072657761726420617661696c61626c65000000000000000000000000005f82015250565b5f612cc0601383612a4f565b9150612ccb82612c8c565b602082019050919050565b5f6020820190508181035f830152612ced81612cb4565b9050919050565b7f746f6b656e20616d6f756e74206e6f7420656e6f7567680000000000000000005f82015250565b5f612d28601783612a4f565b9150612d3382612cf4565b602082019050919050565b5f6020820190508181035f830152612d5581612d1c565b9050919050565b7f6f6e6c792077697468647261774d616e616765722063616e2063616c6c0000005f82015250565b5f612d90601d83612a4f565b9150612d9b82612d5c565b602082019050919050565b5f6020820190508181035f830152612dbd81612d84565b9050919050565b7f746f6b656e206e6f7420656e6f756768000000000000000000000000000000005f82015250565b5f612df8601083612a4f565b9150612e0382612dc4565b602082019050919050565b5f6020820190508181035f830152612e2581612dec565b9050919050565b7f657468206e6f7420656e6f7567680000000000000000000000000000000000005f82015250565b5f612e60600e83612a4f565b9150612e6b82612e2c565b602082019050919050565b5f6020820190508181035f830152612e8d81612e54565b9050919050565b5f819050919050565b5f612eb7612eb2612ead84612561565b612e94565b612561565b9050919050565b5f612ec882612e9d565b9050919050565b5f612ed982612ebe565b9050919050565b612ee981612ecf565b82525050565b5f604082019050612f025f830185612ee0565b612f0f6020830184612999565b9392505050565b5f606082019050612f295f83018661263e565b612f36602083018561263e565b612f436040830184612999565b949350505050565b5f612f55826125bb565b9150612f60836125bb565b9250828201905080821115612f7857612f77612c2c565b5b92915050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f612fb4612faf612faa84612f7e565b612e94565b612f87565b9050919050565b612fc481612f9a565b82525050565b5f602082019050612fdd5f830184612fbb565b92915050565b5f604082019050612ff65f83018561263e565b61300360208301846127f1565b939250505056fea2646970667358221220ae11213193380bb19852756da891e8203d30cc478a7525877ef0dd5f2b267d6764736f6c634300081a0033",
}

// TreasureManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TreasureManagerMetaData.ABI instead.
var TreasureManagerABI = TreasureManagerMetaData.ABI

// TreasureManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TreasureManagerMetaData.Bin instead.
var TreasureManagerBin = TreasureManagerMetaData.Bin

// DeployTreasureManager deploys a new Ethereum contract, binding an instance of TreasureManager to it.
func DeployTreasureManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TreasureManager, error) {
	parsed, err := TreasureManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TreasureManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TreasureManager{TreasureManagerCaller: TreasureManagerCaller{contract: contract}, TreasureManagerTransactor: TreasureManagerTransactor{contract: contract}, TreasureManagerFilterer: TreasureManagerFilterer{contract: contract}}, nil
}

// TreasureManager is an auto generated Go binding around an Ethereum contract.
type TreasureManager struct {
	TreasureManagerCaller     // Read-only binding to the contract
	TreasureManagerTransactor // Write-only binding to the contract
	TreasureManagerFilterer   // Log filterer for contract events
}

// TreasureManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TreasureManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasureManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TreasureManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasureManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TreasureManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasureManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TreasureManagerSession struct {
	Contract     *TreasureManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TreasureManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TreasureManagerCallerSession struct {
	Contract *TreasureManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TreasureManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TreasureManagerTransactorSession struct {
	Contract     *TreasureManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TreasureManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TreasureManagerRaw struct {
	Contract *TreasureManager // Generic contract binding to access the raw methods on
}

// TreasureManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TreasureManagerCallerRaw struct {
	Contract *TreasureManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TreasureManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TreasureManagerTransactorRaw struct {
	Contract *TreasureManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasureManager creates a new instance of TreasureManager, bound to a specific deployed contract.
func NewTreasureManager(address common.Address, backend bind.ContractBackend) (*TreasureManager, error) {
	contract, err := bindTreasureManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TreasureManager{TreasureManagerCaller: TreasureManagerCaller{contract: contract}, TreasureManagerTransactor: TreasureManagerTransactor{contract: contract}, TreasureManagerFilterer: TreasureManagerFilterer{contract: contract}}, nil
}

// NewTreasureManagerCaller creates a new read-only instance of TreasureManager, bound to a specific deployed contract.
func NewTreasureManagerCaller(address common.Address, caller bind.ContractCaller) (*TreasureManagerCaller, error) {
	contract, err := bindTreasureManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerCaller{contract: contract}, nil
}

// NewTreasureManagerTransactor creates a new write-only instance of TreasureManager, bound to a specific deployed contract.
func NewTreasureManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TreasureManagerTransactor, error) {
	contract, err := bindTreasureManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerTransactor{contract: contract}, nil
}

// NewTreasureManagerFilterer creates a new log filterer instance of TreasureManager, bound to a specific deployed contract.
func NewTreasureManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TreasureManagerFilterer, error) {
	contract, err := bindTreasureManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerFilterer{contract: contract}, nil
}

// bindTreasureManager binds a generic wrapper to an already deployed contract.
func bindTreasureManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TreasureManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasureManager *TreasureManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TreasureManager.Contract.TreasureManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasureManager *TreasureManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.Contract.TreasureManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasureManager *TreasureManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasureManager.Contract.TreasureManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TreasureManager *TreasureManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TreasureManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TreasureManager *TreasureManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TreasureManager *TreasureManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TreasureManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TreasureManager *TreasureManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TreasureManager *TreasureManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TreasureManager.Contract.DEFAULTADMINROLE(&_TreasureManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TreasureManager *TreasureManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TreasureManager.Contract.DEFAULTADMINROLE(&_TreasureManager.CallOpts)
}

// EthAddress is a free data retrieval call binding the contract method 0x41398b15.
//
// Solidity: function ethAddress() view returns(address)
func (_TreasureManager *TreasureManagerCaller) EthAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "ethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthAddress is a free data retrieval call binding the contract method 0x41398b15.
//
// Solidity: function ethAddress() view returns(address)
func (_TreasureManager *TreasureManagerSession) EthAddress() (common.Address, error) {
	return _TreasureManager.Contract.EthAddress(&_TreasureManager.CallOpts)
}

// EthAddress is a free data retrieval call binding the contract method 0x41398b15.
//
// Solidity: function ethAddress() view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) EthAddress() (common.Address, error) {
	return _TreasureManager.Contract.EthAddress(&_TreasureManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TreasureManager *TreasureManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TreasureManager *TreasureManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TreasureManager.Contract.GetRoleAdmin(&_TreasureManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TreasureManager *TreasureManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TreasureManager.Contract.GetRoleAdmin(&_TreasureManager.CallOpts, role)
}

// GetTokenWhiteList is a free data retrieval call binding the contract method 0x23ecdbee.
//
// Solidity: function getTokenWhiteList() view returns(address[])
func (_TreasureManager *TreasureManagerCaller) GetTokenWhiteList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "getTokenWhiteList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokenWhiteList is a free data retrieval call binding the contract method 0x23ecdbee.
//
// Solidity: function getTokenWhiteList() view returns(address[])
func (_TreasureManager *TreasureManagerSession) GetTokenWhiteList() ([]common.Address, error) {
	return _TreasureManager.Contract.GetTokenWhiteList(&_TreasureManager.CallOpts)
}

// GetTokenWhiteList is a free data retrieval call binding the contract method 0x23ecdbee.
//
// Solidity: function getTokenWhiteList() view returns(address[])
func (_TreasureManager *TreasureManagerCallerSession) GetTokenWhiteList() ([]common.Address, error) {
	return _TreasureManager.Contract.GetTokenWhiteList(&_TreasureManager.CallOpts)
}

// GranterRewardAmount is a free data retrieval call binding the contract method 0x519acebf.
//
// Solidity: function granterRewardAmount(address , address ) view returns(uint256)
func (_TreasureManager *TreasureManagerCaller) GranterRewardAmount(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "granterRewardAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GranterRewardAmount is a free data retrieval call binding the contract method 0x519acebf.
//
// Solidity: function granterRewardAmount(address , address ) view returns(uint256)
func (_TreasureManager *TreasureManagerSession) GranterRewardAmount(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.GranterRewardAmount(&_TreasureManager.CallOpts, arg0, arg1)
}

// GranterRewardAmount is a free data retrieval call binding the contract method 0x519acebf.
//
// Solidity: function granterRewardAmount(address , address ) view returns(uint256)
func (_TreasureManager *TreasureManagerCallerSession) GranterRewardAmount(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.GranterRewardAmount(&_TreasureManager.CallOpts, arg0, arg1)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TreasureManager *TreasureManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TreasureManager *TreasureManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TreasureManager.Contract.HasRole(&_TreasureManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TreasureManager *TreasureManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TreasureManager.Contract.HasRole(&_TreasureManager.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TreasureManager *TreasureManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TreasureManager *TreasureManagerSession) Owner() (common.Address, error) {
	return _TreasureManager.Contract.Owner(&_TreasureManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) Owner() (common.Address, error) {
	return _TreasureManager.Contract.Owner(&_TreasureManager.CallOpts)
}

// QueryReward is a free data retrieval call binding the contract method 0xc37220ea.
//
// Solidity: function queryReward(address tokenAddress) view returns(uint256)
func (_TreasureManager *TreasureManagerCaller) QueryReward(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "queryReward", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryReward is a free data retrieval call binding the contract method 0xc37220ea.
//
// Solidity: function queryReward(address tokenAddress) view returns(uint256)
func (_TreasureManager *TreasureManagerSession) QueryReward(tokenAddress common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.QueryReward(&_TreasureManager.CallOpts, tokenAddress)
}

// QueryReward is a free data retrieval call binding the contract method 0xc37220ea.
//
// Solidity: function queryReward(address tokenAddress) view returns(uint256)
func (_TreasureManager *TreasureManagerCallerSession) QueryReward(tokenAddress common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.QueryReward(&_TreasureManager.CallOpts, tokenAddress)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TreasureManager *TreasureManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TreasureManager *TreasureManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TreasureManager.Contract.SupportsInterface(&_TreasureManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TreasureManager *TreasureManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TreasureManager.Contract.SupportsInterface(&_TreasureManager.CallOpts, interfaceId)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_TreasureManager *TreasureManagerCaller) TokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_TreasureManager *TreasureManagerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.TokenBalances(&_TreasureManager.CallOpts, arg0)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_TreasureManager *TreasureManagerCallerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _TreasureManager.Contract.TokenBalances(&_TreasureManager.CallOpts, arg0)
}

// TokenWhiteList is a free data retrieval call binding the contract method 0x0ec08b00.
//
// Solidity: function tokenWhiteList(uint256 ) view returns(address)
func (_TreasureManager *TreasureManagerCaller) TokenWhiteList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "tokenWhiteList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenWhiteList is a free data retrieval call binding the contract method 0x0ec08b00.
//
// Solidity: function tokenWhiteList(uint256 ) view returns(address)
func (_TreasureManager *TreasureManagerSession) TokenWhiteList(arg0 *big.Int) (common.Address, error) {
	return _TreasureManager.Contract.TokenWhiteList(&_TreasureManager.CallOpts, arg0)
}

// TokenWhiteList is a free data retrieval call binding the contract method 0x0ec08b00.
//
// Solidity: function tokenWhiteList(uint256 ) view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) TokenWhiteList(arg0 *big.Int) (common.Address, error) {
	return _TreasureManager.Contract.TokenWhiteList(&_TreasureManager.CallOpts, arg0)
}

// TreasureManager is a free data retrieval call binding the contract method 0x03186d22.
//
// Solidity: function treasureManager() view returns(address)
func (_TreasureManager *TreasureManagerCaller) TreasureManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "treasureManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasureManager is a free data retrieval call binding the contract method 0x03186d22.
//
// Solidity: function treasureManager() view returns(address)
func (_TreasureManager *TreasureManagerSession) TreasureManager() (common.Address, error) {
	return _TreasureManager.Contract.TreasureManager(&_TreasureManager.CallOpts)
}

// TreasureManager is a free data retrieval call binding the contract method 0x03186d22.
//
// Solidity: function treasureManager() view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) TreasureManager() (common.Address, error) {
	return _TreasureManager.Contract.TreasureManager(&_TreasureManager.CallOpts)
}

// WithdrawManager is a free data retrieval call binding the contract method 0xec3e9da5.
//
// Solidity: function withdrawManager() view returns(address)
func (_TreasureManager *TreasureManagerCaller) WithdrawManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TreasureManager.contract.Call(opts, &out, "withdrawManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawManager is a free data retrieval call binding the contract method 0xec3e9da5.
//
// Solidity: function withdrawManager() view returns(address)
func (_TreasureManager *TreasureManagerSession) WithdrawManager() (common.Address, error) {
	return _TreasureManager.Contract.WithdrawManager(&_TreasureManager.CallOpts)
}

// WithdrawManager is a free data retrieval call binding the contract method 0xec3e9da5.
//
// Solidity: function withdrawManager() view returns(address)
func (_TreasureManager *TreasureManagerCallerSession) WithdrawManager() (common.Address, error) {
	return _TreasureManager.Contract.WithdrawManager(&_TreasureManager.CallOpts)
}

// ClaimAllToken is a paid mutator transaction binding the contract method 0x084cfb27.
//
// Solidity: function claimAllToken() returns()
func (_TreasureManager *TreasureManagerTransactor) ClaimAllToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "claimAllToken")
}

// ClaimAllToken is a paid mutator transaction binding the contract method 0x084cfb27.
//
// Solidity: function claimAllToken() returns()
func (_TreasureManager *TreasureManagerSession) ClaimAllToken() (*types.Transaction, error) {
	return _TreasureManager.Contract.ClaimAllToken(&_TreasureManager.TransactOpts)
}

// ClaimAllToken is a paid mutator transaction binding the contract method 0x084cfb27.
//
// Solidity: function claimAllToken() returns()
func (_TreasureManager *TreasureManagerTransactorSession) ClaimAllToken() (*types.Transaction, error) {
	return _TreasureManager.Contract.ClaimAllToken(&_TreasureManager.TransactOpts)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x32f289cf.
//
// Solidity: function claimToken(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerTransactor) ClaimToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "claimToken", tokenAddress)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x32f289cf.
//
// Solidity: function claimToken(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerSession) ClaimToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.ClaimToken(&_TreasureManager.TransactOpts, tokenAddress)
}

// ClaimToken is a paid mutator transaction binding the contract method 0x32f289cf.
//
// Solidity: function claimToken(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerTransactorSession) ClaimToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.ClaimToken(&_TreasureManager.TransactOpts, tokenAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address tokenAddress, uint256 amount) returns()
func (_TreasureManager *TreasureManagerTransactor) DepositERC20(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "depositERC20", tokenAddress, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address tokenAddress, uint256 amount) returns()
func (_TreasureManager *TreasureManagerSession) DepositERC20(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.DepositERC20(&_TreasureManager.TransactOpts, tokenAddress, amount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x97feb926.
//
// Solidity: function depositERC20(address tokenAddress, uint256 amount) returns()
func (_TreasureManager *TreasureManagerTransactorSession) DepositERC20(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.DepositERC20(&_TreasureManager.TransactOpts, tokenAddress, amount)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns(bool)
func (_TreasureManager *TreasureManagerTransactor) DepositETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "depositETH")
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns(bool)
func (_TreasureManager *TreasureManagerSession) DepositETH() (*types.Transaction, error) {
	return _TreasureManager.Contract.DepositETH(&_TreasureManager.TransactOpts)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns(bool)
func (_TreasureManager *TreasureManagerTransactorSession) DepositETH() (*types.Transaction, error) {
	return _TreasureManager.Contract.DepositETH(&_TreasureManager.TransactOpts)
}

// GrantReward is a paid mutator transaction binding the contract method 0x02fef78b.
//
// Solidity: function grantReward(address tokenAddress, address granter, uint256 amount) returns()
func (_TreasureManager *TreasureManagerTransactor) GrantReward(opts *bind.TransactOpts, tokenAddress common.Address, granter common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "grantReward", tokenAddress, granter, amount)
}

// GrantReward is a paid mutator transaction binding the contract method 0x02fef78b.
//
// Solidity: function grantReward(address tokenAddress, address granter, uint256 amount) returns()
func (_TreasureManager *TreasureManagerSession) GrantReward(tokenAddress common.Address, granter common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.GrantReward(&_TreasureManager.TransactOpts, tokenAddress, granter, amount)
}

// GrantReward is a paid mutator transaction binding the contract method 0x02fef78b.
//
// Solidity: function grantReward(address tokenAddress, address granter, uint256 amount) returns()
func (_TreasureManager *TreasureManagerTransactorSession) GrantReward(tokenAddress common.Address, granter common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.GrantReward(&_TreasureManager.TransactOpts, tokenAddress, granter, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.GrantRole(&_TreasureManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.GrantRole(&_TreasureManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initOwner, address _treasureManager, address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerTransactor) Initialize(opts *bind.TransactOpts, initOwner common.Address, _treasureManager common.Address, _withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "initialize", initOwner, _treasureManager, _withdrawManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initOwner, address _treasureManager, address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerSession) Initialize(initOwner common.Address, _treasureManager common.Address, _withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.Initialize(&_TreasureManager.TransactOpts, initOwner, _treasureManager, _withdrawManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address initOwner, address _treasureManager, address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerTransactorSession) Initialize(initOwner common.Address, _treasureManager common.Address, _withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.Initialize(&_TreasureManager.TransactOpts, initOwner, _treasureManager, _withdrawManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TreasureManager *TreasureManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TreasureManager *TreasureManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _TreasureManager.Contract.RenounceOwnership(&_TreasureManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TreasureManager *TreasureManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TreasureManager.Contract.RenounceOwnership(&_TreasureManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TreasureManager *TreasureManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TreasureManager *TreasureManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.RenounceRole(&_TreasureManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TreasureManager *TreasureManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.RenounceRole(&_TreasureManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.RevokeRole(&_TreasureManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TreasureManager *TreasureManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.RevokeRole(&_TreasureManager.TransactOpts, role, account)
}

// SetTokenWhiteList is a paid mutator transaction binding the contract method 0xa2672ad7.
//
// Solidity: function setTokenWhiteList(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerTransactor) SetTokenWhiteList(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "setTokenWhiteList", tokenAddress)
}

// SetTokenWhiteList is a paid mutator transaction binding the contract method 0xa2672ad7.
//
// Solidity: function setTokenWhiteList(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerSession) SetTokenWhiteList(tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.SetTokenWhiteList(&_TreasureManager.TransactOpts, tokenAddress)
}

// SetTokenWhiteList is a paid mutator transaction binding the contract method 0xa2672ad7.
//
// Solidity: function setTokenWhiteList(address tokenAddress) returns()
func (_TreasureManager *TreasureManagerTransactorSession) SetTokenWhiteList(tokenAddress common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.SetTokenWhiteList(&_TreasureManager.TransactOpts, tokenAddress)
}

// SetWithdrawManager is a paid mutator transaction binding the contract method 0x17e3e2e8.
//
// Solidity: function setWithdrawManager(address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerTransactor) SetWithdrawManager(opts *bind.TransactOpts, _withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "setWithdrawManager", _withdrawManager)
}

// SetWithdrawManager is a paid mutator transaction binding the contract method 0x17e3e2e8.
//
// Solidity: function setWithdrawManager(address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerSession) SetWithdrawManager(_withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.SetWithdrawManager(&_TreasureManager.TransactOpts, _withdrawManager)
}

// SetWithdrawManager is a paid mutator transaction binding the contract method 0x17e3e2e8.
//
// Solidity: function setWithdrawManager(address _withdrawManager) returns()
func (_TreasureManager *TreasureManagerTransactorSession) SetWithdrawManager(_withdrawManager common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.SetWithdrawManager(&_TreasureManager.TransactOpts, _withdrawManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TreasureManager *TreasureManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TreasureManager *TreasureManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.TransferOwnership(&_TreasureManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TreasureManager *TreasureManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TreasureManager.Contract.TransferOwnership(&_TreasureManager.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address tokenAddress, address withdrawAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerTransactor) WithdrawERC20(opts *bind.TransactOpts, tokenAddress common.Address, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "withdrawERC20", tokenAddress, withdrawAddress, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address tokenAddress, address withdrawAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerSession) WithdrawERC20(tokenAddress common.Address, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.WithdrawERC20(&_TreasureManager.TransactOpts, tokenAddress, withdrawAddress, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x44004cc1.
//
// Solidity: function withdrawERC20(address tokenAddress, address withdrawAddress, uint256 amount) returns(bool)
func (_TreasureManager *TreasureManagerTransactorSession) WithdrawERC20(tokenAddress common.Address, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.WithdrawERC20(&_TreasureManager.TransactOpts, tokenAddress, withdrawAddress, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address withdrawAddress, uint256 amount) payable returns(bool)
func (_TreasureManager *TreasureManagerTransactor) WithdrawETH(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.contract.Transact(opts, "withdrawETH", withdrawAddress, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address withdrawAddress, uint256 amount) payable returns(bool)
func (_TreasureManager *TreasureManagerSession) WithdrawETH(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.WithdrawETH(&_TreasureManager.TransactOpts, withdrawAddress, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address withdrawAddress, uint256 amount) payable returns(bool)
func (_TreasureManager *TreasureManagerTransactorSession) WithdrawETH(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TreasureManager.Contract.WithdrawETH(&_TreasureManager.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TreasureManager *TreasureManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TreasureManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TreasureManager *TreasureManagerSession) Receive() (*types.Transaction, error) {
	return _TreasureManager.Contract.Receive(&_TreasureManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TreasureManager *TreasureManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _TreasureManager.Contract.Receive(&_TreasureManager.TransactOpts)
}

// TreasureManagerDepositTokenIterator is returned from FilterDepositToken and is used to iterate over the raw logs and unpacked data for DepositToken events raised by the TreasureManager contract.
type TreasureManagerDepositTokenIterator struct {
	Event *TreasureManagerDepositToken // Event containing the contract specifics and raw log

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
func (it *TreasureManagerDepositTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerDepositToken)
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
		it.Event = new(TreasureManagerDepositToken)
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
func (it *TreasureManagerDepositTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerDepositTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerDepositToken represents a DepositToken event raised by the TreasureManager contract.
type TreasureManagerDepositToken struct {
	TokenAddress common.Address
	Sender       common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositToken is a free log retrieval operation binding the contract event 0x4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57.
//
// Solidity: event DepositToken(address indexed tokenAddress, address indexed sender, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) FilterDepositToken(opts *bind.FilterOpts, tokenAddress []common.Address, sender []common.Address) (*TreasureManagerDepositTokenIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "DepositToken", tokenAddressRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerDepositTokenIterator{contract: _TreasureManager.contract, event: "DepositToken", logs: logs, sub: sub}, nil
}

// WatchDepositToken is a free log subscription operation binding the contract event 0x4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57.
//
// Solidity: event DepositToken(address indexed tokenAddress, address indexed sender, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) WatchDepositToken(opts *bind.WatchOpts, sink chan<- *TreasureManagerDepositToken, tokenAddress []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "DepositToken", tokenAddressRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerDepositToken)
				if err := _TreasureManager.contract.UnpackLog(event, "DepositToken", log); err != nil {
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

// ParseDepositToken is a log parse operation binding the contract event 0x4b3f81827ede20c81afbf1bb77b954afcdcae24d391d99042310cb1d9210dd57.
//
// Solidity: event DepositToken(address indexed tokenAddress, address indexed sender, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) ParseDepositToken(log types.Log) (*TreasureManagerDepositToken, error) {
	event := new(TreasureManagerDepositToken)
	if err := _TreasureManager.contract.UnpackLog(event, "DepositToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerGrantRewardTokenAmountIterator is returned from FilterGrantRewardTokenAmount and is used to iterate over the raw logs and unpacked data for GrantRewardTokenAmount events raised by the TreasureManager contract.
type TreasureManagerGrantRewardTokenAmountIterator struct {
	Event *TreasureManagerGrantRewardTokenAmount // Event containing the contract specifics and raw log

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
func (it *TreasureManagerGrantRewardTokenAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerGrantRewardTokenAmount)
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
		it.Event = new(TreasureManagerGrantRewardTokenAmount)
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
func (it *TreasureManagerGrantRewardTokenAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerGrantRewardTokenAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerGrantRewardTokenAmount represents a GrantRewardTokenAmount event raised by the TreasureManager contract.
type TreasureManagerGrantRewardTokenAmount struct {
	TokenAddress common.Address
	Granter      common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGrantRewardTokenAmount is a free log retrieval operation binding the contract event 0x10621458f3ad2a9cfcb87c216122570629e44079d6af4d717035eb55d2c60424.
//
// Solidity: event GrantRewardTokenAmount(address indexed tokenAddress, address granter, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) FilterGrantRewardTokenAmount(opts *bind.FilterOpts, tokenAddress []common.Address) (*TreasureManagerGrantRewardTokenAmountIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "GrantRewardTokenAmount", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerGrantRewardTokenAmountIterator{contract: _TreasureManager.contract, event: "GrantRewardTokenAmount", logs: logs, sub: sub}, nil
}

// WatchGrantRewardTokenAmount is a free log subscription operation binding the contract event 0x10621458f3ad2a9cfcb87c216122570629e44079d6af4d717035eb55d2c60424.
//
// Solidity: event GrantRewardTokenAmount(address indexed tokenAddress, address granter, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) WatchGrantRewardTokenAmount(opts *bind.WatchOpts, sink chan<- *TreasureManagerGrantRewardTokenAmount, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "GrantRewardTokenAmount", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerGrantRewardTokenAmount)
				if err := _TreasureManager.contract.UnpackLog(event, "GrantRewardTokenAmount", log); err != nil {
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

// ParseGrantRewardTokenAmount is a log parse operation binding the contract event 0x10621458f3ad2a9cfcb87c216122570629e44079d6af4d717035eb55d2c60424.
//
// Solidity: event GrantRewardTokenAmount(address indexed tokenAddress, address granter, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) ParseGrantRewardTokenAmount(log types.Log) (*TreasureManagerGrantRewardTokenAmount, error) {
	event := new(TreasureManagerGrantRewardTokenAmount)
	if err := _TreasureManager.contract.UnpackLog(event, "GrantRewardTokenAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TreasureManager contract.
type TreasureManagerInitializedIterator struct {
	Event *TreasureManagerInitialized // Event containing the contract specifics and raw log

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
func (it *TreasureManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerInitialized)
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
		it.Event = new(TreasureManagerInitialized)
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
func (it *TreasureManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerInitialized represents a Initialized event raised by the TreasureManager contract.
type TreasureManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TreasureManager *TreasureManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*TreasureManagerInitializedIterator, error) {

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TreasureManagerInitializedIterator{contract: _TreasureManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TreasureManager *TreasureManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TreasureManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerInitialized)
				if err := _TreasureManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TreasureManager *TreasureManagerFilterer) ParseInitialized(log types.Log) (*TreasureManagerInitialized, error) {
	event := new(TreasureManagerInitialized)
	if err := _TreasureManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TreasureManager contract.
type TreasureManagerOwnershipTransferredIterator struct {
	Event *TreasureManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TreasureManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerOwnershipTransferred)
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
		it.Event = new(TreasureManagerOwnershipTransferred)
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
func (it *TreasureManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerOwnershipTransferred represents a OwnershipTransferred event raised by the TreasureManager contract.
type TreasureManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TreasureManager *TreasureManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TreasureManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerOwnershipTransferredIterator{contract: _TreasureManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TreasureManager *TreasureManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TreasureManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerOwnershipTransferred)
				if err := _TreasureManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TreasureManager *TreasureManagerFilterer) ParseOwnershipTransferred(log types.Log) (*TreasureManagerOwnershipTransferred, error) {
	event := new(TreasureManagerOwnershipTransferred)
	if err := _TreasureManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TreasureManager contract.
type TreasureManagerRoleAdminChangedIterator struct {
	Event *TreasureManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TreasureManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerRoleAdminChanged)
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
		it.Event = new(TreasureManagerRoleAdminChanged)
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
func (it *TreasureManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerRoleAdminChanged represents a RoleAdminChanged event raised by the TreasureManager contract.
type TreasureManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TreasureManager *TreasureManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TreasureManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerRoleAdminChangedIterator{contract: _TreasureManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TreasureManager *TreasureManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TreasureManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerRoleAdminChanged)
				if err := _TreasureManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TreasureManager *TreasureManagerFilterer) ParseRoleAdminChanged(log types.Log) (*TreasureManagerRoleAdminChanged, error) {
	event := new(TreasureManagerRoleAdminChanged)
	if err := _TreasureManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TreasureManager contract.
type TreasureManagerRoleGrantedIterator struct {
	Event *TreasureManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *TreasureManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerRoleGranted)
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
		it.Event = new(TreasureManagerRoleGranted)
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
func (it *TreasureManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerRoleGranted represents a RoleGranted event raised by the TreasureManager contract.
type TreasureManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TreasureManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerRoleGrantedIterator{contract: _TreasureManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TreasureManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerRoleGranted)
				if err := _TreasureManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) ParseRoleGranted(log types.Log) (*TreasureManagerRoleGranted, error) {
	event := new(TreasureManagerRoleGranted)
	if err := _TreasureManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TreasureManager contract.
type TreasureManagerRoleRevokedIterator struct {
	Event *TreasureManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TreasureManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerRoleRevoked)
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
		it.Event = new(TreasureManagerRoleRevoked)
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
func (it *TreasureManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerRoleRevoked represents a RoleRevoked event raised by the TreasureManager contract.
type TreasureManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TreasureManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerRoleRevokedIterator{contract: _TreasureManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TreasureManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerRoleRevoked)
				if err := _TreasureManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TreasureManager *TreasureManagerFilterer) ParseRoleRevoked(log types.Log) (*TreasureManagerRoleRevoked, error) {
	event := new(TreasureManagerRoleRevoked)
	if err := _TreasureManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerWithdrawManagerUpdateIterator is returned from FilterWithdrawManagerUpdate and is used to iterate over the raw logs and unpacked data for WithdrawManagerUpdate events raised by the TreasureManager contract.
type TreasureManagerWithdrawManagerUpdateIterator struct {
	Event *TreasureManagerWithdrawManagerUpdate // Event containing the contract specifics and raw log

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
func (it *TreasureManagerWithdrawManagerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerWithdrawManagerUpdate)
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
		it.Event = new(TreasureManagerWithdrawManagerUpdate)
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
func (it *TreasureManagerWithdrawManagerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerWithdrawManagerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerWithdrawManagerUpdate represents a WithdrawManagerUpdate event raised by the TreasureManager contract.
type TreasureManagerWithdrawManagerUpdate struct {
	WithdrawManger common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawManagerUpdate is a free log retrieval operation binding the contract event 0x799e16a314d482c87bc47fd219011aaf33f4f9c7e302be5d7a0af286a52998b9.
//
// Solidity: event WithdrawManagerUpdate(address indexed withdrawManger)
func (_TreasureManager *TreasureManagerFilterer) FilterWithdrawManagerUpdate(opts *bind.FilterOpts, withdrawManger []common.Address) (*TreasureManagerWithdrawManagerUpdateIterator, error) {

	var withdrawMangerRule []interface{}
	for _, withdrawMangerItem := range withdrawManger {
		withdrawMangerRule = append(withdrawMangerRule, withdrawMangerItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "WithdrawManagerUpdate", withdrawMangerRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerWithdrawManagerUpdateIterator{contract: _TreasureManager.contract, event: "WithdrawManagerUpdate", logs: logs, sub: sub}, nil
}

// WatchWithdrawManagerUpdate is a free log subscription operation binding the contract event 0x799e16a314d482c87bc47fd219011aaf33f4f9c7e302be5d7a0af286a52998b9.
//
// Solidity: event WithdrawManagerUpdate(address indexed withdrawManger)
func (_TreasureManager *TreasureManagerFilterer) WatchWithdrawManagerUpdate(opts *bind.WatchOpts, sink chan<- *TreasureManagerWithdrawManagerUpdate, withdrawManger []common.Address) (event.Subscription, error) {

	var withdrawMangerRule []interface{}
	for _, withdrawMangerItem := range withdrawManger {
		withdrawMangerRule = append(withdrawMangerRule, withdrawMangerItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "WithdrawManagerUpdate", withdrawMangerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerWithdrawManagerUpdate)
				if err := _TreasureManager.contract.UnpackLog(event, "WithdrawManagerUpdate", log); err != nil {
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

// ParseWithdrawManagerUpdate is a log parse operation binding the contract event 0x799e16a314d482c87bc47fd219011aaf33f4f9c7e302be5d7a0af286a52998b9.
//
// Solidity: event WithdrawManagerUpdate(address indexed withdrawManger)
func (_TreasureManager *TreasureManagerFilterer) ParseWithdrawManagerUpdate(log types.Log) (*TreasureManagerWithdrawManagerUpdate, error) {
	event := new(TreasureManagerWithdrawManagerUpdate)
	if err := _TreasureManager.contract.UnpackLog(event, "WithdrawManagerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasureManagerWithdrawTokenIterator is returned from FilterWithdrawToken and is used to iterate over the raw logs and unpacked data for WithdrawToken events raised by the TreasureManager contract.
type TreasureManagerWithdrawTokenIterator struct {
	Event *TreasureManagerWithdrawToken // Event containing the contract specifics and raw log

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
func (it *TreasureManagerWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasureManagerWithdrawToken)
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
		it.Event = new(TreasureManagerWithdrawToken)
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
func (it *TreasureManagerWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasureManagerWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasureManagerWithdrawToken represents a WithdrawToken event raised by the TreasureManager contract.
type TreasureManagerWithdrawToken struct {
	TokenAddress    common.Address
	Sender          common.Address
	WithdrawAddress common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawToken is a free log retrieval operation binding the contract event 0x9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746.
//
// Solidity: event WithdrawToken(address indexed tokenAddress, address indexed sender, address withdrawAddress, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) FilterWithdrawToken(opts *bind.FilterOpts, tokenAddress []common.Address, sender []common.Address) (*TreasureManagerWithdrawTokenIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.FilterLogs(opts, "WithdrawToken", tokenAddressRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TreasureManagerWithdrawTokenIterator{contract: _TreasureManager.contract, event: "WithdrawToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawToken is a free log subscription operation binding the contract event 0x9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746.
//
// Solidity: event WithdrawToken(address indexed tokenAddress, address indexed sender, address withdrawAddress, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) WatchWithdrawToken(opts *bind.WatchOpts, sink chan<- *TreasureManagerWithdrawToken, tokenAddress []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TreasureManager.contract.WatchLogs(opts, "WithdrawToken", tokenAddressRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasureManagerWithdrawToken)
				if err := _TreasureManager.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
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

// ParseWithdrawToken is a log parse operation binding the contract event 0x9ca7c1e047552a8048d924a5a8d3c150eb861086a72a9100e5f19d1176c1b746.
//
// Solidity: event WithdrawToken(address indexed tokenAddress, address indexed sender, address withdrawAddress, uint256 amount)
func (_TreasureManager *TreasureManagerFilterer) ParseWithdrawToken(log types.Log) (*TreasureManagerWithdrawToken, error) {
	event := new(TreasureManagerWithdrawToken)
	if err := _TreasureManager.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
