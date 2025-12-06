package userview

import database "github.com/yihao03/reminding/internal/database/sqlc"

type CreateUserView struct {
	FirebaseUID string `json:"firebase_uid" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Name        string `json:"name" validate:"required,max=100"`
}

func (v *CreateUserView) ToCreateUserParams() *database.CreateUserParams {
	return &database.CreateUserParams{
		FirebaseUid: v.FirebaseUID,
		Email:       v.Email,
		UserName:    v.Name,
	}
}
