package main

import (
	controller "github.com/tccip/simpleblog/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	//设置默认的服务器配置
	router := gin.Default()

	//加载静态资源
	router.Static("/static", "./static")
	//加载网页文件
	router.LoadHTMLGlob("views/*")

	/***注册控制器***/
	//登录主页
	router.GET("/", controller.Home)

	/**文章操作**/
	//文章列表
	router.GET("/topic", controller.Topic)
	//浏览文章
	router.GET("/topic/view", controller.TopicView)
	//添加文章
	router.GET("/topic/add", controller.TopicAdd)
	//添加文章后加入数据库
	router.POST("/topic/addIn", controller.TopicAddIn)
	//删除文章
	router.GET("/topic/del", controller.TopicDel)
	//修改文章
	router.GET("/topic/modify", controller.TopicModify)
	//修改文章后更新数据库
	router.POST("/topic/modifyIn", controller.TopicModifyIn)

	//启动服务器
	router.Run(":8080")
}
