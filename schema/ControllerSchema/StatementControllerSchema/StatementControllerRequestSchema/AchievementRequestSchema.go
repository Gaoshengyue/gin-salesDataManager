package StatementControllerRequestSchema

// 业绩请求结构体
type AchievementControllerRequest struct {
	StartTime string `json:"start_time"` //开始时间
	EndTime   string `json:"end_time"`   //结束时间
	PageSize  int    `json:"page_size"`  //翻页参数
	Current   int    `json:"current"`    //翻页参数
}
