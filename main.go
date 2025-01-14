package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	// TODO replace below sandbox app creds with your own obtained from recipient hub
	ClientId = "tsys"
)

func main() {
	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/"),
		gin.Recovery(),
	)
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":            "Getting an auth code",
			"authCodeRedirect": fmt.Sprintf("https://sandbox-idp.ddp.akoya.com/auth?connector=mikomo&response_type=code&client_id=%s&redirect_uri=https://recipient.ddp.akoya.com/flow/callback&scope=openid&offline_access&state=appstate", ClientId),
		})
	})
	r.Run(":8123")
}
