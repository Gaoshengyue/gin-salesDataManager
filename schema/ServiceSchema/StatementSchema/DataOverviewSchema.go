package StatementSchema

//　数据概览基础数据
type DataOverviewBaseDataSchema struct {
	TotalCallCount     int32   `json:"total_call_count"`     //总外呼量
	TotalCustomerCount int32   `json:"total_customer_count"` //总客户量
	CallTimeMean       float64 `json:"call_time_mean"`       //平均外呼时长
	CallConnectRate    float64 `json:"call_connect_rate"`    //接通率
}

//　数据概览营销产品分布数据
type DataOverviewProductDistribution struct {
	ProductName  string `json:"product_name"`  //产品名称
	ProductCount int32  `json:"product_count"` //产品数量
}

//　数据概览订单类型分布
type DataOverviewOrderTypeDistribution struct {
	TypeName  string `json:"type_name"`  //类型名称
	TypeCount int32  `json:"type_count"` //类型数量
}

//　数据概览项目业绩统计  暂空
type DataOverviewAchievement struct {
}
