package modules_user

func NewUser() User {
	return User{}
}

// Getters

func (u User) GetUserId() int {
	return u.UserId
}

func (u User) GetUserTgId() string {
	return u.UserTgId
}

func (u User) GetUserName() string {
	return u.UserName
}

func (u User) GetUserAddressList() []string {
	return u.UserAddress
}

func (u User) GetUserAddress(n int) string {
	return u.UserAddress[n]
}

// Setters

func (u *User) SetUserId(id int) {
	u.UserId = id
}

func (u *User) SetUserTgId(id string) {
	u.UserTgId = id
}

func (u *User) SetUserName(name string) {
	u.UserName = name
}

func (u *User) AddUserAddress(address string) {
	u.UserAddress = append(u.UserAddress, address)
}
