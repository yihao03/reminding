package userview

import (
	"github.com/yihao03/reminding/internal/database"
	"github.com/yihao03/reminding/internal/database/sqlc"
)

type CreateUserView struct {
	FirebaseUID string `json:"uid" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	DisplayName string `json:"displayName" validate:"required,max=100"`
	State       string `json:"state" validate:"oneof='Johor' 'Kedah' 'Kelantan' 'Melaka' 'Negeri Sembilan' 'Pahang' 'Perak' 'Perlis' 'Penang' 'Sabah' 'Sarawak' 'Selangor' 'Terengganu'"`
	Age         int32  `json:"age" validate:"min=0,max=120"`
}

func (v *CreateUserView) ToCreateUserParams() *sqlc.CreateUserParams {
	return &sqlc.CreateUserParams{
		FirebaseUid: v.FirebaseUID,
		Email:       v.Email,
		DisplayName: v.DisplayName,
		Age:         database.ToPGInt4(&v.Age),
	}
}
