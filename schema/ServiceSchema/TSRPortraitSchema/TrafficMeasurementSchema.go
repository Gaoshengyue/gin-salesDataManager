package TSRPortraitSchema

import (
	"dolphin/salesManager/pkg/gmongo"
	"dolphin/salesManager/pkg/logging"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"math"
)

type TrafficMeasurementSchema struct {
	TSRId                      int32  `bson:"小宝工号"`
	TSRName                    string `bson:"坐席"`
	DateTime                   string `bson:"日期"`
	CallTimeTotal              string `bson:"通话总时长"`
	CallCountTotal             int32  `bson:"呼出总次数"`
	CallConnectCount           int32  `bson:"呼出接通数"`
	CallConnectRate            string `bson:"呼出接通率"`
	CallTime                   int32  `bson:"呼出时长"`
	ZeroBetweenForty           int32  `bson:"0秒~40秒"`
	FortyBetweenSixty          int32  `bson:"40秒~60秒"`
	OneMinuteBetweenTwoMinute  int32  `bson:"1分钟~2分钟"`
	TwoMinuteBetweenFiveMinute int32  `bson:"2分钟~5分钟"`
	FiveMinuteBetweenTenMinute int32  `bson:"5分钟~10分钟"`
	MoreThanTenMinute          int32  `bson:"10分钟以上"`
	MaxCallTime                int32  `bson:"最大通时"`
	MeanCallTime               int32  `bson:"通话均时"`
	CollectionTime             int32  `bson:"采集时间"`
}

type TSRTrafficDetail struct {
	CallTimeTotal      int32   `json:"call_time_total"`       //呼叫总时长
	CallConnectCount   int32   `json:"call_connect_count"`    //呼叫接通次数
	RecordCount        int32   `json:"record_count"`          //记录次数(出勤次数)
	CallCountGrade     float64 `json:"call_count_grade"`      //通次评分
	CallTimeGrade      float64 `json:"call_time_grade"`       //通时评分
	CallTimeDayMean    float64 `json:"call_time_day_mean"`    //日均通时
	CallCountDayMean   float64 `json:"call_count_day_mean"`   //日均通次
	CallCountTimeMean  float64 `json:"call_count_time_mean"`  //次均通时
	CallCountTimeGrade float64 `json:"call_count_time_grade"` //次均通时评分
}

// 获取所有话务记录
func GetAllCallTimeRecord(filter bson.D) ([]TrafficMeasurementSchema, error) {
	collection := gmongo.Collection("话务统计")
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		logging.Error(err)
	}
	if err := cur.Err(); err != nil {
		logging.Error(err)
	}
	var TrafficMeasurementList []TrafficMeasurementSchema

	err = cur.All(context.Background(), &TrafficMeasurementList)
	return TrafficMeasurementList, err
}

