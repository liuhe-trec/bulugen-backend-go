package main

import (
	"bulugen-backend-go/utils"
	"fmt"
)

// @title Go-Gin-Web
// @version 1.0.0
// @description learning something about golang.
func main() {
	// defer cmd.Clean()
	// cmd.Start()
	token, _ := utils.GenerateToken(1, "zs")
	fmt.Println(token)

	iJwtCustClaims, _ := utils.ParseToken(token)
	fmt.Println(iJwtCustClaims)
}
