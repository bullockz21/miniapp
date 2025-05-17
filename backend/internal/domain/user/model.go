package modules_user

type User struct {
	UserId      int      `json:"user_id"`
	UserTgId    string   `json:"user_telegram_id"`
	UserName    string   `json:"user_name"`
	UserAddress []string `json:"user_address"`
}
