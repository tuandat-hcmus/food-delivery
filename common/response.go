package common

// model response: neu co data thi data nam trong "data", tuong tu paging va filter
// VD: {
// 		"data" : {
//			"name" : "Ivy League"
// 		}
//}
type successRes struct {
	Data interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return &successRes{Data: data, Paging: nil, Filter: nil}
}