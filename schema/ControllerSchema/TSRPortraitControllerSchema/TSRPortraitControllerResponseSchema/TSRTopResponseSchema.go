package TSRPortraitControllerResponseSchema

import (
	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/schema/ServiceSchema/TSRPortraitSchema"
)

type TSRTopControllerResponseSchema struct {
	app.Response
	Data []TSRPortraitSchema.TsrStarPageResponseSchema
}
