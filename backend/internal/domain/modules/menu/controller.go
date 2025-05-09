package modules_menu

func NewMenu() Menu {
	return Menu{}
}

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
