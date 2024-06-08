package response

type BaseResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
	Meta   interface{} `json:"meta,omitempty"`
}

type Status struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func SuccessResponse(code, message string, data interface{}) BaseResponse {
	return BaseResponse{
		Status: Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	}
}

func SuccessResponseWithMeta(code, message string, data, meta interface{}) BaseResponse {
	return BaseResponse{
		Status: Status{
			Code:    code,
			Message: message,
		},
		Data: data,
		Meta: meta,
	}
}

func ErrorResponse(code, message string) BaseResponse {
	return BaseResponse{
		Status: Status{
			Code:    code,
			Message: message,
		},
	}
}

func ErrorResponseWithDetail(code, message string, errors interface{}) BaseResponse {
	return BaseResponse{
		Status: Status{
			Code:    code,
			Message: message,
		},
		Error: errors,
	}
}

func UnauthorizedResponse() BaseResponse {
	return BaseResponse{
		Status: Status{
			Code:    "010100",
			Message: "Unauthorized",
		},
	}
}
