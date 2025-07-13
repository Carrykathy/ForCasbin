package handler

import (
	"fmt"
	"log"

	"ForCasbin/component"
)

func login(c *gin.Context) {
	username, password := c.PostForm("username"), c.PostForm("password")
	if username == "admin" && password == "admin" {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "登录成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "用户名或密码错误",
		})
	}
	u.err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	sessionId := fmt.Sprintf("%s-%s", u.String(), username)

	component.GlobalCache.Set(sessionId, []byte(username))
	c.SetCookie("current_subject", sessionId, 30*60, "/resources", "", false, true)
	c.JSON(200, component.RestResponse{code: 1, Message: username + "登录成功"})
}
