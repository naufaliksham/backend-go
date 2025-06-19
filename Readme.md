**Nama:** Naufal Iksham  
**Email:** naufal.iksham@gmail.com

---
Soal No
1. Folder `System Design Test`
2. File `Database Design Test.png`
3. Skill Test 
---

## 📌 Deskripsi

Proyek ini merupakan solusi backend untuk platform pembelian tiket bioskop online yang terdiri dari:

- System Design (Flowchart & Penjelasan)
- Database Design (ERD Diagram)
- RESTful API menggunakan Golang (Login, CRUD Jadwal Tayang)
- Autentikasi menggunakan JWT
- Dokumentasi via Postman Collection

---

## 🚀 Cara Menjalankan Aplikasi

### 🔧 1. Clone Repository
```bash
git clone https://github.com/naufaliksham/backend-go.git
cd backend-go
```

### 📦 2. Install Dependency Go
```bash
go mod tidy
```

### 🗄️ 3. Buat Database Baru
Contoh: cinema

### 📄 4. Buat File .env
⚠️ Ubah sesuai konfigurasi PostgreSQL Anda dengan database yang baru saja dibuat, untuk JWT secret key bebas.

### ▶️ 5. Jalankan Aplikasi
```bash
go run main.go
```
Server akan berjalan di: http://localhost:8080

## 🧪 Tes API dengan Postman

### 📤 1. Import Collection
File `Cinema.postman_collection.json`

### 🔐 2. Login
Endpoint `POST /login`
Body 
```
{
  "email": "naufalbackend@gmail.com",
  "password": "naufalbackend"
}
```
Response 
```
{
    "token": "Bearer <token>"
}
```

### 🪪 3. Gunakan Bearer token untuk endpoint yang butuh auth:
Masukan Header pada kolom `Authorization`

Contoh = `Bearer eyJhbGc....`

### 🎬 4. Endpoint Jadwal Tayang
| Method | Endpoint         | Keterangan        |
| ------ | ---------------- | ----------------- |
| GET    | `/schedules`     | List semua jadwal |
| POST   | `/schedules`     | Tambah jadwal     |
| PUT    | `/schedules/id` | Edit jadwal       |
| DELETE | `/schedules/id` | Hapus jadwal      |

## ⚙️ Teknologi yang Digunakan
- Golang (Gin, GORM)
- PostgreSQL
- JWT Authentication
- Postman (API Testing)
- dotenv (.env)
