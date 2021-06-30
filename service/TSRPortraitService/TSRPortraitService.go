package TSRPortraitService

import (
	"dolphin/salesManager/pkg/logging"
	"dolphin/salesManager/schema/ControllerSchema/TSRPortraitControllerSchema/TSRPortraitControllerRequestSchema"
	"dolphin/salesManager/schema/ServiceSchema/TSRPortraitSchema"
	"go.mongodb.org/mongo-driver/bson"
)

//计算在职TSR画像评分
func CalculationTSRGrade(TSRGradeQueryParams TSRPortraitControllerRequestSchema.TSRGradeControllerRequest) ([]TSRPortraitSchema.TSRGradeDetail, error) {
	TsrDetailList, err := TSRPortraitSchema.GetUserDetailAll(bson.D{{"在职状态", "在职"}}) //{"联系电话1", bson.D{{"$in", TSRGradeQueryParams.PhoneList}}},

	if err != nil {
		logging.Error(err)
		return nil, err
	}
	var TSRGradeArray []TSRPortraitSchema.TSRGradeDetail
	for i := 0; i < len(TsrDetailList); i++ {
		//实例化坐席评分对象
		TSRGradeDetailObj := TSRPortraitSchema.TSRGradeDetail{TSRId: TsrDetailList[i].TSRId, TSRName: TsrDetailList[i].TSRName}
		//计算坐席评分相关结果
		err = TSRGradeDetailObj.GradeInit(TSRGradeQueryParams)
		if err != nil {
			logging.Error(err)
			return nil, err
		}
		//添加坐席评分结果到数组
		TSRGradeArray = append(TSRGradeArray, TSRGradeDetailObj)
	}
	return TSRGradeArray, nil

}
