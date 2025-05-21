package modules_menu

func NewMenu() Menu {
	return &MenuImpl{}
}

// Getters

func (m MenuImpl) GetMenuId() int {
	return m.MenuId
}

func (m MenuImpl) GetMenuName() string {
	return m.MenuName
}

func (m MenuImpl) GetMenuCategory() int {
	return m.MenuCategory
}

func (m MenuImpl) GetMenuPrice() int {
	return m.MenuPrice
}

//Setters

func (m *MenuImpl) SetMenuId(id int) {
	m.MenuId = id
}

func (m *MenuImpl) SetMenuName(name string) {
	m.MenuName = name
}

func (m *MenuImpl) SetMenuCategory(cat int) {
	m.MenuCategory = cat
}

func (m *MenuImpl) SetMenuPrice(pri int) {
	m.MenuPrice = pri
}
