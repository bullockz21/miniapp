package modules_menu

type Menu interface {
}

type MenuImpl struct {
	MenuId       int    `json:"menu_id"`
	MenuName     string `json:"menu_name"`
	MenuCategory int    `json:"menu_category"`
	MenuPrice    int    `json:"menu_price"`
}
