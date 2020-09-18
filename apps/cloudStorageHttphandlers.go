package apps

import (
	csLog "CloudStorage/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpHandlerTest(c *gin.Context) {
	csLog.Logger.Error("test")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello yinuo",
	})
}
