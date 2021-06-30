package clientSchema

import (
	"dolphin/salesManager/models"
	"dolphin/salesManager/pkg/e"
	"dolphin/salesManager/pkg/gredis"
	"dolphin/salesManager/pkg/logging"
	"encoding/json"
	"strconv"
	"strings"
)

type AddClientForm struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	AppName   string `json:"app_name"`
	Desc      string `json:"desc"`
}

type Client struct {
	Id        int    `json:"id"`
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	AppName   string `json:"app_name"`
	PageNum   int    `json:"page_num"`
	PageSize  int    `json:"page_size"`
	Desc      string `json:"desc"`
}

func (a *Client) GetClientKey() string {
	return e.CACHE_ARTICLE + "_" + strconv.Itoa(a.Id)
}

func (a *Client) GetClientsKey() string {
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
	}

	if a.Id > 0 {
		keys = append(keys, strconv.Itoa(a.Id))
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}

func (a *Client) Add() error {
	appClient := map[string]interface{}{
		"app_id":     a.AppId,
		"app_name":   a.AppName,
		"desc":       a.Desc,
		"app_secret": a.AppSecret,
	}

	if err := models.AddAppClient(appClient); err != nil {
		return err
	}

	return nil
}
func (a *Client) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if a.Id != -1 {
		maps["id"] = a.Id
	}

	return maps
}

func (a *Client) GetAll() ([]*models.AppClient, error) {
	var (
		clients, cacheAppclients []*models.AppClient
	)

	cache := Client{
		Id:       a.Id,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetClientsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheAppclients)
			return cacheAppclients, nil
		}
	}

	clients, err := models.GetAppClients(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, clients, 3600)
	return clients, nil

}

func (a *Client) Delete() error {
	return models.DeleteAppClient(a.Id)
}

func (a *Client) ExistByID() (bool, error) {
	return models.ExistAppClientByID(a.AppId)
}

func (a *Client) Count() (int, error) {
	return models.GetAppClientTotal(a.getMaps())
}

func (a *Client) Get() (*models.AppClient, error) {
	var cacheArticle *models.AppClient

	cache := Client{Id: a.Id}
	key := cache.GetClientKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	client, err := models.GetAppClient(a.Id)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, client, 3600)
	return client, nil
}

func (a *Client) Edit() error {
	return models.EditAppClient(a.Id, map[string]interface{}{
		"app_name":   a.AppName,
		"desc":       a.Desc,
		"app_secret": a.AppSecret,
	})
}
