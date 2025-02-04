
# **User Onboarding Service - Design Document**

## **1. Overview**
The **User Onboarding Service** is a RESTful API built using Go (Golang) and MySQL for managing users.

---

## **2. Design Decisions**

### **2.1. Project Structure**
The project follows a clean, layered architecture with separate modules for:
- **Models** – Contains the `User` struct that represents the database schema.
- **Controller** – Handles incoming HTTP requests and maps them to service methods.
- **Service** – Contains the business logic.
- **Repository** – Manages database operations.
- **Config** – Centralized configuration management, including database connection.
- **Router** – Sets up routes using `gin-gonic`.

**Why?**  
- This design provides clear separation of concerns, making the code easier to maintain and extend.
- Unit testing is simpler because each layer can be tested independently.

---

### **2.2. Database Connection and Retry Logic**
The `LoadConfig` function in `config.go` ensures the database is available before proceeding, with a retry mechanism of 10 attempts.

**Why?**  
- Avoids race conditions between application startup and MySQL initialization.
- Ensures robust connection handling for containerized environments.

---

## **3. Running the Project**

### **3.1. Prerequisites**
- Docker & Docker Compose installed
- Go (if running without Docker)
- MySQL 8.0 (if running without Docker)

---

### **3.2. Steps to Run (Dockerized Setup)**

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-repo/user-onboarding-service.git
   cd user-onboarding-service
   ```

2. **Build and Start Services**:
   ```bash
   docker-compose up --build
   ```

3. **Access the Service**:
   The application will be available at `http://localhost:8080`.

---

## **4. API Endpoints and cURL Commands**

### **4.1. Create a User**
**POST** `/api/users`  
**Request Body**:
```json
{
  "id": "1",
  "name": "John Doe",
  "signupTime": 1627683600000
}
```
**cURL Command**:
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"id":"1","name":"John Doe","signupTime":1627683600000}'
```

---

### **4.2. Get a User by ID**
**GET** `/api/users/:id`  
**Example**: `/api/users/1`

**cURL Command**:
```bash
curl -X GET http://localhost:8080/api/users/1
```

---

### **4.3. List All Users**
**GET** `/api/users`

**cURL Command**:
```bash
curl -X GET http://localhost:8080/api/users
```

---

## **5. Future Improvements**
- **Add Authentication & Authorization**: Use JWT-based authentication.
- **Database Migrations**: Use tools like `go-migrate` for managing schema changes.
- **Logging & Monitoring**: Integrate centralized logging and monitoring (e.g., Prometheus, Grafana).
- **Error Handling & Validation Improvements**: Enhance validation rules and provide more detailed error messages.

