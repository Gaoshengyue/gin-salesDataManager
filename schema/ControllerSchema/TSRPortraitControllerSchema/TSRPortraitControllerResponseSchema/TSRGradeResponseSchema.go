package TSRPortraitControllerResponseSchema

import (
	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/schema/ServiceSchema/TSRPortraitSchema"
)

// 业绩评星
type TSRGradeControllerResponseSchema struct {
	app.Response
	Data []TSRPortraitSchema.TSRGradeDetail
}
