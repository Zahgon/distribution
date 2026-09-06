package errcode

type ErrorCoder interface {
	ErrorCode() ErrorCode
}

type ErrorCode int

var _ error = ErrorCode(0)

func (ec ErrorCode) ErrorCode() ErrorCode { _ = "STUB: not implemented"; return *new(ErrorCode) }

func (ec ErrorCode) Error() string { _ = "STUB: not implemented"; return "" }

func (ec ErrorCode) Descriptor() ErrorDescriptor {
	_ = "STUB: not implemented"
	return *new(ErrorDescriptor)
}

func (ec ErrorCode) String() string { _ = "STUB: not implemented"; return "" }

func (ec ErrorCode) Message() string { _ = "STUB: not implemented"; return "" }

func (ec ErrorCode) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ec *ErrorCode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (ec ErrorCode) WithMessage(message string) Error {
	_ = "STUB: not implemented"
	return *new(Error)
}

func (ec ErrorCode) WithDetail(detail any) Error { _ = "STUB: not implemented"; return *new(Error) }

func (ec ErrorCode) WithArgs(args ...any) Error { _ = "STUB: not implemented"; return *new(Error) }

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Detail  any       `json:"detail,omitempty"`
}

var _ error = Error{}

func (e Error) ErrorCode() ErrorCode { _ = "STUB: not implemented"; return *new(ErrorCode) }

func (e Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e Error) WithDetail(detail any) Error { _ = "STUB: not implemented"; return *new(Error) }

func (e Error) WithArgs(args ...any) Error { _ = "STUB: not implemented"; return *new(Error) }

type ErrorDescriptor struct {
	Code ErrorCode

	Value string

	Message string

	Description string

	HTTPStatusCode int
}

func ParseErrorCode(value string) ErrorCode { _ = "STUB: not implemented"; return *new(ErrorCode) }

type Errors []error

var _ error = Errors{}

func (errs Errors) Error() string { _ = "STUB: not implemented"; return "" }

func (errs Errors) Len() int { _ = "STUB: not implemented"; return 0 }

func (errs Errors) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (errs *Errors) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
