# FinalTask Back-End (Rakamin × Evermos)

A clean-architecture REST API built with Go, Fiber, GORM, and MySQL. Supports:

- 🔑 Authentication (JWT) & role-based access (admin vs user)  
- 👤 User profile & address management  
- 🏪 Store (toko) management  
- 🗂️ Category management (admin only)  
- 📦 Product CRUD + image upload + pagination & filtering  
- 💰 Transaction handling + audit snapshots (`log_produk`)  
- 🌏 Public region lookup (Province & Regency) via Emsifa API  

---

## 🖥️ Prerequisites

- Go 1.25.4  
- MySQL 5.7+ (or compatible, XAMPP recommended for local setup)  
- [Postman](https://www.postman.com/) (for testing)  

---

## ⚙️ Setup

1. **Clone repository**  
   ```bash
   git clone https://github.com/enggardyh57/evermos-task.git
   cd evermos-task



### Create Database

Buka MySQL / MariaDB (misal via phpMyAdmin atau terminal) dan jalankan:

```sql
CREATE DATABASE toko_db;
