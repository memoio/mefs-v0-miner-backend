package eth

import(
	"testing"
)

func TestGet(t *testing.T) {
	data := get("asdf","123123",123)
	t.Log(data)
}
func TestGetTransactionReceipt(t *testing.T) {
	data := GetTransactionReceipt("0x8cc0bb2af869f05ebf59a37644e8fb1c4876ff5890b72b05ebe2771f0d7280c2")
	t.Log(data)
}