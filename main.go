package main

import (
	"context"

	"github.com/doffy007/user-login-register/config"
	_ "github.com/doffy007/user-login-register/docs"
	"github.com/doffy007/user-login-register/package/server"
)

// @title 		Tag User Login And Register API
// @version 	1.0
// @description A Tag service API in Golang

// @host 		localhost:8081
// @BasePath 	/api
func main() {
	config.Init()

	ctx := context.Background()
	server.NewApp(ctx).Start()
}
