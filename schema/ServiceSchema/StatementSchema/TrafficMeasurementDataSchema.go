package StatementSchema

type TrafficMeasurementSchema struct {
	TSRId            int32  `bson:"小宝工号"`  //小宝工号
	TSRName          string `bson:"坐席"`    //坐席
	CallTimeTotal    string `bson:"通话总时长"` //通话总时长
	CallCountTotal   int32  `bson:"呼出总次数"` //呼出总次数
	CallConnectCount int32  `bson:"呼出接通数"` //呼出接通数
	CallConnectRate  string `bson:"呼出接通率"` //呼出接通率
	MaxCallTime      int32  `bson:"最大通时"`  //最大通时
	MeanCallTime     int32  `bson:"通话均时"`  //通话均时
}
