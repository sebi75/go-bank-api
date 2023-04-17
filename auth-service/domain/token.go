package domain

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	CustomerId string `json:"customer_id"`
	Role       string `json:"role"`
	Username   string `json:"username"`
	Expiry	 int64  `json:"expiry"`
	Accounts  []string `json:"accounts"`
}

func (c Claims) IsUserRole() bool {
	return c.Role == "USER"
}

func (c Claims) IsAdminRole() bool {
	return c.Role == "ADMIN"
}

func (c Claims) IsRequestVerifiedWithTokenClaims(urlParams map[string]string) bool {
	if c.IsUserRole() {
		return c.CustomerId == urlParams["customer_id"]
	}

	if !c.IsValidAccountId(urlParams["account_id"]) {
		return false
	}

	return true
}

func (c Claims) IsValidAccountId(accountId string) bool {
	for _, account := range c.Accounts {
		if account == accountId {
			return true
		}
	}
	return false
}

func (c Claims) IsAccountOwner(accountId string) bool {
	return c.IsValidAccountId(accountId)
}

func BuildClaimsFromJwtMapClaims(mapClaims jwt.MapClaims) (*Claims, error) {
	claims := Claims{
		CustomerId: mapClaims["customer_id"].(string),
		Role:       mapClaims["role"].(string),
		Username:   mapClaims["username"].(string),
		Expiry:     int64(mapClaims["expiry"].(float64)),
		Accounts:   mapClaims["accounts"].([]string),
	}
	return &claims, nil
}