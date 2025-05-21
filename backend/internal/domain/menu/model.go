package modules_menu

type Menu interface {
	GetMenuId() int
	GetMenuName() string
	GetMenuCategory() int
	GetMenuPrice() int
	SetMenuId(id int)
	SetMenuName(name string)
	SetMenuCategory(cat int)
	SetMenuPrice(pri int)
}

type MenuImpl struct {
	MenuId       int    `json:"menu_id"`
	MenuName     string `json:"menu_name"`
	MenuCategory int    `json:"menu_category"`
	MenuPrice    int    `json:"menu_price"`
}
