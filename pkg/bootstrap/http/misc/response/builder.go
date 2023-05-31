package response

// Build — ...
func Build(err error) *Response {
	return &Response{
		Message: err.Error(),
		Payload: nil,
	}
}
