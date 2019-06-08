package core

type ErrorCode int

const (
	BadRequest ErrorCode = iota + 1
	NotFound
	Internal
	ApiInternal
)

type IdentixOneError struct {
	Code ErrorCode
	Data *[]byte
	err  error
}

func (ie IdentixOneError) Error() string {
	if ie.err == nil {
		if ie.Data != nil {
			return string(*ie.Data)
		} else {
			switch ie.Code {
			case BadRequest:
				return "Bad request"
			case NotFound:
				return "Not found"
			case Internal:
				return "Oops"
			case ApiInternal:
				return "identix.one service error"
			default:
				return "Something wrong"
			}
		}

	}
	return ie.err.Error()
}

func NewError(code ErrorCode, err error, data []byte) IdentixOneError {
	return IdentixOneError{
		Code: code,
		err:  err,
		Data: &data,
	}
}
