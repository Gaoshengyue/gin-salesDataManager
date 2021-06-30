package StatementControllerResponseSchema

import (
	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/schema/ServiceSchema/StatementSchema"
)

// 业绩返回结构体
type AchievementPageResponse struct {
	app.Response
	PageSize int                                     `json:"pageSize"` //翻页参数
	Current  int                                     `json:"current"`  //翻页参数
	Total    int                                     `json:"total"`    //翻页参数
	Data     []StatementSchema.AchievementPageSchema `json:"data"`     //翻页数据
}


// 业绩摘要结构体
type AchievementSummaryResponse struct {
	Data  StatementSchema.AchievementSummarySchema `json:"data"`     //翻页数据
}