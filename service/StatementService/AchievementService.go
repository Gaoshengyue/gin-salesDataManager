package StatementService

import (
	"dolphin/salesManager/schema/ControllerSchema/StatementControllerSchema/StatementControllerRequestSchema"
	"dolphin/salesManager/schema/ControllerSchema/StatementControllerSchema/StatementControllerResponseSchema"
	"dolphin/salesManager/schema/ServiceSchema/TSRPortraitSchema"

	"go.mongodb.org/mongo-driver/bson"
)

//业绩翻页
func AchievementMoonPageFunc(achievementControllerRequest StatementControllerRequestSchema.AchievementControllerRequest) (StatementControllerResponseSchema.AchievementPageResponse, error) {

	matchStage := []bson.M{{"$match": bson.M{"承保时间": bson.M{"$gt": "2021-01-03 19:57:45", "$lte": "2021-01-04 21:57:45"}}}}
	achievementPageSchema, err := TSRPortraitSchema.AchievementMoonPage(matchStage)
	achievementPageResponse := StatementControllerResponseSchema.AchievementPageResponse{
		PageSize: achievementControllerRequest.PageSize,
		Current:  achievementControllerRequest.Current,
		Data:     achievementPageSchema,
	}
	if err != nil {
		return achievementPageResponse, err
	}

	return achievementPageResponse, nil
}

func AchievementSummaryFunc(achievementControllerRequest StatementControllerRequestSchema.AchievementControllerRequest) (StatementControllerResponseSchema.AchievementSummaryResponse, error) {
	matchStage := []bson.M{{"$match": bson.M{"承保时间": bson.M{"$gt": "2021-01-03 19:57:45", "$lte": "2021-01-04 21:57:45"}}}}
	achievementSummarySchema, err := TSRPortraitSchema.AchievementSummary(matchStage)
	achievementSummaryResponse := StatementControllerResponseSchema.AchievementSummaryResponse{Data: achievementSummarySchema}
	if err != nil {
		return achievementSummaryResponse, err
	}
	return achievementSummaryResponse, nil
}