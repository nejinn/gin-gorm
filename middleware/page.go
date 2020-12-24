package middleware

import (
	"github.com/gin-gonic/gin"
	"note/global"
	"note/model"
	"strings"
)

// 记录用户访问页面和接口数据
func VisitLog() func(c *gin.Context) {
	return func(c *gin.Context) {
		user := GetJwtUser(c)
		if url, isExist := c.GetQuery("url"); isExist == true {
			var VisitPageData model.VisitPage
			var VisitNav model.Nav
			VisitPageData.UserId = user.ID
			VisitPageData.VisitIp = c.ClientIP()
			err := global.NLY_DB.Where("router = ?", url).First(&VisitNav).Error
			if err == nil {
				VisitPageData.UrlId = VisitNav.ID
				if err := global.NLY_DB.Create(&VisitPageData).Error; err != nil {
					c.Abort()
				}
			}
		}
		method := c.Request.Method
		if method != "OPTIONS" {
			visitApiUrl := c.Request.RequestURI
			visitApiUrlArray := strings.SplitN(visitApiUrl, "?", 2)
			var VisitApiData model.VisitApi
			VisitApiData.Url = visitApiUrlArray[0]
			VisitApiData.VisitIp = c.ClientIP()
			VisitApiData.UserId = user.ID
			_ = global.NLY_DB.Create(&VisitApiData).Error
		}
		c.Next()
	}
}
