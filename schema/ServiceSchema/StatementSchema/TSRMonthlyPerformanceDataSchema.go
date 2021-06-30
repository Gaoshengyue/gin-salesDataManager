package StatementSchema

type TSRMonthlyPerformance struct {
	ActiveMonth        float64 `bson:"活动月"`     //活动月
	TSRName            string  `bson:"TSR姓名"`   //TSR姓名
	TSRId              int32   `bson:"TSR小宝工号"` //TSR小宝工号
	CallCount          int32   `bson:"拨打次数"`    //拨打次数
	ConnectCount       int32   `bson:"接通次数"`    //接通次数
	FirstConnectCount  int32   `bson:"首拨接通量"`   //首拨接通量
	SecondConnectCount int32   `bson:"二次接通量"`   //二次接通量
	ContactListCount   int32   `bson:"接触名单量"`   //接触名单量
}
