package response

import (
	"user-service/internal/core/entity"
)

type Response struct {
	Data     interface{}      `json:"data"`
	Status   bool             `json:"status"`
	Err_code entity.ErrorCode `json:"err_code"`
	Err_msg  string           `json:"err_msg"`
}
