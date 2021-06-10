package route

import (
	"awesomeProject/handle"
	"awesomeProject/jsonDeal"
	"awesomeProject/soldierInfo"
	"github.com/gin-gonic/gin"
)


func SetUpRoute() *gin.Engine{
	r := gin.Default()
	// GET：请求方式；/getinfo：请求的路径
	// 当客户端以GET方法请求/getinfo，会执行后面的匿名函数



	r.GET("/getstage", getStage)
	r.GET("/getArmyAssembled", getArmyAssembled)
	r.GET("/idGetRarity", idGetRarity)
	r.GET("/idGetCombatPoints", idGetCombatPoints)
	r.GET("/getSoldiersJson", getSoldiersJson)

	return r
}

func getStage(c *gin.Context)  {
	stage := c.Query("stage")
	retMap,retInfo :=handle.HandleGetStage(stage)
	c.JSON(200, gin.H{
		"condition":retMap["condition"],
		"stage" : retMap["stage"],
		"data" : retInfo,
	})
}
func getArmyAssembled(c *gin.Context){
	rarity := c.Query("rarity")
	stage  := c.Query("stage")
	if len(stage) != 0 && len(rarity) != 0{
		var dataRes [] *soldierInfo.Info = jsonDeal.GetArmyAssembled(rarity,stage,*handle.UMAP)
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
	ret := handle.HandleIdGetRarity(id)

	c.JSON(200, gin.H{
		"condition":ret["condition"],
		"id":ret["id"],
		"Rarity" :ret["Rarity"],
	})
}
func idGetCombatPoints(c *gin.Context){
	id := c.Query("id")
	ret := handle.HandleIdGetCombatPoints(id)

	c.JSON(200, gin.H{
		"condition":ret["condition"],
		"id":ret["id"],
		"CombatPoints" :ret["CombatPoints"],
	})

}

func getSoldiersJson(c *gin.Context){
	stage := c.Query("stage")
	retMap , retInfo := handle.HandleGetSoldiersJson(stage)

	c.JSON(200, gin.H{
			"condition":retMap["condition"],
			"stage" : retMap["stage"],
			"data" : retInfo,
	})
}