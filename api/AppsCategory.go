package api

type AppsCategory struct {
	Id     string         `json:"id"`
	Name   string         `json:"name"`
	Tabs   map[string]Tab `json:"tabs"`
	Status int            `json:"status"`
}

type Tab struct {
	Name           string `json:"name"`
	TotalItemCount int    `json:"total_item_count"`
	Items          []Item `json:"items"`
}

type Item struct {
	ItemType int `json:"type"`
	Id       int `json:"id"`
}
