package modules_customer

type Customer interface {
}

type CustomerImpl struct {
	CustomerId      int      `json:"customer_id"`
	CustomerTgId    string   `json:"customer_telegram_id"`
	CustomerName    string   `json:"customer_name"`
	CustomerAddress []string `json:"customer_address"`
}
