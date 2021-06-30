package TSRPortraitSchema

import (
	"dolphin/salesManager/pkg/gmongo"
	"dolphin/salesManager/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"math"
)

type TSRMonthlyPerformance struct {
	ActiveMonth        float64 `bson:"活动月"`
	BatchName          string  `bson:"批次名称"`
	City               string  `bson:"城市"`
	ListType           string  `bson:"名单类型"`
	Place              string  `bson:"区部"`
	Group              string  `bson:"小组"`
	TSRName            string  `bson:"TSR姓名"`
	TSRId              int32   `bson:"TSR小宝工号"`
	ListDistribution   int32   `bson:"名单分配量"`
	CallListCount      int32   `bson:"拨打量"`
	CallCount          int32   `bson:"拨打次数"`
	ConnectCount       int32   `bson:"接通次数"`
	CallTime           int32   `bson:"通话时长"`
	FirstConnectCount  int32   `bson:"首拨接通量"`
	SecondConnectCount int32   `bson:"二次接通量"`
	ContactListCount   int32   `bson:"接触名单量"`
	CollectionTime     string  `bson:"采集时间"`
}

type TSRMonthlyDetail struct {
	NameListCallCount     int32   `json:"name_list_call_count"`      //名单分配量
	NameListCustomerCount int32   `json:"name_list_customer_count"`  //名单客户拨打次数
	NameListCallCountMean float64 `json:"name_list_call_count_mean"` //批次名单平均拨打次数
	NameListCallGrade     float64 `json:"name_list_call_grade"`      //名单拨打次数评分
}

//　更新坐席通时参数
func GetTSRNameListCallDetail(matchStage bson.D, TSRMonthlyCallGradeStandard float64, TSRMonthlyCallGradeRate float64) (*TSRMonthlyDetail, error) {
	collection := gmongo.Collection("签单月拨打表现")
	groupStage := bson.D{{"$group" +
		"", bson.D{{"_id", "$TSR小宝工号"},
		{"NameListCount", bson.D{{"$sum", "$名单分配量"}}},
		{"BatchCallCount", bson.D{{"$sum", "$拨打次数"}}},
	}}}

	showInfoCursor, err := collection.Aggregate(context.Background(), mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	var showWithInfo []bson.M
	var NameListCount int32
	var BatchCallCount int32
	var TSRMonthlyDetailObj TSRMonthlyDetail
	var NameListCallCountMean float64
	var CalculationNameListCallCountMean float64
	if err = showInfoCursor.All(context.Background(), &showWithInfo); err != nil {
		logging.Error(err)
		return nil, err
	}
	if showWithInfo != nil {
		NameListCountNum, ok := showWithInfo[0]["NameListCount"].(int32)
		if ok {
			NameListCount = NameListCountNum
		}
		BatchCallCountNum, ok := showWithInfo[0]["BatchCallCount"].(int32)
		if ok {
			BatchCallCount = BatchCallCountNum
		}
		TSRMonthlyDetailObj.NameListCallCount = NameListCount
		TSRMonthlyDetailObj.NameListCustomerCount = BatchCallCount
	} else {
		TSRMonthlyDetailObj.NameListCallCount = 0
		TSRMonthlyDetailObj.NameListCustomerCount = 0
	}
	//计算坐席通时通次标准得分
	NameListCallCountMean = float64(BatchCallCount) / float64(NameListCount)
	if math.IsNaN(NameListCallCountMean) {
		NameListCallCountMean = 0
	}
	TSRMonthlyDetailObj.NameListCallCountMean = NameListCallCountMean

	if NameListCallCountMean > TSRMonthlyCallGradeStandard {
		CalculationNameListCallCountMean = TSRMonthlyCallGradeStandard
	} else {
		CalculationNameListCallCountMean = NameListCallCountMean
	}

	TSRMonthlyDetailObj.NameListCallGrade = (CalculationNameListCallCountMean / TSRMonthlyCallGradeStandard) * TSRMonthlyCallGradeRate
	if math.IsNaN(TSRMonthlyDetailObj.NameListCallGrade) {
		TSRMonthlyDetailObj.NameListCallGrade = 0
	}
	return &TSRMonthlyDetailObj, nil

}
