# Smart Lock Management System

A robust and scalable backend system for managing smart locks, built with Go and integrated with Tuya IoT platform. This project demonstrates modern software engineering practices and enterprise-level architecture.

## 🚀 Features

- **Smart Lock Integration**: Seamless integration with Tuya IoT platform for smart lock management
- **Role-Based Access Control**: Comprehensive user role management system
- **Space Management**: Hierarchical space organization for lock management
- **Whitelist System**: Advanced access control through whitelist management
- **Audit Logging**: Detailed activity logging for security and monitoring
- **RESTful API**: Well-structured API endpoints following REST principles
- **JWT Authentication**: Secure authentication using JWT tokens
- **CORS Support**: Built-in CORS middleware for cross-origin requests
- **Docker Support**: Containerized application for easy deployment
- **Environment Configuration**: Flexible configuration through environment variables

## 🛠️ Technical Stack

- **Language**: Go
- **Framework**: Gin (Web Framework)
- **Database**: SQL (with custom database package)
- **Authentication**: JWT
- **IoT Integration**: Tuya IoT Platform
- **Containerization**: Docker & Docker Compose
- **API Documentation**: RESTful API
- **Environment Management**: godotenv

## 🏗️ Architecture

The project follows a clean architecture pattern with the following components:

- **Controllers**: Handle HTTP requests and responses
- **Services**: Implement business logic
- **Models**: Define data structures
- **Middleware**: Handle cross-cutting concerns
- **Utils**: Shared utilities and helpers

## 🚀 Getting Started

### Prerequisites

- Go 1.x
- Docker and Docker Compose
- Tuya IoT Platform credentials

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/lock-system.git
cd lock-system
```

2. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

3. Run with Docker:
```bash
docker-compose up --build
```

Or run locally:
```bash
go mod download
go run main.go
```

## 🔒 Security Features

- JWT-based authentication
- Role-based access control
- Whitelist management
- Audit logging
- Secure environment variable handling
- CORS protection

## 📝 API Endpoints

The system provides endpoints for:
- User management
- Device control
- Space management
- Whitelist operations
- Role management
- Activity logging

## 🧪 Testing

```bash
go test ./...
```

## 📦 Deployment

The application is containerized using Docker and can be deployed using Docker Compose:

```bash
docker-compose up -d
```

## 🔧 Configuration

The application can be configured through environment variables:
- `DSN`: Database connection string
- `TUYA_ACCESS_ID`: Tuya IoT platform access ID
- `TUYA_ACCESS_KEY`: Tuya IoT platform access key
- `TUYA_APP_NAME`: Tuya IoT platform application name
- `GIN_MODE`: Gin framework mode (debug/release)

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👤 Author

Your Name - [@yourusername](https://github.com/yourusername)

## 🙏 Acknowledgments

- Tuya IoT Platform
- Gin Framework
- Go Community
