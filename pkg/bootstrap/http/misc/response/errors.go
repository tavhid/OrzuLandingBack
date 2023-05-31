package response

import "errors"

var (
	// ErrAccessDenied ...
	ErrAccessDenied = errors.New("access is denied")
	// ErrAlreadyRegistered ...
	ErrAlreadyRegistered = errors.New("this phone number is already registered")
	// ErrBadRequest ...
	ErrBadRequest = errors.New("bad request")
	// ErrDataNotFound ...
	ErrDataNotFound = errors.New("not found")
	// ErrIncorrectOtpCode ...
	ErrIncorrectOtpCode = errors.New("incorrect otp code")
	// ErrInternalServer ...
	ErrInternalServer = errors.New("internal server error")
	// ErrInvalidPasswordFormat ...
	ErrInvalidPasswordFormat = errors.New("invalid password format")
	// ErrInvalidPhoneFormat ...
	ErrInvalidPhoneFormat = errors.New("invalid phone number format")
	// ErrLimitExceeded ...
	ErrLimitExceeded = errors.New("retry limit exceeded")
	// ErrLogin ...
	ErrLogin = errors.New("wrong login or password")
	// ErrNoContent ...
	ErrNoContent = errors.New("no content")
	// ErrNonExistentLocation ...
	ErrNonExistentLocation = errors.New("non-existent country or city")
	// ErrNotImplementation ...
	ErrNotImplementation = errors.New("not implementation")
	// ErrSimilarPasswords ...
	ErrSimilarPasswords = errors.New("the new password must be different from the old one")
	// ErrSomethingWentWrong ...
	ErrSomethingWentWrong = errors.New("something went wrong")
	// ErrSuccess ...
	ErrSuccess = errors.New("success")
	// ErrTokenIsExpired ...
	ErrTokenIsExpired = errors.New("token is expired")
	// ErrUnauthorized ...
	ErrUnauthorized = errors.New("unauthorized")
	ErrEmptyFileds  = errors.New("one or more fields are empty")
	// ErrIncorrectFillFields ...
	ErrIncorrectFillFields = errors.New("fill in the correct field or fields")
)

var errorCode = map[string]int{
	ErrAccessDenied.Error():          403,
	ErrAlreadyRegistered.Error():     409,
	ErrBadRequest.Error():            400,
	ErrDataNotFound.Error():          404,
	ErrIncorrectOtpCode.Error():      400,
	ErrInternalServer.Error():        500,
	ErrInvalidPasswordFormat.Error(): 400,
	ErrInvalidPhoneFormat.Error():    400,
	ErrLimitExceeded.Error():         419,
	ErrLogin.Error():                 401,
	ErrNoContent.Error():             201,
	ErrNotImplementation.Error():     501,
	ErrSimilarPasswords.Error():      400,
	ErrSomethingWentWrong.Error():    500,
	ErrSuccess.Error():               200,
	ErrTokenIsExpired.Error():        401,
	ErrUnauthorized.Error():          401,
	ErrEmptyFileds.Error():           400,
	ErrIncorrectFillFields.Error():   400,
}
