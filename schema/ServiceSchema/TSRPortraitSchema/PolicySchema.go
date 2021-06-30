package TSRPortraitSchema

import (
	"dolphin/salesManager/pkg/gmongo"
	"dolphin/salesManager/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

type PolicyStatusIdList struct {
	PolicyId string `bson:"保单号"`
}

type PolicyStatusSchema struct {
	PolicyStatusIdList
	PaymentCount       float64 `bson:"缴费次数"`
	UnderWritingDate   string  `bson:"承保日期"`
	PolicyStatus       string  `bson:"保单状态"`
	ShouldPaymentCount int32   `bson:"应缴费次数"`
	SurrenderTime      string  `bson:"退保时间"`
}

//获取所有保单状态
func GetTSRSurrenderPolicyList(filter bson.D) ([]PolicyStatusSchema, error) {
	collection := gmongo.Collection("保单状态")
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	if err := cur.Err(); err != nil {
		logging.Error(err)
		return nil, err
	}
	var SurrenderPolicyList []PolicyStatusSchema
	err = cur.All(context.Background(), &SurrenderPolicyList)

	return SurrenderPolicyList, err

}

//获取所有保单状态id
func GetTSRSurrenderPolicyIdList(filter bson.D) ([]string, error) {
	collection := gmongo.Collection("保单状态")
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	if err := cur.Err(); err != nil {
		logging.Error(err)
		return nil, err
	}
	var SurrenderPolicyIdObjList []PolicyStatusIdList
	err = cur.All(context.Background(), &SurrenderPolicyIdObjList)
	if err != nil {
		return nil, err
	}
	var TSRSurrenderPolicyIdList []string
	for i := 0; i < len(SurrenderPolicyIdObjList); i++ {
		TSRSurrenderPolicyIdList = append(TSRSurrenderPolicyIdList, SurrenderPolicyIdObjList[i].PolicyId)
	}
	return TSRSurrenderPolicyIdList, err

}
