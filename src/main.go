package main

import (
	"htmx/src/shared/infra/env"
	"htmx/src/shared/infra/server"
	"htmx/src/shared/utils/logger"
	user_model "htmx/src/users/domain/models"
)

func main() {
	env, _ := env.NewConfig()

	user := user_model.UserModel{Id: 1, Name: "Victor", Email: "victor@htmx", Status: "Active"}

	logger.Default.Log(user.String())

	server.Serve(&env)
}
