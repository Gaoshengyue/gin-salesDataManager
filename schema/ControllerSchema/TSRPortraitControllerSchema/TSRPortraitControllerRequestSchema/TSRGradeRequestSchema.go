package TSRPortraitControllerRequestSchema

import "dolphin/salesManager/pkg/util"

type TSRGradeControllerRequest struct {
	StartTime                   string  `json:"start_time"`                      //开始时间
	EndTime                     string  `json:"end_time"`                        //结束时间
	PhoneList                   []int64 `json:"phone_list"`                      //手机数组
	PolicyPremiumMeanStandard   float64 `json:"policy_premium_mean_standard"`    //平均保费标准
	PolicyPremiumMeanRate       float64 `json:"policy_premium_mean_rate"`        //平均保费比率
	PremiumStandard             float64 `json:"premium_standard"`                //净承保费标准
	PremiumGradeRate            float64 `json:"premium_grade_rate"`              //净承保费比率
	CallCountStandard           float64 `json:"call_count_standard"`             //呼叫次数标准
	CallCountRate               float64 `json:"call_count_rate"`                 //呼叫次数比率
	CallTimeStandard            float64 `json:"call_time_standard"`              //呼叫通时标准
	CallTimeRate                float64 `json:"call_time_rate"`                  //呼叫通时比率
	CallCountTimeStandard       float64 `json:"call_count_time_standard"`        //次均通时标准
	CallCountTimeRate           float64 `json:"call_count_time_rate"`            //次均通时比率
	TSRMonthlyCallGradeStandard float64 `json:"tsr_monthly_call_grade_standard"` //签单月拨打表现标准
	TSRMonthlyCallGradeRate     float64 `json:"tsr_monthly_call_grade_rate"`     //签单月拨打表现比率
}

func (TSRGradeControllerRequestObj *TSRGradeControllerRequest) InitDefaultRequest() {
	TSRGradeControllerRequestObj.PolicyPremiumMeanStandard = 3500
	TSRGradeControllerRequestObj.PolicyPremiumMeanRate = 0.2
	TSRGradeControllerRequestObj.PremiumStandard = 15000
	TSRGradeControllerRequestObj.PremiumGradeRate = 0.7
	TSRGradeControllerRequestObj.CallCountStandard = 40
	TSRGradeControllerRequestObj.CallCountRate = 0.2
	TSRGradeControllerRequestObj.CallTimeStandard = 3
	TSRGradeControllerRequestObj.CallTimeRate = 0.5
	TSRGradeControllerRequestObj.CallCountTimeStandard = 300
	TSRGradeControllerRequestObj.CallCountTimeRate = 0.1
	_, _, day := util.GetToday()
	if day <= 10 {
		TSRGradeControllerRequestObj.TSRMonthlyCallGradeStandard = 2
	} else if day > 10 && day <= 20 {
		TSRGradeControllerRequestObj.TSRMonthlyCallGradeStandard = 2.5
	} else {
		TSRGradeControllerRequestObj.TSRMonthlyCallGradeStandard = 3.5
	}

	TSRGradeControllerRequestObj.TSRMonthlyCallGradeRate = 0.3
	TSRGradeControllerRequestObj.StartTime = util.GetMonthFirstDay()
	TSRGradeControllerRequestObj.EndTime = util.GetTodayDateStr()
	TSRGradeControllerRequestObj.PhoneList = make([]int64, 0)
}
