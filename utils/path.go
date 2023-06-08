package utils

const (
	BasePath  = "/v1"
	AuthPath  = "/v1/auth"
	UserPath  = "/v1/users"
	TablePath = "/v1/tables"
	MenuPath  = "/v1/menus"
	FoodPath  = "/v1/foods"
	OrderPath = "/v1/orders"

	PolicyMethodGet    = "GET"
	PolicyMethodPost   = "POST"
	PolicyMethodPatch  = "PATCH"
	PolicyMethodAll    = "(GET)|(POST)|(DELETE)|(PATCH)|(PUT)"
	PolicyMethodDelete = "DELETE"

	PolicyEffectAllow = "allow"
	PolicyEffectDeny  = "deny"
)
