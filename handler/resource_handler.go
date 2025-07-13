package handler

func ReadResource(c *gin.Context){
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})


func WriteResource(c *gin.Context){
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
