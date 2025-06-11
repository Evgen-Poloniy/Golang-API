package handler

// var allowedAttributes = [4]string{
// 	"user_id",
// 	"username",
// 	"name",
// 	"surname",
// }

const countAllowedAttributes int = 4

type allowedAttributes struct {
	ID       uint32 `json:"user_id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}
