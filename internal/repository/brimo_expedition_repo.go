package repositories

import (
	"gorm.io/gorm"
	"ragnarok/internal/dto"
	brimoInterface "ragnarok/internal/interface"
	"ragnarok/internal/model"
	"ragnarok/internal/pkg/util"
	"strings"
)

type brimoExpeditionRepo struct {
	db *gorm.DB
}

func NewBrimoExpeditionRepo() brimoInterface.BrimoExpeditionRepository {
	return &brimoExpeditionRepo{
		db: util.DB,
	}
}

func (repo *brimoExpeditionRepo) GetListBrimoExpedition(payload dto.BrimoExpeditionRequest) ([]model.BrimoExpedition, int, error) {
	var (
		brimoExpedition []model.BrimoExpedition
		totalRecords    int
		args            []interface{}
		conditions      []string
		limit           = 10
	)

	if payload.PostalCode != "" {
		conditions = append(conditions, "postal_code LIKE ?")
		args = append(args, "%"+payload.PostalCode+"%")
	}

	if payload.City != "" {
		conditions = append(conditions, "city LIKE ?")
		args = append(args, "%"+payload.City+"%")
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}

	query := "SELECT * FROM brimo_expedition" + whereClause
	countQuery := "SELECT COUNT(*) FROM brimo_expedition" + whereClause

	err := repo.db.Raw(countQuery, args...).Scan(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (payload.Page - 1) * limit
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	err = repo.db.Raw(query, args...).Scan(&brimoExpedition).Error
	if err != nil {
		return nil, 0, err
	}

	return brimoExpedition, totalRecords, nil
}
