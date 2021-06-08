package main

import (
	"testing"
)

func Test_readConfigPort(t *testing.T) {
	port := readConfigPort("./app.ini")

	if port != "8000" {
		t.Errorf("Test_readConfigPort error")
	}else {
		t.Log("Test_readConfigPort pass")
	}
}

func Test_getRarity(t *testing.T){
	var Umap map[string]*Info = make(map[string]*Info,1)
	Umap["10101"] = &Info{
		Id:"10101",
		Name:"Swordsman",
		UnlockArena:"0",
		Rarity:"1",
		CombatPoints:"167",
		Desc:"army_desc_10101",
	}

	raeity := getRarity("10101",Umap)
	if raeity != "1" {
		t.Errorf("Test_getRarity error")
	}else {
		t.Log("Test_getRarity pass")
	}
}

func Test_getCombatPoints(t *testing.T){
	var Umap map[string]*Info = make(map[string]*Info,1)
	Umap["10101"] = &Info{
		Id:"10101",
		Name:"Swordsman",
		UnlockArena:"0",
		Rarity:"1",
		CombatPoints:"167",
		Desc:"army_desc_10101",
	}
	comPointer := getCombatPoints("10101",Umap)
	if comPointer != "167" {
		t.Errorf("Test_getRarity error")
	}else {
		t.Log("Test_getRarity pass")
	}
}

