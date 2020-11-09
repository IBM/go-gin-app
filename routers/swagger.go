package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SwaggerDocRedirect(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/swagger/api-docs/index.html")
}

func SwaggerAPI(c *gin.Context) {
	c.File("./public/swagger.yaml")
}
