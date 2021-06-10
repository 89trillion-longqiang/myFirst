package jsonDeal

import (
	"awesomeProject/jsonConfig"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"awesomeProject/soldierInfo"
)

func Init() map[string]*soldierInfo.Info {

	dataJson := jsonConfig.ParseConfigFile()
	var Umap map[string]*soldierInfo.Info

	err := json.Unmarshal(dataJson, &Umap)
	if err != nil {
		fmt.Println("unmarshal json file error")
	}

	da := JsonMarshal(Umap)  ///将需要转化的数据通过json.Marshal() 转化为[]byte 格式
	JsonWrite(da,"data.json")		  ///将[]byte写入到文件中
	return  Umap
}
// JsonMarshal /将数据转化为[]byte
func JsonMarshal(info map[string]*soldierInfo.Info) []byte {

	data, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// JsonWrite /将数据写入json文件中，生成新的json文件
func JsonWrite(data []byte,filename string) {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

// GetArmyAssembled /通过稀有度，阶段来获取已解锁的所有士兵的信息
func GetArmyAssembled(rarity string , stage string , Umap map[string]*soldierInfo.Info) (data []*soldierInfo.Info){
	if rarity == "" || stage == "" {
		return
	}
	for _,v :=range Umap{
		if v.Rarity == rarity  && v.Desc <= stage {
			if v.UnlockArena == "1"{
				data = append(data,v)
			}
		}
	}
	return
}

// GetRarity /通过id来获取稀有度
func GetRarity(id string,Umap map[string]*soldierInfo.Info) string{
	for k,v:=range Umap{
		if id == k{
			return  v.Rarity
		}
	}

	return ""
}

// GetCombatPoints /通过id来获取战力
func GetCombatPoints(id string,Umap map[string]*soldierInfo.Info)  string{
	for k,v:=range Umap{
		if id == k{
			return  v.CombatPoints
		}
	}

	return ""
}

// GetSoldiersJson /通过阶段来获取已解锁的士兵信息
func GetSoldiersJson(stage string,Umap map[string]*soldierInfo.Info) (data []*soldierInfo.Info) {
	if stage == "" {
		return
	}
	for _, v := range Umap {
		if v.Desc <= stage {
			data = append(data, v)
		}
	}
	return
}