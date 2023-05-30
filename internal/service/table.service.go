package service

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"

	"github.com/docker/distribution/uuid"
)

type TableService interface {
	Add(req *forms.Table) (*models.Table, *se.ServiceError)
	Get(tableId string) (*models.Table, *se.ServiceError)
	GetAll(role string) ([]*models.Table, *se.ServiceError)
	Edit(tableId string, req *forms.EditTable) (*models.Table, *se.ServiceError)
	Delete(tableId string) *se.ServiceError
}

type tableSrv struct {
	repo repo.TableRepo
}

func (ta *tableSrv) Add(req *forms.Table) (*models.Table, *se.ServiceError) {
	if err := Validator.Validate(req); err != nil {
		return nil, se.Validating(err)
	}

	var table models.Table

	table.Booked = false
	table.Seats = req.Seats

	tabl, err := ta.repo.Add(&table)
	if err != nil {
		return nil, se.Internal(err, "error when creating table")
	}

	return tabl, nil
}

func (ta *tableSrv) Get(tableId string) (*models.Table, *se.ServiceError) {
	if _, err := uuid.Parse(tableId); err != nil {
		return nil, se.Internal(err, "invalid table id")
	}

	tabl, err := ta.repo.Get(tableId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "table not found")
	}

	return tabl, nil
}

func (ta *tableSrv) GetAll(role string) ([]*models.Table, *se.ServiceError) {
	tables, err := ta.repo.GetAll(role)
	if err != nil {
		return nil, se.Internal(err)
	}

	return tables, nil
}

func (ta *tableSrv) Edit(tableId string, req *forms.EditTable) (*models.Table, *se.ServiceError) {
	if err := Validator.Validate(req); err != nil {
		return nil, se.Validating(err)
	}

	if _, err := uuid.Parse(tableId); err != nil {
		return nil, se.Internal(err, "invalid table id")
	}

	tabl, err := ta.repo.Get(tableId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "user not found")
	}

	table := ta.getEdit(req, tabl)

	table, err = ta.repo.Edit(tableId, table)
	if err != nil {
		return nil, se.Internal(err, "error when editing user")
	}

	return table, nil
}

func (ta *tableSrv) Delete(tableId string) *se.ServiceError {
	if _, err := uuid.Parse(tableId); err != nil {
		return se.Internal(err, "invalid table id")
	}

	err := ta.repo.Delete(tableId)
	if err != nil {
		return se.NotFoundOrInternal(err, "table not found")
	}

	return nil
}

func NewTableService(repo repo.TableRepo) TableService {
	return &tableSrv{repo: repo}
}

func (ta *tableSrv) getEdit(req *forms.EditTable, table *models.Table) *models.Table {
	if req.Booked != table.Booked {
		table.Booked = req.Booked
	}

	if req.Seats != 0 && req.Seats != table.Seats {
		table.Seats = req.Seats
	}

	return table
}
