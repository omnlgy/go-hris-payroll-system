# Test Results - Positive Cases

**Date:** 2026-06-29
**Base URL:** `http://localhost:8080`
**Status:** Partial execution (Auth endpoints tested)

---

## 1. Auth Routes (`/api/auth`) ✅ TESTED

### 1.1 Register Employee
**Endpoint:** `POST /api/auth/register`
**Status:** ✅ PASS
**Request:**
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
**Response:** 200 OK
```json
{
  "message": "Employee registered successfully",
  "data": {
    "id": 7,
    "nik": "EMP001",
    "full_name": "John Doe",
    "email": "john@company.co.id",
    "department_id": 1,
    "position_id": 1,
    "role": "EMPLOYEE",
    "status": "ACTIVE",
    "department": {...},
    "position": {...},
    "created_at": "2026-06-29T19:19:31.750206715+07:00",
    "updated_at": "2026-06-29T19:19:31.750206715+07:00"
  }
}
```

### 1.2 Register Employee (Invalid Email - Negative Test for Reference)
**Endpoint:** `POST /api/auth/register`
**Status:** ✅ VALIDATION WORKS (Expected 400)
**Request:** Email: `jane@gmail.com`
**Response:** 400 Bad Request
```json
{
  "message": "Invalid request body",
  "errors": [
    {
      "field": "body",
      "message": "Key: 'CreateEmployeeRequest.Email' Error:Field validation for 'Email' failed on the 'companyemail' tag"
    }
  ]
}
```

### 1.3 Login
**Endpoint:** `POST /api/auth/login`
**Status:** ✅ PASS
**Request:**
```json
{
  "email": "john@company.co.id",
  "password": "password123"
}
```
**Response:** 200 OK
```json
{
  "message": "Login successful",
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 1.4 Logout
**Endpoint:** `POST /api/auth/logout`
**Status:** ⏳ NOT EXECUTED
**Notes:** Requires valid token from login

---

## 2. Department Routes (`/api/departments`) ⏳ NOT EXECUTED

| Test Case | Status | Notes |
|-----------|--------|-------|
| 2.1 Create Department | ⏳ PENDING | |
| 2.2 Get All Departments | ⏳ PENDING | |
| 2.3 Update Department | ⏳ PENDING | |
| 2.4 Delete Department | ⏳ PENDING | |

---

## 3. Position Routes (`/api/positions`) ⏳ NOT EXECUTED

| Test Case | Status | Notes |
|-----------|--------|-------|
| 3.1 Create Position | ⏳ PENDING | |
| 3.2 Get All Positions | ⏳ PENDING | |
| 3.3 Update Position | ⏳ PENDING | |
| 3.4 Delete Position (HRD) | ⏳ PENDING | Requires HRD token |

---

## 4. Employee Routes (`/api/employees`) ⏳ NOT EXECUTED

| Test Case | Status | Notes |
|-----------|--------|-------|
| 4.1 Create Employee (HRD) | ⏳ PENDING | Requires HRD token |
| 4.2 Get All Employees | ⏳ PENDING | |
| 4.3 Update Employee (HRD) | ⏳ PENDING | Requires HRD token |
| 4.4 Delete Employee (HRD) | ⏳ PENDING | Requires HRD token |

---

## 5. Attendance Routes (`/api/attendance`) ⏳ NOT EXECUTED

| Test Case | Status | Notes |
|-----------|--------|-------|
| 5.1 Record Attendance | ⏳ PENDING | Requires auth token |

---

## 6. Leave Routes (`/api/leaves`) ⏳ NOT EXECUTED

| Test Case | Status | Notes |
|-----------|--------|-------|
| 6.1 Request Leave | ⏳ PENDING | Requires auth token |
| 6.2 Approve Leave (HRD) | ⏳ PENDING | Requires HRD token |

---

## 7. Salary Routes (`/api/salaries`) ⏳ NOT EXECUTED

| Test Case | Status | Notes |
|-----------|--------|-------|
| 7.1 Get Salaries by Period | ⏳ PENDING | |
| 7.2 Calculate Salary | ⏳ PENDING | |

---

## Summary

| Category | Total | Passed | Failed | Pending |
|----------|-------|--------|--------|---------|
| Auth | 4 | 3* | 0 | 1 |
| Departments | 4 | 0 | 0 | 4 |
| Positions | 4 | 0 | 0 | 4 |
| Employees | 4 | 0 | 0 | 4 |
| Attendance | 1 | 0 | 0 | 1 |
| Leaves | 2 | 0 | 0 | 2 |
| Salaries | 2 | 0 | 0 | 2 |
| **Total** | **21** | **3** | **0** | **18** |

*Includes validation test for invalid email

---

## Next Steps
1. Create HRD user for testing HRD-required endpoints
2. Seed Department and Position data (IDs 1)
3. Execute remaining test cases in order
4. Document any failures or issues found