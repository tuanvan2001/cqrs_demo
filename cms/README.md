# Go CQRS Demo API

Ứng dụng Go REST API sử dụng Gin framework và MySQL database với GORM ORM.

## Cấu trúc project

```
cqrs_demo/
├── main.go              # Entry point
├── go.mod              # Go modules
├── models/             # Database models
│   └── models.go
├── controllers/        # API controllers
│   ├── location_controller.go
│   ├── game_controller.go
│   ├── news_controller.go
│   ├── setting_controller.go
│   └── banner_controller.go
├── database/           # Database connection
│   └── database.go
├── routes/             # API routes
│   └── routes.go
├── tables.sql          # Database schema
└── .env.example        # Environment variables example
```

## Cài đặt

1. Clone repository và di chuyển vào thư mục:
```bash
cd /home/tuantech/codews/cqrs_demo
```

2. Initialize Go module và cài đặt dependencies:
```bash
go mod tidy
```

3. Tạo database MySQL:
```sql
CREATE DATABASE demo;
```

4. Chạy SQL script để tạo tables:
```bash
mysql -u root -p demo < tables.sql
```

5. Copy environment file và cấu hình:
```bash
cp .env.example .env
```

6. Chỉnh sửa file `.env` với thông tin database của bạn.

## Chạy ứng dụng

```bash
go run main.go
```

Server sẽ chạy trên port 8080 (hoặc PORT trong file .env).

## API Endpoints

### Health Check
- `GET /health` - Kiểm tra server status

### Locations
- `GET /api/v1/locations` - Lấy tất cả locations
- `GET /api/v1/locations/:id` - Lấy location theo ID
- `POST /api/v1/locations` - Tạo location mới
- `PUT /api/v1/locations/:id` - Cập nhật location
- `DELETE /api/v1/locations/:id` - Xóa location

### Games
- `GET /api/v1/games` - Lấy tất cả games
- `GET /api/v1/games/:id` - Lấy game theo ID
- `POST /api/v1/games` - Tạo game mới
- `PUT /api/v1/games/:id` - Cập nhật game
- `DELETE /api/v1/games/:id` - Xóa game

### News
- `GET /api/v1/news` - Lấy tất cả news
- `GET /api/v1/news/:id` - Lấy news theo ID
- `POST /api/v1/news` - Tạo news mới
- `PUT /api/v1/news/:id` - Cập nhật news
- `DELETE /api/v1/news/:id` - Xóa news

### Settings
- `GET /api/v1/settings` - Lấy tất cả settings
- `GET /api/v1/settings/:id` - Lấy setting theo ID
- `POST /api/v1/settings` - Tạo setting mới
- `PUT /api/v1/settings/:id` - Cập nhật setting
- `DELETE /api/v1/settings/:id` - Xóa setting

### Banners
- `GET /api/v1/banners` - Lấy tất cả banners
- `GET /api/v1/banners/:id` - Lấy banner theo ID
- `POST /api/v1/banners` - Tạo banner mới
- `PUT /api/v1/banners/:id` - Cập nhật banner
- `DELETE /api/v1/banners/:id` - Xóa banner

## Ví dụ sử dụng API

### Tạo Location mới
```bash
curl -X POST http://localhost:8080/api/v1/locations \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Location 1",
    "address": "123 Main St",
    "status": 1
  }'
```

### Tạo Game mới
```bash
curl -X POST http://localhost:8080/api/v1/games \
  -H "Content-Type: application/json" \
  -d '{
    "location_id": 1,
    "name": "Game 1",
    "description": "Game description",
    "status": 1
  }'
```

### Lấy tất cả Games
```bash
curl http://localhost:8080/api/v1/games
```

## Environment Variables

- `DB_HOST` - MySQL host (default: localhost)
- `DB_PORT` - MySQL port (default: 3306)
- `DB_USER` - MySQL username (default: root)
- `DB_PASSWORD` - MySQL password
- `DB_NAME` - Database name (default: demo)
- `PORT` - Server port (default: 8080)

## Features

- ✅ CRUD operations cho tất cả tables
- ✅ Foreign key relationships
- ✅ CORS support
- ✅ Error handling
- ✅ JSON response format
- ✅ Auto migration
- ✅ Environment configuration
