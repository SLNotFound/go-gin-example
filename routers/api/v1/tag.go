package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-gin-example/models"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/setting"
	"go-gin-example/pkg/util"
	"net/http"
)

func GetTags(c *gin.Context) {
	// c.Query 获取 ?name=test&state=1 URL参数
	// c.DefaultQuery 获取默认值
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	// code 使用e 定义的错误编码
	code := e.SUCCESS

	// util.GetPage 保证各接口的page处理是一致的
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagsTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func AddTag(c *gin.Context) {

}
func EditTag(c *gin.Context)   {}
func DeleteTag(c *gin.Context) {}
