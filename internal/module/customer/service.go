package customer

import (
	"context"
	"time"

	"github.com/khairulharu/miniapps/domain"
)

type service struct {
	customerRepository domain.CustomerRepository
}

func NewService(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &service{customerRepository}
}

func (s service) All(ctx context.Context) domain.ApiResponse {
	customers, err := s.customerRepository.FindAll(ctx)
	if err != nil {
		return domain.ApiResponse{
			Code:    "911",
			Message: "SYSTEM FIND IN DATABASE NOT FOUND",
		}
	}

	var customersData []domain.CustomerData

	for _, v := range customers {
		customersData = append(customersData, domain.CustomerData{
			ID:    v.ID,
			Name:  v.Name,
			Phone: v.Phone,
		})
	}

	return domain.ApiResponse{
		Code:    "00",
		Message: "APPROVE",
		Data:    customersData,
	}
}

func (s service) Save(ctx context.Context, customerData domain.CustomerData) domain.ApiResponse {
	customer := domain.Customer{
		Name:      customerData.Name,
		Phone:     customerData.Phone,
		CreatedAt: time.Now(),
	}

	err := s.customerRepository.Insert(ctx, &customer)
	if err != nil {
		return domain.ApiResponse{
			Code:    "911",
			Message: "SYSTEM MALFUNCTION SAVE TO DATABASE",
		}
	}

	return domain.ApiResponse{
		Code:    "00",
		Message: "APPROVE",
	}
}
