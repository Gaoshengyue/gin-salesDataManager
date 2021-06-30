package StatementService

import (
	"context"
	"dolphin/salesManager/pkg/gmongo"
	"dolphin/salesManager/pkg/logging"
	"dolphin/salesManager/schema/ControllerSchema/StatementControllerSchema/StatementControllerRequestSchema"
	"dolphin/salesManager/schema/ServiceSchema/StatementSchema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetNetPowerPolicyBaseDataArray(PageNationObj StatementControllerRequestSchema.PageNationColumn) ([]StatementSchema.NetPowerPolicySchema, error) {
	collection := gmongo.Collection("网电保单")
	findOptions := options.Find()

	findOptions.SetSkip(PageNationObj.PageSize * (PageNationObj.Current - 1))
	findOptions.SetLimit(PageNationObj.PageSize)
	cur, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	if err := cur.Err(); err != nil {
		logging.Error(err)
		return nil, err
	}
	var NetPowerPolicyBaseArray []StatementSchema.NetPowerPolicySchema
	err = cur.All(context.Background(), &NetPowerPolicyBaseArray)
	if err != nil {
		return nil, err
	}
	return NetPowerPolicyBaseArray, nil
}
