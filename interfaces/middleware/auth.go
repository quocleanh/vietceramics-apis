package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "vietceramics-apis/infrastructure/jwt"
)

// Context keys used to store values in the request context.
const (
    ContextUserIDKey   = "userID"
    ContextUsernameKey = "username"
    ContextRoleKey     = "role"
)

// AuthMiddleware verifies a JWT token and extracts the claims into the gin context.
func AuthMiddleware(jwtService *jwt.Service) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
            return
        }
        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header format"})
            return
        }
        tokenStr := strings.TrimSpace(parts[1])
        claims, err := jwtService.ValidateToken(tokenStr)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }
        // Attach claims to context for downstream handlers.
        c.Set(ContextUserIDKey, claims.UserID)
        c.Set(ContextUsernameKey, claims.Username)
        c.Set(ContextRoleKey, claims.Role)
        c.Next()
    }
}
