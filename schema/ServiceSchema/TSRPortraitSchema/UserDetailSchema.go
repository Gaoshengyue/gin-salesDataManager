package TSRPortraitSchema

import (
	"context"
	"dolphin/salesManager/pkg/dictData"
	"dolphin/salesManager/pkg/gmongo"
	"dolphin/salesManager/pkg/logging"
	"dolphin/salesManager/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
)

type UserDetail struct {
	TSRId             int32   `bson:"员工工号"`
	TSRName           string  `bson:"员工姓名"`
	Gender            string  `bson:"性别"`
	IdCard            string  `bson:"身份证号码"`
	Department        string  `bson:"部门"`
	WorkCardId        string  `bson:"展业证号码"`
	WorkCardStatus    string  `bson:"展业证状态"`
	PhoneNumberFirst  int64   `bson:"联系电话1"`
	PhoneNumberSecond float64 `bson:"联系电话2"`
	PhoneNumberThird  float64 `bson:"联系电话3"`
	DefaultMaterCall  int64   `bson:"默认外呼主叫"`
	CallStatus        string  `bson:"外呼"`
	WorkStatus        string  `bson:"在职状态"`
	Type              string  `bson:"类型"`
	CollectionTime    string  `bson:"采集时间"`
	AnnualizedPremium float64 `bson:"年化保费"`
	Star              int     `bson:"评星"`
}

//GetUserDetailAll 获取所有员工信息
func GetUserDetailAll(filter bson.D) ([]UserDetail, error) {
	return GetUserDetailHandler(filter)
}

//GetUserDetailHandler 获取所有员工信息
func GetUserDetailHandler(filter bson.D) ([]UserDetail, error) {
	collection := gmongo.Collection("员工信息")
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	if err := cur.Err(); err != nil {
		logging.Error(err)
		return nil, err
	}
	var UserDetailList []UserDetail

	err = cur.All(context.Background(), &UserDetailList)
	return UserDetailList, err
}

//GetUserDetailTSRIds 获取所有员工的id信息
func GetUserDetailTSRIds(filter bson.D) ([]int32, error) {
	UserDetails, err := GetUserDetailHandler(filter)
	if err != nil {
		logging.Error(err)
		return nil, err
	}
	var TSRIds []int32
	for _, UserDetail := range UserDetails {
		TSRIds = append(TSRIds, UserDetail.TSRId)
	}
	return TSRIds, nil
}

//绑定每个员工数据对象通时部分数据
func (TsrObj *UserDetail) GetTSRCallTimeDetail(CallCountStandard float64, CallCountRate float64, CallTimeStandard float64, CallTimeRate float64, CallCountTimeStandard float64, CallCountTimeRate float64) (*TSRTrafficDetail, error) {
	matchStage := bson.D{{"$match", bson.D{{"小宝工号", TsrObj.TSRId}}}}
	return GetTSRCallTimeRecordTotal(matchStage, CallCountStandard, CallCountRate, CallTimeStandard, CallTimeRate, CallCountTimeStandard, CallCountTimeRate)

}

//获取坐席保单id列表
func (TsrObj *UserDetail) GetTSRPolicyIdList(startTime string, endTime string) ([]string, error) {
	filter := bson.D{{"坐席ID", TsrObj.TSRId}, {"承保时间", bson.D{{"$gte", startTime}, {"$lte", endTime}}}}
	var TSRPolicyIdList []string
	TSRPolicyIdObjList, err := GetTSRPolicyIdList(filter)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(TSRPolicyIdObjList); i++ {
		TSRPolicyIdList = append(TSRPolicyIdList, TSRPolicyIdObjList[i].PolicyId)
	}
	return TSRPolicyIdList, nil
}

//获取坐席保单对象列表
func (TsrObj *UserDetail) GetTSRPolicyObjList(startTime string, endTime string) ([]NetPowerPolicySchema, error) {
	filter := bson.D{{"坐席ID", TsrObj.TSRId}, {"承保时间", bson.D{{"$gte", startTime}, {"$lte", endTime}}}}
	return GetTSRPolicyIdList(filter)
}

//获取坐席已失效状态下的保单
func (TsrObj *UserDetail) GetTSRSurrenderPolicyList() ([]PolicyStatusSchema, error) {
	startTime := util.GetNowTimeOffset(0, 0, -15)
	endTime := util.GetTodayDateStr()
	PolicyList, err := TsrObj.GetTSRPolicyIdList(startTime, endTime)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"保单号", bson.D{{"$in", bson.A{PolicyList}}}},
		{"保单状态", bson.D{{"$in", bson.A{"停效", "契撤", "犹退终止", "解约", "退保终止"}}}},
	}
	return GetTSRSurrenderPolicyList(filter)
}

