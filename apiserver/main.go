package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin default running in "debug" mode.Switch to "release" mode in production using following code.
	// gin.SetMode(gin.ReleaseMode)

	g := gin.New()

	fmt.Printf("%s", http.ListenAndServe(":8088", g).Error())
}
