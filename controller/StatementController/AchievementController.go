package StatementController

import (
	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/pkg/e"
	"dolphin/salesManager/schema/ControllerSchema/StatementControllerSchema/StatementControllerRequestSchema"
	"dolphin/salesManager/service/StatementService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary AchievementMoon  按月份统计
// @Failure 500 {object} app.Response
// @Router /projectData/AchievementMoon [Get]
func AchievementMoonController(c *gin.Context) {
	appG := app.Gin{C: c}
	var achievementControllerRequest StatementControllerRequestSchema.AchievementControllerRequest
	// TSRTopRequest.InitDefaultRequest()
	// //获取Query参数，根据地址复制到对象属性中　对应form

	if c.BindJSON(&achievementControllerRequest) != nil {
		appG.Response(http.StatusUnprocessableEntity, e.INVALID_PARAMS, nil)
		return
	}
	//　测试获取全部外呼量
	achievementPageResponse, err := StatementService.AchievementMoonPageFunc(achievementControllerRequest)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_CALCULATION_TSR_GRADE, nil)
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, achievementPageResponse)
	}

}

// @Summary AchievementSummary  按月份统计
// @Failure 500 {object} app.Response
// @Router /projectData/AchievementSummary [Get]
func AchievementSummaryController(c *gin.Context) {
	appG := app.Gin{C: c}
	var achievementControllerRequest StatementControllerRequestSchema.AchievementControllerRequest
	// TSRTopRequest.InitDefaultRequest()
	// //获取Query参数，根据地址复制到对象属性中　对应form

	if c.BindJSON(&achievementControllerRequest) != nil {
		appG.Response(http.StatusUnprocessableEntity, e.INVALID_PARAMS, nil)
		return
	}
	//　测试获取全部外呼量
	achievementPageResponse, err := StatementService.AchievementSummaryFunc(achievementControllerRequest)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_CALCULATION_TSR_GRADE, nil)
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, achievementPageResponse)
	}

}