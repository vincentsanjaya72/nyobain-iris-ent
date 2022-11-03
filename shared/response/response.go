package response

type Response[T interface{}] struct {
	ErrorCode *string `json:"errorCode"`
	Data      T       `json:"data"`
}

func Wrap[T interface{}](data T, errorCode *string) Response[T] {
	return Response[T]{
		ErrorCode: errorCode,
		Data:      data,
	}
}
