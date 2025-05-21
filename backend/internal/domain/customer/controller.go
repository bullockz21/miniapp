package modules_customer

func NewCustomer() Customer {
	return &CustomerImpl{}
}

// Getters

func (u CustomerImpl) GetCustomerId() int {
	return u.CustomerId
}

func (u CustomerImpl) GetCustomerTgId() string {
	return u.CustomerTgId
}

func (u CustomerImpl) GetCustomerName() string {
	return u.CustomerName
}

func (u CustomerImpl) GetCustomerAddressList() []string {
	return u.CustomerAddress
}

func (u CustomerImpl) GetCustomerAddress(n int) string {
	return u.CustomerAddress[n]
}

// Setters

func (u *CustomerImpl) SetCustomerId(id int) {
	u.CustomerId = id
}

func (u *CustomerImpl) SetCustomerTgId(id string) {
	u.CustomerTgId = id
}

func (u *CustomerImpl) SetCustomerName(name string) {
	u.CustomerName = name
}

func (u *CustomerImpl) AddCustomerAddress(address string) {
	u.CustomerAddress = append(u.CustomerAddress, address)
}
