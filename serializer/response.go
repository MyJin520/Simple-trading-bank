package serializer

import "go-store/pkg/e"

// Response 定义一个序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

func BuildListCarouse(item interface{}, total int64) Response {
	return Response{
		Status: 200,
		Data: DataList{
			item,
			total,
		},
		Msg: e.GetMSG(200),
	}
}
