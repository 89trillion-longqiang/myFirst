package route

import (
	"awesomeProject/jsonDeal"
	"awesomeProject/soldierInfo"
	"github.com/gin-gonic/gin"
)
var UMAP *map[string]*soldierInfo.Info

func SetUpRoute(Umap map[string]*soldierInfo.Info) *gin.Engine{
	UMAP = &Umap
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

	r.GET("/getstage", getStage)
	r.GET("/getArmyAssembled", getArmyAssembled)
	r.GET("/idGetRarity", idGetRarity)
	r.GET("/idGetCombatPoints", idGetCombatPoints)
	r.GET("/getSoldiersJson", getSoldiersJson)

	return r
}
func getStage(c *gin.Context)  {
	stage := c.Query("stage")
	if len(stage) != 0 {
		var dataRes [] *soldierInfo.Info = jsonDeal.GetSoldiersJson(stage ,*UMAP)
		if dataRes != nil{
			c.JSON(200, gin.H{
				"stage" : stage,
				"data" : dataRes,
			})
		}
	}else {
		c.JSON(200, gin.H{
			"condition":"error",
			"Rarity" :"error",
		})
	}
}
func getArmyAssembled(c *gin.Context){
	rarity := c.Query("rarity")
	stage  := c.Query("stage")
	if len(stage) != 0 && len(rarity) != 0{
		var dataRes [] *soldierInfo.Info = jsonDeal.GetArmyAssembled(rarity,stage,*UMAP)
		if dataRes != nil{
			c.JSON(200, gin.H{
				"stage" : stage,
				"data" : dataRes,
			})
		}
	}else {
		c.JSON(200, gin.H{
			"condition":"error",
			"data" :"error",
		})
	}
}
func idGetRarity(c *gin.Context){
	id := c.Query("id")
	if len(id) != 0 {
		ret := jsonDeal.GetRarity(id,*UMAP)
		if  ret != ""{
			c.JSON(200, gin.H{
				"condition":"success",
				"id":id,
				"Rarity" :ret,
			})
		}else {
			c.JSON(200, gin.H{
				"condition":"not_exist",
				"id":id,
				"Rarity" :"error",
			})
		}
	}else {
		c.JSON(200, gin.H{
			"condition":"error",
			"id":id,
			"Rarity" :"error",
		})
	}
}
func idGetCombatPoints(c *gin.Context){
	id := c.Query("id")
	if len(id) != 0 {
		ret := jsonDeal.GetCombatPoints(id,*UMAP)
		if  ret != ""{
			c.JSON(200, gin.H{
				"condition":"success",
				"id":id,
				"CombatPoints" :ret,
			})
		}else {
			c.JSON(200, gin.H{
				"condition":"not_exist",
				"id":id,
				"CombatPoints" :"error",
			})
		}
	}else {
		c.JSON(200, gin.H{
			"condition":"error",
			"id":id,
			"CombatPoints" :"error",
		})
	}
}
func getSoldiersJson(c *gin.Context){
	stage := c.Query("stage")
	if len(stage) != 0 {
		var dataRes [] *soldierInfo.Info = jsonDeal.GetSoldiersJson(stage ,*UMAP)
		if dataRes != nil{
			c.JSON(200, gin.H{
				"condition":"success",
				"stage" : stage,
				"data" : dataRes,
			})
		}else {
			c.JSON(200, gin.H{
				"condition":"not_exist",
				"stage" : stage,
				"data" : "",
			})
		}
	}else {
		c.JSON(200, gin.H{
			"condition":"error",
			"stage" : stage,
			"data" : "",
		})
	}
}