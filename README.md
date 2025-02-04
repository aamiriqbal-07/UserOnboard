Here's a fresh **README.md** for your project:

---

# **User Onboarding Service**

A simple RESTful service for user onboarding built with **Golang**, **MySQL**, and containerized using **Docker**. This service supports basic CRUD operations for managing users.

---

## **1. How to Run the Application**

### **Prerequisites**

- Docker and Docker Compose installed  
- (Optional) Go 1.23+ and MySQL 8.0 if running without Docker  

---

### **Run the Application with Docker**  

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-repo/user-onboard-service.git
   cd user-onboard-service
   ```

2. **Build and Start Services**:
   ```bash
   docker-compose up --build
   ```

3. **Access the Service**:  
   The service will be available at `http://localhost:8080`.

4. **Stopping the Services**:
   ```bash
   docker-compose down
   ```

---

## **2. API Endpoints and cURL Commands**

### **Create a User**  
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

### **Get a User by ID**  
**GET** `/api/users/:id`  
**Example**: `/api/users/1`  

**cURL Command**:
```bash
curl -X GET http://localhost:8080/api/users/1
```

---

### **List All Users**  
**GET** `/api/users`  

**cURL Command**:
```bash
curl -X GET http://localhost:8080/api/users
```

---

## **3. Design Decisions**

### **3.1. Project Structure**  

- **Models**: Contains the `User` struct representing the database schema.  
- **Controller**: Handles HTTP requests and maps them to service methods.  
- **Service**: Contains business logic.  
- **Repository**: Manages database operations.  
- **Config**: Centralized configuration for environment variables and database setup.  
- **Router**: Sets up the routes using `gin-gonic`.

**Why?**  

- **Separation of Concerns**: Each layer is responsible for a specific functionality.  
- **Testability**: Each component can be tested independently.  
- **Maintainability**: Easier to modify and extend the service.

---

### **3.2. Retry Logic for Database Connection**  

The application retries the database connection 10 times before exiting.  

**Why?**  

- **Avoids Race Conditions**: Ensures MySQL is ready before the application starts.  
- **Reliability**: Prevents immediate failure during container startup.

---

## **4. Future Improvements**

1. **Authentication & Authorization**  
   Add JWT-based authentication to secure the API.

2. **Database Migrations**  
   Use a migration tool like `golang-migrate` for managing schema changes.

3. **Error Handling & Validation**  
   Improve input validation and error responses.

4. **Centralized Logging & Monitoring**  
   Integrate tools like Prometheus and Grafana for monitoring and ELK stack for logging.

5. **API Documentation**  
   Use Swagger or Postman to provide interactive API documentation.

---

### **Why Future Improvements?**  

- **Security**: Authentication and validation enhance security.  
- **Maintainability**: Migration tools simplify database management.  
- **Monitoring & Logging**: Improves observability and troubleshooting.  
- **Documentation**: Reduces the learning curve for new developers.
