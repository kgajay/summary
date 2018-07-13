package services

import (
	"model"
	"utils"
	"net/http"
)

func CreateUser(req *model.User) (int, interface{}) {

	var err map[string]interface{}
	err = utils.ValidateStruct(req)
	if err != nil {
		return http.StatusBadRequest, err
	}

	u := model.CreateEntry(req)

	return http.StatusOK, u
}

func GetUser(id int) (int, interface{}) {
	//var u = make(map[string]interface{})
	r := model.FetchById(&model.User{}, id)
	if r.RowsAffected == 0 {
		return http.StatusNotFound, nil
	}
	return http.StatusOK, r.Value
}
