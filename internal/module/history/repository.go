package history

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/khairulharu/miniapps/domain"
)

type repository struct {
	db *goqu.Database
}

func NewRepository(con *sql.DB) domain.HistoryRepository {
	return &repository{
		db: goqu.New("default", con),
	}
}

func (r repository) FindByVehicle(ctx context.Context, id int) (histories []domain.HistoryDetail, err error) {
	dataset := r.db.From("history_details").Where(goqu.Ex{
		"vehicle_id": id,
	}).Order(goqu.I("id").Asc())
	err = dataset.ScanStructsContext(ctx, &histories)
	return
}

func (r repository) Insert(ctx context.Context, history *domain.HistoryDetail) error {
	history.CreatedAt = time.Now()
	executor := r.db.Insert("history_details").Rows(goqu.Record{
		"vehicle_id":   history.VehicleID,
		"pic":          history.PIC,
		"notes":        history.Notes,
		"customer_id":  history.CustomerID,
		"plate_number": history.PlateNumber,
		"created_at":   history.CreatedAt,
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, history)
	return err
}
