package service

import (
	"ragnarok/internal/dto"
	"ragnarok/internal/interface"
)

type BrimoExpeditionSvc struct {
	brimoRepo _interface.BrimoExpeditionRepository
}

func NewBrimoExpeditionService(repo _interface.BrimoExpeditionRepository) _interface.BrimoExpeditionService {
	return &BrimoExpeditionSvc{
		brimoRepo: repo,
	}
}

func (svc *BrimoExpeditionSvc) GetListBrimoExpedition(payload dto.BrimoExpeditionRequest) ([]dto.BrimoExpeditionResponse, int, error) {
	data, totalRecords, err := svc.brimoRepo.GetListBrimoExpedition(payload)
	if err != nil {
		return nil, 0, err
	}

	var results []dto.BrimoExpeditionResponse
	for _, item := range data {
		results = append(results, dto.BrimoExpeditionResponse{
			Id:         item.Id,
			PostalCode: item.PostalCode,
			TimeZone:   item.TimeZone,
			Province:   item.Province,
			City:       item.City,
			Expedition: item.Expedition,
			Sla:        item.Sla,
		})
	}

	return results, totalRecords, nil
}
