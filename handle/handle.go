package handle

import (
	"awesomeProject/jsonDeal"
	"awesomeProject/soldierInfo"
)
var UMAP *map[string]*soldierInfo.Info
func Init(Umap map[string]*soldierInfo.Info){
	UMAP = &Umap
}
func HandleIdGetRarity(id string)  map[string]string{
	retMap := make(map[string]string,3)

	if len(id) != 0 {
		ret := jsonDeal.GetRarity(id,*UMAP)
		if  ret != ""{
			retMap["condition"]="success"
			retMap["id"]=id
			retMap["Rarity"]=ret
		}else {
			retMap["condition"] = "not_exist"
			retMap["id"] = id
			retMap["Rarity"] = "error"
		}
	}else {
		retMap["condition"] = "error"
		retMap["id"] = id
		retMap["Rarity"] = "error"
	}
	return retMap
}
func HandleIdGetCombatPoints(id string)map[string]string{

	retMap := make(map[string]string,3)
	if len(id) != 0 {
		ret := jsonDeal.GetCombatPoints(id,*UMAP)
		if  ret != ""{
			retMap["condition"]="success"
			retMap["id"]=id
			retMap["CombatPoints"]=ret
		}else {
			retMap["condition"]="not_exist"
			retMap["id"]=id
			retMap["CombatPoints"]="error"
		}
	}else {
		retMap["condition"]="error"
		retMap["id"] = id
		retMap["CombatPoints"] = "error"

	}
	return retMap
}
func HandleGetStage(stage string)(map[string]string,[] *soldierInfo.Info){

	var dataRes [] *soldierInfo.Info
	retMap := make(map[string]string,2)
	if len(stage) != 0 {
		dataRes  = jsonDeal.GetSoldiersJson(stage ,*UMAP)
		if dataRes != nil{
			retMap["condition"]="pass"
			retMap["stage"] = stage
		}
	}else {
		retMap["condition"]="error"
		retMap["stage"] = stage
	}
	return retMap,dataRes
}
func HandleGetSoldiersJson(stage string)(map[string]string,[] *soldierInfo.Info){
	retMap := make(map[string]string,3)
	var dataRes [] *soldierInfo.Info
	if len(stage) != 0 {
		dataRes = jsonDeal.GetSoldiersJson(stage ,*UMAP)
		if dataRes != nil{
			retMap["condition"]="success"
			retMap["stage"] = stage

		}else {

			retMap["condition"] = "not_exist"
			retMap ["stage"]  = stage
			retMap["data"]=""
		}
	}else {
		retMap["condition"] = "error"
		retMap["stage" ]=  stage
		retMap["data" ]=  ""
	}
	return  retMap,dataRes
}