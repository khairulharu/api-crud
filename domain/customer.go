package domain

import (
	"context"
	"time"
)

type Customer struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindById(ctx context.Context, id int) (Customer, error)
	FindByIds(ctx context.Context, ids int) ([]Customer, error)
	FindByPhone(ctx context.Context, phone string) (Customer, error)
	Insert(ctx context.Context, customer *Customer) error
}

type CustomerService interface {
	All(ctx context.Context) ApiResponse
	Save(ctx context.Context, customer CustomerData) ApiResponse
}

type CustomerData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
