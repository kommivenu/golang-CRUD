package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:root@tcp(localhost:3306)/venu"

type StudentDetails struct {
	gorm.Model
	Name     string `json:"name"`
	Mobileno int    `json:"mobileno"`
	Email    string `json:"email"`
	Course   string `json:"course"`
}

func InitialMigration() {

	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println("err.Error()")
		panic("can not connect to DB")
	}
	DB.AutoMigrate(&StudentDetails{})
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var users []StudentDetails
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)

}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	var user StudentDetails
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var user StudentDetails
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	var user StudentDetails
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	var user StudentDetails
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode(user)

}
