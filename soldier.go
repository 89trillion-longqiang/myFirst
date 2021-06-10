package main

import (
	"encoding/json"
	"fmt"

	"awesomeProject/initDeal"
	"awesomeProject/jsonConfig"
	"awesomeProject/jsonDeal"
	"awesomeProject/route"
	"awesomeProject/soldierInfo"
)




func main() {

	dataJson := jsonConfig.ParseConfigFile()

	var Umap map[string]*soldierInfo.Info

	err := json.Unmarshal(dataJson, &Umap)
	if err != nil {
		fmt.Println("unmarshal json file error")
	}
	da := jsonDeal.JsonMarshal(Umap)  ///将需要转化的数据通过json.Marshal() 转化为[]byte 格式
	jsonDeal.JsonWrite(da,"data.json")		  ///将[]byte写入到文件中

	port := initDeal.ReadConfigPort("./app.ini")  ///通过配置文件获取端口号

	r := route.SetUpRoute(Umap)
	r.Run(":"+port)
}



