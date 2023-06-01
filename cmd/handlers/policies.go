package handlers

const (
	BasePath  = "/api/v1"
	AuthPath  = "/api/v1/auth"
	UserPath  = "/api/v1/users"
	TablePath = "/api/v1/tables"
	MenuPath  = "/api/v1/menus"
	FoodPath  = "/api/v1/foods"

	PolicyMethodGet    = "GET"
	PolicyMethodPost   = "POST"
	PolicyMethodPatch  = "PATCH"
	PolicyMethodAll    = "(GET)|(POST)|(DELETE)|(PATCH)|(PUT)"
	PolicyMethodDelete = "DELETE"

	PolicyEffectAllow = "allow"
	PolicyEffectDeny  = "deny"
)
