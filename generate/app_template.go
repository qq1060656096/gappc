package generate

const appTemplate = `package {{packageName}}
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func DemoAppCreate(c *gin.Context) {
	data := gin.H{
		"title":   "create",
		"nowDate": time.Now().Format("2006-03-01 00:00:00"),
		"content": "POST请求演示",
	}
	c.HTML(http.StatusOK, "default/demo/create.tmpl", data)
}

func DemoAppDelete(c *gin.Context) {
	data := gin.H{
		"title":   "delete",
		"nowDate": time.Now().Format("2006-03-01 00:00:00"),
		"content": "DELETE 请求演示",
	}
	c.HTML(http.StatusOK, "default/demo/delete.tmpl", data)
}

func DemoAppUpdate(c *gin.Context) {
	data := gin.H{
		"title":   "put",
		"nowDate": time.Now().Format("2006-03-01 00:00:00"),
		"content": "PUT 请求演示",
	}
	c.HTML(http.StatusOK, "default/demo/update.tmpl", data)}

func DemoAppGet(c *gin.Context) {
	data := gin.H{
		"title":   "get",
		"nowDate": time.Now().Format("2006-03-01 00:00:00"),
		"content": "GET 请求演示",
	}
	c.HTML(http.StatusOK, "default/demo/list.tmpl", data)
}

func DemoAppGetDetail(c *gin.Context) {
	// db, _ := common.GetDefaultDbConn()
	data := gin.H{
		"title":   "get",
		"nowDate": time.Now().Format("2006-03-01 00:00:00"),
		"content": "GET 请求演示",
	}
	c.HTML(http.StatusOK, "default/demo/detail.tmpl", data)
}

`