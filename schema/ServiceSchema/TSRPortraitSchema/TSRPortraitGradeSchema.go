package TSRPortraitSchema

import (
	"dolphin/salesManager/pkg/logging"
	"dolphin/salesManager/pkg/util"
	"dolphin/salesManager/schema/ControllerSchema/TSRPortraitControllerSchema/TSRPortraitControllerRequestSchema"
	"go.mongodb.org/mongo-driver/bson"
)

type TSRGradeDetail struct {
	TSRTrafficDetail
	TSRPolicyGradeDetail
	TSRMonthlyDetail
	TSRId          int32   `json:"tsr_id"`           //坐席ID
	TSRName        string  `json:"tsr_name"`         //坐席名称
	TSRStatusGrade float64 `json:"tsr_status_grade"` //坐席状态评分
	TSRSkillGrade  float64 `json:"tsr_skill_grade"`  //坐席技能评分
	TSRFinalGrade  float64 `json:"tsr_grade"`        //坐席最终评分
}

func (TSRGradeObj *TSRGradeDetail) GradeInit(TSRGradeQueryParams TSRPortraitControllerRequestSchema.TSRGradeControllerRequest) error {
	err := TSRGradeObj.GetTSRCallTimeDetail(TSRGradeQueryParams.CallCountStandard,
		TSRGradeQueryParams.CallCountRate, TSRGradeQueryParams.CallTimeStandard,
		TSRGradeQueryParams.CallTimeRate, TSRGradeQueryParams.CallCountTimeStandard,
		TSRGradeQueryParams.CallCountTimeRate, TSRGradeQueryParams.StartTime, TSRGradeQueryParams.EndTime)
	if err != nil {
		logging.Error(err)
		return err
	}
	err = TSRGradeObj.GetTSRPolicyGradeDetail(TSRGradeQueryParams.PolicyPremiumMeanStandard,
		TSRGradeQueryParams.PolicyPremiumMeanRate, TSRGradeQueryParams.PremiumStandard,
		TSRGradeQueryParams.PremiumGradeRate, TSRGradeQueryParams.StartTime, TSRGradeQueryParams.EndTime)
	if err != nil {
		logging.Error(err)
		return err
	}
	err = TSRGradeObj.GetTSRNameListGradeDetail(TSRGradeQueryParams.TSRMonthlyCallGradeStandard, TSRGradeQueryParams.TSRMonthlyCallGradeRate)
	if err != nil {
		logging.Error(err)
		return err
	}
	//计算坐席状态评分
	TSRGradeObj.TSRStatusGrade = TSRGradeObj.CallCountGrade + TSRGradeObj.CallTimeGrade + TSRGradeObj.NameListCallGrade
	//计算坐席技能评分
	TSRGradeObj.TSRSkillGrade = TSRGradeObj.TSRPolicyGrade + TSRGradeObj.TSRPolicyPremiumMeanGrade + TSRGradeObj.CallCountTimeGrade
	//计算坐席总评分
	TSRGradeObj.TSRFinalGrade = TSRGradeObj.TSRStatusGrade + TSRGradeObj.TSRSkillGrade
	return nil
}

//绑定每个员工数据对象通时部分数据
func (TSRGradeObj *TSRGradeDetail) GetTSRCallTimeDetail(CallCountStandard float64, CallCountRate float64, CallTimeStandard float64, CallTimeRate float64, CallCountTimeStandard float64, CallCountTimeRate float64, StartTime string, EndTime string) error {
	matchStage := bson.D{{"$match", bson.D{{"小宝工号", TSRGradeObj.TSRId}, {"日期", bson.D{{"$gte", StartTime}, {"$lte", EndTime}}}}}}
	TSRTrafficObj, err := GetTSRCallTimeRecordTotal(matchStage, CallCountStandard, CallCountRate, CallTimeStandard, CallTimeRate, CallCountTimeStandard, CallCountTimeRate)
	if err != nil {
		logging.Error(err)
		return err
	}
	TSRGradeObj.CallTimeTotal = TSRTrafficObj.CallTimeTotal
	TSRGradeObj.CallConnectCount = TSRTrafficObj.CallConnectCount
	TSRGradeObj.RecordCount = TSRTrafficObj.RecordCount
	TSRGradeObj.CallCountGrade = TSRTrafficObj.CallCountGrade
	TSRGradeObj.CallTimeGrade = TSRTrafficObj.CallTimeGrade
	TSRGradeObj.CallTimeDayMean = TSRTrafficObj.CallTimeDayMean
	TSRGradeObj.CallCountDayMean = TSRTrafficObj.CallCountDayMean
	TSRGradeObj.CallCountTimeMean = TSRTrafficObj.CallCountTimeMean
	TSRGradeObj.CallCountTimeGrade = TSRTrafficObj.CallCountTimeGrade

	return nil

}

