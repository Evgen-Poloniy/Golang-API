package repository

type Users struct {
	ID           uint32  `json:"user_id"`
	Username     string  `json:"username"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	PasswordHash string  `json:"password_hash"`
	Coins        float64 `json:"coins"`
}

// type Messages struct {
// 	ID     uint32 `json:"message_id"`
// 	Text   string `json:"text"`
// 	Time   string `json:"sending_time"`
// 	Status int    `json:"message_status"`
// 	ChatID uint64 `json:"chat_id"`
// }

// type Chats struct {
// 	ID      uint32   `json:"chat_id"`
// 	UsersID []uint32 `json:"users_id"`
// }

type Transactions struct {
	ID              uint32  `json:"transaction_id"`
	SenderID        uint32  `json:"sender_id"`
	RecipientID     uint32  `json:"recipient_id"`
	Amount          float64 `json:"amount"`
	TransactionTime string  `json:"transaction_time"`
}
