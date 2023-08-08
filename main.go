package main

import "bulugen-backend-go/cmd"

// @title Go-Gin-Web
// @version 1.0.0
// @description learning something about golang.
func main() {
	defer cmd.Clean()
	cmd.Start()
}
