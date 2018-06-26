package oauth

type (
	// TokenModel sent to oauth/token path
	TokenModel struct {
		GrantType string `json:"grant_type" binding:"required"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Scope     string `json:"scope"`
	}

	// TransformedToken response
	TransformedToken struct {
		AccessToken string `json:"access_token"`
	}
)
