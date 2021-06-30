package StatementControllerResponseSchema

import (
	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/schema/ServiceSchema/StatementSchema"
)

// 话务统计
type TrafficMeasurementDataResponseSchema struct {
	app.Response
	Data []StatementSchema.TrafficMeasurementSchema
}