//　更新坐席通时参数
func GetTSRCallTimeRecordTotal(matchStage bson.D, CallCountStandard float64, CallCountRate float64, CallTimeStandard float64, CallTimeRate float64, CallCountTimeStandard float64, CallCountTimeRate float64) (*TSRTrafficDetail, error) {
	collection := gmongo.Collection("话务统计")
	groupStage := bson.D{{"$group", bson.D{{"_id", "$小宝工号"},
		{"CallTimeTotal", bson.D{{"$sum", "$呼出时长"}}},
		{"CallConnectCount", bson.D{{"$sum", "$呼出接通数"}}},
		{"RecordCount", bson.D{{"$sum", 1}}},
	}}}
	showInfoCursor, err := collection.Aggregate(context.Background(), mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	var showWithInfo []bson.M
	var CallTimeTotal int32
	var CallConnectCount int32
	var RecordCount int32
	var TSRTrafficObj TSRTrafficDetail
	var CallTimeDayMean float64
	var CallCountDayMean float64
	var CalculationCallTimeDayMean float64
	var CalculationCallCountDayMean float64
	var CallCountTimeMean float64
	var CalculationCountTimeMean float64
	if err = showInfoCursor.All(context.Background(), &showWithInfo); err != nil {
		logging.Error(err)
		return nil, err
	}
	if showWithInfo != nil {
		CallTimeTotalNum, ok := showWithInfo[0]["CallTimeTotal"].(int32)
		if ok {
			CallTimeTotal = CallTimeTotalNum
		}
		CallConnectCountNum, ok := showWithInfo[0]["CallConnectCount"].(int32)
		if ok {
			CallConnectCount = CallConnectCountNum
		}
		RecordCountNum, ok := showWithInfo[0]["RecordCount"].(int32)
		if ok {
			RecordCount = RecordCountNum
		}
		TSRTrafficObj.RecordCount = RecordCount
		TSRTrafficObj.CallTimeTotal = CallTimeTotal
		TSRTrafficObj.CallConnectCount = CallConnectCount
	} else {
		TSRTrafficObj.CallConnectCount = 0
		TSRTrafficObj.CallTimeTotal = 0
		TSRTrafficObj.RecordCount = 0
	}
	//计算坐席通时通次标准得分
	CallTimeDayMean = (float64(CallTimeTotal) / float64(RecordCount)) / 3600
	if math.IsNaN(CallTimeDayMean) {
		CallTimeDayMean = 0
	}
	if CallTimeDayMean > CallTimeStandard {
		CalculationCallTimeDayMean = CallTimeStandard
	} else {
		CalculationCallTimeDayMean = CallTimeDayMean
	}

	CallCountDayMean = float64(CallConnectCount) / float64(RecordCount)
	if math.IsNaN(CallCountDayMean) {
		CallCountDayMean = 0
	}
	if CallCountDayMean > CallCountStandard {
		CalculationCallCountDayMean = CallCountStandard
	} else {
		CalculationCallCountDayMean = CallCountDayMean
	}
	//按照比率计算评分
	TSRTrafficObj.CallTimeGrade = (CalculationCallTimeDayMean / CallTimeStandard) * CallTimeRate * 100
	if math.IsNaN(TSRTrafficObj.CallTimeGrade) {
		TSRTrafficObj.CallTimeGrade = 0
	}
	TSRTrafficObj.CallCountGrade = (CalculationCallCountDayMean / CallCountStandard) * CallCountRate * 100
	if math.IsNaN(TSRTrafficObj.CallCountGrade) {
		TSRTrafficObj.CallCountGrade = 0
	}
	TSRTrafficObj.CallTimeDayMean = CallTimeDayMean
	TSRTrafficObj.CallCountDayMean = CallCountDayMean
	CallCountTimeMean = CallTimeDayMean / CallCountDayMean
	if math.IsNaN(CallCountTimeMean) {
		CallCountTimeMean = 0
	}
	if CallCountTimeMean > CallCountTimeStandard {
		CalculationCountTimeMean = CallCountTimeStandard
	} else {
		CalculationCountTimeMean = CallCountTimeMean
	}
	TSRTrafficObj.CallCountTimeMean = CallCountTimeMean
	TSRTrafficObj.CallCountTimeGrade = (CalculationCountTimeMean * 3600 / CallCountTimeStandard) * CallCountTimeRate * 100
	if math.IsNaN(TSRTrafficObj.CallCountTimeGrade) {
		TSRTrafficObj.CallCountTimeGrade = 0
	}
	return &TSRTrafficObj, nil

}

//获取话务基础统计数据
func GetCallRecordBaseData() (int32, float64, float64, error) {
	collection := gmongo.Collection("话务统计")
	groupStage := bson.D{{"$group", bson.D{{"_id", "Total"},
		{"CallConnectCount", bson.D{{"$sum", "$呼出接通数"}}},
		{"CallCount", bson.D{{"$sum", "$呼出总次数"}}},
		{"CallTime", bson.D{{"$sum", "$呼出时长"}}},
	}}}
	showInfoCursor, err := collection.Aggregate(context.Background(), mongo.Pipeline{groupStage})
	if err != nil {
		logging.Error(err)
		return 0, 0, 0, err
	}
	var showWithInfo []bson.M
	var TotalCallCount int32
	var CallConnectRate float64
	var CallTimeMean float64

	if err = showInfoCursor.All(context.Background(), &showWithInfo); err != nil {
		logging.Error(err)
		return 0, 0, 0, err
	}
	if showWithInfo != nil {
		CallCountNum, ok := showWithInfo[0]["CallCount"].(int32)
		if ok {
			TotalCallCount = CallCountNum
		}
		CallConnectCountNum, ok := showWithInfo[0]["CallConnectCount"].(int32)
		if ok {
			CallConnectRate = float64(CallConnectCountNum) / float64(TotalCallCount)
			if math.IsNaN(CallConnectRate) {
				CallConnectRate = 0
			}
		}
		CallTimeNum, ok := showWithInfo[0]["CallTime"].(int32)
		if ok {
			fmt.Println(CallConnectCountNum, CallTimeNum)
			CallTimeMean = float64(CallTimeNum) / float64(CallConnectCountNum)
			if math.IsNaN(CallTimeMean) {
				CallTimeMean = 0
			}
		}

	}
	return TotalCallCount, CallConnectRate, CallTimeMean, err
}
