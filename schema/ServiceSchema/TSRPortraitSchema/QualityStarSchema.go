package TSRPortraitSchema

import (
	"dolphin/salesManager/pkg/gmongo"
	"dolphin/salesManager/pkg/logging"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

// 质检评星结构体
type QualityStarSchema struct {
	TSRId       int32  `bson:"坐席工号"`
	QualityType string `bson:"质检类型"`
	PolicyId    string `bson:"投注单号"`
}

//质检评星查询
func QualityStarHandler(filter bson.D) ([]QualityStarSchema, error) {
	collection := gmongo.Collection("质检报表")
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	if err := cur.Err(); err != nil {
		logging.Error(err)
		return nil, err
	}
	var QualityStar []QualityStarSchema

	err = cur.All(context.Background(), &QualityStar)
	return QualityStar, err
}
