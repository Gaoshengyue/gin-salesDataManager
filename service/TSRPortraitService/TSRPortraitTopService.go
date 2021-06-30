package TSRPortraitService

import (
	"dolphin/salesManager/pkg/logging"
	"dolphin/salesManager/pkg/util"
	"dolphin/salesManager/schema/ControllerSchema/TSRPortraitControllerSchema/TSRPortraitControllerRequestSchema"
	"dolphin/salesManager/schema/ServiceSchema/TSRPortraitSchema"
	"fmt"
	"sort"
	"time"

	"github.com/shopspring/decimal"

	"go.mongodb.org/mongo-driver/bson"
)

type NetPowerPolicySortSchema = TSRPortraitSchema.NetPowerPolicySortSchema
type QualityResponseSchema = TSRPortraitSchema.QualityResponseSchema
type TsrStarResponseSchema = TSRPortraitSchema.TsrStarResponseSchema
type TSRTopControllerRequest = TSRPortraitControllerRequestSchema.TSRTopControllerRequest
type TsrStarPageResponseSchema = TSRPortraitSchema.TsrStarPageResponseSchema

func GetTSRTop(TSRTopQueryParams TSRTopControllerRequest) (TsrStarPageResponseSchema, error) {
	TsrDetailList, err := TSRPortraitSchema.GetUserDetailAll(bson.D{{"在职状态", "在职"}}) //{"联系电话1", bson.D{{"$in", TSRTopQueryParams.PhoneList}}},

	if TsrDetailList == nil || err != nil {
		logging.Error(err)
		return TsrStarPageResponseSchema{PageSize: TSRTopQueryParams.PageSize, Current: TSRTopQueryParams.Current, Total: 0}, err
	}
	DictTSR := make(map[int32]TsrStarResponseSchema)
	for _, TsrDetail := range TsrDetailList {
		if _, ok := DictTSR[TsrDetail.TSRId]; !ok {
			DictTSR[TsrDetail.TSRId] = TsrStarResponseSchema{TSRId: TsrDetail.TSRId, TSRName: TsrDetail.TSRName, PremiumsStar: 1, PolicyStar: 1, QualifiedStar: 1}
		}
	}
	// 净承保保费
	err = GetTSRNetPremiums(TsrDetailList, DictTSR, TSRTopQueryParams)
	if err != nil {
		return TsrStarPageResponseSchema{PageSize: TSRTopQueryParams.PageSize, Current: TSRTopQueryParams.Current, Total: 0}, err
	}
	// 坐席继续率
	err = GetTSRRenewal(TsrDetailList, DictTSR, TSRTopQueryParams)
	if err != nil {
		return TsrStarPageResponseSchema{PageSize: TSRTopQueryParams.PageSize, Current: TSRTopQueryParams.Current, Total: 0}, err
	}
	// 品质
	err = GetTSRQuality(TsrDetailList, DictTSR, TSRTopQueryParams)
	if err != nil {
		return TsrStarPageResponseSchema{PageSize: TSRTopQueryParams.PageSize, Current: TSRTopQueryParams.Current, Total: 0}, err
	}
	var TsrStars []TsrStarResponseSchema
	for _, TsrStar := range DictTSR {
		TsrStar.TotalStar = TsrStar.PolicyStar + TsrStar.PremiumsStar + TsrStar.QualifiedStar
		TsrStars = append(TsrStars, TsrStar)
	}
	if TsrStars == nil {
		return TsrStarPageResponseSchema{PageSize: TSRTopQueryParams.PageSize, Current: TSRTopQueryParams.Current, Total: 0}, nil
	}

	SortTsrStars(TsrStars)
	return TotalTsrStars(TsrStars, TSRTopQueryParams.Current, TSRTopQueryParams.PageSize), nil

}

// 排序
func SortTsrStars(TsrStars []TsrStarResponseSchema) {
	sort.Slice(TsrStars, func(i, j int) bool { // desc
		return TsrStars[i].PolicyStar+TsrStars[i].PremiumsStar+TsrStars[i].QualifiedStar > TsrStars[j].PolicyStar+TsrStars[j].PremiumsStar+TsrStars[j].QualifiedStar
	})
}

