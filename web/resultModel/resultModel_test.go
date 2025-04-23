package resultModel

import (
	"testing"
)

func Test1(t *testing.T) {
	t.Log(Ok("success"))
	t.Log(Error("fail"))
}