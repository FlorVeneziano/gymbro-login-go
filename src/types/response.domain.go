package types

type Response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BadRequestResponse struct {
	Response
	Data string `json:"data"`
}

type UnauthorizedResponse struct {
	Response
}

type InternalServerErrorResponse struct {
	Response
}
