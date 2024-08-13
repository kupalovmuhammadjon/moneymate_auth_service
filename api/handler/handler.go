package handler

import (
	"auth_service/models"
	"auth_service/pkg/logger"
	"auth_service/storage"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	log     logger.ILogger
	storage storage.IStorage
}

func NewHandler(log logger.ILogger, storage storage.IStorage) *Handler {
	return &Handler{
		log:     log,
		storage: storage,
	}
}

func handleResponse(ctx *gin.Context, log logger.ILogger, msg string, statusCode int, data interface{}) {

	resp := models.Response{}

	switch {
	case statusCode >= 200 && statusCode < 300:
		resp.Description = "OK"
		log.Info("Response OK", logger.String("msg", msg), logger.Int("status", statusCode))
	case statusCode == 400:
		resp.Description = "Bad Request"
		log.Warn("Bad Request", logger.String("msg", msg), logger.Int("status", statusCode), logger.Any("error", data))
	case statusCode == 401:
		resp.Description = "Unauthorized"
		log.Warn("Unauthorized", logger.String("msg", msg), logger.Int("status", statusCode))
	case statusCode >= 500:
		resp.Description = "Internal Server Error"
		log.Error("Internal Server Error", logger.String("msg", msg), logger.Int("status", statusCode), logger.Any("error", data))
	default:
		resp.Description = "Unknown Error"
		log.Error("Unknown Error", logger.String("msg", msg), logger.Int("status", statusCode), logger.Any("error", data))
	}

	resp.StatusCode = statusCode
	resp.Data = data

	ctx.JSON(resp.StatusCode, resp)
}
