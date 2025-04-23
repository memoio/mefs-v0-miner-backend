package json
import(
	"strings"
	"strconv"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)
type Json struct {
	data []byte
}

func NewByStr(str string) Json{
	return Json{
		data:[]byte(str),
	}
}
func (j Json) GetString(selector string) string{
	return jsoniter.Get(j.data,format(selector)...).ToString()
}

func split(data string) []string {
  res := ""
  c := ""
  
  for i:=0; i<len(data); i++ {
    c = string(data[i])
    if c == "." || c == "[" {
      res += ","
    }else if c == "]" {
      res += ""
    }else{
      res += c
    }
  }
  
  return strings.Split(res,",")
}
// transform result.logs[0].topics[0] to ["result","logs",0,"topics",0]
func format(data string) []interface{} {
	arr := split(data)
	res := make([]interface{}, len(arr))

	for i:=0; i<len(arr); i++ {
		if val, err := strconv.Atoi(arr[i]); err == nil {
			res[i] = val
		}else{
			res[i] = arr[i]
		}
	}
	return res
}

func Compress(data string) string{
	flag := false
  res := ""
  c := ""
  for i:=0; i<len(data); i++ {
    c = string(data[i])
    if c == `"` {
      if flag {
        flag = false
      }else{
        flag = true
      }
    }
    if !flag {
      if !(c == " " || c == "\n" || c == "\t") {
        res += c
      }
    }else{
      res += c
    }
  }
  return res
}

func ToStandard(data string) string{
	data = Compress(data)

  res := ""
  c := ""

  flag := false

  rightQuotes := false
  for i:=0; i<len(data)-1; i++ {
    c = string(data[i])
    res += c
    if c == `[` {
      flag = true
    }else if c == `]` {
      flag = false
    }else if (c == `{` || (c == `,` && !flag)) && string(data[i+1]) != `"` {
      res += `"`
      rightQuotes = true
    }else if rightQuotes && string(data[i+1]) == `:`{
      res += `"`
    }
  }

  res += string(data[len(data)-1])
  
  return res
}

func ToMap(data string) map[string]interface{}{
	var res map[string]interface{}
	err := jsoniter.Unmarshal([]byte(data), &res)
	if err != nil {
		fmt.Println("json to map failed", err)
	}
	return res
}