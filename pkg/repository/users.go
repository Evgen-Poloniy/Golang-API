package repository

type Users struct {
	ID       int     `json:"user_id"`
	Niсkname string  `json:"nickname"`
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Password string  `json:"password"`
	Coins    float64 `json:"coins"`
}

type Messages struct {
	ID           int    `json:"message_id"`
	Text         string `json:"text"`
	Time         string `json:"sending_time"`
	Status       int    `json:"message_status"`
	ID_sender    int    `json:"sender_id"`
	ID_recipient int    `json:"id_recipient"`
}

type Transactions struct {
	ID           int     `json:"transaction_id"`
	Amount       float64 `json:"amount"`
	ID_sender    int     `json:"sender_id"`
	ID_recipient int     `json:"id_recipient"`
}
