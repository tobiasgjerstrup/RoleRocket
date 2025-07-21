package routes

import (
	"context"
	sqlite "rolerocket/internal/db"
)

func GetUsers(ctx context.Context) ([]string, error) {
	return sqlite.DBInstance.GetUsers(ctx)
}

func InsertUser(ctx context.Context) {
	sqlite.DBInstance.InsertUser(ctx)
}

func UpdateUser() {

}

func DeleteUser() {

}
