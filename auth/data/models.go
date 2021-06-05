package data

import (
	"github.com/Projector-Solutions/Pharaon-tools/time"
	"github.com/google/uuid"
)

type Credentials struct {
	Id       *uuid.UUID `db:"auth_id" json:"id"`
	Login    string     `db:"login" json:"login"`
	Password string     `db:"password" json:"password"`
	Email    string     `db:"email" json:"email"`
	Method   string     `db:"method" json:"-"`
	JWT      string     `json:"jwt"`
}

func (c *Credentials) Hide() {
	c.Password = ""
}

type History struct {
	Id     *uuid.UUID     `db:"history_id" json:"id"`
	AuthId *uuid.UUID     `db:"auth_id" json:"-"`
	Ip     string         `db:"ip" json:"ip"`
	Time   *time.DateTime `db:"time" json:"time"`
}

type RestoreToken struct {
	Id      *uuid.UUID     `db:"token_id" json:"id"`
	AuthId  *uuid.UUID     `db:"auth_id" json:"-"`
	Token   string         `db:"token" json:"token"`
	Expires *time.DateTime `db:"expires" json:"expires"`
}
