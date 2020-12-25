package v1

import (
	"github.com/gin-gonic/gin"
	"note/response"
	"note/view"
)

// infobox

func InfoBox(c *gin.Context) {
	err, data := view.InfoBox()
	if err != nil {
		response.ErrorMsgWithResponseMsg(c, response.InfoBoxError)
		return
	}
	response.OkData(c, response.Ok, data)

}
