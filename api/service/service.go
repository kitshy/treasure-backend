package service

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/kitshy/treasure-backend/api/models"
	"github.com/kitshy/treasure-backend/api/validator"
	"github.com/kitshy/treasure-backend/database/eventlog"
	"strconv"
)

type HandlerService struct {
	v                      *validator.Validator
	depositTokensViews     eventlog.DepositTokensViews
	grantRewardTokensViews eventlog.GrantRewardTokensViews
}

type Service interface {
	// deposit token
	GetDepositTokensList(*models.QueryListParams) (*models.DepositTokensResponse, error)

	// grant reward token
	GetGrantRewardTokensList(*models.QueryListParams) (*models.GrantRewardTokensResponse, error)

	QueryListParams(page string, pageSize string, order string) (*models.QueryListParams, error)
	QueryAddressListParams(address string, page string, pageSize string, order string) (*models.QueryAddressListParams, error)
}

func NewService(v *validator.Validator, depositTokensViews eventlog.DepositTokensViews) Service {
	return &HandlerService{
		v:                  v,
		depositTokensViews: depositTokensViews,
	}
}

func (h *HandlerService) QueryListParams(page string, pageSize string, order string) (*models.QueryListParams, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	pageVal := h.v.ValidatePage(pageInt)

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, err
	}
	pageSizeVal := h.v.ValidatePageSize(pageSizeInt)
	orderBy := h.v.ValidateOrder(order)

	return &models.QueryListParams{
		Page:     pageVal,
		PageSize: pageSizeVal,
		Order:    orderBy,
	}, nil
}

func (h *HandlerService) QueryAddressListParams(address string, page string, pageSize string, order string) (*models.QueryAddressListParams, error) {
	var paraAddress string
	if address == "0x00" {
		paraAddress = "0x00"
	} else {
		addr, err := h.v.ParseValidateAddress(address)
		if err != nil {
			log.Error("invalid address param", "address", address, "err", err)
			return nil, err
		}
		paraAddress = addr.String()
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	pageVal := h.v.ValidatePage(pageInt)

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return nil, err
	}
	pageSizeVal := h.v.ValidatePageSize(pageSizeInt)
	orderBy := h.v.ValidateOrder(order)

	return &models.QueryAddressListParams{
		Address:  paraAddress,
		Page:     pageVal,
		PageSize: pageSizeVal,
		Order:    orderBy,
	}, nil
}
