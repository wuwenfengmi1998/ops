package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	fmt.Println("http://" + "127.0.0.1:8080")
	r.Run("127.0.0.1:8080")

}
