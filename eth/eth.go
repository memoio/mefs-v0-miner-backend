package eth
import(
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	jsoniter "github.com/json-iterator/go"

	"longchain.com/memoriae/orePool/config"
	"longchain.com/memoriae/orePool/log"
	"longchain.com/memoriae/orePool/json"
	"longchain.com/memoriae/orePool/web/http"
)

var(
	url = config.EthUrl
	callAddr = "0eb5b66c31b3c5a12aae81a9d629540b6433cac6"
)

func getClient() *ethclient.Client{
	client, err := rpc.Dial(url)
	if err != nil {
		log.Error(err)
	}
	return ethclient.NewClient(client)
}

func get(method string, params ...interface{}) string {
	str,err := jsoniter.MarshalToString(params)
	if err != nil {
		log.Error(err)
	}
	data := fmt.Sprintf(`
		{
			jsonrpc:"2.0",
			id:1,
			method:"%s",
			params:%s
		}
	`,method,str)
	data = json.ToStandard(data)
  // fmt.Println(data)
  
	return http.PostByJson(url,data)
}
func GetTransactionReceipt(hash string) string {
  return get("eth_getTransactionReceipt",hash)
}