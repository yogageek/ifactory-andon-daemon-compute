package validator

import (
	"encoding/json"
	"reflect"
	"strings"

	ensaasErrors "iii/ifactory/compute/util/cch/errors"

	"iii/ifactory/compute/util/cch/logger"

	"gopkg.in/go-playground/validator.v9"
)

var v *validator.Validate

func init() {
	v = validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func Struct(s interface{}) error {
	err := v.Struct(s)
	if err != nil {
		var errorMsg string
		for _, e := range err.(validator.ValidationErrors) {
			logger.Debug(e.Field())
			errorMsg += e.Tag() + ":" + e.Field() + "; "
		}
		ensaasErr := ensaasErrors.NewEnsaasError(ensaasErrors.Json_Validate_Error, "", errorMsg)

		return ensaasErr
	}

	return nil
}

func JsonStruct(data []byte, v interface{}) error {
	//	json.Unmarshal(data, v)
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	if err := Struct(v); err != nil {
		logger.Debug(err.Error())
		return err
	}
	return nil
}
