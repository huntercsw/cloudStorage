package apps

import "github.com/gin-gonic/gin"

func RouterSetUp(r *gin.Engine) {
	r.GET("/api/test", HttpHandlerTest)

}
