package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// RBACMiddleware checks if the user has the required role. It should be
// used after AuthMiddleware, which sets the user role in the context.
func RBACMiddleware(requiredRole string) gin.HandlerFunc {
    return func(c *gin.Context) {
        roleVal, exists := c.Get(ContextRoleKey)
        if !exists {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "role not found in context"})
            return
        }
        role, ok := roleVal.(string)
        if !ok || role != requiredRole {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
            return
        }
        c.Next()
    }
}
