package weemesh

import (
	"github.com/google/uuid"
)

type CommonResult[T any] struct {
	/** 错误编码
	 * 1xx：信息 - 收到请求，继续处理；
	 * 2xx：成功——动作被成功接收、理解和接受；
	 * 3xx：重定向 - 必须采取进一步的行动才能完成请求；
	 * 4xx：客户端错误 - 请求包含错误的语法或无法实现；
	 * 5xx：服务器错误——服务器未能完成明显有效的请求；
	 * */
	Code int `json:"code"`

	Message string `json:"message"`

	Data T `json:"data"`

	RequestID uuid.UUID `json:"request_id"`
}
