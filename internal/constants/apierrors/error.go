package apierrors

// Коды ошибок в поле error ответа API.
const (
	CodeUnauthorized   = "unauthorized"
	CodeInvalidRequest = "invalid_request"
	CodeLimitReached   = "limit_reached"
	CodeNotFound       = "not_found"
	CodeForbidden      = "forbidden"
	CodeAlreadyExists  = "already_exists"
)

// Общие сообщения в поле message ответа API.
const (
	MessageAuthHeaderRequired    = "Authorization header required"
	MessageInvalidAuth           = "Invalid or expired token"
	MessageAdvertisementNotFound = "Advertisement not found"
	MessageFavoriteNotFound      = "Favorite entry not found"
	MessageOwnAdvertisement      = "Cannot add own advertisement to favorites"
	MessageAlreadyInFavorites    = "Advertisement is already in favorites"
	MessageInvalidUUIDLength     = "invalid UUID length: %d"
)
