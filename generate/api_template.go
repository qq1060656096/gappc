package generate

const apiTemplate = `package {{packageName}}

import (
	"github.com/gin-gonic/gin"
	"github.com/qq1060656096/go-api"
	// "github.com/qq1060656096/common"
	"net/http"
	"strconv"
)

func {{apiName}}Create(c *gin.Context) {
	// db, _ := common.GetDefaultDbConn()
	c.JSON(http.StatusOK, api.NewResult("200", "success").Simple(nil))
}

func {{apiName}}Delete(c *gin.Context) {
	// db, _ := common.GetDefaultDbConn()
	c.JSON(http.StatusOK, api.NewResult("200", "success").Simple(nil))
}

func {{apiName}}Update(c *gin.Context) {
	// db, _ := common.GetDefaultDbConn()
	c.JSON(http.StatusOK, api.NewResult("200", "success").Simple(nil))
}

func {{apiName}}Get(c *gin.Context) {
	count := 0
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	offset := (page - 1) * pageSize
	c.JSON(http.StatusOK, api.NewResult("200", "success").Paging(nil, count, page, pageSize))
}

func {{apiName}}GetDetail(c *gin.Context) {
	// db, _ := common.GetDefaultDbConn()
	c.JSON(http.StatusOK, api.NewResult("200", "success").Simple(nil))
}
`