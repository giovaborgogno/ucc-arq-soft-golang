package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var migrations = []*gormigrate.Migration{
	{
		ID: "202203141800",
		Migrate: func(tx *gorm.DB) error {
			type CartItem struct {
				gorm.Model
				ID        uint    `gorm:"primaryKey;autoIncrement"`
				CartID    uint    `gorm:"not null;constraint:onDelete:CASCADE"`
				ProductID uint    `gorm:"not null"`
				Quantity  uint    `gorm:"not null"`
				Price     float64 `gorm:"not null"`
			}
			type Cart struct {
				gorm.Model
				ID         uint       `gorm:"primaryKey;autoIncrement"`
				UserID     uint       `gorm:"not null;constraint:onDelete:CASCADE"`
				CartsItems []CartItem `gorm:"foreignKey:CartID;constraint:onDelete:CASCADE"`
			}
			type User struct {
				gorm.Model
				ID    uint   `gorm:"primaryKey;autoIncrement"`
				Name  string `gorm:"not null"`
				Email string `gorm:"not null;unique"`
				Cart  Cart   `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE"`
			}
			type Product struct {
				gorm.Model
				ID       uint     `gorm:"primaryKey;autoIncrement"`
				Name     string   `gorm:"not null"`
				Price    float64  `gorm:"not null"`
				CartItem CartItem `gorm:"foreignKey:ProductID;constraint:onDelete:CASCADE"`
			}

			return tx.AutoMigrate(&Product{}, &User{}, &Cart{}, &CartItem{})

		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("products", "cart_items", "carts", "users")

		},
	},
}

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	return m.Migrate()
}
