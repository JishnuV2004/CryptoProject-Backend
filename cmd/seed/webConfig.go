package seed

import (
	webconfiguration "cryptox/internal/modules/webConfiguration"
	"fmt"

	"gorm.io/gorm"
)

func WebConfig(db *gorm.DB) {

	db.Create(&webconfiguration.FeatureFlag{
		FeatureName: "registration",
		Enabled:     true,
	})

	db.Create(&webconfiguration.FeatureFlag{
		FeatureName: "login",
		Enabled:     true,
	})

	db.Create(&webconfiguration.FeatureFlag{
		FeatureName: "market",
		Enabled:     true,
	})

	db.Create(&webconfiguration.FeatureFlag{
		FeatureName: "trading",
		Enabled:     true,
	})

	db.Create(&webconfiguration.FeatureFlag{
		FeatureName: "wallet",
		Enabled:     true,
	})

	db.Create(&webconfiguration.FeatureFlag{
		FeatureName: "kyc",
		Enabled:     true,
	})

	fmt.Println("WebConfiguration Seeding Success")
}
