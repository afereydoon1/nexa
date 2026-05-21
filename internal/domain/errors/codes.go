package errors

const (

	// Validation Errors
	ErrNameRequired     = "ERROR_NAME_REQUIRED"
	ErrNameMin          = "ERROR_NAME_MIN"
	ErrSlugRequired     = "ERROR_SLUG_REQUIRED"
	ErrImageRequired    = "ERROR_IMAGE_REQUIRED"
	ErrEmailRequired    = "ERROR_EMAIL_REQUIRED"
	ErrEmailInvalid     = "ERROR_EMAIL_INVALID"
	ErrPasswordRequired = "ERROR_PASSWORD_REQUIRED"
	ErrPasswordMin      = "ERROR_PASSWORD_MIN"

	// System Errors
	ErrUploadFailed = "ERROR_UPLOAD_FAILED"

	// Business Errors
	ErrGenreCreateFailed  = "ERROR_GENRE_CREATE_FAILED"
	ErrGenreNotFound      = "ERROR_GENRE_NOT_FOUND"
	ErrInvalidCredentials = "ERROR_INVALID_CREDENTIALS"
	ErrUserCreateFailed   = "ERROR_USER_CREATE_FAILED"

	ErrInvalidID = "ERROR_INVALID_ID"

	// Success Messages
	SuccessGenreCreated  = "SUCCESS_GENRE_CREATED"
	SuccessGenreUpdated  = "SUCCESS_GENRE_UPDATED"
	SuccessGenresFetched = "SUCCESS_GENRES_FETCHED"
	SuccessGenreFetched  = "SUCCESS_GENRE_FETCHED"
	SuccessGenreDeleted  = "SUCCESS_GENRE_DELETED"
	SuccessUserCreated   = "SUCCESS_USER_CREATED"
	SuccessLogin         = "SUCCESS_LOGIN"
)
