package userview

import database "github.com/yihao03/reminding/internal/database/sqlc"

type CreateUserView struct {
	FirebaseUid string `json:"firebase_uid"`
	Email       string `json:"email"`
	Name        string `json:"name"`
}

func (v *CreateUserView) ToCreateUserParams() *database.CreateUserParams {
	return &database.CreateUserParams{
		FirebaseUid: v.FirebaseUid,
		Email:       v.Email,
		UserName:    v.Name,
	}
}
