package response

var (
	ErrSuccess = newError(200, "ok")
	// more biz errors
	ErrUpdateUser          = newError(402, "Update user info failure")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")
	ErrUsernameAlreadyUse  = newError(1001, "The username is already in use.")
	ErrEncryptPassword     = newError(1002, "failed to hash password")
	ErrUserPassWordEncrypt = newError(1003, "failed to hash password")
	ErrGenJWT              = newError(1004, "failed to generate JWT token")
	ErrUserNotFound        = newError(1005, "failed to generate user ID")
	ErrCreateUser          = newError(1006, "failed to create user")
)
