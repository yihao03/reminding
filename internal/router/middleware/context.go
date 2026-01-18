package middleware

type contextKey string

const (
	UserUIDKey contextKey = "userID"
	DBPoolKey  contextKey = "dbPool"
)
