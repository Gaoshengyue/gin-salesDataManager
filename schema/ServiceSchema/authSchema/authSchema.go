package authSchema

import "dolphin/salesManager/models"

type Auth struct {
	AppId     string `valid:"Required; MaxSize(50)" json:"app_id"`
	AppSecret string `valid:"Required; MaxSize(50)" json:"app_secret"`
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.AppId, a.AppSecret)
}
