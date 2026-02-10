package seeds

import (
	"fmt"
	"time"

	"github.com/mohmdsaalim/ecommerce-Gin/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedProducts(db *gorm.DB) error {
	fmt.Println("🌱 Seeding products...")

	products := []models.Product{
		{
			Name:           "FC Barcelona Home Kit 2024/25",
			Description:    "Official FC Barcelona home jersey for the 2024/25 season. Features the iconic Blaugrana stripes with modern design elements. Made with advanced moisture-wicking fabric technology and the official Nike Dri-FIT system. Includes authentic club crest and sponsor logos.",
			Category:       "Kits",
			SubCategory:    "Home",
			BasePrice:      89.99,
			SKU:            "KIT-HOME-2425-001",
			Season:         "2024/25",
			PrimaryImage:   "https://store.fcbarcelona.com/cdn/shop/files/HJ4590-456_415227879_D_A_1X1_2laliga_92b83d62-53fe-4728-b143-3b8653e39427.jpg?v=1763654921&width=1200",
			SecondaryImage: "https://store.fcbarcelona.com/cdn/shop/files/HJ4590-456_415227879_D_E_1X1_0fb3accd-440b-43d9-b5c8-db59754589ae.jpg?v=1763654921&width=1200",
			ThumbnailImage: "https://store.fcbarcelona.com/cdn/shop/files/HJ4590-456_415227879_D_D_1X1_e49a5dea-e5e6-46b2-b1cf-ae3500e62d90.jpg?v=1763654921&width=1200",
			IsActive:       true,
			Variants: []models.ProductVariant{
				{
					Size:     "S",
					Stock:    120,
					SKU:      "KIT-HOME-2425-001-S",
					IsActive: true,
				},
				{
					Size:     "M",
					Stock:    200,
					SKU:      "KIT-HOME-2425-001-M",
					IsActive: true,
				},
				{
					Size:     "L",
					Stock:    180,
					SKU:      "KIT-HOME-2425-001-L",
					IsActive: true,
				},
				{
					Size:     "XL",
					Stock:    150,
					SKU:      "KIT-HOME-2425-001-XL",
					IsActive: true,
				},
				{
					Size:     "XXL",
					Stock:    100,
					SKU:      "KIT-HOME-2425-001-XXL",
					IsActive: true,
				},
			},
		},
		{
			Name:           "FC Barcelona Away Kit 2024/25",
			Description:    "Official FC Barcelona away jersey for the 2024/25 season. Stunning design with vibrant colors inspired by club heritage. Premium quality construction with breathable mesh panels and Nike's latest performance technology. Perfect for match day or casual wear.",
			Category:       "Kits",
			SubCategory:    "Away",
			BasePrice:      89.99,
			SKU:            "KIT-AWAY-2425-002",
			Season:         "2024/25",
			PrimaryImage:   "https://store.fcbarcelona.com/cdn/shop/files/base-ambilight_1.webp?v=6857459708159031844",
			SecondaryImage: "https://store.fcbarcelona.com/cdn/shop/files/fcbarcelonanikefootball-29_2.jpg?v=1753762193&width=1200",
			ThumbnailImage: "https://store.fcbarcelona.com/cdn/shop/files/HJ4554-784_431735159_D_C_1X1_ddb660d8-72dc-4d0f-9c48-845bc43d0031.jpg?v=1753762193&width=1200",
			IsActive:       true,
			Variants: []models.ProductVariant{
				{
					Size:     "S",
					Stock:    100,
					SKU:      "KIT-AWAY-2425-002-S",
					IsActive: true,
				},
				{
					Size:     "M",
					Stock:    160,
					SKU:      "KIT-AWAY-2425-002-M",
					IsActive: true,
				},
				{
					Size:     "L",
					Stock:    140,
					SKU:      "KIT-AWAY-2425-002-L",
					IsActive: true,
				},
				{
					Size:     "XL",
					Stock:    120,
					SKU:      "KIT-AWAY-2425-002-XL",
					IsActive: true,
				},
				{
					Size:     "XXL",
					Stock:    80,
					SKU:      "KIT-AWAY-2425-002-XXL",
					IsActive: true,
				},
			},
		},
		{
			Name:           "FC Barcelona Third Kit 2024/25",
			Description:    "Official FC Barcelona third jersey for the 2024/25 season. Bold and contemporary design that stands out on the pitch. Features cutting-edge fabric technology for optimal performance and comfort. Showcases the legendary Barça spirit in a unique colorway.",
			Category:       "Kits",
			SubCategory:    "Third",
			BasePrice:      89.99,
			SKU:            "KIT-THIRD-2425-003",
			Season:         "2024/25",
			PrimaryImage:   "https://store.fcbarcelona.com/cdn/shop/files/HM3193-855_415228861_D_A_1X1_a5a536e9-71ea-4bcf-9723-08165902f26f.jpg?v=1755587326&width=1200",
			SecondaryImage: "https://store.fcbarcelona.com/cdn/shop/files/Imagen3_0e6e6a49-38c6-4f84-b6b3-123831c6e498.jpg?v=1755587326&width=1200",
			ThumbnailImage: "https://store.fcbarcelona.com/cdn/shop/files/HM3193-855_415228861_D_B_1X1_858efe83-fb6b-45b2-83cf-9065ec597f56.jpg?v=1755587326&width=1200",
			IsActive:       true,
			Variants: []models.ProductVariant{
				{
					Size:     "S",
					Stock:    90,
					SKU:      "KIT-THIRD-2425-003-S",
					IsActive: true,
				},
				{
					Size:     "M",
					Stock:    140,
					SKU:      "KIT-THIRD-2425-003-M",
					IsActive: true,
				},
				{
					Size:     "L",
					Stock:    120,
					SKU:      "KIT-THIRD-2425-003-L",
					IsActive: true,
				},
				{
					Size:     "XL",
					Stock:    100,
					SKU:      "KIT-THIRD-2425-003-XL",
					IsActive: true,
				},
				{
					Size:     "XXL",
					Stock:    70,
					SKU:      "KIT-THIRD-2425-003-XXL",
					IsActive: true,
				},
			},
		},
		{
			Name:           "FC Barcelona Goalkeeper Kit 2024/25",
			Description:    "Official FC Barcelona goalkeeper jersey for the 2024/25 season. Specially designed with padded elbows and enhanced durability for shot-stopping action. Features bright, distinctive colors for maximum visibility and Nike's GK-specific technology.",
			Category:       "Kits",
			SubCategory:    "Goalkeeper",
			BasePrice:      84.99,
			SKU:            "KIT-GK-2425-004",
			Season:         "2024/25",
			PrimaryImage:   "https://store.fcbarcelona.com/cdn/shop/files/HQ9290-511_419456431_D_C_1X1_b7bd6a5d-4f17-4b98-8c6a-acdfa158b23a.jpg?v=1753764164&width=1200",
			SecondaryImage: "https://store.fcbarcelona.com/cdn/shop/files/HQ9290-511_419456431_D_D_1X1_37d599fa-7a8f-42c0-907b-032535d7aec7.jpg?v=1753764164&width=1200",
			ThumbnailImage: "https://store.fcbarcelona.com/cdn/shop/files/HQ9290-511_419456431_D_B_1X1_d905d573-d384-4190-b60d-0e11842d55ae.jpg?v=1753764164&width=1200",
			IsActive:       true,
			Variants: []models.ProductVariant{
				{
					Size:     "S",
					Stock:    40,
					SKU:      "KIT-GK-2425-004-S",
					IsActive: true,
				},
				{
					Size:     "M",
					Stock:    60,
					SKU:      "KIT-GK-2425-004-M",
					IsActive: true,
				},
				{
					Size:     "L",
					Stock:    55,
					SKU:      "KIT-GK-2425-004-L",
					IsActive: true,
				},
				{
					Size:     "XL",
					Stock:    45,
					SKU:      "KIT-GK-2425-004-XL",
					IsActive: true,
				},
				{
					Size:     "XXL",
					Stock:    30,
					SKU:      "KIT-GK-2425-004-XXL",
					IsActive: true,
				},
			},
		},
		{
			Name:           "FC Barcelona Retro 1999 Home Kit",
			Description:    "Classic FC Barcelona home jersey from the legendary 1999 season. Authentic reproduction of the iconic design worn during memorable victories. Premium quality materials with vintage styling and the nostalgic club crest. A must-have for collectors and longtime fans.",
			Category:       "Kits",
			SubCategory:    "Retro",
			BasePrice:      79.99,
			SKU:            "KIT-RETRO-1999-005",
			Season:         "fourth",
			PrimaryImage:   "https://store.fcbarcelona.com/cdn/shop/files/4Kit_SergiAlcazar_Ecom_044_1_e0846158-4a88-4306-a2fd-2b3e19943a11.jpg?v=1767964096&width=1200",
			SecondaryImage: "https://store.fcbarcelona.com/cdn/shop/files/16_1a4ee356-5197-4c11-b91e-d8ace6ddd127.jpg?v=1767964096&width=1200",
			ThumbnailImage: "https://store.fcbarcelona.com/cdn/shop/files/15_94c2120f-2a60-4a26-a00a-cbae1b986211.jpg?v=1767964096&width=1200",
			IsActive:       true,
			Variants: []models.ProductVariant{
				{
					Size:     "S",
					Stock:    50,
					SKU:      "KIT-RETRO-1999-005-S",
					IsActive: true,
				},
				{
					Size:     "M",
					Stock:    75,
					SKU:      "KIT-RETRO-1999-005-M",
					IsActive: true,
				},
				{
					Size:     "L",
					Stock:    65,
					SKU:      "KIT-RETRO-1999-005-L",
					IsActive: true,
				},
				{
					Size:     "XL",
					Stock:    55,
					SKU:      "KIT-RETRO-1999-005-XL",
					IsActive: true,
				},
				{
					Size:     "XXL",
					Stock:    35,
					SKU:      "KIT-RETRO-1999-005-XXL",
					IsActive: true,
				},
			},
		},
		{
			Name:           "FC Barcelona Premium Hoodie",
			Description:    "Premium FC Barcelona hoodie with modern streetwear design. Crafted from soft cotton blend fleece with an adjustable drawstring hood and kangaroo pocket. Features embroidered club crest and Nike Swoosh. Perfect for cold days at Camp Nou or everyday wear.",
			Category:       "Lifestyles",
			SubCategory:    "Hoodie",
			BasePrice:      69.99,
			SKU:            "LIFE-HOODIE-001",
			Season:         "All Season",
			PrimaryImage:   "https://store.fcbarcelona.com/cdn/shop/files/Retro_Players_Baixa--2.jpg?v=1763462759&width=1200",
			SecondaryImage: "https://store.fcbarcelona.com/cdn/shop/files/BZ3A1281.jpg?v=1763462759&width=1200",
			ThumbnailImage: "https://store.fcbarcelona.com/cdn/shop/files/BZ3A1289.jpg?v=1763462759&width=1200",
			IsActive:       true,
			Variants: []models.ProductVariant{
				{
					Size:     "S",
					Stock:    80,
					SKU:      "LIFE-HOODIE-001-S",
					IsActive: true,
				},
				{
					Size:     "M",
					Stock:    130,
					SKU:      "LIFE-HOODIE-001-M",
					IsActive: true,
				},
				{
					Size:     "L",
					Stock:    110,
					SKU:      "LIFE-HOODIE-001-L",
					IsActive: true,
				},
				{
					Size:     "XL",
					Stock:    90,
					SKU:      "LIFE-HOODIE-001-XL",
					IsActive: true,
				},
				{
					Size:     "XXL",
					Stock:    60,
					SKU:      "LIFE-HOODIE-001-XXL",
					IsActive: true,
				},
			},
		},
		{
			Name:           "FC Barcelona Classic Polo Shirt",
			Description:    "Elegant FC Barcelona polo shirt perfect for smart-casual occasions. Made from breathable piqué cotton with a comfortable regular fit. Features embroidered club crest, contrasting collar and sleeve trim. Ideal for match days or business casual settings.",
			Category:       "Lifestyles",
			SubCategory:    "Polo",
			BasePrice:      54.99,
			SKU:            "LIFE-POLO-002",
			Season:         "All Season",
			PrimaryImage:   "https://store.fcbarcelona.com/cdn/shop/files/Bolet_Baixa-11795.jpg?v=1765970839&width=1200",
			SecondaryImage: "https://store.fcbarcelona.com/cdn/shop/files/Bolet_Baixa-11825.jpg?v=1765970839&width=1200",
			ThumbnailImage: "https://store.fcbarcelona.com/cdn/shop/files/BLMP898003_1_1d16cce7-d927-4e3a-8b0f-4a383ab21674.jpg?v=1765970839&width=1200",
			IsActive:       true,
			Variants: []models.ProductVariant{
				{
					Size:     "S",
					Stock:    70,
					SKU:      "LIFE-POLO-002-S",
					IsActive: true,
				},
				{
					Size:     "M",
					Stock:    110,
					SKU:      "LIFE-POLO-002-M",
					IsActive: true,
				},
				{
					Size:     "L",
					Stock:    95,
					SKU:      "LIFE-POLO-002-L",
					IsActive: true,
				},
				{
					Size:     "XL",
					Stock:    75,
					SKU:      "LIFE-POLO-002-XL",
					IsActive: true,
				},
				{
					Size:     "XXL",
					Stock:    50,
					SKU:      "LIFE-POLO-002-XXL",
					IsActive: true,
				},
			},
		},
		{
			Name:           "FC Barcelona Retro Track Jacket",
			Description:    "Vintage-inspired FC Barcelona track jacket celebrating the club's golden era. Premium quality tricot fabric with authentic retro styling and classic Barça colors. Features full zip closure, side pockets, and embroidered badges. A timeless piece of football fashion.",
			Category:       "Lifestyles",
			SubCategory:    "Retro",
			BasePrice:      74.99,
			SKU:            "LIFE-RETRO-003",
			Season:         "All Season",
			PrimaryImage:   "https://store.fcbarcelona.com/cdn/shop/files/IMG_6325.jpg?v=1737554661&width=1200",
			SecondaryImage: "https://store.fcbarcelona.com/cdn/shop/files/IMG_7943.jpg?v=1737554661&width=1200",
			ThumbnailImage: "https://store.fcbarcelona.com/cdn/shop/files/BLMP740011_1.jpg?v=1737471543&width=1200",
			IsActive:       true,
			Variants: []models.ProductVariant{
				{
					Size:     "S",
					Stock:    60,
					SKU:      "LIFE-RETRO-003-S",
					IsActive: true,
				},
				{
					Size:     "M",
					Stock:    95,
					SKU:      "LIFE-RETRO-003-M",
					IsActive: true,
				},
				{
					Size:     "L",
					Stock:    85,
					SKU:      "LIFE-RETRO-003-L",
					IsActive: true,
				},
				{
					Size:     "XL",
					Stock:    70,
					SKU:      "LIFE-RETRO-003-XL",
					IsActive: true,
				},
				{
					Size:     "XXL",
					Stock:    45,
					SKU:      "LIFE-RETRO-003-XXL",
					IsActive: true,
				},
			},
		},
		{
			Name:           "FC Barcelona Training Polo 2024/25",
			Description:    "Official FC Barcelona training polo shirt used by the players. Advanced Nike Dri-FIT technology keeps you dry and comfortable during intense sessions. Features mesh panels for enhanced ventilation and the authentic club crest. Professional quality for dedicated fans.",
			Category:       "Lifestyles",
			SubCategory:    "Polo",
			BasePrice:      59.99,
			SKU:            "LIFE-POLO-TRAIN-004",
			Season:         "2024/25",
			PrimaryImage:   "https://store.fcbarcelona.com/cdn/shop/products/700x1060-BLMP0008120102-1.jpg?v=1680014731&width=1200",
			SecondaryImage: "https://store.fcbarcelona.com/cdn/shop/files/BLMP0008120102_3_1.jpg?v=1698158764&width=1200",
			ThumbnailImage: "https://store.fcbarcelona.com/cdn/shop/files/BLMP0008120102_3.jpg?v=1698158763&width=1200",
			IsActive:       true,
			Variants: []models.ProductVariant{
				{
					Size:     "S",
					Stock:    65,
					SKU:      "LIFE-POLO-TRAIN-004-S",
					IsActive: true,
				},
				{
					Size:     "M",
					Stock:    105,
					SKU:      "LIFE-POLO-TRAIN-004-M",
					IsActive: true,
				},
				{
					Size:     "L",
					Stock:    90,
					SKU:      "LIFE-POLO-TRAIN-004-L",
					IsActive: true,
				},
				{
					Size:     "XL",
					Stock:    70,
					SKU:      "LIFE-POLO-TRAIN-004-XL",
					IsActive: true,
				},
				{
					Size:     "XXL",
					Stock:    45,
					SKU:      "LIFE-POLO-TRAIN-004-XXL",
					IsActive: true,
				},
			},
		},
		{
			Name:           "FC Barcelona Fleece Hoodie - Blaugrana Edition",
			Description:    "Exclusive FC Barcelona fleece hoodie in the iconic Blaugrana colors. Ultra-soft brushed fleece interior provides maximum warmth and comfort. Features large embroidered club crest, front pocket, and ribbed cuffs. Essential for every Culé's wardrobe.",
			Category:       "Lifestyles",
			SubCategory:    "Hoodie",
			BasePrice:      64.99,
			SKU:            "LIFE-HOODIE-FLEECE-005",
			Season:         "Fall/Winter",
			PrimaryImage:   "https://store.fcbarcelona.com/cdn/shop/files/2023-04-24-BLM-TOMASSATORANSKY-2218.jpg?v=1727089507&width=1200",
			SecondaryImage: "https://store.fcbarcelona.com/cdn/shop/products/700x1060-BLM2BHC-1.jpg?v=1689944517&width=1200",
			ThumbnailImage: "https://store.fcbarcelona.com/cdn/shop/files/FCB_13_04_23_Dia_2_S_022_1740.jpg?v=1689944509&width=1200",
			IsActive:       true,
			Variants: []models.ProductVariant{
				{
					Size:     "S",
					Stock:    75,
					SKU:      "LIFE-HOODIE-FLEECE-005-S",
					IsActive: true,
				},
				{
					Size:     "M",
					Stock:    120,
					SKU:      "LIFE-HOODIE-FLEECE-005-M",
					IsActive: true,
				},
				{
					Size:     "L",
					Stock:    100,
					SKU:      "LIFE-HOODIE-FLEECE-005-L",
					IsActive: true,
				},
				{
					Size:     "XL",
					Stock:    80,
					SKU:      "LIFE-HOODIE-FLEECE-005-XL",
					IsActive: true,
				},
				{
					Size:     "XXL",
					Stock:    55,
					SKU:      "LIFE-HOODIE-FLEECE-005-XXL",
					IsActive: true,
				},
			},
		},
	}

	// Upsert products with variants using SKU as the conflict key
	for i := range products {

		product := &products[i]

		//  Upsert Product (WITHOUT auto-saving variants)
		if err := db.
			Omit("Variants").
			Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "sku"}},
				DoUpdates: clause.Assignments(map[string]interface{}{
					"name":            product.Name,
					"description":     product.Description,
					"category":        product.Category,
					"sub_category":    product.SubCategory,
					"base_price":      product.BasePrice,
					"season":          product.Season,
					"primary_image":   product.PrimaryImage,
					"secondary_image": product.SecondaryImage,
					"thumbnail_image": product.ThumbnailImage,
					"is_active":       product.IsActive,
					"updated_at":      time.Now(),
				}),
			}).Create(product).Error; err != nil {
			return fmt.Errorf("failed to seed product %s: %w", product.SKU, err)
		}

		// Get Product ID
		var existingProduct models.Product
		if err := db.Where("sku = ?", product.SKU).
			First(&existingProduct).Error; err != nil {
			return fmt.Errorf("failed to fetch product %s: %w", product.SKU, err)
		}

		//  Upsert Variants
		for j := range product.Variants {

			variant := &product.Variants[j]
			variant.ProductID = existingProduct.ID

			if err := db.
				Clauses(clause.OnConflict{
					Columns: []clause.Column{{Name: "sku"}},
					DoUpdates: clause.Assignments(map[string]interface{}{
						"product_id": variant.ProductID,
						"size":       variant.Size,
						"stock":      variant.Stock,
						"is_active":  variant.IsActive,
						"updated_at": time.Now(),
					}),
				}).Create(variant).Error; err != nil {
				return fmt.Errorf("failed to seed variant %s: %w", variant.SKU, err)
			}
		}

		fmt.Printf("✓ Seeded/Updated product: %s (SKU: %s)\n", product.Name, product.SKU)
	}

	fmt.Println("✅ Products seeding completed!")
	return nil
}