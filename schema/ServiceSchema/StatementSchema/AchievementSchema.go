package StatementSchema

type AchievementPageSchema struct {
	UnderwritingTime string `json:"underwriting_time"` //承保时间
	TotalPremium     int32  `json:"total_premium"`     //承保人数 当日/当月
}

type AchievementSummarySchema struct {
	TotalPolicy int32   `json:"total_policy"` //总数
	AvgPremium  float64 `json:"avg_premium"`  //平均保费
	SumPremium  float64 `json:"sum_premium"`  //总保费

}
