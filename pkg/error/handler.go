package localError

type GlobalError *errorEntity

type errorEntity struct {
	Code    int
	Message string
	Error   error
}

// Create errrorEntity from available exposed method
func newErrorEntity(code int, msg string, err error) *errorEntity {
	return &errorEntity{
		Code:    code,
		Message: msg,
		Error:   err,
	}
}

// Return the error instance of Internal Server Error
func ErrorInternalServer(msg string, err error) *errorEntity {
	return newErrorEntity(500, msg, err)
}

// Return the error instance of Not FOund
func ErrorNotFound(msg string, err error) *errorEntity {
	return newErrorEntity(404, msg, err)
}

// Return the error instance of Unauthorized
func ErrorUnauthorized(msg string, err error) *errorEntity {
	return newErrorEntity(401, msg, err)
}

// Return the error instance of Bad Request
func ErrorBadRequest(msg string, err error) *errorEntity {
	return newErrorEntity(400, msg, err)
}

// Return the error instance of Forbidden
func ErrorForbidden(msg string, err error) *errorEntity {
	return newErrorEntity(403, msg, err)
}

// Return the error instance of Conflict
func ErrorConflict(msg string, err error) *errorEntity {
	return newErrorEntity(409, msg, err)
}
