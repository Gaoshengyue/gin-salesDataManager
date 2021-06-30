package models

import "github.com/jinzhu/gorm"

//此部分留为例子
type AppClient struct {
	Model

	AppId string `json:"app_id" gorm:"index"`

	AppName   string `json:"app_name"`
	Desc      string `json:"desc"`
	AppSecret string `json:"app_secret"`
	State     int    `json:"state" gorm:"default:1"`
}

// ExistAppClientByID checks if an article exists based on ID
func ExistAppClientByID(appId string) (bool, error) {
	var appClient AppClient
	err := db.Select("id").Where("app_id = ? AND deleted_on = ? ", appId, 0).First(&appClient).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if appClient.Id > 0 {
		return true, nil
	}

	return false, nil
}

// GetAppClientTotal gets the total number of appClient based on the constraints
func GetAppClientTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&AppClient{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// GetAppClients gets a list of appClients based on paging constraints
func GetAppClients(pageNum int, pageSize int, maps interface{}) ([]*AppClient, error) {
	var appClients []*AppClient
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&appClients).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return appClients, nil
}

// GetAppClient Get a single article based on ID
func GetAppClient(id int) (*AppClient, error) {
	var appCLient AppClient
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&appCLient).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &appCLient, nil
}

// EditAppClient modify a single appClient
func EditAppClient(id int, data interface{}) error {
	if err := db.Model(&AppClient{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddAppClient add a single appClient
func AddAppClient(data map[string]interface{}) error {
	appClient := AppClient{
		Desc:      data["desc"].(string),
		AppId:     data["app_id"].(string),
		AppName:   data["app_name"].(string),
		AppSecret: data["app_secret"].(string),
	}
	if err := db.Create(&appClient).Error; err != nil {
		return err
	}

	return nil
}

// DeleteAppClient delete a single appClient
func DeleteAppClient(id int) error {
	if err := db.Where("id = ?", id).Delete(AppClient{}).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllAppClient clear all appClient
func CleanAllAppClient() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&AppClient{}).Error; err != nil {
		return err
	}

	return nil
}

// CheckAuth checks if authentication information exists
func CheckAuth(appId string, appSecret string) (bool, error) {
	var auth AppClient
	err := db.Select("id").Where(AppClient{AppId: appId, AppSecret: appSecret}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.Id > 0 {
		return true, nil
	}

	return false, nil
}
