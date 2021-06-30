package StatementControllerResponseSchema

import (
	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/schema/ServiceSchema/StatementSchema"
)

// 签单月拨打表现
type TSRMonthlyPerformanceResponseSchema struct {
	app.Response
	Data []StatementSchema.TSRMonthlyPerformance
}
