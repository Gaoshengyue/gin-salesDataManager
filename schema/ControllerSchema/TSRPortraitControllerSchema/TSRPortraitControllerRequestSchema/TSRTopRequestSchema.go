package TSRPortraitControllerRequestSchema

import "dolphin/salesManager/pkg/util"

type TSRTopControllerRequest struct {
	StartTime         string  `json:"start_time"`          //开始时间
	EndTime           string  `json:"end_time"`            //结束时间
	PhoneList         []int64 `json:"phone_list"`          //手机数组
	PremiumsFiveStar  float64 `json:"premiums_five_star"`  //净承保费五星排名前多少
	PremiumsFourStar  float64 `json:"premiums_four_star"`  //净承保费四星排名前多少
	PremiumsThreeStar float64 `json:"premiums_three_star"` //净承保费三星排名前多少
	PremiumsTwoStar   float64 `json:"premiums_two_star"`   //净承保费两星排名前多少
	PremiumsOneStar   float64 `json:"premiums_one_star"`   //净承保费一星排名前多少
	RenewalFiveStar   float64 `json:"renewal_five_star"`   //继续率五星排名前多少
	RenewalFourStar   float64 `json:"renewal_four_star"`   //继续率四星排名前多少
	RenewalThreeStar  float64 `json:"renewal_three_star"`  //继续率三星排名前多少
	RenewalTwoStar    float64 `json:"renewal_two_star"`    //继续率两星排名前多少
	RenewalOneStar    float64 `json:"renewal_one_star"`    //继续率一星排名前多少
	QualityFiveStar   float64 `json:"quality_five_star"`   //质检率五星排名前多少
	QualityFourStar   float64 `json:"quality_four_star"`   //质检率四星排名前多少
	QualityThreeStar  float64 `json:"quality_three_star"`  //质检率三星排名前多少
	QualityTwoStar    float64 `json:"quality_two_star"`    //质检率两星排名前多少
	QualityOneStar    float64 `json:"quality_one_star"`    //质检率一星排名前多少
	ProductType       string  `json:"product_type"`        //产品类型
	OrderPlace        string  `json:"order_place"`         //成单区域
	NameListType      string  `json:"name_list_type"`      //名单类型
	CustomerAgeBottom int64   `json:"customer_age_bottom"` //客户年龄下限
	CustomerAgeTop    int64   `json:"customer_age_top"`    //客户年龄上限
	PageSize          int     `json:"page_size"`           // 翻页参数
	Current           int     `json:"current"`             // 翻页参数
}

func (TSRTopControllerRequestObj *TSRTopControllerRequest) InitDefaultRequest() {
	TSRTopControllerRequestObj.StartTime = util.GetMonthFirstDay()
	TSRTopControllerRequestObj.EndTime = util.GetTodayDateStr()
	TSRTopControllerRequestObj.PremiumsFiveStar = 0.1
	TSRTopControllerRequestObj.PremiumsFourStar = 0.3
	TSRTopControllerRequestObj.PremiumsThreeStar = 0.5
	TSRTopControllerRequestObj.PremiumsTwoStar = 0.8
	TSRTopControllerRequestObj.RenewalFiveStar = 0.1
	TSRTopControllerRequestObj.RenewalFourStar = 0.2
	TSRTopControllerRequestObj.RenewalThreeStar = 0.4
	TSRTopControllerRequestObj.RenewalTwoStar = 0.5
	TSRTopControllerRequestObj.QualityFiveStar = 0.2
	TSRTopControllerRequestObj.QualityFourStar = 0.3
	TSRTopControllerRequestObj.QualityThreeStar = 0.4
	TSRTopControllerRequestObj.QualityTwoStar = 0.5
	TSRTopControllerRequestObj.PhoneList = make([]int64, 0)
	// TSRTopControllerRequestObj.PhoneList = []int64{15138972662, 15838214862}
	TSRTopControllerRequestObj.ProductType = "nil"
	TSRTopControllerRequestObj.OrderPlace = "nil"
	TSRTopControllerRequestObj.NameListType = "nil"
	TSRTopControllerRequestObj.CustomerAgeBottom = 0
	TSRTopControllerRequestObj.CustomerAgeTop = 0
	TSRTopControllerRequestObj.PageSize = 10
	TSRTopControllerRequestObj.Current = 1
}
