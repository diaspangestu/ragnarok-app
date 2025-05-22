package _interface

import (
	"ragnarok/internal/dto"
	"ragnarok/internal/model"
)

type BrimoExpeditionRepository interface {
	GetListBrimoExpedition(payload dto.BrimoExpeditionRequest) ([]model.BrimoExpedition, int, error)
}

type BrimoExpeditionService interface {
	GetListBrimoExpedition(payload dto.BrimoExpeditionRequest) ([]dto.BrimoExpeditionResponse, int, error)
}
