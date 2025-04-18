package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rilgilang/sticker-collection-api/config/dotenv"
	"github.com/rilgilang/sticker-collection-api/internal/api/presenter"
	"github.com/rilgilang/sticker-collection-api/internal/consts"
	"github.com/rilgilang/sticker-collection-api/internal/helper"
	"github.com/rilgilang/sticker-collection-api/internal/pkg/logger"
	"github.com/rilgilang/sticker-collection-api/internal/service"
)

func GetProfile(cfg *dotenv.Config, service service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			ctx = c.Context()
			log = logger.NewLog("get_profile_handler", cfg.LoggerEnable)
		)

		//log.Info(fmt.Sprintf(`start service login for user %s`, requestBody.Email))
		serv := service.GetProfile(ctx, helper.InterfaceToString(c.Locals(consts.UserId)))
		if serv.Code != 200 {
			//log.Error(fmt.Sprintf(`error on service login got %s`, serv.Errors))
			c.Status(serv.Code)
			return c.JSON(presenter.ErrorResponse(serv.Errors))
		}

		log.Info("get profile success")

		c.Status(200)
		return c.JSON(presenter.SuccessResponse(serv.Data))
	}
}
