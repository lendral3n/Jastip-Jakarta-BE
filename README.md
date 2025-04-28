# Jastip Jakarta - Backend

Backend API untuk aplikasi **Jastip Jakarta** — layanan titip beli barang dari Jakarta.

---

## 📂 Struktur Folder

```bash
.
├── .github/workflows/    # CI/CD workflows (belum lengkap)
├── app/                  # Struktur aplikasi utama (controller, model, repository, service)
├── features/             # Endpoint dan fitur utama (auth, products, transactions, users)
├── utils/                # Utility umum (JWT, Hash, Logger)
├── main.go               # Entry point aplikasi
├── Dockerfile            # Docker setup
├── go.mod, go.sum        # Dependency manager (Golang modules)
└── README.md             # Dokumentasi ini
```

---

## 🚀 Cara Menjalankan

### 1. Clone Repository
```bash
git clone https://github.com/lendral3n/Jastip-Jakarta-BE.git
cd Jastip-Jakarta-BE
```

### 2. Setup Environment
Buat file `.env` dengan isi seperti berikut:

```dotenv
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=jastip
JWT_SECRET=your_jwt_secret
```

Contoh file `.env.example` sudah tersedia:
```dotenv
# .env.example
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=jastip
JWT_SECRET=your_secret_key
```

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Run Aplikasi
```bash
go run main.go
```

### 5. (Opsional) Build dan Jalankan dengan Docker
```bash
docker build -t jastip-be .
docker run -p 8080:8080 jastip-be
```

---

## 📚 Fitur Utama

- **Authentication**
  - Register
  - Login
  - JWT Authorization
- **Product Management**
  - CRUD Produk
- **Transaction Management**
  - CRUD Transaksi
- **User Management**
  - Update Profile
  - View Profile

---

## ⚙️ Teknologi yang Digunakan

| Teknologi | Keterangan |
| :-------- | :--------- |
| Golang    | Bahasa pemrograman backend utama |
| PostgreSQL| Database relational utama |
| Gin       | Web framework untuk HTTP router |
| JWT       | JSON Web Token untuk autentikasi |
| Docker    | Containerization untuk deployment |

---

## 🔒 Autentikasi

- Semua endpoint sensitif dilindungi oleh middleware JWT.
- Token harus dikirimkan di header:

```http
Authorization: Bearer <your_token>
```

---

## 🛠️ API Endpoints (Contoh)

| Method | Endpoint              | Keterangan                   |
| :----: | :-------------------- | :---------------------------- |
| POST   | `/api/auth/register`   | Register user baru            |
| POST   | `/api/auth/login`      | Login dan mendapatkan token   |
| GET    | `/api/products`        | List semua produk             |
| POST   | `/api/products`        | Tambah produk (perlu token)    |
| PUT    | `/api/products/:id`    | Update produk                 |
| DELETE | `/api/products/:id`    | Hapus produk                  |
| POST   | `/api/transactions`    | Buat transaksi baru            |

*(Endpoint lengkap bisa dilihat di masing-masing folder `features/`)*

---

## 📈 Database Schema (Ringkasan)

- **Users**
  - `id` (PK)
  - `name`
  - `email`
  - `password`
  - `role`

- **Products**
  - `id` (PK)
  - `name`
  - `description`
  - `price`
  - `stock`

- **Transactions**
  - `id` (PK)
  - `user_id` (FK -> Users)
  - `product_id` (FK -> Products)
  - `quantity`
  - `total_price`
  - `status`

> **Catatan**: Skema lengkap sebaiknya divisualisasikan dalam ERD.

Contoh sederhana ERD:

```plaintext
Users (1) ─→ (N) Transactions (N) ←─ (1) Products
```

---

## 📋 TODO / Roadmap

- [ ] Implementasi refresh token.
- [ ] Integrasi Payment Gateway.
- [ ] Admin Panel untuk approval transaksi.
- [ ] Unit Testing dan Integration Testing.
- [ ] CI/CD Workflow untuk deployment otomatis.

---

## 👨‍💼 Kontribusi

Pull Request, Issue, dan Saran sangat diterima!

Cara kontribusi:
1. Fork repo ini.
2. Buat branch baru (`git checkout -b fitur-baru`)
3. Commit perubahanmu (`git commit -m 'Add fitur baru'`)
4. Push ke branch (`git push origin fitur-baru`)
5. Buat Pull Request

---

## 📢 Catatan Tambahan

- Error handling masih bisa diperbaiki.
- Struktur folder cukup scalable untuk pengembangan microservices di masa depan.
- `.env` file sebaiknya tidak di-commit ke repository publik.
- Perlu dibuat dokumentasi API lengkap menggunakan Postman Collection.

Contoh struktur dokumentasi Postman Collection:

- Folder: **Authentication**
  - Register User [POST] `/api/auth/register`
  - Login User [POST] `/api/auth/login`

- Folder: **Products**
  - List Products [GET] `/api/products`
  - Create Product [POST] `/api/products`
  - Update Product [PUT] `/api/products/:id`
  - Delete Product [DELETE] `/api/products/:id`

- Folder: **Transactions**
  - Create Transaction [POST] `/api/transactions`
  - List Transactions [GET] `/api/transactions`

---

# Terima Kasih! 🚀

Yuk bantu kembangkan proyek ini bersama!
