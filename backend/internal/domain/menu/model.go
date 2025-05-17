package modules_menu

type Menu struct {
	MenuId       int    `json:"menu_id"`
	MenuName     string `json:"menu_name"`
	MenuCategory int    `json:"menu_category"`
	MenuPrice    int    `json:"menu_price"`
}
