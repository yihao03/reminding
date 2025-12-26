package userview

import database "github.com/yihao03/reminding/internal/database/sqlc"

type CreateUserView struct {
	FirebaseUID string `json:"uid" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	DisplayName string `json:"displayName" validate:"required,max=100"`
}

func (v *CreateUserView) ToCreateUserParams() *database.CreateUserParams {
	return &database.CreateUserParams{
		FirebaseUid: v.FirebaseUID,
		Email:       v.Email,
		DisplayName: v.DisplayName,
	}
}
