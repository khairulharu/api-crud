package customer

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/khairulharu/miniapps/domain"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.CustomerRepository {
	return &repository{db: goqu.New("default", con)}
}

func (r repository) FindAll(ctx context.Context) (customers []domain.Customer, err error) {
	dataset := r.db.From("customers").
		Order(goqu.I("name").Asc())
	if err := dataset.ScanStructsContext(ctx, &customers); err != nil {
		return nil, err
	}
	return
}
func (r repository) FindById(ctx context.Context, id int) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{
		"id": id,
	})
	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, err
	}
	return
}

func (r repository) FindByIds(ctx context.Context, ids int) (customer []domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{
		"id": ids,
	})
	if err := dataset.ScanStructsContext(ctx, &customer); err != nil {
		return nil, err
	}
	return
}

func (r repository) FindByPhone(ctx context.Context, phone string) (customer domain.Customer, err error) {
	dataset := r.db.From("customers").Where(goqu.Ex{
		"phone": phone,
	})
	if _, err := dataset.ScanStructContext(ctx, &customer); err != nil {
		return domain.Customer{}, err
	}
	return
}

func (r repository) Insert(ctx context.Context, customer *domain.Customer) error {
	executor := r.db.Insert("customers").Rows(goqu.Ex{
		"name":       customer.Name,
		"phone":      customer.Phone,
		"created_at": customer.CreatedAt,
	}).Returning("id").Executor()

	var customerDb domain.Customer

	_, err := executor.ScanStructContext(ctx, &customerDb)
	if err != nil {
		return err
	}
	customer.ID = customerDb.ID
	return err
}