// 数据整理
func TotalTsrStars(TsrStars []TsrStarResponseSchema, Current int, PageSize int) TsrStarPageResponseSchema {
	for i := 0; i < len(TsrStars); i++ {
		TsrStars[i].TotalStar = TsrStars[i].PolicyStar + TsrStars[i].PremiumsStar + TsrStars[i].QualifiedStar
		TsrStars[i].Top = i + 1
	}
	var limitSize int32
	limitSize = int32((Current-1)*PageSize + PageSize)
	if len(TsrStars) < PageSize {
		limitSize = int32(len(TsrStars))
	}
	if int32(len(TsrStars)) < limitSize {
		limitSize = int32(len(TsrStars))
	}
	var TSRList = TsrStars[(Current-1)*PageSize : limitSize]
	return TsrStarPageResponseSchema{PageSize: PageSize, Current: Current, Total: len(TsrStars), TSRList: TSRList}
}

//
func GetTSRNetPremiums(TsrDetailList []TSRPortraitSchema.UserDetail, DictTSR map[int32]TsrStarResponseSchema, TSRTopQueryParams TSRTopControllerRequest) error {
	// 获取每个tsr的用户列表 是否含有退保记录
	for i := 0; i < len(TsrDetailList); i++ {
		// 先按单条记录操作 性能不够时候再批量
		// startTime := util.GetNowTimeOffset(0, 0, -15)
		// endTime := util.GetTodayDateStr()
		PolicyTotalPremium, err := TsrDetailList[i].GetTSRPolicyPremiumTotalNumber(TSRTopQueryParams.StartTime, TSRTopQueryParams.EndTime,
			TSRTopQueryParams.ProductType, TSRTopQueryParams.OrderPlace,
			TSRTopQueryParams.NameListType, TSRTopQueryParams.CustomerAgeTop, TSRTopQueryParams.CustomerAgeBottom)

		SurrenderPolicyPremium, err := TsrDetailList[i].GetTSRSurrenderPolicyPremiumTotalNumber()

		if err != nil {
			logging.Error(err)
			return err
		}
		// 如果存在 退保 放入字典中
		TsrDetailList[i].AnnualizedPremium = PolicyTotalPremium - SurrenderPolicyPremium
		// fmt.Println(TsrDetailList[i].TSRName, TsrDetailList[i].AnnualizedPremium)
	}
	SortGetTSRNetPremiums(TsrDetailList)
	StarTSRNetPremiums(TsrDetailList, TSRTopQueryParams.PremiumsFiveStar,
		TSRTopQueryParams.PremiumsFourStar, TSRTopQueryParams.PremiumsThreeStar,
		TSRTopQueryParams.PremiumsTwoStar, TSRTopQueryParams.PremiumsOneStar)

	for _, TsrDetail := range TsrDetailList {
		if v1, ok := DictTSR[TsrDetail.TSRId]; ok {
			v1.PremiumsStar = TsrDetail.Star
			DictTSR[TsrDetail.TSRId] = v1
		}
	}

	return nil
}

// 排序
func SortGetTSRNetPremiums(user []TSRPortraitSchema.UserDetail) {
	sort.Slice(user, func(i, j int) bool { // desc
		return user[i].AnnualizedPremium > user[j].AnnualizedPremium
	})
}

// 显示星
func StarTSRNetPremiums(users []TSRPortraitSchema.UserDetail, PremiumsFiveStar float64, PremiumsFourStar float64,
	PremiumsThreeStar float64, PremiumsTwoStar float64, PremiumsOneStar float64) {
	var max int = 5
	//如果不够无位
	if len(users) < max {
		for i := 0; i < len(users); i++ {
			// 如果 0保单 1星
			if users[i].AnnualizedPremium == 0 {
				users[i].Star = 1
			} else {
				users[i].Star = max - i
			}

		}
	} else {
		for i := 0; i < len(users); i++ {
			// 如果 0保单 1星
			if users[i].AnnualizedPremium == 0 {
				users[i].Star = 1
			} else if float32(i) <= float32(len(users))*float32(PremiumsFiveStar) {
				users[i].Star = 5
			} else if float32(i) <= float32(len(users))*float32(PremiumsFourStar) {
				users[i].Star = 4
			} else if float32(i) <= float32(len(users))*float32(PremiumsThreeStar) {
				users[i].Star = 3
			} else if float32(i) <= float32(len(users))*float32(PremiumsTwoStar) {
				users[i].Star = 2
			} else {
				users[i].Star = 1
			}

		}
	}
	// 如果星相同 与上排名同星
	for i := 0; i < len(users); i++ {
		if i > 0 {
			if users[i].AnnualizedPremium == users[i-1].AnnualizedPremium {
				users[i].Star = users[i-1].Star
			}

		}
		//fmt.Println(users[i].TSRName, users[i].AnnualizedPremium, users[i].Star)
	}
}

