package user_model

import (
	"htmx/src/shared/utils"
)

type UserModel struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func (u UserModel) String() string {
	return utils.JsonifyPretty(u)
}
