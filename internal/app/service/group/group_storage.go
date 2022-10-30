package group

import (
	"github.com/vaberof/TelegramBotUniversitySchedule/internal/infra/storage/postgres/grouppg"
	"gorm.io/gorm"
)

type GroupStorage interface {
	CreateGroup(id string, name string, externalId string) error
	UpdateGroup(id string, name string, externalId string, newId string, newName string, newExternalId string) error
	DeleteGroup(id string, name string, externalId string) error
}

type GroupStoragePostgres struct {
	GroupStorage
}

func NewGroupStoragePostgres(db *gorm.DB) *GroupStoragePostgres {
	return &GroupStoragePostgres{GroupStorage: grouppg.NewGroupStoragePostgres(db)}
}