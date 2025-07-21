package routes

import sqlite "rolerocket/internal/db"

func GetUsers() ([]string, error) {
	return sqlite.DBInstance.GetUsers()
}

func InsertUser() {
	sqlite.DBInstance.InsertUser()
}

func UpdateUser() {

}

func DeleteUser() {

}
