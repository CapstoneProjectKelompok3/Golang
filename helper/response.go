package helper

type MapResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WebResponse(code int, message string, data interface{}) MapResponse {
	return MapResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

type MapResponsePagination struct {
	Code      int         `json:"code"`
	TotalData int         `json:"total_data"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
}

func WebResponsePagination(code int, totalData int, message string, data interface{}) MapResponsePagination {
	return MapResponsePagination{
		Code:      code,
		TotalData: totalData,
		Message:   message,
		Data:      data,
	}
}
