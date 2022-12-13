package exception

type BusinessException struct {
	error
	ID  int64
	Msg string
}

func (e *BusinessException) Error() string {
	return e.Msg
}

type ClientException struct {
	BusinessException
	ID  int64
	Msg string
}

func (e *ClientException) Error() string {
	return e.Msg
}

type ServerException struct {
	BusinessException
	ID  int64
	Msg string
}

func (e *ServerException) Error() string {
	return e.Msg
}