//绑定每个员工数据对象获取保单部分评分
func (TSRGradeObj *TSRGradeDetail) GetTSRPolicyGradeDetail(PolicyPremiumMeanStandard float64, PolicyPremiumMeanRate float64, PremiumStandard float64, PremiumGradeRate float64, StartTime string, EndTime string) error {

	matchStage := bson.D{{"$match", bson.D{{"坐席ID", TSRGradeObj.TSRId}, {"承保时间", bson.D{{"$gte", StartTime}, {"$lte", EndTime}}}}}}
	TotalPremium, err := TSRGradeObj.GetTSRPolicyPremiumTotalNumber(StartTime, EndTime)
	if err != nil {
		logging.Error(err)
		return err
	}
	SurrenderPremium, err := TSRGradeObj.GetTSRSurrenderPolicyPremiumTotalNumber()
	if err != nil {
		logging.Error(err)
		return err
	}
	Premium := TotalPremium - SurrenderPremium
	TSRPolicyGradeObj, err := GetTSRPolicyDetailAndGrade(matchStage, PolicyPremiumMeanStandard, PolicyPremiumMeanRate, PremiumStandard, PremiumGradeRate, Premium)
	if err != nil {
		logging.Error(err)
		return err
	}
	TSRGradeObj.PolicyTotalPremium = TSRPolicyGradeObj.PolicyTotalPremium
	TSRGradeObj.PolicyCount = TSRPolicyGradeObj.PolicyCount
	TSRGradeObj.TSRPolicyPremiumMean = TSRPolicyGradeObj.TSRPolicyPremiumMean
	TSRGradeObj.TSRPolicyGrade = TSRPolicyGradeObj.TSRPolicyGrade
	TSRGradeObj.TSRPolicyPremiumMeanGrade = TSRPolicyGradeObj.TSRPolicyPremiumMeanGrade
	return nil
}

//绑定每个员工数据对象获取名单拨打部分评分
func (TSRGradeObj *TSRGradeDetail) GetTSRNameListGradeDetail(TSRMonthlyCallGradeStandard float64, TSRMonthlyCallGradeRate float64) error {

	matchStage := bson.D{{"$match", bson.D{{"TSR小宝工号", TSRGradeObj.TSRId},
		{"采集时间", bson.D{{"$gte", util.GetMonthFirstDay()}}}}}}
	matchStage.Map()
	TSRMonthlyDetailObj, err := GetTSRNameListCallDetail(matchStage, TSRMonthlyCallGradeStandard, TSRMonthlyCallGradeRate)
	if err != nil {
		logging.Error(err)
		return err
	}
	TSRGradeObj.NameListCallCount = TSRMonthlyDetailObj.NameListCallCount
	TSRGradeObj.NameListCustomerCount = TSRMonthlyDetailObj.NameListCustomerCount
	TSRGradeObj.NameListCallCountMean = TSRMonthlyDetailObj.NameListCallCountMean
	TSRGradeObj.NameListCallGrade = TSRMonthlyDetailObj.NameListCallGrade
	return nil
}

//获取坐席保单id列表
func (TSRGradeObj *TSRGradeDetail) GetTSRPolicyIdList(StartTime string, EndTime string) ([]string, error) {
	filter := bson.D{{"坐席ID", TSRGradeObj.TSRId}, {"承保时间", bson.D{{"$gte", StartTime}, {"$lte", EndTime}}}}
	var TSRPolicyIdList []string
	TSRPolicyIdObjList, err := GetTSRPolicyIdList(filter)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	for i := 0; i < len(TSRPolicyIdObjList); i++ {
		TSRPolicyIdList = append(TSRPolicyIdList, TSRPolicyIdObjList[i].PolicyId)
	}
	return TSRPolicyIdList, nil
}

//获取坐席已失效状态下的保单id列表
func (TSRGradeObj *TSRGradeDetail) GetTSRSurrenderPolicyIdList() ([]string, error) {
	startTime := util.GetNowTimeOffset(0, 0, -15)
	endTime := util.GetTodayDateStr()
	PolicyList, err := TSRGradeObj.GetTSRPolicyIdList(startTime, endTime)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	filter := bson.D{{"保单号", bson.D{{"$in", bson.A{PolicyList}}}},
		{"保单状态", bson.D{{"$in", bson.A{"停效", "契撤", "犹退终止", "解约", "退保终止"}}}},
	}
	return GetTSRSurrenderPolicyIdList(filter)
}

//获取坐席成单保费金额
func (TSRGradeObj *TSRGradeDetail) GetTSRPolicyPremiumTotalNumber(StartTime string, EndTime string) (float64, error) {
	matchStage := bson.D{{"$match", bson.D{{"坐席ID", TSRGradeObj.TSRId}, {"承保时间", bson.D{{"$gte", StartTime}, {"$lte", EndTime}}}}}}
	return GetTSRPolicyPremiumTotal(matchStage)
}

//获取坐席当月退单保费金额
func (TSRGradeObj *TSRGradeDetail) GetTSRSurrenderPolicyPremiumTotalNumber() (float64, error) {
	SurrenderPolicyIdList, err := TSRGradeObj.GetTSRSurrenderPolicyIdList()
	if err != nil {
		logging.Error(err)
		return 0, err
	}
	matchStage := bson.D{{"$match", bson.D{{"坐席ID", TSRGradeObj.TSRId},
		{"保单号", bson.D{{"$in", bson.A{SurrenderPolicyIdList}}}},
	}}}
	return GetTSRPolicyPremiumTotal(matchStage)
}
