# KrakenD JWT API Gateway Project

This project is a **JWT-authenticated API** using **KrakenD API Gateway** and a **Golang backend**. KrakenD acts as an **API Gateway** that validates JWTs, while the Golang backend handles authentication and data processing.

---

## 🚀 Features

- **API Gateway with KrakenD**
- **JWT (JSON Web Token) authentication**
- **Backend API written in Golang**
- **Easy deployment with Docker Compose**
- **Secure API calls via KrakenD**

---

## 📁 Project Structure

```
krakend-jwt-gateway/
├── backend/               # Golang backend API
│   ├── main.go            # Backend source code
│   ├── go.mod             # Go module file
│   ├── go.sum             # Go dependencies
├── krakend/               # KrakenD configurations
│   ├── krakend.json       # KrakenD config file
├── docker-compose.yml     # Docker Compose setup
└── README.md              # Project documentation
```

---

## 🔧 Setup & Run

### 1️⃣ **Clone the Repository**

```sh
git clone https://github.com/user/krakend-jwt-gateway.git
cd krakend-jwt-gateway
```

### 2️⃣ **Run the Application with Docker**

```sh
docker-compose up
```

This command starts both **KrakenD Gateway** and the **Golang backend API**.

---

## 📌 API Usage

### 1️⃣ **Login & Get a Token**

```sh
curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{"username": "admin", "password": "password"}'
```

Response:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI..."
}
```

### 2️⃣ **Fetch Orders with JWT Token**

```sh
curl -X GET http://localhost:8080/orders \
     -H "Authorization: Bearer <YOUR_TOKEN_HERE>"
```

or via **KrakenD Gateway**:

```sh
curl -X GET http://localhost:8081/api/orders \
     -H "Authorization: Bearer <YOUR_TOKEN_HERE>"
```

---

## 📌 Technologies

- **Golang** - Backend API development
- **KrakenD** - API Gateway management
- **JWT (JSON Web Token)** - Authentication
- **Docker & Docker Compose** - Containerization
- **Echo Framework** - Golang web framework

---

## ✨ Contributing

Feel free to submit **pull requests**! 🛠️

If you encounter any issues, open an **issue** on GitHub. 🚀

---

**📌 License:** MIT

This project is open-source and free to use! 😊

