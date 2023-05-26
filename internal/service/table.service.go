package service

import (
	"restaurant-management/internal/forms"
	"restaurant-management/internal/models"
	repo "restaurant-management/internal/repository"
	"restaurant-management/internal/se"

	"github.com/casbin/casbin/v2"
	"github.com/docker/distribution/uuid"
)

type TableService interface {
	Add(req *forms.Table) (*models.Table, *se.ServiceError)
	Get(tableId string) (*models.Table, *se.ServiceError)
	GetAll() ([]*models.Table, *se.ServiceError)
	Delete(tableId string) *se.ServiceError
}

type tableSrv struct {
	repo         repo.TableRepo
	validatorSrv ValidationService
	cashbin      *casbin.Enforcer
}

func (ta *tableSrv) Add(req *forms.Table) (*models.Table, *se.ServiceError) {
	if err := ta.validatorSrv.Validate(req); err != nil {
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
	_, err := uuid.Parse(tableId)
	if err != nil {
		return nil, se.NotFound("table not found")
	}

	tabl, err := ta.repo.Get(tableId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err, "table not found")
	}

	return tabl, nil
}

func (ta *tableSrv) GetAll() ([]*models.Table, *se.ServiceError) {
	tables, err := ta.repo.GetAll()
	if err != nil {
		return nil, se.Internal(err)
	}

	return tables, nil
}

func (ta *tableSrv) Delete(tableId string) *se.ServiceError {
	_, err := uuid.Parse(tableId)
	if err != nil {
		return se.NotFound("table not found")
	}

	err = ta.repo.Delete(tableId)
	if err != nil {
		return se.NotFoundOrInternal(err, "table not found")
	}

	return nil
}

func NewTableService(repo repo.TableRepo, validatorSrv ValidationService, cashbin *casbin.Enforcer) TableService {
	return &tableSrv{repo: repo, validatorSrv: validatorSrv, cashbin: cashbin}
}
