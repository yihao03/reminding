package userview

type AuthView struct {
	User      UserView
	UserToken string `json:"idToken" validate:"required"`
}
