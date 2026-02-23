package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	var dpiUint uint
	router := gin.Default()

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dpiString := os.Getenv("PDF_DPI")

	if dpiString != "" {
		dpi, err := strconv.Atoi(dpiString)
		if err != nil {
			log.Fatal(err)
		}
		dpiUint = uint(dpi)
	} else {
		dpiUint = 1200
	}

	fmt.Printf("Info: PDF dpi set to: %d\n", dpiUint)

	var authorized *gin.RouterGroup
	if user != "" && password != "" {
		fmt.Println("* PROTECTED BY USER AND PASSWORD *")

		authorized = router.Group("/", gin.BasicAuth(gin.Accounts{
			user: password,
		}))
	} else {
		fmt.Println("* OPEN ACCESS *")

		authorized = router.Group("/")
	}

	// / open endpoint
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "你好\nHello\nHola")
	})

	// /pdf protected endpoint
	authorized.POST("/pdf", func(c *gin.Context) {
		html := c.PostForm("html")

		pdf, err := GeneratePDF(html, dpiUint)
		if err != nil {
			fmt.Println(err)
			c.Writer.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		c.String(http.StatusOK, base64.StdEncoding.EncodeToString(pdf))
	})

	return router
}
