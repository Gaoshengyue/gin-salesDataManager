package StatementService

import (
	"dolphin/salesManager/pkg/gmongo"
	"dolphin/salesManager/pkg/logging"
	"dolphin/salesManager/schema/ControllerSchema/StatementControllerSchema/StatementControllerRequestSchema"
	"dolphin/salesManager/schema/ServiceSchema/StatementSchema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func GetTrafficMeasurementBaseDataArray(PageNationObj StatementControllerRequestSchema.PageNationColumn) ([]StatementSchema.TrafficMeasurementSchema, error) {
	collection := gmongo.Collection("话务统计")
	findOptions := options.Find()
	findOptions.SetLimit(PageNationObj.PageSize)
	findOptions.SetSkip(PageNationObj.PageSize * (PageNationObj.Current - 1))
	cur, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	if err := cur.Err(); err != nil {
		logging.Error(err)
		return nil, err
	}
	var TrafficMeasurementBaseArray []StatementSchema.TrafficMeasurementSchema
	err = cur.All(context.Background(), &TrafficMeasurementBaseArray)
	if err != nil {
		return nil, err
	}
	return TrafficMeasurementBaseArray, nil
}
