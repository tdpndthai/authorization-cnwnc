package controllers

import (
	"admin-go/models"
	"admin-go/services/serviceImplement"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var loginData models.LoginModel
	err := json.NewDecoder(r.Body).Decode(&loginData)

	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Đăng nhập không thành công"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	userServices := serviceImplement.NewUserService()

	resp := userServices.Login(&loginData)
	json.NewEncoder(w).Encode(resp)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	userServices := serviceImplement.NewUserService()
	resp := userServices.GetAll()
	json.NewEncoder(w).Encode(resp)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Id không hợp lệ"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	userServices := serviceImplement.NewUserService()
	resp := userServices.GetById(id)
	json.NewEncoder(w).Encode(resp)

}

func GetDetailUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Id không hợp lệ"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	userServices := serviceImplement.NewUserService()
	resp := userServices.GetDetailUser(id)
	json.NewEncoder(w).Encode(resp)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var dataUpdate models.UpdateUserModel

	//decode dữ liệu từ body
	err := json.NewDecoder(r.Body).Decode(&dataUpdate)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Dữ liệu cập nhật không hợp lệ"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	userServices := serviceImplement.NewUserService()
	resp := userServices.UpdateUser(&dataUpdate)
	json.NewEncoder(w).Encode(resp)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var dataAdd models.CreateUserModel

	//decode dữ liệu từ body
	err := json.NewDecoder(r.Body).Decode(&dataAdd)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Dữ liệu tạo mới không hợp lệ"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	userServices := serviceImplement.NewUserService()
	resp := userServices.CreateUser(&dataAdd)
	json.NewEncoder(w).Encode(resp)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "ID người dùng ko hợp lệ"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	userServices := serviceImplement.NewUserService()
	resp := userServices.DeleteUser(id)
	json.NewEncoder(w).Encode(resp)
}