// 坐席继续率
func GetTSRRenewal(TsrDetailList []TSRPortraitSchema.UserDetail, DictTSR map[int32]TsrStarResponseSchema, TSRGradeQueryParams TSRTopControllerRequest) error {
	//filter := bson.D{{"在职状态", "在职"}}
	//_, err := TSRPortraitSchema.GetUserDetailTSRIds(filter)
	//if err != nil {
	//	logging.Error(err)
	//	return nil, err
	//}
	var TsrIds []int32
	for _, TsrDetail := range TsrDetailList {
		TsrIds = append(TsrIds, TsrDetail.TSRId)
	}
	// 如有其他条件另拼写
	var filter1 bson.D
	filter1 = append(filter1, bson.E{Key: "坐席ID", Value: bson.D{{"$in", TsrIds}}})

	NetPowerRenewalSchemas, err := TSRPortraitSchema.GetTSRRenewalSchemaList(filter1)

	if err != nil {
		logging.Error(err)
		return err
	}
	NetPowerPolicySortSchema, err := NetPowerPolicysStatistics(NetPowerRenewalSchemas)
	if err != nil {
		logging.Error(err)
		return err
	}
	SortRenewal(NetPowerPolicySortSchema)
	StarRenewal(NetPowerPolicySortSchema, TSRGradeQueryParams.RenewalFiveStar, TSRGradeQueryParams.RenewalFourStar,
		TSRGradeQueryParams.RenewalThreeStar, TSRGradeQueryParams.RenewalTwoStar, TSRGradeQueryParams.RenewalOneStar)
	for _, NetPowerPolicy := range NetPowerPolicySortSchema {
		if v1, ok := DictTSR[NetPowerPolicy.TSRId]; ok {
			v1.PolicyStar = NetPowerPolicy.Star
			DictTSR[NetPowerPolicy.TSRId] = v1
		}
	}
	return nil
}

//
func NetPowerPolicysStatistics(NetPowerPolicySchemas []TSRPortraitSchema.NetPowerRenewalSchema) ([]NetPowerPolicySortSchema, error) {
	DictNetOrder := make(map[int32]NetPowerPolicySortSchema)
	var NetPowerPolicySortList []TSRPortraitSchema.NetPowerPolicySortSchema
	for _, NetPowerPolicy := range NetPowerPolicySchemas {
		var NetPowerPolicySort NetPowerPolicySortSchema
		if v1, ok := DictNetOrder[NetPowerPolicy.TSRId]; ok {
			NetPowerPolicySort = v1
		} else {
			NetPowerPolicySort = NetPowerPolicySortSchema{}
			NetPowerPolicySort.TSRId = NetPowerPolicy.TSRId
			NetPowerPolicySort.TSRName = NetPowerPolicy.TSRName

		}
		count, err := ComputePolicy(&NetPowerPolicy)
		if err != nil {
			logging.Error(err)
			return nil, err
		}

		NetPowerPolicySort.OrderForm = NetPowerPolicySort.OrderForm + count
		NetPowerPolicySort.OrderReceiving = NetPowerPolicySort.OrderReceiving + 1

		DictNetOrder[NetPowerPolicy.TSRId] = NetPowerPolicySort
		//年华保费
	}
	fmt.Println(DictNetOrder)
	for _, value := range DictNetOrder {
		NetPowerPolicySortList = append(NetPowerPolicySortList, value)
	}
	return NetPowerPolicySortList, nil
}

