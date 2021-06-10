package main

import (
	"awesomeProject/handle"
	"awesomeProject/initDeal"
	"awesomeProject/jsonDeal"
	"awesomeProject/route"
)




func main() {
	Umap := jsonDeal.Init()

	handle.Init(Umap)

	port := initDeal.ReadConfigPort("./app.ini")  ///通过配置文件获取端口号

	r := route.SetUpRoute()
	r.Run(":"+port)
}



