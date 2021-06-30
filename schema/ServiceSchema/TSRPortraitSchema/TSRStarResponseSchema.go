package TSRPortraitSchema

type TsrStarResponseSchema struct {
	TSRId         int32  `json:"tsrId"`         //员工ID
	TSRName       string `json:"tsrName"`       //员工姓名
	PremiumsStar  int    `json:"premiumsStar"`  //业绩评星
	PolicyStar    int    `json:"policyStar"`    //续期评星
	QualifiedStar int    `json:"qualifiedStar"` //质检评星
	TotalStar     int    `json:"totalStar"`     //总星数
	Top           int    `json:"Top"`           //排行
}

type TsrStarPageResponseSchema struct {
	PageSize int                     `json:"pageSize"` //翻页参数
	Current  int                     `json:"current"`  //翻页参数
	Total    int                     `json:"total"`    //翻页参数
	TSRList  []TsrStarResponseSchema `json:"tsrList"`  //翻页数据
}

// 续期评星
type NetPowerPolicySortSchema struct {
	TSRName        string
	TSRId          int32
	Star           int
	OrderReceiving int32
	OrderForm      int32
}

// 质检评星
type QualityResponseSchema struct {
	TSRId     int32 `json:"tsrId"`     //员工ID
	Star      int   `json:"star"`      //评星
	Qualified int32 `json:"qualified"` //合格的
	Total     int32 `json:"total"`     //总数
}
