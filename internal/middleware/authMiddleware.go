package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/go-hris-payroll-system/internal/service"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			// controller.MapError(ctx, &service.UnauthorizedError{Message: "missing authorization header"})
			// ctx.Abort()
			ctx.AbortWithStatusJSON(401, gin.H{"error": "missing authorization header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			// controller.MapError(ctx, &service.UnauthorizedError{Message: "invalid authorization format"})
			// ctx.Abort()
			ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid authorization format"})
			return
		}

		claim, err := service.ValidateToken(tokenString)
		if err != nil {
			// controller.MapError(ctx, err)
			// ctx.Abort()
			ctx.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("employee_id", claim.EmployeeID)
		ctx.Set("email", claim.Email)
		ctx.Set("department_id", claim.DepartmentID)
		ctx.Set("role", claim.Role)
		ctx.Set("token_string", tokenString)
		ctx.Next()
	}
}

func RequireRole(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.GetString("role")
		for _, r := range roles {
			if r == role {
				ctx.Next()
				return
			}
		}
		// controller.MapError(ctx, &service.ForbiddenError{Message: "Access Forbidden"})
		// ctx.Abort()
		ctx.AbortWithStatusJSON(403, gin.H{"error": "Access Forbidden"})
	}
}
