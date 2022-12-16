package exception

type BusinessException struct {
	Msg string
}

func (e *BusinessException) Error() string {
	return e.Msg
}

type ClientException struct {
	Msg string
}

func (e *ClientException) Error() string {
	return e.Msg
}

type ServerException struct {
	Msg string
}

func (e *ServerException) Error() string {
	return e.Msg
}
