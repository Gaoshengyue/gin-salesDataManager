package StatementSchema

type NetPowerPolicySchema struct {
	AnnualizedPremium float64 `bson:"年化保费"` //年化保费
	BatchName         string  `bson:"批次名称"` //批次名称
	ProductName       string  `bson:"产品名称"` //产品名称
	PolicyStatus      string  `bson:"保单状态"` //保单状态
	InsuredAmount     float64 `bson:"保额"`   //保额
	UnderwritingTime  string  `bson:"承保时间"` //承保时间
	TSRId             int32   `bson:"坐席ID"` //坐席ID
	ActiveMonth       int32   `bson:"活动月"`  //活动月
}
