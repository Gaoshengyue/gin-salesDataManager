package StatementService

import (
	"context"
	"dolphin/salesManager/pkg/gmongo"
	"dolphin/salesManager/pkg/logging"
	"dolphin/salesManager/schema/ServiceSchema/StatementSchema"
	"dolphin/salesManager/schema/ServiceSchema/TSRPortraitSchema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 获取数据概览基础数据
func GetBaseData() (StatementSchema.DataOverviewBaseDataSchema, error) {
	var DataOverviewBaseData StatementSchema.DataOverviewBaseDataSchema
	//　获取话务统计基础指标
	TotalCallCount, CallConnectRate, CallTimeMean, err := TSRPortraitSchema.GetCallRecordBaseData()
	if err != nil {
		return DataOverviewBaseData, err
	}
	DataOverviewBaseData.TotalCallCount = TotalCallCount
	DataOverviewBaseData.CallConnectRate = CallConnectRate
	DataOverviewBaseData.CallTimeMean = CallTimeMean
	//　获取全部成单量
	TotalPolicyCount, err := TSRPortraitSchema.GetTotalPolicyCount()
	if err != nil {
		return DataOverviewBaseData, err
	}
	DataOverviewBaseData.TotalCustomerCount = TotalPolicyCount
	return DataOverviewBaseData, nil
}

// 获取产品分布数据
func GetProductDistribution() ([]StatementSchema.DataOverviewProductDistribution, error) {
	var showWithInfo []bson.M
	var DistributionArray []StatementSchema.DataOverviewProductDistribution
	collection := gmongo.Collection("网电保单")
	groupStage := bson.D{{"$group", bson.D{{"_id", "$产品名称"},
		{"TotalCount", bson.D{{"$sum", 1}}},
	}}}
	showInfoCursor, err := collection.Aggregate(context.Background(), mongo.Pipeline{groupStage})
	if err != nil {
		logging.Error(err)
		return DistributionArray, err
	}

	if err = showInfoCursor.All(context.Background(), &showWithInfo); err != nil {
		logging.Error(err)
		return DistributionArray, err
	}
	if showWithInfo != nil {
		for i := 0; i < len(showWithInfo); i++ {
			var productName string
			Name, ok := showWithInfo[i]["_id"].(string)
			if ok {
				productName = Name
			}
			var productCount int32
			Count, ok := showWithInfo[i]["TotalCount"].(int32)
			if ok {
				productCount = Count
			}
			distribution := StatementSchema.DataOverviewProductDistribution{ProductName: productName, ProductCount: productCount}
			DistributionArray = append(DistributionArray, distribution)
		}
	}
	return DistributionArray, nil

}

// 获取
