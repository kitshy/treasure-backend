package models

import "github.com/kitshy/treasure-backend/database/eventlog"

type QueryAddressListParams struct {
	Address  string
	Page     int
	PageSize int
	Order    string
}

type QueryListParams struct {
	Page     int
	PageSize int
	Order    string
}

type DepositTokensResponse struct {
	Current int                      `json:"Current"`
	Size    int                      `json:"Size"`
	Total   int64                    `json:"Total"`
	Result  []eventlog.DepositTokens `json:"result"`
}

type GrantRewardTokensResponse struct {
	Current int                          `json:"Current"`
	Size    int                          `json:"Size"`
	Total   int64                        `json:"Total"`
	Result  []eventlog.GrantRewardTokens `json:"result"`
}
