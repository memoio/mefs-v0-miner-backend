package service
import(
	// "fmt"

	"longchain.com/memoriae/orePool/config"
	// "longchain.com/memoriae/orePool/json"
	// "longchain.com/memoriae/orePool/log"
	"longchain.com/memoriae/orePool/eth"
)
var(
	pk = config.Pk
	indexerAddr = config.IndexerAddr
)

// [
// 	{
// 		keeper:'',
// 		providerList:[
// 			{
// 				addr:'',
// 				money:'',
// 				size:'',
// 				time:'',
// 			}
// 		]
// 	}
// ]
func GetList() interface{} {
	var res []eth.KeeperInfo
	indexer := eth.NewIndexer(indexerAddr,pk)
	kpmAddr := indexer.GetResolver("keeperProviderMap")
	providerAddr := indexer.GetResolver("provider")
	kp := eth.NewKeeperProviderMap(kpmAddr,pk)
	p := eth.NewProvider(providerAddr,pk)
	keeperList := kp.GetAllKeeper()
	for _,keeper := range keeperList {
		keeperInfo := eth.KeeperInfo{Addr:keeper}
		providerList := kp.GetProvider(keeper)
		for _,provider := range providerList {
			pi := p.Info(provider)
			keeperInfo.ProviderList = append(keeperInfo.ProviderList,pi)
		}
		res = append(res,keeperInfo)
	}
	return res
}