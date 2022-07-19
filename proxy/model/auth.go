package model

type LoginSignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginSignupResponse struct {
	Token string `json:"token"`
}

type VerifyTokenResponse struct {
	IsAuth bool `json:"isAuth"`
}
