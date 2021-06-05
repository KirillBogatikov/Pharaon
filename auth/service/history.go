package service

import (
	ptime "github.com/Projector-Solutions/Pharaon-tools/time"
	"github.com/google/uuid"
	"pharaon-auth/data"
	"time"
)

func SaveToHistory(id uuid.UUID, ip string) (*data.History, error) {
	historyId := uuid.New()
	history := data.History{
		Id:     &historyId,
		AuthId: &id,
		Ip:     ip,
		Time:   ptime.NewDateTime(time.Now().UTC()),
	}

	err := historyRepo.Insert(history)
	if err != nil {
		return nil, err
	}

	return &history, nil
}

func GetHistory(id uuid.UUID) (*data.History, error) {
	return historyRepo.Get(id)
}

func ListHistory(id uuid.UUID) ([]data.History, error) {
	list, err := historyRepo.List(id)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func DeleteFromHistory(id uuid.UUID) (bool, error) {
	found, err := historyRepo.Delete(id)
	if err != nil {
		return false, err
	}

	return found, nil
}
