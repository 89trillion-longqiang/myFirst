package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	goini "github.com/clod-moon/goconf"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

type Info struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	UnlockArena string `json:"UnlockArena"`
	Rarity       string `json:"rarity"`
	CombatPoints string `json:"combatPoints"`
	Desc         string `json:"desc"`
}

func main() {

	var fileLocker sync.Mutex // file locker

	var jsonPath = pflag.StringP("config.army.model.json	-path", "p", "", "Input JsonPath")
	// 把用户传递的命令行参数解析为对应变量的值
	pflag.Parse()

	filename := *jsonPath
	fileLocker.Lock()
	dataJson, err := ioutil.ReadFile(filename) //读取json文件
	fileLocker.Unlock()
	if err != nil {
		fmt.Println("read file error")
	}

	var Umap map[string]*Info

	err = json.Unmarshal(dataJson, &Umap)
	if err != nil {
		fmt.Println("unmarshal json file error")
	}

	da := jsonMarshal(Umap)  ///将需要转化的数据通过json.Marshal() 转化为[]byte 格式
	jsonWrite(da,"data.json")		  ///将[]byte写入到文件中

	port := readConfigPort("./app.ini")  ///通过配置文件获取端口号
	
	// 创建一个默认的路由引擎
	r := gin.Default()
	// GET：请求方式；/getinfo：请求的路径
	// 当客户端以GET方法请求/getinfo，会执行后面的匿名函数

	r.GET("/getinfo", func(c *gin.Context) {
		id := c.Query("id")
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"data": Umap[id],
		})
	})

	r.GET("/getstage", func(c *gin.Context) {
		stage := c.Query("stage")
		//获取相应的数据
		var dataRes []*Info = getSoldiersJson(stage,Umap)
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"data": dataRes,
		})
	})

	r.GET("/getArmyAssembled", func(c *gin.Context) {
		rarity := c.Query("rarity")
		stage  := c.Query("stage")
		var dataRes [] *Info = getArmyAssembled(rarity,stage,Umap)
			c.JSON(200,gin.H{
				"data" : dataRes,
			})
	})

	r.GET("/idGetRarity", func(c *gin.Context) {
		id := c.Query("id")
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"id":id,
			"Rarity":getRarity(id,Umap),
		})
	})

	r.GET("/idGetCombatPoints", func(c *gin.Context) {
		id := c.Query("id")
		c.JSON(200,gin.H{
			"id":id,
			"CombatPoints":getCombatPoints(id,Umap),
		})
	})

	r.GET("/getSoldiersJson", func(c *gin.Context) {
		stage := c.Query("stage")
		var dataRes [] *Info = getSoldiersJson(stage ,Umap)
		c.JSON(200,gin.H{
			"stage" : stage,
			"data" : dataRes,
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":" + port)

}

///读取配置文件来获取端口号
func readConfigPort(filePath string)(port string){
	conf := goini.InitConfig("./app.ini")
	port = conf.GetValue("server", "HttpPort")
	fmt.Println("port:", port)
	return 
}
///将数据转化为[]byte
func jsonMarshal(info map[string]*Info) []byte {

	data, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
//将数据写入json文件中，生成新的json文件
func jsonWrite(data []byte,filename string) {
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
///通过稀有度，阶段来获取已解锁的所有士兵的信息
func getArmyAssembled(rarity string , stage string , Umap map[string]*Info) (data []*Info){
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
///通过id来获取稀有度
func getRarity(id string,Umap map[string]*Info) string{
	info := Umap[id]
	if info.Rarity == "" {
		return ""
	}
	return  info.Rarity
}
///通过id来获取战力
func getCombatPoints(id string,Umap map[string]*Info)  string{
	info := Umap[id]
	if info.CombatPoints == "" {
		return ""
	}
	return  info.CombatPoints
}
///通过阶段来获取已解锁的士兵信息
func getSoldiersJson(stage string,Umap map[string]*Info) (data []*Info) {
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