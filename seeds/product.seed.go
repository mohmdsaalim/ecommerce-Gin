package seeds

import (
	"fmt"
	"log"

	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"gorm.io/gorm"
	// Update this import path to match your project
	// "your-project-name/models"
)

// SeedProducts seeds the products table with FC Barcelona store items
func SeedProducts(db *gorm.DB) error {
	fmt.Println("🌱 Seeding products...")

	products := []models.Product{
		// ==================== KITS - HOME ====================
		{
			Name:        "FC Barcelona Home Kit 2024/25 - L",
			Description: "Official FC Barcelona home jersey for the 2024/25 season. Blaugrana colors with club crest.",
			Price:       89.99,
			Category:    "kits",
			SubCategory: "home",
			Stock:       80,
			Size:        "L",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/15_94c2120f-2a60-4a26-a00a-cbae1b986211.jpg?v=1767964096&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Home Kit 2024/25 - XL",
			Description: "Official FC Barcelona home jersey for the 2024/25 season. Blaugrana colors with club crest.",
			Price:       89.99,
			Category:    "kits",
			SubCategory: "home",
			Stock:       70,
			Size:        "XL",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/HJ4611-456-PHSYMG01.jpg?v=1751430940&width=1200",
			IsAvailable: true,
		},
		
		// ==================== KITS - AWAY ====================
		{
			Name:        "FC Barcelona Away Kit 2024/25 - M",
			Description: "Official FC Barcelona away jersey with modern design and club badge.",
			Price:       89.99,
			Category:    "kits",
			SubCategory: "away",
			Stock:       75,
			Size:        "M",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/HJ5210-455_419384899_D_A_1X1_6174322e-594d-4542-91f2-6e6fc229931e.jpg?v=1751433731&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Away Kit 2024/25 - L",
			Description: "Official FC Barcelona away jersey with modern design and club badge.",
			Price:       89.99,
			Category:    "kits",
			SubCategory: "away",
			Stock:       65,
			Size:        "L",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/BARCA1-24662.jpg?v=1753951926&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Away Kit 2024/25 - XL",
			Description: "Official FC Barcelona away jersey with modern design and club badge.",
			Price:       89.99,
			Category:    "kits",
			SubCategory: "away",
			Stock:       60,
			Size:        "XL",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/HJ4603-784_431735711_D_A_1X1_48cfc706-a037-4073-aa5a-8abbde823c12.jpg?v=1753951926&width=1200",
			IsAvailable: true,
		},
		
		// ==================== KITS - GOALKEEPER ====================
		{
			Name:        "FC Barcelona Goalkeeper Kit 2024/25 - M",
			Description: "Official FC Barcelona goalkeeper jersey with protective padding areas.",
			Price:       94.99,
			Category:    "kits",
			SubCategory: "goalkeeper",
			Stock:       30,
			Size:        "M",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/HQ9290-511_419456431_D_C_1X1_b7bd6a5d-4f17-4b98-8c6a-acdfa158b23a.jpg?v=1753764164&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Goalkeeper Kit 2024/25 - L",
			Description: "Official FC Barcelona goalkeeper jersey with protective padding areas.",
			Price:       94.99,
			Category:    "kits",
			SubCategory: "goalkeeper",
			Stock:       25,
			Size:        "L",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/HQ0477-311_419456390_D_A_1X1_0607bc28-a02f-49a2-82ea-94c2554211ba.jpg?v=1751430429&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Goalkeeper Kit 2024/25 - XL",
			Description: "Official FC Barcelona goalkeeper jersey with protective padding areas.",
			Price:       94.99,
			Category:    "kits",
			SubCategory: "goalkeeper",
			Stock:       20,
			Size:        "XL",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/IM0463-511_419308725_D_D_4X5ucl.jpg?v=1753763789&width=1200",
			IsAvailable: true,
		},
		
		// ==================== LIFESTYLE - HOODIE ====================
		{
			Name:        "FC Barcelona Classic Hoodie - M",
			Description: "Comfortable FC Barcelona hoodie with club logo. Perfect for casual wear.",
			Price:       64.99,
			Category:    "lifestyle",
			SubCategory: "hoodie",
			Stock:       120,
			Size:        "M",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/Organic_Masc_Baixa--20_f35413c7-2b3f-4f42-9f45-e3115b85670c.jpg?v=1763463147&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Classic Hoodie - L",
			Description: "Comfortable FC Barcelona hoodie with club logo. Perfect for casual wear.",
			Price:       64.99,
			Category:    "lifestyle",
			SubCategory: "hoodie",
			Stock:       95,
			Size:        "L",
			ImageURL:    "https://example.com/images/hoodie-l.jpg",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Classic Hoodie - XL",
			Description: "Comfortable FC Barcelona hoodie with club logo. Perfect for casual wear.",
			Price:       64.99,
			Category:    "lifestyle",
			SubCategory: "hoodie",
			Stock:       80,
			Size:        "XL",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/IMG_8055_1d6bc7b5-de38-4c44-ad2f-0eeeb767a121.jpg?v=1757318666&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Home Kit 2024/25 - S",
			Description: "Official FC Barcelona home jersey for the 2024/25 season. Blaugrana colors with club crest.",
			Price:       89.99,
			Category:    "kits",
			SubCategory: "home",
			Stock:       50,
			Size:        "S",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/HJ4590-456_415227879_D_A_1X1_2laliga_92b83d62-53fe-4728-b143-3b8653e39427.jpg?v=1763654921&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Home Kit 2024/25 - M",
			Description: "Official FC Barcelona home jersey for the 2024/25 season. Blaugrana colors with club crest.",
			Price:       89.99,
			Category:    "kits",
			SubCategory: "home",
			Stock:       100,
			Size:        "M",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/base-ambilight.webp?v=12057333375252784091",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Premium Zip Hoodie - M",
			Description: "Premium quality zip-up hoodie with embroidered Barça crest.",
			Price:       79.99,
			Category:    "lifestyle",
			SubCategory: "hoodie",
			Stock:       60,
			Size:        "M",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/FCB_ECOMMERCE_2025408-2.jpg?v=1752663102&width=1200",
			IsAvailable: true,
		},

		// ==================== LIFESTYLE - RETRO ====================
		{
			Name:        "FC Barcelona Retro Jersey 1999 - M",
			Description: "Vintage FC Barcelona jersey from the legendary 1999 season.",
			Price:       79.99,
			Category:    "lifestyle",
			SubCategory: "retro",
			Stock:       45,
			Size:        "M",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/10RETRO1718_5.jpg?v=1707728541&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Retro Jersey 1999 - L",
			Description: "Vintage FC Barcelona jersey from the legendary 1999 season.",
			Price:       79.99,
			Category:    "lifestyle",
			SubCategory: "retro",
			Stock:       35,
			Size:        "L",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/700x1060-BLMP0007401707-1_0c38b1bd-737e-466c-bf55-b3ea7923fb29.jpg?v=1743685367&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Retro Jersey 2006 - M",
			Description: "Classic FC Barcelona jersey from the 2006 Champions League winning season.",
			Price:       84.99,
			Category:    "lifestyle",
			SubCategory: "retro",
			Stock:       30,
			Size:        "M",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/Retro_Players_Baixa-8064.jpg?v=1763462853&width=460",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Retro Jersey 2011 - L",
			Description: "Iconic FC Barcelona jersey from the historic 2011 season.",
			Price:       84.99,
			Category:    "lifestyle",
			SubCategory: "retro",
			Stock:       25,
			Size:        "L",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/2023-09-28-BLM-GAVI-AC203159.webp?v=1737471758&width=1200",
			IsAvailable: true,
		},

		// ==================== LIFESTYLE - CASUALS ====================
		{
			Name:        "FC Barcelona Casual Joggers - L",
			Description: "Comfortable joggers with FC Barcelona branding.",
			Price:       54.99,
			Category:    "lifestyle",
			SubCategory: "casuals",
			Stock:       85,
			Size:        "L",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/Bolet_Baixa-11795.jpg?v=1765970839&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Casual T-Shirt - M",
			Description: "Everyday FC Barcelona t-shirt with minimalist club logo.",
			Price:       34.99,
			Category:    "lifestyle",
			SubCategory: "casuals",
			Stock:       200,
			Size:        "M",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/PLAYERS_DAY_RETRO_0302.jpg?v=1765279539&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Casual POLO T-Shirt - L",
			Description: "Everyday FC Barcelona t-shirt with minimalist club logo.",
			Price:       34.99,
			Category:    "lifestyle",
			SubCategory: "casuals",
			Stock:       180,
			Size:        "L",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/BLM1PLTX_1_b3874d2b-7ef0-4e67-a28e-0d26ca867d94.jpg?v=1720606505&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Casual T-Shirt - XL",
			Description: "Everyday FC Barcelona t-shirt with minimalist club logo.",
			Price:       34.99,
			Category:    "lifestyle",
			SubCategory: "casuals",
			Stock:       160,
			Size:        "XL",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/products/700x1060-BLMP000813003-3.jpg?v=1743685526&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Casual Polo - M",
			Description: "Smart casual FC Barcelona polo shirt, perfect for any occasion.",
			Price:       49.99,
			Category:    "lifestyle",
			SubCategory: "casuals",
			Stock:       150,
			Size:        "M",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/FN8286-011_1.jpg?v=1738926528&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Casual Polo - L",
			Description: "Smart casual FC Barcelona polo shirt, perfect for any occasion.",
			Price:       49.99,
			Category:    "lifestyle",
			SubCategory: "casuals",
			Stock:       130,
			Size:        "L",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/FN8286-066_1.jpg?v=1723110386&width=1200",
			IsAvailable: true,
		},
		{
			Name:        "FC Barcelona Polo - M",
			Description: "Comfortable joggers with FC Barcelona branding.",
			Price:       54.99,
			Category:    "lifestyle",
			SubCategory: "casuals",
			Stock:       100,
			Size:        "M",
			ImageURL:    "https://store.fcbarcelona.com/cdn/shop/files/BLMP884046_1.jpg?v=1743584974&width=1200",
			IsAvailable: true,
		},
	}

	// Create products in database
	for _, product := range products {
	if err := db.
		Where("name = ? AND size = ?", product.Name, product.Size).
		Assign(models.Product{
			Description: product.Description,
			Price:       product.Price,
			Category:    product.Category,
			SubCategory: product.SubCategory,
			Stock:       product.Stock,
			ImageURL:    product.ImageURL,
			IsAvailable: product.IsAvailable,
		}).
		FirstOrCreate(&product).
		Error; err != nil {

		log.Printf("❌ Error seeding product '%s': %v", product.Name, err)
	}
}
return nil
}