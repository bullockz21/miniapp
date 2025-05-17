package modules_customer

func NewCustomer() Customer {
	return Customer{}
}

// Getters

func (u Customer) GetCustomerId() int {
	return u.CustomerId
}

func (u Customer) GetCustomerTgId() string {
	return u.CustomerTgId
}

func (u Customer) GetCustomerName() string {
	return u.CustomerName
}

func (u Customer) GetCustomerAddressList() []string {
	return u.CustomerAddress
}

func (u Customer) GetCustomerAddress(n int) string {
	return u.CustomerAddress[n]
}

// Setters

func (u *Customer) SetCustomerId(id int) {
	u.CustomerId = id
}

func (u *Customer) SetCustomerTgId(id string) {
	u.CustomerTgId = id
}

func (u *Customer) SetCustomerName(name string) {
	u.CustomerName = name
}

func (u *Customer) AddCustomerAddress(address string) {
	u.CustomerAddress = append(u.CustomerAddress, address)
}
