package http

import(
	"testing"
)

func TestGet(t *testing.T) {
	data := Get("https://api.etherscan.io/api?module=proxy&action=eth_getTransactionByHash&txhash=0x1e2910a262b1008d0616a0beb24c1a491d78771baa54a33e66065e03b1f46bc1")
	t.Log(data)
}