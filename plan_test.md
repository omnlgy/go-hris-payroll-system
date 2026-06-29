# Test Plan - Positive Cases Only

## Base URL
```
http://localhost:8080
```

## Test Data Setup (Prerequisites)
- Database seeded with:
  - Department ID: 1 (e.g., "Engineering")
  - Position ID: 1 (e.g., "Software Engineer")

---

## 1. Auth Routes (`/api/auth`)

### 1.1 Register Employee
**Endpoint:** `POST /api/auth/register`
**Headers:** `Content-Type: application/json`
**Body:**
```json
{
  "nik": "EMP001",
  "full_name": "John Doe",
  "email": "john@company.co.id",
  "department_id": 1,
  "position_id": 1,
  "role": "EMPLOYEE",
  "password": "password123",
  "status": "ACTIVE"
}
```
**Expected:** 200 OK, returns employee data with ID

### 1.2 Login
**Endpoint:** `POST /api/auth/login`
**Headers:** `Content-Type: application/json`
**Body:**
```json
{
  "email": "john@company.co.id",
  "password": "password123"
}
```
**Expected:** 200 OK, returns JWT token
**Save token for authenticated requests**

### 1.3 Logout
**Endpoint:** `POST /api/auth/logout`
**Headers:** `Authorization: Bearer <token>`
**Expected:** 200 OK, token blacklisted

---

## 2. Department Routes (`/api/departments`)

### 2.1 Create Department
**Endpoint:** `POST /api/departments`
**Headers:** `Content-Type: application/json`
**Body:**
```json
{
  "name": "Marketing",
  "code": "MKT"
}
```
**Expected:** 200 OK, returns created department

### 2.2 Get All Departments
**Endpoint:** `GET /api/departments`
**Expected:** 200 OK, returns array of departments

### 2.3 Update Department
**Endpoint:** `PUT /api/departments/:id` (use ID from 2.1)
**Headers:** `Content-Type: application/json`
**Body:**
```json
{
  "name": "Marketing & Sales",
  "code": "MKT"
}
```
**Expected:** 200 OK, returns updated department

### 2.4 Delete Department
**Endpoint:** `DELETE /api/departments/:id` (use ID from 2.1)
**Expected:** 200 OK, department deleted

---

## 3. Position Routes (`/api/positions`)

### 3.1 Create Position
**Endpoint:** `POST /api/positions`
**Headers:** `Content-Type: application/json`
**Body:**
```json
{
  "title": "Senior Engineer",
  "base_salary": 15000000
}
```
**Expected:** 200 OK, returns created position

### 3.2 Get All Positions
**Endpoint:** `GET /api/positions`
**Expected:** 200 OK, returns array of positions

### 3.2 Update Position
**Endpoint:** `PUT /api/positions/:id` (use ID from 3.1)
**Headers:** `Content-Type: application/json`
**Body:**
```json
{
  "title": "Lead Engineer",
  "base_salary": 20000000
}
```
**Expected:** 200 OK, returns updated position

### 3.3 Delete Position (Requires HRD)
**Endpoint:** `DELETE /api/positions/:id` (use ID from 3.1)
**Headers:** `Authorization: Bearer <HRD_token>`
**Expected:** 200 OK, position deleted

---

## 4. Employee Routes (`/api/employees`)

### 4.1 Create Employee (Requires HRD)
**Endpoint:** `POST /api/employees`
**Headers:** `Authorization: Bearer <HRD_token>`, `Content-Type: application/json`
**Body:**
```json
{
  "nik": "EMP002",
  "full_name": "Jane Smith",
  "email": "jane@company.co.id",
  "department_id": 1,
  "position_id": 1,
  "role": "MANAGER",
  "password": "password123",
  "status": "ACTIVE"
}
```
**Expected:** 200 OK, returns created employee

### 4.2 Get All Employees
**Endpoint:** `GET /api/employees`
**Expected:** 200 OK, returns array of employees

### 4.3 Update Employee (Requires HRD)
**Endpoint:** `PUT /api/employees/:id` (use ID from 4.1)
**Headers:** `Authorization: Bearer <HRD_token>`, `Content-Type: application/json`
**Body:**
```json
{
  "nik": "EMP002",
  "full_name": "Jane Smith Updated",
  "email": "jane@company.co.id",
  "department_id": 1,
  "position_id": 1,
  "role": "MANAGER",
  "password": "newpassword123",
  "status": "ACTIVE"
}
```
**Expected:** 200 OK, returns updated employee

### 4.4 Delete Employee (Requires HRD)
**Endpoint:** `DELETE /api/employees/:id` (use ID from 4.1)
**Headers:** `Authorization: Bearer <HRD_token>`
**Expected:** 200 OK, employee deleted

---

## 5. Attendance Routes (`/api/attendance`)

### 5.1 Record Attendance (Requires Auth)
**Endpoint:** `POST /api/attendance`
**Headers:** `Authorization: Bearer <token>`, `Content-Type: application/json`
**Body:**
```json
{
  "employee_id": 1,
  "date": "2026-06-29",
  "check_in": "08:00:00",
  "check_out": "17:00:00"
}
```
**Expected:** 200 OK, returns attendance record

---

## 6. Leave Routes (`/api/leaves`)

### 6.1 Request Leave (Requires Auth)
**Endpoint:** `POST /api/leaves`
**Headers:** `Authorization: Bearer <token>`, `Content-Type: application/json`
**Body:**
```json
{
  "employee_id": 1,
  "start_date": "2026-07-01",
  "end_date": "2026-07-05",
  "reason": "Annual vacation"
}
```
**Expected:** 200 OK, returns leave request

### 6.2 Approve Leave (Requires HRD)
**Endpoint:** `PATCH /api/leaves/:id/approve` (use ID from 6.1)
**Headers:** `Authorization: Bearer <HRD_token>`, `Content-Type: application/json`
**Body:**
```json
{
  "status": "APPROVED"
}
```
**Expected:** 200 OK, returns approved leave

---

## 7. Salary Routes (`/api/salaries`)

### 7.1 Get Salaries by Period
**Endpoint:** `GET /api/salaries/period/2026-06`
**Expected:** 200 OK, returns salary records for June 2026

### 7.2 Calculate Salary
**Endpoint:** `POST /api/salaries/calculate`
**Headers:** `Content-Type: application/json`
**Body:**
```json
{
  "employee_id": 1,
  "period": "2026-06"
}
```
**Expected:** 200 OK, returns calculated salary

---

## Test Execution Order
1. Register HRD user first (for HRD-required endpoints)
2. Login as HRD, save HRD token
3. Create Department & Position (prerequisites)
4. Run all other tests using appropriate tokens

## Notes
- All timestamps in ISO 8601 format
- Passwords must be min 6 characters
- Emails must contain `@company.co.id`
- Roles: ADMIN, MANAGER, EMPLOYEE, HRD
- Statuses: ACTIVE, SUSPENDED, TERMINATED