package infra

type Response struct {
	Meta Meta
	Data interface{}
}

type ResponseList struct {
	Meta Meta
	Data interface{}
	Total int
}

type Meta struct {
	Message string
	Code int
	Status string
}

func NewResponseAPI(message string, code int, status string, data interface{}) Response {
	return Response{
		Meta: Meta{
			Message: message,
			Code: code,
			Status: status,
		},
		Data: data,
	}
}

func NewResponseListAPI(message string, code int, status string, data interface{}, total int) ResponseList {
	return ResponseList{
		Meta: Meta{
			Message: message,
			Code: code,
			Status: status,
		},
		Data: data,
		Total: total,
	}
}