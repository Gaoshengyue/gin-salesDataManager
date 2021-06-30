package StatementControllerResponseSchema

import (
	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/schema/ServiceSchema/StatementSchema"
)

// 网电保单
type NetPowerPolicyDataResponseSchema struct {
	app.Response
	Data []StatementSchema.NetPowerPolicySchema
}
