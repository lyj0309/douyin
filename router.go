package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lyj0309/douyin/controller"
	"github.com/lyj0309/douyin/utils"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

	apiRouter.Use(utils.JWTAuthMiddleware())

	apiRouter.GET("/user/", controller.GetUserInfo)
	//apiRouter.POST("/publish/action/", controller.Publish)
	//apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	//apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
