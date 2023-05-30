package handlers

const (
	UserPath = "/api/v1/users"

	PolicyMethodGet    = "GET"
	PolicyMethodPost   = "POST"
	PolicyMethodPatch  = "PATCH"
	PolicyMethodAll    = "(GET)|(POST)|(DELETE)|(PATCH)|(PUT)"
	PolicyMethodDelete = "DELETE"

	PolicyEffectAllow = "allow"
	PolicyEffectDeny  = "deny"
)
