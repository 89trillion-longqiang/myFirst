package main

import (
	"awesomeProject/initDeal"
	"awesomeProject/jsonDeal"
	"awesomeProject/soldierInfo"
	"testing"
)

func Test_readConfigPort(t *testing.T) {
	port := initDeal.ReadConfigPort("./app.ini")

	if port != "8000" {
		t.Errorf("Test_readConfigPort error")
	}else {
		t.Log("Test_readConfigPort pass")
	}
}

func Test_getRarity(t *testing.T){
	var Umap map[string]*soldierInfo.Info = make(map[string]*soldierInfo.Info,1)
	Umap["10101"] = &soldierInfo.Info{
		Id:"10101",
		Name:"Swordsman",
		UnlockArena:"0",
		Rarity:"1",
		CombatPoints:"167",
		Desc:"army_desc_10101",
	}

	raeity := jsonDeal.GetRarity("10101",Umap)
	if raeity != "1" {
		t.Errorf("Test_getRarity error")
	}else {
		t.Log("Test_getRarity pass")
	}
}

func Test_getCombatPoints(t *testing.T){
	var Umap map[string]*soldierInfo.Info = make(map[string]*soldierInfo.Info,1)
	Umap["10101"] = &soldierInfo.Info{
		Id:"10101",
		Name:"Swordsman",
		UnlockArena:"0",
		Rarity:"1",
		CombatPoints:"167",
		Desc:"army_desc_10101",
	}
	comPointer := jsonDeal.GetCombatPoints("10101",Umap)
	if comPointer != "167" {
		t.Errorf("Test_getRarity error")
	}else {
		t.Log("Test_getRarity pass")
	}
}

