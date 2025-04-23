package eth
import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"

	"longchain.com/memoriae/orePool/log"
	"longchain.com/memoriae/orePool/contracts"
)

type KeeperProviderMap struct{
	ct *contracts.KeeperProviderMap
	auth *bind.TransactOpts
	call *bind.CallOpts
}
func NewKeeperProviderMap(addr string, pk string) KeeperProviderMap {
	ct, err := contracts.NewKeeperProviderMap(common.HexToAddress(addr), getClient())
	if err != nil {
		log.Error(err)
	}
	key, err := crypto.HexToECDSA(pk[2:])
	if err != nil {
		log.Error(err)
	}
	auth := bind.NewKeyedTransactor(key)
	call := &bind.CallOpts{
		From:common.HexToAddress(callAddr),
	}
	return KeeperProviderMap{
		ct:ct,
		auth:auth,
		call:call,
	}
}
func (this KeeperProviderMap) GetAllKeeper() []string {
	var res []string
	addrList,err := this.ct.GetAllKeeper(this.call)
	if err != nil {
		log.Error(err)
		return res
	}
	for _,addr := range addrList {
		res = append(res,addr.String())
	}

	return res
}
func (this KeeperProviderMap) GetProvider(keeper string) []string {
	var res []string
	addrList,err := this.ct.GetProvider(this.call,common.HexToAddress(keeper))
	if err != nil {
		log.Error(err)
		return res
	}
	for _,addr := range addrList {
		res = append(res,addr.String())
	}

	return res
}