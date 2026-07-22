package auth

import (
	middleware "cryptox/internal/middleWare"
	webconfiguration "cryptox/internal/modules/webConfiguration"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func AuthRoutes(r fiber.Router, db *gorm.DB, redis *redis.Client, jwtSecret string, featureService webconfiguration.FeatureService) {

	repo := NewRepo(db)
	service := NewAuthService(repo, redis, jwtSecret)
	controller := NewAuthController(service)

	auth := r.Group("/auth")

	auth.Post("/register", middleware.FeatureMiddleware(featureService, "registration"), controller.Register)
	auth.Post("/login", middleware.FeatureMiddleware(featureService, "login"), controller.Login)
	auth.Post("/sendotp", controller.SendOTP)
	auth.Post("/forgototp", controller.ForgotPassWordOTP)
	auth.Post("/verifyotp", controller.VerifyOTP)
	auth.Post("/changepassword", controller.ForgotPassWordNewCreation)


	auth.Use(middleware.AuthMiddleWare(jwtSecret))
	auth.Post("/logout", controller.Logout)
	auth.Post("/refresh", controller.Refresh)
	
	////////// Admin Routes \\\\\\\\\\

	admin := r.Group("/admin")

	admin.Get("/getallusers", controller.GetAllUsers)
	admin.Get("/getbyid/:id", controller.GetByID)
	admin.Post("/editprofile/:id", controller.EditProfile)
	admin.Post("/blockunblock/:id", controller.BlockUnblock)
}