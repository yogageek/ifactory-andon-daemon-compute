package util

import (
	"encoding/json"

	ensaasErr "iii/ifactory/compute/util/cch/errors"
	"iii/ifactory/compute/util/cch/logger"
	"iii/ifactory/compute/util/cch/validator"
)

//導入validor包 可用Tag去檢查前端參數是否有帶

type Json struct{}

func (p *Json) ToJson(v interface{}) (string, error) {
	return ToJson(v)
}

func (p *Json) FromJson(in string, v interface{}) error {
	return FromJson(in, v)
}

func (p *Json) FromJsonV(in string, v interface{}) error {
	return FromJsonV(in, v)
}

func ToJson(v interface{}) (string, error) {
	rsp, err := json.Marshal(v)
	if err != nil {
		err := ensaasErr.NewEnsaasError(ensaasErr.Json_Parser_Error, err.Error())
		return "", err
	}

	return string(rsp), nil
}

//把json轉成指定物件 v放你要轉換成的物件類型 且 tag放validate:"required"會自動檢查
func FromJson(in string, v interface{}) error {
	if err := json.Unmarshal([]byte(in), v); err != nil {
		logger.Error(err.Error())
		err := ensaasErr.NewEnsaasError(ensaasErr.Json_Parser_Error, err.Error())
		return err
	}

	if err := validator.Struct(v); err != nil {
		err := ensaasErr.NewEnsaasError(ensaasErr.Json_Validate_Error, err.Error())
		return err
	}
	return nil
}

//把json轉成指定物件  v放你要轉換成的物件類型
func FromJsonV(in string, v interface{}) error {
	if err := validator.JsonStruct([]byte(in), v); err != nil {
		return err
	}
	return nil
}

func FromJsonNoV(in string, v interface{}) error {
	if err := json.Unmarshal([]byte(in), v); err != nil {
		return err
	}
	return nil
}

/*
等於省下了
var v interface{}
if err := json.Unmarshal(body, &v); err != nil {
	glog.Error(err)
}

var v xxxstruct{}
if err := json.Unmarshal(body, &v); err != nil {
	glog.Error(err)
}
*/
