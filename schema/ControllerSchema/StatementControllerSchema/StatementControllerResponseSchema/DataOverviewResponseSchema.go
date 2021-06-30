package StatementControllerResponseSchema

import (
	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/schema/ServiceSchema/StatementSchema"
)

// 数据概览基础看板
type BaseDataResponseSchema struct {
	app.Response
	Data StatementSchema.DataOverviewBaseDataSchema
}

// 数据概览产品分布
type ProductDataResponseSchema struct {
	app.Response
	Data []StatementSchema.DataOverviewProductDistribution
}
