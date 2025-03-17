package service

import "github.com/kitshy/treasure-backend/api/models"

func (h *HandlerService) GetDepositTokensList(request *models.QueryListParams) (*models.DepositTokensResponse, error) {
	return &models.DepositTokensResponse{
		Current: request.Page,
		Size:    request.PageSize,
	}, nil
}
