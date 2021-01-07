package admin

import "fmt"

type Result struct {
	Code int `json:"code"` //0登陆成功
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

type DataList []interface{}
type DataMap map[string]interface{}

func GetResult() *Result {
	return &Result{
		Code:0,
		Msg:"Success",
		Data:nil,
	}
}

func CheckErr(err error, result *Result)  {
	result.Code = 1
	result.Msg = fmt.Sprintf("%v", err)
}