package response

type AuthResponse struct {
	AccessToken  string `json:"acsess_token"`
	RefreshToken string `json:"refresh_token"`
}
