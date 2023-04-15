package response

type AuthCallbackResponse struct {
	Token string `json:"token"`
	IsNew bool   `json:"is_new"`
}
