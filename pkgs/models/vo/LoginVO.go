package vo

type LoginVO struct {
	Token string `json:"token"`
	User  UserVO `json:"user"`
}
