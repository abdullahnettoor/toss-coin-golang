package handlers

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func TossCoin(ctx *gin.Context) {
	res := rand.Int()%2 == 0
	fmt.Println(res)
	ctx.JSON(200, gin.H{"res": res})
}
