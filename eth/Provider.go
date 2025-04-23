package eth
import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"

	"longchain.com/memoriae/orePool/log"
	"longchain.com/memoriae/orePool/contracts"
)

type Provider struct{
	ct *contracts.Provider
	auth *bind.TransactOpts
	call *bind.CallOpts
}
type ProviderInfo struct{
	Addr string `json:"addr"`
	Money string `json:"money"`
	Size string `json:"size"`
	Time string `json:"time"`
}
func NewProvider(addr string, pk string) Provider {
	ct, err := contracts.NewProvider(common.HexToAddress(addr), getClient())
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
	return Provider{
		ct:ct,
		auth:auth,
		call:call,
	}
}
func (this Provider) Info(addr string) ProviderInfo {
	_, money, size, time, err := this.ct.Info(this.call,common.HexToAddress(addr))
	if err != nil {
		log.Error(err)
		return ProviderInfo{}
	}
	
	return ProviderInfo{addr,money.Text(10),size.Text(10),time.Text(10),}
}