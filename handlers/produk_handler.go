package handlers

import (
	"time"
	"strconv"
    "math"
	"strings"
	"fmt"
    "os"

	"github.com/gofiber/fiber/v2"
	"evermos-task/database"
	"evermos-task/models"
)

func GetAllProduk(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	if page < 1 { page = 1 }
	if limit < 1 { limit = 10 }

	offset := (page - 1) * limit

	query := database.DB.Model(&models.Produk{}).
		Preload("Toko").
		Preload("Kategori").
		Preload("Photos")

	if nama := c.Query("nama_produk"); nama != "" {
		query = query.Where("nama_produk LIKE ?", "%"+nama+"%")
	}

	if kategori := c.Query("category_id"); kategori != "" {
		query = query.Where("id_kategori = ?", kategori)
	}

	if toko := c.Query("toko_id"); toko != "" {
		query = query.Where("id_toko = ?", toko)
	}

	if minHarga := c.Query("min_harga"); minHarga != "" {
    if v, err := strconv.Atoi(minHarga); err == nil {
        query = query.Where("harga_konsumen >= ?", v)
    }
}

if maxHarga := c.Query("max_harga"); maxHarga != "" {
    if v, err := strconv.Atoi(maxHarga); err == nil {
        query = query.Where("harga_konsumen <= ?", v)
    }
}

if kategori := c.Query("category_id"); kategori != "" {
    if v, err := strconv.Atoi(kategori); err == nil {
        query = query.Where("id_kategori = ?", v)
    }
}

if toko := c.Query("toko_id"); toko != "" {
    if v, err := strconv.Atoi(toko); err == nil {
        query = query.Where("id_toko = ?", v)
    }
}


	var total int64
	query.Count(&total)

	var produk []models.Produk
	if err := query.Limit(limit).Offset(offset).Find(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data produk"})
	}

	return c.JSON(fiber.Map{
		"page":       page,
		"limit":      limit,
		"total":      total,
		"total_page": int(math.Ceil(float64(total) / float64(limit))),
		"data":       produk,
	})
}




func GetProdukByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var produk models.Produk

	if err := database.DB.
		Preload("Toko").
		Preload("Kategori").
		Preload("Photos").
		First(&produk, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
	}

	return c.JSON(produk)
}


func CreateProduk(c *fiber.Ctx) error {
    contentType := c.Get("Content-Type")

    var produk models.Produk

    if strings.Contains(contentType, "application/json") {
        if err := c.BodyParser(&produk); err != nil {
            return c.Status(400).JSON(fiber.Map{
                "error":  "Request JSON tidak valid",
                "detail": err.Error(),
            })
        }

        produk.CreatedAt = time.Now()
        produk.UpdatedAt = time.Now()

        if err := database.DB.Create(&produk).Error; err != nil {
            return c.Status(500).JSON(fiber.Map{
                "error":  "Gagal menambahkan produk (JSON)",
                "detail": err.Error(),
            })
        }

        return c.JSON(fiber.Map{
            "message": "Produk berhasil ditambahkan",
            "produk":  produk,
        })
    }

    if strings.Contains(contentType, "multipart/form-data") {

        produk.NamaProduk = c.FormValue("nama_produk")
        produk.Slug = c.FormValue("slug")
        produk.HargaReseller = c.FormValue("harga_reseller")
        produk.HargaKonsumen = c.FormValue("harga_konsumen")
        produk.Deskripsi = c.FormValue("deskripsi")

        stok, _ := strconv.Atoi(c.FormValue("stok"))
        produk.Stok = uint(stok)

        idToko, _ := strconv.Atoi(c.FormValue("id_toko"))
        produk.IDToko = uint(idToko)

        idKategori, _ := strconv.Atoi(c.FormValue("id_kategori"))
        produk.IDKategori = uint(idKategori)

        produk.CreatedAt = time.Now()
        produk.UpdatedAt = time.Now()

        if err := database.DB.Create(&produk).Error; err != nil {
            return c.Status(500).JSON(fiber.Map{
                "error":  "Gagal menambahkan produk (FormData)",
                "detail": err.Error(),
            })
        }

       
        os.MkdirAll("uploads", os.ModePerm)

        
        form, err := c.MultipartForm()
        if err == nil && form.File != nil {

            files := form.File["photos"]
            for _, file := range files {

                filename := fmt.Sprintf("uploads/%d_%s", time.Now().UnixNano(), file.Filename)

                
                if err := c.SaveFile(file, filename); err == nil {
                    foto := models.Foto_Produk{
                        ID_Produk: produk.ID,
                        URLFoto:   filename,
                        CreatedAt: time.Now(),
                        UpdatedAt: time.Now(),
                    }
                    database.DB.Create(&foto)
                }
            }
        }

        return c.JSON(fiber.Map{
            "message": "Produk  berhasil ditambahkan",
            "produk":  produk,
        })
    }

    return c.Status(400).JSON(fiber.Map{
        "error": "Content-Type tidak dikenali. Gunakan JSON atau Form-Data.",
    })
}





func UpdateProduk(c *fiber.Ctx) error {
	id := c.Params("id")
	var produk models.Produk

	if err := database.DB.First(&produk, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request tidak valid"})
	}

	data["updated_at"] = time.Now()
	if err := database.DB.Model(&produk).Updates(data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal memperbarui produk"})
	}

	return c.JSON(fiber.Map{
		"message": "Produk berhasil diperbarui",
		"produk":  produk,
	})
}


func DeleteProduk(c *fiber.Ctx) error {
	id := c.Params("id")
	var produk models.Produk

	if err := database.DB.First(&produk, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
	}

	if err := database.DB.Delete(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus produk"})
	}

	return c.JSON(fiber.Map{"message": "Produk berhasil dihapus"})
}

