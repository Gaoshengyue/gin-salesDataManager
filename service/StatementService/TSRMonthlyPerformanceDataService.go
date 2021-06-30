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

func GetTSRMonthlyPerformanceBaseDataArray(PageNationObj StatementControllerRequestSchema.PageNationColumn) ([]StatementSchema.TSRMonthlyPerformance, error) {
	collection := gmongo.Collection("签单月拨打表现")
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
	var TSRMonthlyPerformanceBaseArray []StatementSchema.TSRMonthlyPerformance
	err = cur.All(context.Background(), &TSRMonthlyPerformanceBaseArray)
	if err != nil {
		return nil, err
	}
	return TSRMonthlyPerformanceBaseArray, nil
}
