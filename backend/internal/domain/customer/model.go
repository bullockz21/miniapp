package modules_customer

type Customer interface {
	GetCustomerId() int
	GetCustomerTgId() string
	GetCustomerName() string
	GetCustomerAddressList() []string
	GetCustomerAddress(n int) string
	SetCustomerId(id int)
	SetCustomerTgId(id string)
	SetCustomerName(name string)
	AddCustomerAddress(address string)
}

type CustomerImpl struct {
	CustomerId      int      `json:"customer_id"`
	CustomerTgId    string   `json:"customer_telegram_id"`
	CustomerName    string   `json:"customer_name"`
	CustomerAddress []string `json:"customer_address"`
}
