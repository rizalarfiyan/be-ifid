package response

type AuthTokenResponse struct {
	Token string `json:"token"`
	IsNew bool   `json:"is_new"`
}
