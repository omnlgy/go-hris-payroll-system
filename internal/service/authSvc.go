package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/omnlgy/go-hris-payroll-system/internal/domain"
	"github.com/omnlgy/go-hris-payroll-system/internal/utils"
)

type AuthService struct {
	employeeRepo  domain.EmployeeRepository
	blackListRepo domain.BlacklistedTokenRepository
}

func NewAuthService(employeeRepo domain.EmployeeRepository, blackListRepo domain.BlacklistedTokenRepository) *AuthService {
	return &AuthService{
		employeeRepo:  employeeRepo,
		blackListRepo: blackListRepo,
	}
}

func (s *AuthService) Login(email, password string) (string, error) {
	employee, err := s.employeeRepo.GetEmployeeByEmail(email)
	if err != nil {
		return "", err
	}

	if err := utils.ComparePassword(employee.Password, password); err != nil {
		return "", err
	}

	token, err := GenerateJWT(employee.ID, employee.Email, employee.DepartmentID, employee.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Logout(tokenString string) error {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return err
	}

	if isBlacklisted, err := s.blackListRepo.IsBlacklisted(tokenString); err != nil {
		return err
	} else if isBlacklisted {
		return nil
	}

	expiredAt := claims.ExpiresAt.Time

	_, err = s.blackListRepo.Blacklist(tokenString, expiredAt)
	return err
}

type JWTClaims struct {
	EmployeeID   uint   `json:"employee_id"`
	Email        string `json:"email"`
	DepartmentID uint   `json:"department_id"`
	Role         string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(employeeId uint, email string, departmentId uint, role string) (string, error) {
	claims := JWTClaims{
		EmployeeID:   employeeId,
		Email:        email,
		DepartmentID: departmentId,
		Role:         role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your-secret-key"))
}

func ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenMalformed
	}

	return claims, nil
}