//获取坐席已失效状态下的保单id列表
func (TsrObj *UserDetail) GetTSRSurrenderPolicyIdList() ([]string, error) {
	startTime := util.GetNowTimeOffset(0, 0, -15)
	endTime := util.GetTodayDateStr()
	PolicyList, err := TsrObj.GetTSRPolicyIdList(startTime, endTime)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"保单号", bson.D{{"$in", bson.A{PolicyList}}}},
		{"保单状态", bson.D{{"$in", bson.A{"停效", "契撤", "犹退终止", "解约", "退保终止"}}}},
	}
	return GetTSRSurrenderPolicyIdList(filter)
}

//获取坐席成单保费金额
func (TsrObj *UserDetail) GetTSRPolicyPremiumTotalNumber(StartTime string, EndTime string, ProductType string,
	OrderPlace string, NameListType string, CustomerAgeTop int64, CustomerAgeBottom int64) (float64, error) {
	var filterStage bson.D
	filterStage = append(filterStage, bson.E{Key: "坐席ID", Value: TsrObj.TSRId})
	filterStage = append(filterStage, bson.E{Key: "承保时间",
		Value: bson.D{{"$gte", StartTime}, {"$lte", EndTime}}})
	if ProductType != "0" {
		productObj := dictData.ProductObj()
		switch ProductType {
		case "意外险":
			filterStage = append(filterStage, bson.E{Key: "产品名称", Value: bson.D{{"$in", productObj.Accident}}})
		case "重疾险":
			filterStage = append(filterStage, bson.E{Key: "产品名称", Value: bson.D{{"$in", productObj.SeriousIllness}}})
		case "年金险":
			filterStage = append(filterStage, bson.E{Key: "产品名称", Value: bson.D{{"$in", productObj.Annuity}}})
		case "寿险":
			filterStage = append(filterStage, bson.E{Key: "产品名称", Value: bson.D{{"$in", productObj.LifeInsurance}}})
		}

	}
	if OrderPlace != "0" {
		provinceObj := dictData.ProvincePlaceObj()
		switch OrderPlace {
		case "西南":
			filterStage = append(filterStage, bson.E{Key: "省份", Value: bson.D{{"$in", provinceObj.SouthWest}}})
		case "西北":
			filterStage = append(filterStage, bson.E{Key: "省份", Value: bson.D{{"$in", provinceObj.NorthWest}}})
		case "华中":
			filterStage = append(filterStage, bson.E{Key: "省份", Value: bson.D{{"$in", provinceObj.CentralChina}}})
		case "华南":
			filterStage = append(filterStage, bson.E{Key: "省份", Value: bson.D{{"$in", provinceObj.SouthChina}}})
		case "华东":
			filterStage = append(filterStage, bson.E{Key: "省份", Value: bson.D{{"$in", provinceObj.EastChina}}})
		case "北方":
			filterStage = append(filterStage, bson.E{Key: "省份", Value: bson.D{{"$in", provinceObj.North}}})

		}
		if NameListType != "0" {
			filterStage = append(filterStage, bson.E{Key: "名单大类", Value: NameListType})
		}
		if CustomerAgeTop != 0 && CustomerAgeBottom != 0 {
			filterStage = append(filterStage, bson.E{Key: "年龄", Value: bson.D{{"$gte", CustomerAgeBottom},
				{"$lte", CustomerAgeTop}}})
		}

	}

	matchStage := bson.D{{"$match", filterStage}}

	return GetTSRPolicyPremiumTotal(matchStage)
}

//获取坐席当月退单保费金额
func (TsrObj *UserDetail) GetTSRSurrenderPolicyPremiumTotalNumber() (float64, error) {
	SurrenderPolicyIdList, err := TsrObj.GetTSRSurrenderPolicyIdList()
	if err != nil {
		return 0, err
	}
	matchStage := bson.D{{"$match", bson.D{{"坐席ID", TsrObj.TSRId},
		{"保单号", bson.D{{"$in", bson.A{SurrenderPolicyIdList}}}},
	}}}
	return GetTSRPolicyPremiumTotal(matchStage)
}
