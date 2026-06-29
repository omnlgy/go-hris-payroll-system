# go-hris-payroll-system
ujian week 3

## API Endpoints

### Department Routes (`/api/departments`)
*   `GET /api/departments` - Get all departments
*   `POST /api/departments` - Create a new department
*   `PUT /api/departments/:id` - Update an existing department
*   `DELETE /api/departments/:id` - Delete a department

### Position Routes (`/api/positions`)
*   `GET /api/positions` - Get all positions
*   `POST /api/positions` - Create a new position
*   `PUT /api/positions/:id` - Update an existing position
*   `DELETE /api/positions/:id` - Delete a position (Requires HRD role)

### Employee Routes (`/api/employees`)
*   `POST /api/employees` - Create a new employee (Requires HRD role)
*   `GET /api/employees` - Get all employees
*   `DELETE /api/employees/:id` - Delete an employee (Requires HRD role)
*   `PUT /api/employees/:id` - Update an employee (Requires HRD role)

### Attendance Routes (`/api/attendance`)
*   `POST /api/attendance` - Record attendance (Requires authentication)

### Leave Routes (`/api/leaves`)
*   `POST /api/leaves` - Request leave (Requires authentication)
*   `PATCH /api/leaves/:id/approve` - Approve a leave request (Requires HRD role)

### Salary Routes (`/api/salaries`)
*   `GET /api/salaries/period/:period` - Get employee salaries by period
*   `POST /api/salaries/calculate` - Calculate salaries

### Auth Routes (`/api/auth`)
*   `POST /api/auth/login` - User login
*   `POST /api/auth/register` - Register new employee
*   `POST /api/auth/logout` - User logout
