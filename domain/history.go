package domain

import (
	"context"
	"time"
)

type HistoryDetail struct {
	ID          int       `db:"id"`
	VehicleID   int       `db:"vehicle_id"`
	Notes       string    `db:"notes"`
	PIC         string    `db:"pic"`
	CustomerID  int       `db:"customer_id"`
	PlateNumber string    `db:"plate_number"`
	CreatedAt   time.Time `db:"created_at"`
}

type HistoryRepository interface {
	FindByVehicle(ctx context.Context, id int) ([]HistoryDetail, error)
	Insert(ctx context.Context, history *HistoryDetail) error
}

type HistoryService interface {
}

type HistoricalData struct {
	VehicleID   int    `json:"vehicle_id"`
	Notes       string `json:"notes"`
	PIC         string `json:"pic"`
	CustomerID  int    `json:"customers_id"`
	PlateNumber string `json:"plate_number"`
	ComeAt      string `json:"come_at"`
}
