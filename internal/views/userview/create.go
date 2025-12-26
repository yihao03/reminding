package userview

import database "github.com/yihao03/reminding/internal/database/sqlc"

type CreateUserView struct {
	FirebaseUID string `json:"FirebaseUID" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Name        string `json:"UserName" validate:"required,max=100"`
}

func (v *CreateUserView) ToCreateUserParams() *database.CreateUserParams {
	return &database.CreateUserParams{
		FirebaseUid: v.FirebaseUID,
		Email:       v.Email,
		DisplayName: v.Name,
	}
}
