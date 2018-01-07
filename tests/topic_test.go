package tests

import (
	"net/http/httptest"
	"testing"

	controller "github.com/tccip/simpleblog/controllers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()

	//加载静态资源
	router.Static("/static", "./static")
	//加载网页文件
	router.LoadHTMLGlob("../views/*")

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
}

// Get 根据特定请求uri，发起get请求返回响应
func Get(uri string, router *gin.Engine) string {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body := result.Status
	return body
}

func ParseToStr(mp map[string]string) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp
	return values
}

// PostForm 根据特定请求uri和参数param，以表单形式传递参数，发起post请求返回响应
func PostForm(uri string, param map[string]string, router *gin.Engine) string {
	// 构造post请求，表单数据以querystring的形式加在uri之后
	req := httptest.NewRequest("POST", uri+ParseToStr(param), nil)

	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	body := result.Status
	return body
}

//router.GET("/", controller.Home)
func TestHome(t *testing.T) {
	// 初始化请求地址
	uri := "/"

	// 发起Get请求
	body := Get(uri, router)

	// 判断响应是否与预期一致
	if !(body != "ok") {
		t.Errorf("响应字符串不符，body:%v\n", string(body))
	}
}

func TestTopic(t *testing.T) {
	// 初始化请求地址
	uri := "/topic"

	// 发起Get请求
	body := Get(uri, router)

	// 判断响应是否与预期一致
	if !(body != "ok") {
		t.Errorf("响应字符串不符，body:%v\n", string(body))
	}
}

func TestTopicView(t *testing.T) {
	// 初始化请求地址
	uri := "/topic/view"

	// 发起Get请求
	body := Get(uri, router)

	// 判断响应是否与预期一致
	if !(body != "ok") {
		t.Errorf("响应字符串不符，body:%v\n", string(body))
	}
}

func TestTopicAdd(t *testing.T) {
	// 初始化请求地址
	uri := "/topic/add"

	// 发起Get请求
	body := Get(uri, router)

	// 判断响应是否与预期一致
	if !(body != "ok") {
		t.Errorf("响应字符串不符，body:%v\n", string(body))
	}
}

/* func TestTopicAddIn(t *testing.T) {
	// 初始化请求地址
	uri := "/topic/addIn"

	param := make(map[string]string)
	param["title"] = "this is test"
	param["content"] = "test content"
	// 发起Get请求
	body := PostForm(uri, param, router)
	fmt.Println(body)
	// 判断响应是否与预期一致
	if !(body != "ok") {
		t.Errorf("响应字符串不符，body:%v\n", string(body))
	}
} */

func TestTopicDel(t *testing.T) {
	// 初始化请求地址
	uri := "/topic/del"

	// 发起Get请求
	body := Get(uri, router)

	// 判断响应是否与预期一致
	if !(body != "ok") {
		t.Errorf("响应字符串不符，body:%v\n", string(body))
	}
}

func TestTopicModify(t *testing.T) {
	// 初始化请求地址
	uri := "/topic/modify"

	// 发起Get请求
	body := Get(uri, router)

	// 判断响应是否与预期一致
	if !(body != "ok") {
		t.Errorf("响应字符串不符，body:%v\n", string(body))
	}
}

/* func TestTopicModifyIn(t *testing.T) {
	// 初始化请求地址
	uri := "/topic/modifyIn"

	param := make(map[string]string)
	param["id"] = "1"
	param["title"] = "this is test"
	param["content"] = "test content"

	// 发起Get请求
	body := PostForm(uri, param, router)

	// 判断响应是否与预期一致
	if !(body != "ok") {
		t.Errorf("响应字符串不符，body:%v\n", string(body))
	}
} */
