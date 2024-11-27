# Go Pharmacy API

API RESTful yang dibangun dengan Go (Golang) untuk mengelola produk apotek.

## Fitur

- Operasi CRUD (Create, Read, Update, Delete) untuk produk.
- Data produk disimpan dalam database SQLite.
- Endpoint API untuk:
  - Mengambil semua produk.
  - Menambahkan produk baru.
  - Memperbarui produk yang ada.
  - Menghapus produk.
- Penanganan error dan validasi untuk permintaan API.

## Memulai

### Prasyarat

- [Go (Golang)](https://golang.org/doc/install) terinstal di perangkat Anda.
- Database SQLite terinstal di perangkat Anda.

### Menjalankan API

1. Cloning repository:

   ```bash
   git clone https://github.com/adityamaulanazidqy/go-pharmacy-api.git
   ```

2. Masuk ke direktori proyek:

   ```bash
   cd go-pharmacy-api
   ```

3. Pastikan untuk menginstal beberapa package yang diperlukan sebelum menjalankan proyek ini.

- **Gorilla Mux**: Untuk routing HTTP.

  ```bash
  go get -u github.com/gorilla/mux
  ```

- **SQLite Driver**: Untuk berinteraksi dengan database SQLite.

  ```bash
  go get -u modernc.org/sqlite
  ```

4. Jalankan API:

   ```bash
   go run main.go
   ```

5. API akan tersedia di `http://localhost:8080`

### Endpoint API

| Metode   | Endpoint         | Description                                           |
| :------- | :--------------- | :---------------------------------------------------- |
| `GET`    | `/product`       | Mengambil semua produk.                               |
| `POST`   | `/product`       | Menambah product baru.                                |
| `PUT`    | `/product/{id}`  | Memperbarui data product berdasarkan ID.              |
| `DELETE` | `/product/{ID}`  | Menghapus product berdasarkan ID.                     |
| `POST`   | `/product/reset` | Menghapus semua product dan mereset ID autoincrement. |

## Kontak

Jika Anda memiliki pertanyaan atau saran, silakan hubungi saya di adityamaullana234@gmail.com.
