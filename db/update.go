package db

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/letanthang/mongo/sequence"
	"github.com/letanthang/my_framework/db/types"
	"github.com/letanthang/my_framework/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertAccount(req types.AccountAddReq) (interface{}, error) {
	acc, _ := GetOneAccount(req.Phone, req.Password)
	if acc != nil {
		return nil, errors.New("Existed user")
	}

	newID, _ := sequence.GetNextID(Client.Database(dbName).Collection("counter"),
		"account_id_seq")
	account := types.Account{
		ID:        newID,
		Phone:     req.Phone,
		Email:     req.Email,
		Password:  utils.GetMD5(req.Password),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	res, err := Client.Database(dbName).Collection("account").InsertOne(context.TODO(), account)
	id := res.InsertedID
	return id, err
}

func UpdateAccount(claims *types.MyClaims, req types.AccountUpdateReq) (interface{}, error) {

	filter := bson.M{"phone": claims.ClientPhone}

	var data map[string]interface{}
	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &data)

	//md5 password
	if p, ok := data["password"]; ok == true {
		data["password"] = utils.GetMD5(p.(string))
	}

	update := bson.M{"$set": data}

	res, err := Client.Database(dbName).Collection("account").UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Error(err)
		return nil, errors.New("update account error")
	}

	return res, nil
}

func DeleteAccount(id int) (*mongo.DeleteResult, error) {
	res, err := Client.Database(dbName).Collection("account").DeleteOne(context.TODO(), bson.D{{"id", id}})
	return res, err
}