// ComputePolicy
func ComputePolicy(NetPowerPolicy *TSRPortraitSchema.NetPowerRenewalSchema) (int32, error) {
	// 获取月平均保费

	MonthPremium := decimal.NewFromFloat(NetPowerPolicy.AnnualizedPremium).Div(decimal.NewFromFloat(12))
	// 完成缴费金额/ 已缴费月份
	CompleteMonth := decimal.NewFromFloat(NetPowerPolicy.InitialPremium).Div(MonthPremium)
	// 承保日期
	UnderwritingTime, err := util.ParseDateTime(NetPowerPolicy.UnderwritingTime)

	if err != nil {
		logging.Error(err)
		//return nil, err
	}
	// 系统当前时间 - 承保时间
	DiffDay := util.GetDiffDay(time.Now(), UnderwritingTime)
	//fmt.Println("MonthPremium",NetPowerPolicy.AnnualizedPremium,MonthPremium,CompleteMonth,int64(DiffDay)-CompleteMonth.IntPart())
	//相差时间
	if int64(DiffDay)-CompleteMonth.IntPart() > 75 {
		return 1, nil
	} else {
		return 0, nil
	}
}

// 排序
func SortRenewal(user []NetPowerPolicySortSchema) {
	sort.Slice(user, func(i, j int) bool { // desc
		return user[i].OrderForm/user[i].OrderReceiving > user[j].OrderForm/user[j].OrderReceiving
	})
}

// 显示星
func StarRenewal(users []NetPowerPolicySortSchema, RenewalFiveStar float64, RenewalFourStar float64,
	RenewalThreeStar float64, RenewalTwoStar float64, RenewalOneStar float64) {
	var max int = 5
	//如果不够无位
	if len(users) < max {
		for i := 0; i < len(users); i++ {
			users[i].Star = max - i
		}
	} else {
		for i := 0; i < len(users); i++ {
			// 如果 0保单 1星
			if float32(i) <= float32(len(users))*float32(RenewalFiveStar) {
				users[i].Star = 5
			} else if float32(i) <= float32(len(users))*float32(RenewalFourStar) {
				users[i].Star = 4
			} else if float32(i) <= float32(len(users))*float32(RenewalThreeStar) {
				users[i].Star = 3
			} else if float32(i) <= float32(len(users))*float32(RenewalTwoStar) {
				users[i].Star = 2
			} else {
				users[i].Star = 1
			}
		}
	}
	// 如果星相同 与上排名同星
	for i := 0; i < len(users); i++ {
		if i > 0 {
			if users[i].OrderForm/users[i].OrderReceiving == users[i-1].OrderForm/users[i-1].OrderReceiving {
				users[i].Star = users[i-1].Star
			}

		}
	}
}

// 坐席画像质检
func GetTSRQuality(TsrDetailList []TSRPortraitSchema.UserDetail, DictTSR map[int32]TsrStarResponseSchema, TSRGradeQueryParams TSRTopControllerRequest) error {
	var TsrIds []int32
	for _, TsrDetail := range TsrDetailList {
		TsrIds = append(TsrIds, TsrDetail.TSRId)
	}
	// 2获取电网保单数据  其他条件待追加
	var filter1 bson.D
	filter1 = append(filter1, bson.E{Key: "坐席ID", Value: bson.D{{"$in", TsrIds}}})
	NetPowerRenewalSchemas, err := TSRPortraitSchema.GetTSRRenewalSchemaList(filter1)
	if err != nil {
		return err
	}
	var PolicyIds []string
	for _, NetPowerRenewalSchema := range NetPowerRenewalSchemas {
		PolicyIds = append(PolicyIds, NetPowerRenewalSchema.PolicyId)
	}
	var filter2 bson.D
	filter2 = append(filter2, bson.E{Key: "投保单号", Value: bson.D{{"$in", PolicyIds}}})

	QualityStarSchemas, err := TSRPortraitSchema.QualityStarHandler(filter2)
	if err != nil {
		logging.Error(err)
		return err
	}
	QualityResponseSchemas, err := QualityStatistics(NetPowerRenewalSchemas, QualityStarSchemas)
	if err != nil {
		logging.Error(err)
		return err
	}
	SortQuality(QualityResponseSchemas)
	StarQuality(QualityResponseSchemas, TSRGradeQueryParams.QualityFiveStar,
		TSRGradeQueryParams.QualityFourStar, TSRGradeQueryParams.QualityThreeStar,
		TSRGradeQueryParams.QualityTwoStar, TSRGradeQueryParams.QualityOneStar)
	for _, Quality := range QualityResponseSchemas {
		if v1, ok := DictTSR[Quality.TSRId]; ok {
			v1.QualifiedStar = Quality.Star
			DictTSR[Quality.TSRId] = v1

		}
	}
	return nil

}

