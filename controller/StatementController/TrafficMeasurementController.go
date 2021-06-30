package StatementController

import (
	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/pkg/e"
	"dolphin/salesManager/schema/ControllerSchema/StatementControllerSchema/StatementControllerRequestSchema"
	"dolphin/salesManager/service/StatementService"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary GetTrafficMeasurementData
// @Param page_size query number  false "每页数据大小"
// @Param current query number  false "当前页码"
// @Success 200 {object} StatementControllerResponseSchema.NetPowerPolicyDataResponseSchema
// @Failure 500 {object} app.Response
// @Router /projectData/GetTrafficMeasurementData [Get]
func TrafficMeasurementDataController(c *gin.Context) {
	appG := app.Gin{C: c}
	var PageNationRequest StatementControllerRequestSchema.PageNationColumn
	//初始化默认参数
	PageNationRequest.InitDefaultRequest()
	//获取Query参数，根据地址复制到对象属性中　对应form
	if c.BindQuery(&PageNationRequest) != nil {
		appG.Response(http.StatusUnprocessableEntity, e.INVALID_PARAMS, nil)
		return
	}
	//　坐席话务统计
	TrafficMeasurementArray, err := StatementService.GetTrafficMeasurementBaseDataArray(PageNationRequest)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_CALCULATION_TSR_GRADE, nil)
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, TrafficMeasurementArray)
	}
}
