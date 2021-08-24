package domain

import (
	"migrates3/common"
	"net/http"
)

type BasicResponse struct {
	Status string `json:"status"`
}
type SuccessResponse struct {
	BasicResponse
	Data interface{} `json:"data"`
}

type SuccessResponseWithPageInfo struct {
	BasicResponse
	Data     interface{} `json:"data"`
	PageInfo *PageInfo   `json:"page_info"`
}

type PageInfo struct {
	PageNumber       string `json:"page_number"`
	PageSize         string `json:"page_size"`
	TotalRecordCount string `json:"total_record_count"`
}

type ErrorResponse struct {
	BasicResponse
	Message string `json:"message"`
}

type MixedResponse struct {
	BasicResponse
	Data struct {
		Success interface{} `json:"success"`
		Fail    interface{} `json:"fail"`
	} `json:"data"`
}

func NewSuccessResponse(_data interface{}) SuccessResponse {
	return SuccessResponse{
		BasicResponse: BasicResponse{
			Status: common.HTTP_STATUS_SUCCESS,
		},
		Data: _data,
	}
}

func NewSuccessResponseWithPageInfo(_data interface{}, _page *PageInfo) SuccessResponseWithPageInfo {
	return SuccessResponseWithPageInfo{
		BasicResponse: BasicResponse{
			Status: common.HTTP_STATUS_SUCCESS,
		},
		Data:     _data,
		PageInfo: _page,
	}
}

func NewErrorResponse(_message string) ErrorResponse {
	return ErrorResponse{
		BasicResponse: BasicResponse{
			Status: common.HTTP_STATUS_ERROR,
		},
		Message: _message,
	}
}

func NewMixedResponse(_success, _fail interface{}) MixedResponse {

	return MixedResponse{
		BasicResponse: BasicResponse{
			Status: common.HttpStatusMix,
		},
		Data: struct {
			Success interface{} `json:"success"`
			Fail    interface{} `json:"fail"`
		}{_success, _fail},
	}
}

func NewUnauthorizedResponse() (int, ErrorResponse) {
	return http.StatusUnauthorized, NewErrorResponse("Authorization failed")
}
