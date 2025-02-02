package database

import (
	"context"
	"url-shortner/constant"
	"url-shortner/types"

	"go.mongodb.org/mongo-driver/bson"
)

func (mgr *manager) Insert(data interface{}, collectionName string) (interface{}, error) {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	result, err := inst.InsertOne(context.TODO(), data)
	return result.InsertedID, err
}

func (mgr *manager) GetUrlFromCode(code string, collectionName string) (res types.UrlDb, err error) {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	err = inst.FindOne(context.TODO(), bson.M{"url_code": code}).Decode(&res)
	return res, err
}
