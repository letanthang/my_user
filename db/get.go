package db

import (
	"context"
	"time"

	"github.com/letanthang/my_framework/db/types"
	"github.com/letanthang/my_framework/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetOneAccount(phone, password string) (*types.Account, error) {
	var student types.Account
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	hashPass := utils.GetMD5(password)
	filter := bson.M{"phone": phone, "password": hashPass}
	err := Client.Database(dbName).Collection("account").FindOne(ctx, filter).Decode(&student)

	if err != nil {
		return nil, err
	}
	return &student, nil
}
