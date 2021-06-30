package TSRPortraitSchema

import (
	"dolphin/salesManager/pkg/gmongo"
	"dolphin/salesManager/pkg/logging"
	"dolphin/salesManager/schema/ServiceSchema/StatementSchema"
	"fmt"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type NetPowerPolicyIdList struct {
	PolicyId string `bson:"保单号"`
}
type NetPowerRenewalSchema struct {
	PolicyId          string  `bson:"保单号"`
	TSRName           string  `bson:"坐席"`
	TSRId             int32   `bson:"坐席ID"`
	AnnualizedPremium float64 `bson:"年化保费"`
	InitialPremium    float64 `bson:"首期保费"`
	UnderwritingTime  string  `bson:"承保时间"`
}
type NetPowerPolicySchema struct {
	NetPowerPolicyIdList
	AnnualizedPremium float64 `bson:"年化保费"`
	Age               float64 `bson:"年龄"`
	BatchName         string  `bson:"批次名称"`
	Province          string  `bson:"省份"`
	OrderId           string  `bson:"订单号"`
	CustomerCode      string  `bson:"客户编码"`
	CustomerName      string  `bson:"客户姓名"`
	ProductName       string  `bson:"产品名称"`
	PolicyHolder      string  `bson:"投保人"`
	PaymentFrequency  string  `bson:"缴费频率"`
	InitialPremium    float64 `bson:"首期保费"`
	PolicyStatus      string  `bson:"保单状态"`
	InsuredAmount     float64 `bson:"保额"`
	PaymentPeriod     string  `bson:"缴费年限"`
	PremiumPeriod     string  `bson:"保费年限"`
	SubmitTime        string  `bson:"提交时间"`
	UnderwritingTime  string  `bson:"承保时间"`
	TSRName           string  `bson:"坐席"`
	TSRId             int32   `bson:"坐席ID"`
	Department        string  `bson:"部门"`
	Place             string  `bson:"区部"`
	Group             string  `bson:"团队"`
	ActiveMonth       int32   `bson:"活动月"`
	ListType          string  `bson:"名单大类"`
	ListGroup         string  `bson:"名单系"`
	CollectionTime    string  `bson:"采集时间"`
}

type TSRPolicyGradeDetail struct {
	PolicyTotalPremium        float64 `json:"policy_total_premium"`          //保单总额
	PolicyCount               int32   `json:"policy_count"`                  //保单量
	TSRPolicyPremiumMean      float64 `json:"tsr_policy_premium_mean"`       //坐席平均保单保费
	TSRPolicyGrade            float64 `json:"tsr_policy_grade"`              //坐席月净保费评分
	TSRPolicyPremiumMeanGrade float64 `json:"tsr_policy_premium_mean_grade"` //坐席件均评分
}

//创建TSR保单评分部分返回结构数据
func GetTSRPolicyDetailAndGrade(matchStage bson.D, PolicyPremiumMeanStandard float64, PolicyPremiumMeanRate float64, PremiumStandard float64, PremiumGradeRate float64, Premium float64) (*TSRPolicyGradeDetail, error) {
	/*
		PolicyPremiumMeanStandard 件均标准　float64
		PolicyPremiumMeanRate 件均评分占比   float64
		PremiumStandard	承包保费标准	float64
		PremiumRate	承包保费评分占比	float64
	*/
	collection := gmongo.Collection("网电保单")
	groupStage := bson.D{{"$group", bson.D{{"_id", "$坐席ID"}, {"policyCount", bson.D{{"$sum", 1}}}}}}
	showInfoCursor, err := collection.Aggregate(context.Background(), mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	//获取坐席月净保费、成单件数等
	var showWithInfo []bson.M
	var policyCount int32
	var PolicyPremiumMean float64
	if err = showInfoCursor.All(context.Background(), &showWithInfo); err != nil {
		logging.Error(err)
		return nil, err
	}
	if showWithInfo != nil {
		countNum, ok := showWithInfo[0]["policyCount"].(int32)
		if ok {
			policyCount = countNum
		}

	} else {
		policyCount = 0

	}
	//获取坐席成单件均保费
	PolicyPremiumMean = Premium / float64(policyCount)
	if math.IsNaN(PolicyPremiumMean) {
		PolicyPremiumMean = 0
	}
	var CalculationPremium float64
	var CalculationPolicyPremiumMean float64
	//判断坐席成单月净保费与件均是否超过标准。如超过则截止于标准数值
	if Premium > PremiumStandard {
		CalculationPremium = PremiumStandard
	} else {
		CalculationPremium = Premium
	}
	if PolicyPremiumMean > PolicyPremiumMeanStandard {
		CalculationPolicyPremiumMean = PolicyPremiumMeanStandard

	} else {
		CalculationPolicyPremiumMean = PolicyPremiumMean
	}
	//返回TSR月净保费评分、件均评分以及相关内容
	TSRGradeDetail := TSRPolicyGradeDetail{
		PolicyTotalPremium:        Premium,
		PolicyCount:               policyCount,
		TSRPolicyPremiumMean:      PolicyPremiumMean,
		TSRPolicyGrade:            (CalculationPremium / PremiumStandard) * PremiumGradeRate * 100,
		TSRPolicyPremiumMeanGrade: (CalculationPolicyPremiumMean / PolicyPremiumMeanStandard) * PolicyPremiumMeanRate * 100,
	}
	if math.IsNaN(TSRGradeDetail.TSRPolicyGrade) {
		TSRGradeDetail.TSRPolicyGrade = 0
	}
	if math.IsNaN(TSRGradeDetail.TSRPolicyPremiumMeanGrade) {

		TSRGradeDetail.TSRPolicyPremiumMeanGrade = 0
	}
	return &TSRGradeDetail, nil
}

//获取坐席保单承保总额
func GetTSRPolicyPremiumTotal(matchStage bson.D) (float64, error) {
	collection := gmongo.Collection("网电保单")
	groupStage := bson.D{{"$group", bson.D{{"_id", "$坐席ID"}, {"totalPremium", bson.D{{"$sum", "$年化保费"}}}}}}
	showInfoCursor, err := collection.Aggregate(context.Background(), mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		logging.Error(err)
		return 0, err
	}
	var totalPremium float64
	var showWithInfo []bson.M
	if err = showInfoCursor.All(context.Background(), &showWithInfo); err != nil {
		logging.Error(err)
		return 0, err
	}
	if showWithInfo != nil {
		totalNum, ok := showWithInfo[0]["totalPremium"].(float64)
		fmt.Println(totalNum)
		if ok {
			totalPremium = totalNum
		}
	}
	return totalPremium, err

}

//获取坐席保单id数组
func GetTSRPolicyIdList(filter bson.D) ([]NetPowerPolicySchema, error) {
	collection := gmongo.Collection("网电保单")
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	if err := cur.Err(); err != nil {
		logging.Error(err)
		return nil, err
	}
	var TSRPolicyIdObjList []NetPowerPolicySchema
	err = cur.All(context.Background(), &TSRPolicyIdObjList)

	return TSRPolicyIdObjList, err
}

//获取坐席保单Schema数组
func GetTSRRenewalSchemaList(filter bson.D) ([]NetPowerRenewalSchema, error) {
	collection := gmongo.Collection("网电保单")
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	if err := cur.Err(); err != nil {
		logging.Error(err)
		return nil, err
	}
	var TSRPolicyIdObjList []NetPowerRenewalSchema
	err = cur.All(context.Background(), &TSRPolicyIdObjList)

	return TSRPolicyIdObjList, err
}

//GetUserDetailTSRIds 获取所有员工的id信息
func GetPolicyIds(filter bson.D) ([]string, error) {
	Policys, err := GetTSRPolicyIdList(filter)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	var PolicyIds []string
	for _, Policy := range Policys {
		PolicyIds = append(PolicyIds, Policy.PolicyId)
	}
	return PolicyIds, nil
}

// 获取所有总成单量
func GetTotalPolicyCount() (int32, error) {
	collection := gmongo.Collection("网电保单")
	groupStage := bson.D{{"$group", bson.D{{"_id", "Total"},
		{"TotalPolicyCount", bson.D{{"$sum", 1}}},
	}}}
	showInfoCursor, err := collection.Aggregate(context.Background(), mongo.Pipeline{groupStage})
	if err != nil {
		logging.Error(err)
		return 0, err
	}
	var showWithInfo []bson.M
	var TotalPolicyCount int32
	if err = showInfoCursor.All(context.Background(), &showWithInfo); err != nil {
		logging.Error(err)
		return 0, err
	}
	if showWithInfo != nil {
		TotalPolicyCountNum, ok := showWithInfo[0]["TotalPolicyCount"].(int32)
		if ok {
			TotalPolicyCount = TotalPolicyCountNum
		}
	}
	return TotalPolicyCount, nil
}

// 按月查询 业绩数据 翻页
func AchievementMoonPage(matchQuery []bson.M) ([]StatementSchema.AchievementPageSchema, error) {

	collection := gmongo.Collection("网电保单")

	groupQuery := []bson.M{
		{"$group": bson.M{"_id": bson.M{"$substr": bson.A{"$承保时间", 0, 10}}, "totalPremium": bson.M{"$sum": 1}}},
		{"$sort": bson.M{"_id": -1}},
		{"$limit": 10}, {"$skip": 1},
	}
	query := append(matchQuery, groupQuery...)
	showInfoCursor, err := collection.Aggregate(context.Background(), query)
	var showWithInfo []bson.M
	if err = showInfoCursor.All(context.Background(), &showWithInfo); err != nil {
		logging.Error(err)
		return nil, err
	}
	var achievementSchemas []StatementSchema.AchievementPageSchema
	if showWithInfo != nil {
		for _, info := range showWithInfo {
			achievementSchema := StatementSchema.AchievementPageSchema{UnderwritingTime: info["_id"].(string), TotalPremium: info["totalPremium"].(int32)}

			achievementSchemas = append(achievementSchemas, achievementSchema)
		}
	}
	return achievementSchemas, err
}

// 业绩汇总
func AchievementSummary(matchQuery []bson.M) (StatementSchema.AchievementSummarySchema, error) {

	collection := gmongo.Collection("网电保单")

	groupQuery := []bson.M{
		{"$group": bson.M{"_id": nil, "totalPolicy": bson.M{"$sum": 1}, "avgPremium": bson.M{"$avg": "$保额"}, "sumPremium": bson.M{"$sum": "$保额"}}},
	}
	query := append(matchQuery, groupQuery...)
	showInfoCursor, err := collection.Aggregate(context.Background(), query)
	var showWithInfo []bson.M
	var achievementSchema StatementSchema.AchievementSummarySchema
	if err = showInfoCursor.All(context.Background(), &showWithInfo); err != nil {
		logging.Error(err)
		return achievementSchema, err
	}

	if showWithInfo != nil {
		achievementSchema = StatementSchema.AchievementSummarySchema{
			TotalPolicy: showWithInfo[0]["totalPolicy"].(int32),
			AvgPremium:  showWithInfo[0]["avgPremium"].(float64),
			SumPremium:  showWithInfo[0]["sumPremium"].(float64),
		}
	}
	return achievementSchema, err
}
