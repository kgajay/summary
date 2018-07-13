package utils

import (
	"github.com/asaskevich/govalidator"
	"logger"
	"reflect"
)

func ValidateStruct(i interface{}) (resp map[string]interface{}) {
	res, err := govalidator.ValidateStruct(i)
	if err != nil {
		logger.Log.Infof("govalidator obj: %s, res: %s, err: %s", i, res, err.Error())
	} else {
		logger.Log.Infof("govalidator obj: %s, res: %s", i, res)
	}
	return marshallError(err)
}

func marshallError(err error) (resp map[string]interface{}) {
	if err != nil {
		resp = make(map[string]interface{})
		for _, e := range err.(govalidator.Errors) {
			logger.Log.Infof("type %s", reflect.TypeOf(e))
			switch e.(type) {
			case govalidator.Errors:
				if errs := marshallError(e); errs != nil {
					resp[e.(govalidator.Error).Name] = errs
				}
			case govalidator.Error:
				resp[e.(govalidator.Error).Name] = e.Error()
			}

		}
		return resp
	}
	return nil
}