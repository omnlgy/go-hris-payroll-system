package dto

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BadRequestResponse struct {
	Message string `json:"message"`
	Errors  []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"errors"`
}

type InternalServerErrorResponse struct {
	Message string `json:"message"`
}

type NotFoundResponse struct {
	Message string `json:"message"`
}

type UnauthorizedResponse struct {
	Message string `json:"message"`
}

type ForbiddenResponse struct {
	Message string `json:"message"`
}

type CreateEmployeeRequest struct {
	NIK          string `json:"nik" binding:"required"`
	FullName     string `json:"full_name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID uint   `json:"department_id" binding:"required"`
	PositionID   uint   `json:"position_id" binding:"required"`
	Role         string `json:"role" binding:"required,oneof=ADMIN MANAGER EMPLOYEE HRD"`
	Password     string `json:"password" binding:"required,min=6"`
	Status       string `json:"status" binding:"omitempty,oneof=ACTIVE SUSPENDED TERMINATED"`
}

type UpdateEmployeeRequest struct {
	NIK          string `json:"nik" binding:"required"`
	FullName     string `json:"full_name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	DepartmentID uint   `json:"department_id" binding:"required"`
	PositionID   uint   `json:"position_id" binding:"required"`
	Role         string `json:"role" binding:"required,oneof=ADMIN MANAGER EMPLOYEE HRD"`
	Password     string `json:"password" binding:"omitempty,min=6"`
	Status       string `json:"status" binding:"omitempty,oneof=ACTIVE SUSPENDED TERMINATED"`
}

type CreateDepartmentRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type RequestLeaveRequest struct {
	StartDate string `json:"start_date" binding:"required"`
	EndDate   string `json:"end_date" binding:"required"`
	Reason    string `json:"reason" binding:"required"`
}

type ApproveLeaveRequest struct {
	Status string `json:"status" binding:"required"`
}

type CreatePositionRequest struct {
	Title      string  `json:"title" binding:"required"`
	BaseSalary float64 `json:"base_salary" binding:"required"`
}

type SalaryRequest struct {
	EmployeeID uint   `json:"employee_id" binding:"required"`
	Period     string `json:"period" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