// 质检统计方法
func QualityStatistics(NetPowerRenewalSchemas []TSRPortraitSchema.NetPowerRenewalSchema, QualityStarSchemas []TSRPortraitSchema.QualityStarSchema) ([]QualityResponseSchema, error) {
	DictNetOrder := make(map[int32]QualityResponseSchema)
	var QualityResponseList []QualityResponseSchema
	// 计算分母
	for _, NetPowerRenewalSchema := range NetPowerRenewalSchemas {
		var QualityResponse QualityResponseSchema
		if v1, ok := DictNetOrder[NetPowerRenewalSchema.TSRId]; ok {
			QualityResponse = v1
		} else {
			QualityResponse = QualityResponseSchema{}
			QualityResponse.TSRId = NetPowerRenewalSchema.TSRId
		}
		// 100%的质检单
		QualityResponse.Qualified = QualityResponse.Qualified + 1
		QualityResponse.Total = QualityResponse.Total + 1
		DictNetOrder[QualityResponse.TSRId] = QualityResponse
	}
	// 计算分子
	for _, QualityStarSchema := range QualityStarSchemas {
		var QualityResponse QualityResponseSchema
		if v1, ok := DictNetOrder[QualityStarSchema.TSRId]; ok {
			QualityResponse = v1
			count := ComputeQuality(&QualityStarSchema)
			QualityResponse.Qualified = QualityResponse.Qualified - count
			DictNetOrder[QualityResponse.TSRId] = QualityResponse
		}
		// 累计合格数量
	}
	for _, value := range DictNetOrder {
		QualityResponseList = append(QualityResponseList, value)
	}
	return QualityResponseList, nil
}

// 累计合格数量
func ComputeQuality(QualityStarSchema *TSRPortraitSchema.QualityStarSchema) int32 {
	if QualityStarSchema.QualityType == "nan" {
		return 0
	} else {
		return 1
	}
}

// 排序
func SortQuality(Quality []QualityResponseSchema) {
	sort.Slice(Quality, func(i, j int) bool { // desc
		return Quality[i].Qualified/Quality[i].Total > Quality[j].Qualified/Quality[j].Total
	})
}

// 显示星
func StarQuality(Quality []QualityResponseSchema, QualityFiveStar float64, QualityFourStar float64,
	QualityThreeStar float64, QualityTwoStar float64, QualityOneStar float64) {
	var max int = 5
	//如果不够无位
	if len(Quality) < max {
		for i := 0; i < len(Quality); i++ {
			Quality[i].Star = max - i
		}
	} else {
		for i := 0; i < len(Quality); i++ {
			// 如果 0保单 1星
			if float32(i) <= float32(len(Quality))*float32(QualityFiveStar) {
				Quality[i].Star = 5
			} else if float32(i) <= float32(len(Quality))*float32(QualityFourStar) {
				Quality[i].Star = 4
			} else if float32(i) <= float32(len(Quality))*float32(QualityThreeStar) {
				Quality[i].Star = 3
			} else if float32(i) <= float32(len(Quality))*float32(QualityTwoStar) {
				Quality[i].Star = 2
			} else {
				Quality[i].Star = 1
			}
		}
	}
	// 如果星相同 与上排名同星
	for i := 0; i < len(Quality); i++ {
		if i > 0 {
			if Quality[i].Qualified/Quality[i].Total == Quality[i-1].Qualified/Quality[i-1].Total {
				Quality[i].Star = Quality[i-1].Star
			}

		}
	}
}
