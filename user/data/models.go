package data

import "github.com/google/uuid"

type DBUser struct {
	Id             uuid.UUID `db:"user_id"`
	CredentialsId  uuid.UUID `db:"credentials_id"`
	PersonalDataId uuid.UUID `db:"personal_data_id"`
}

type UserIdList struct {
	Ids []uuid.UUID `db:"user_ids"`
}
