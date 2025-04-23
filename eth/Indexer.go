package eth
import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"

	"longchain.com/memoriae/orePool/log"
	"longchain.com/memoriae/orePool/contracts"
)

type Indexer struct{
	ct *contracts.Indexer
	auth *bind.TransactOpts
	call *bind.CallOpts
}
func NewIndexer(addr string, pk string) Indexer {
	ct, err := contracts.NewIndexer(common.HexToAddress(addr), getClient())
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
	return Indexer{
		ct:ct,
		auth:auth,
		call:call,
	}
}

func (this Indexer) GetResolver(key string) string {
	addr,err := this.ct.GetResolver(this.call,key)
	if err != nil {
		log.Error(err)
	}

	return addr.String()
}