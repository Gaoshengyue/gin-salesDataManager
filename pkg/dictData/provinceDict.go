package dictData

import (
	"dolphin/salesManager/pkg/setting"
	"encoding/json"
	"io/ioutil"
	"os"
)

type ProvinceDict struct {
	SouthWest    []string //西南
	NorthWest    []string //西北
	CentralChina []string //华中
	SouthChina   []string //华南
	EastChina    []string //华东
	North        []string //北方
}

var provinceDictObj *ProvinceDict

func Setup() {

	fp, err := os.Open(setting.AppSetting.ProvinceDictPath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	bytes, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &provinceDictObj)
	if err != nil {
		panic(err)
	}
	//使用cs
}

func ProvincePlaceObj() *ProvinceDict {

	return provinceDictObj
}
