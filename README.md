# Simple REST API - Bioskop

Proyek ini adalah contoh implementasi REST API sederhana menggunakan bahasa pemrograman Go (Golang), framework web Gin, dan database PostgreSQL. Proyek ini di-refactor ke dalam struktur direktori yang modular untuk memudahkan pemeliharaan dan skalabilitas.

## Struktur Direktori

- `config/` : Berisi konfigurasi aplikasi, termasuk format konfigurasi untuk koneksi ke database.
- `database/` : Mengelola inisialisasi dan koneksi ke database PostgreSQL.
- `models/` : Berisi definisi struktur data (struct) yang digunakan dalam aplikasi, seperti model Bioskop.
- `handlers/` : Berisi logika untuk menangani request HTTP dan memberikan response (controller).
- `schema.sql` : Skema database untuk membuat tabel yang dibutuhkan.
- `main.go` : Titik masuk (entry point) aplikasi yang merangkai konfigurasi, koneksi database, dan inisialisasi router.

## Prasyarat

- Go 1.25 atau lebih baru.
- PostgreSQL berjalan di sistem Anda.

## Konfigurasi Database

Secara default, aplikasi mencoba terhubung ke PostgreSQL lokal. Anda bisa menyesuaikan konfigurasi kredensial pada file `main.go`:

```go
dbConf := config.Config{
    Host:     "localhost",
    Port:     "5432",
    User:     "postgres",
    Password: "password", // Ubah dengan password database Anda
    DBName:   "bioskop_db",
}
```

Pastikan tabel `bioskop` sudah ada di dalam database `bioskop_db`. Anda dapat membuat tabel tersebut menggunakan file `schema.sql` yang telah disediakan.
Contoh untuk menjalankan file SQL di PostgreSQL:

```bash
psql -U postgres -d bioskop_db -f schema.sql
```

## Cara Menjalankan

1. Unduh semua dependensi modul:
   ```bash
   go mod tidy
   ```

2. Jalankan aplikasi:
   ```bash
   go run .
   ```
   Atau dengan melakukan build terlebih dahulu:
   ```bash
   go build -o app
   ./app
   ```

Server akan berjalan secara default di port 3000 (`http://localhost:3000`).

## Endpoint API

### Menambahkan Bioskop Baru

- **URL:** `/bioskop`
- **Method:** `POST`
- **Header:** `Content-Type: application/json`
- **Body Request:**
  ```json
  {
      "nama": "Bioskop XXI",
      "lokasi": "Jakarta",
      "rating": 4.5
  }
  ```
- **Response Sukses (200 OK):**
  ```json
  {
      "message": "Bioskop berhasil ditambahkan"
  }
  ```
- **Response Error (400 Bad Request):**
  ```json
  {
      "error": "Nama dan Lokasi tidak boleh kosong"
  }
  ```

### Mengambil Semua Bioskop

- **URL:** `/bioskop`
- **Method:** `GET`
- **Response Sukses (200 OK):**
  ```json
  [
      {
          "id": 1,
          "nama": "Bioskop XXI",
          "lokasi": "Jakarta",
          "rating": 4.5
      }
  ]
  ```

### Mengambil Bioskop Berdasarkan ID

- **URL:** `/bioskop/:id`
- **Method:** `GET`
- **Response Sukses (200 OK):**
  ```json
  {
      "id": 1,
      "nama": "Bioskop XXI",
      "lokasi": "Jakarta",
      "rating": 4.5
  }
  ```
- **Response Error (404 Not Found):**
  ```json
  {
      "error": "Bioskop tidak ditemukan"
  }
  ```

### Memperbarui Bioskop

- **URL:** `/bioskop/:id`
- **Method:** `PUT`
- **Header:** `Content-Type: application/json`
- **Body Request:**
  ```json
  {
      "nama": "Bioskop CGV",
      "lokasi": "Bandung",
      "rating": 4.8
  }
  ```
- **Response Sukses (200 OK):**
  ```json
  {
      "message": "Bioskop berhasil diperbarui"
  }
  ```
- **Response Error (404 Not Found):**
  ```json
  {
      "error": "Bioskop tidak ditemukan"
  }
  ```

### Menghapus Bioskop

- **URL:** `/bioskop/:id`
- **Method:** `DELETE`
- **Response Sukses (200 OK):**
  ```json
  {
      "message": "Bioskop berhasil dihapus"
  }
  ```
- **Response Error (404 Not Found):**
  ```json
  {
      "error": "Bioskop tidak ditemukan"
  }
  ```
