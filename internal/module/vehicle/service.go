package vehicle

import (
	"context"
	"time"

	"github.com/khairulharu/miniapps/domain"
)

type service struct {
	historyRepository domain.HistoryRepository
	vehicleRepository domain.VehicleRepository
}

func NewService(historyRepository domain.HistoryRepository, vehicleRepository domain.VehicleRepository) domain.VehicleService {
	return &service{
		historyRepository,
		vehicleRepository,
	}
}

func (s service) FindHistorical(ctx context.Context, vin string) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindByVIN(ctx, vin)
	if err != nil {
		return domain.ApiResponse{
			Code:    "911",
			Message: err.Error(),
		}
	}
	if vehicle == (domain.Vehicle{}) {
		return domain.ApiResponse{
			Code:    "01",
			Message: "VEHICLE NOT FOUND",
		}
	}

	histories, err := s.historyRepository.FindByVehicle(ctx, vehicle.ID)
	if err != nil {
		return domain.ApiResponse{
			Code:    "911",
			Message: err.Error(),
		}
	}

	var historiesData []domain.HistoricalData

	for _, v := range histories {
		historiesData = append(historiesData, domain.HistoricalData{
			VehicleID:   v.VehicleID,
			Notes:       v.Notes,
			PIC:         v.PIC,
			CustomerID:  v.CustomerID,
			PlateNumber: v.PlateNumber,
			ComeAt:      v.CreatedAt.Format(time.RFC822Z),
		})
	}

	result := domain.VehicleHistorical{
		ID:        vehicle.ID,
		Brand:     vehicle.Brand,
		VIN:       vehicle.VIN,
		Histories: historiesData,
	}

	return domain.ApiResponse{
		Code:    "00",
		Message: "APPROVE",
		Data:    result,
	}
}

func (s service) StoreHistorical(ctx context.Context, requset domain.VehicleHistoricalRequest) domain.ApiResponse {
	vehicle, err := s.vehicleRepository.FindByVIN(ctx, requset.VIN)
	if err != nil {
		return domain.ApiResponse{
			Code:    "911",
			Message: err.Error(),
		}
	}

	if vehicle == (domain.Vehicle{}) {
		vehicle.VIN = requset.VIN
		vehicle.Brand = requset.Brand
		vehicle.UpdatedAt = time.Now()
		err = s.vehicleRepository.Insert(ctx, &vehicle)
		if err != nil {
			return domain.ApiResponse{
				Code:    "911",
				Message: err.Error(),
			}
		}
	}
	history := domain.HistoryDetail{
		VehicleID:   vehicle.ID,
		CustomerID:  requset.CustomerID,
		PIC:         requset.PIC,
		PlateNumber: requset.PlateNumber,
		Notes:       requset.Notes,
	}
	history.CreatedAt = time.Now()
	err = s.historyRepository.Insert(ctx, &history)
	if err != nil {
		return domain.ApiResponse{
			Code:    "911",
			Message: err.Error(),
		}
	}
	return domain.ApiResponse{
		Code:    "00",
		Message: "APPROVED",
	}
}
