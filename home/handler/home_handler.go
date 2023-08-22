package handler

import (
	"github.com/gin-gonic/gin"
)

type Home struct{}

func NewHome() *Home {
	return &Home{}
}

func (this *Home) InitRoute(router *gin.Engine) {
	router.GET("/", this.home)
	router.GET("/user", this.user)
	router.GET("/news", this.news)
}

func (this *Home) home(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"code": 200, "msg": "hello world"})
}

func (this *Home) user(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"code": 200, "msg": "user api"})
}

func (this *Home) news(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"code": 200, "msg": "news api"})
}
