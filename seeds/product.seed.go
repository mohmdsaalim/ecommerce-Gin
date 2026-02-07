package seeds

import (
	"fmt"

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
			PrimaryImage:   "https://example.com/images/products/barca-home-kit-front.jpg",
			SecondaryImage: "https://example.com/images/products/barca-home-kit-back.jpg",
			ThumbnailImage: "https://example.com/images/products/barca-home-kit-thumb.jpg",
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
			PrimaryImage:   "https://example.com/images/products/barca-away-kit-front.jpg",
			SecondaryImage: "https://example.com/images/products/barca-away-kit-back.jpg",
			ThumbnailImage: "https://example.com/images/products/barca-away-kit-thumb.jpg",
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
			PrimaryImage:   "https://example.com/images/products/barca-third-kit-front.jpg",
			SecondaryImage: "https://example.com/images/products/barca-third-kit-back.jpg",
			ThumbnailImage: "https://example.com/images/products/barca-third-kit-thumb.jpg",
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
			PrimaryImage:   "https://example.com/images/products/barca-gk-kit-front.jpg",
			SecondaryImage: "https://example.com/images/products/barca-gk-kit-back.jpg",
			ThumbnailImage: "https://example.com/images/products/barca-gk-kit-thumb.jpg",
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
			Season:         "Retro",
			PrimaryImage:   "https://example.com/images/products/barca-retro-1999-front.jpg",
			SecondaryImage: "https://example.com/images/products/barca-retro-1999-back.jpg",
			ThumbnailImage: "https://example.com/images/products/barca-retro-1999-thumb.jpg",
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
			PrimaryImage:   "https://example.com/images/products/barca-hoodie-front.jpg",
			SecondaryImage: "https://example.com/images/products/barca-hoodie-back.jpg",
			ThumbnailImage: "https://example.com/images/products/barca-hoodie-thumb.jpg",
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
			PrimaryImage:   "https://example.com/images/products/barca-polo-front.jpg",
			SecondaryImage: "https://example.com/images/products/barca-polo-back.jpg",
			ThumbnailImage: "https://example.com/images/products/barca-polo-thumb.jpg",
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
			PrimaryImage:   "https://example.com/images/products/barca-track-jacket-front.jpg",
			SecondaryImage: "https://example.com/images/products/barca-track-jacket-back.jpg",
			ThumbnailImage: "https://example.com/images/products/barca-track-jacket-thumb.jpg",
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
			PrimaryImage:   "https://example.com/images/products/barca-training-polo-front.jpg",
			SecondaryImage: "https://example.com/images/products/barca-training-polo-back.jpg",
			ThumbnailImage: "https://example.com/images/products/barca-training-polo-thumb.jpg",
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
			PrimaryImage:   "https://example.com/images/products/barca-fleece-hoodie-front.jpg",
			SecondaryImage: "https://example.com/images/products/barca-fleece-hoodie-back.jpg",
			ThumbnailImage: "https://example.com/images/products/barca-fleece-hoodie-thumb.jpg",
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
	for _, product := range products {
		// First, upsert the product
		if err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "sku"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"name", "description", "category", "sub_category",
				"base_price", "season", "primary_image", "secondary_image",
				"thumbnail_image", "is_active", "updated_at",
			}),
		}).Create(&product).Error; err != nil {
			return fmt.Errorf("failed to seed product %s: %w", product.SKU, err)
		}

		// Fetch the product ID after upsert
		var existingProduct models.Product
		if err := db.Where("sku = ?", product.SKU).First(&existingProduct).Error; err != nil {
			return fmt.Errorf("failed to fetch product %s: %w", product.SKU, err)
		}

		// Upsert variants
		for _, variant := range product.Variants {
			variant.ProductID = existingProduct.ID
			if err := db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "sku"}},
				DoUpdates: clause.AssignmentColumns([]string{
					"product_id", "size", "stock", "is_active", "updated_at",
				}),
			}).Create(&variant).Error; err != nil {
				return fmt.Errorf("failed to seed variant %s: %w", variant.SKU, err)
			}
		}

		fmt.Printf("  ✓ Seeded/Updated product: %s (SKU: %s)\n", product.Name, product.SKU)
	}

	fmt.Println("✅ Products seeding completed!")
	return nil
}
