package attribute

type AuthField struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
}

type ActionField struct {
	ID       uint32  `json:"user_id"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Coins    float64 `json:"coins"`
}
