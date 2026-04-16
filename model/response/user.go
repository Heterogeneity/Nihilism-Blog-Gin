package response

import (
	"github.com/gofrs/uuid"
	"server/model/database"
)

type Login struct {
	User                database.User `json:"user"`
	AccessToken         string        `json:"access_token"`
	AccessTokenExpireAt int64         `json:"access_token_expire_at"`
}

type UserCard struct {
	UID       uuid.UUID `json:"uid"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar"`
	Address   string    `json:"address"`
	Signature string    `json:"signature"`
}

type UserChart struct {
	DateList     []string `json:"date_list"`
	LoginData    []int    `json:"login_data"`
	RegisterData []int    `json:"register_data"`
}
