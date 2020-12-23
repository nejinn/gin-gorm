package middleware

import (
	"github.com/gin-gonic/gin"
)

type VisitPagegUrl struct {
	Url string `json:"url"`
}

func VisitLog() func(c *gin.Context) {
	return func(c *gin.Context) {
		//pp := VisitPagegUrl{}
		//user := GetJwtUser(c)
		//if err := c.BindJSON(&pp); err == nil {
		//	var VisitPageData model.VisitPage
		//	var VisitNav model.Nav
		//	VisitPageData.UserId = user.ID
		//	VisitPageData.VisitIp = c.ClientIP()
		//	if err = global.NLY_DB.Where("router = ?", pp.Url).First(&VisitNav).Error; err == nil {
		//		VisitPageData.UrlId = VisitNav.ID
		//		if err := global.NLY_DB.Create(&VisitPageData).Error; err != nil {
		//			c.Abort()
		//		}
		//	}
		//}
		//visitApiUrl := c.Request.RequestURI
		//var VisitApiData model.VisitApi
		//VisitApiData.Url = visitApiUrl
		//VisitApiData.VisitIp = c.ClientIP()
		//VisitApiData.UserId = user.ID
		//if err := global.NLY_DB.Create(&VisitApiData).Error; err !=nil{
		//	c.Abort()
		//}
		c.Next()
	}
}
