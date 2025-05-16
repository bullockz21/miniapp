package modules_menu

func NewMenu() Menu {
	return Menu{}
}

// Getters

func (m Menu) GetMenuId() int {
	return m.MenuId
}

func (m Menu) GetMenuName() string {
	return m.MenuName
}

func (m Menu) GetMenuCategory() int {
	return m.MenuCategory
}

func (m Menu) GetMenuPrice() int {
	return m.MenuPrice
}

//Setters

func (m *Menu) SetMenuId(id int) {
	m.MenuId = id
}

func (m *Menu) SetMenuName(name string) {
	m.MenuName = name
}

func (m *Menu) SetMenuCategory(cat int) {
	m.MenuCategory = cat
}

func (m *Menu) SetMenuPrice(pri int) {
	m.MenuPrice = pri
}
