package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rilgilang/sticker-collection-api/config/dotenv"
	"github.com/rilgilang/sticker-collection-api/internal/api/handlers"
	"github.com/rilgilang/sticker-collection-api/internal/middlewares/jwt"
	"github.com/rilgilang/sticker-collection-api/internal/service"
)

func AuthRouter(app fiber.Router, cfg *dotenv.Config, authMiddleware jwt.AuthMiddleware, service service.AuthService) {
	app.Post("/login", handlers.Login(cfg, service))
	app.Get("/profile", authMiddleware.ValidateToken(), handlers.GetProfile(cfg, service))
}
