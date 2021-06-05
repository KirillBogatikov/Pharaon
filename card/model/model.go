package model

import (
	"github.com/Projector-Solutions/Pharaon-api/user"
	"github.com/google/uuid"
)

type Type int
type Priority int

const (
	TypeTask     Type = 1
	TypeError    Type = 2
	TypeReport   Type = 3
	TypeProposal Type = 4

	PriorityLow    Priority = 1
	PriorityNormal Priority = 2
	PriorityHigh   Priority = 3
	PriorityTop    Priority = 4
)

type Tag struct {
	Id   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
}

type Card struct {
	Id          *uuid.UUID  `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Author      user.User   `json:"author"`
	Executor    *user.User  `json:"executor,omitempty"`
	Type        Type        `json:"type"`
	Priority    Priority    `json:"priority"`
	Observers   []user.User `json:"observers"`
	Tags        []Tag       `json:"tags"`
}
