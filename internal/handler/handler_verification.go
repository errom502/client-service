package handler

import (
	"errors"
	"net/http"

	"github.com/errom502/client-service/internal/dto"
	"github.com/errom502/client-service/internal/usecase"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VerificationHandler struct {
	uc     *usecase.VerificationUsecase
	logger *zap.Logger
}

func NewVerificationHandler(uc *usecase.VerificationUsecase, l *zap.Logger) *VerificationHandler {
	return &VerificationHandler{
		uc:     uc,
		logger: l,
	}
}

// CheckStatus GET /api/v1/verification/status?email=xxx
func (h *VerificationHandler) CheckStatus(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	resp, err := h.uc.CheckStatus(c.Request.Context(), email)
	if err != nil {
		h.logger.Error("VerificationHandler.CheckStatus: internal error", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Create POST /api/v1/verification/create
func (h *VerificationHandler) Create(c *gin.Context) {
	var req dto.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	id, err := h.uc.Create(c.Request.Context(), req.ExtUserID, req.Email)
	if err != nil {
		code, msg := h.grpcErrToHTTP(err)
		c.JSON(code, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Retry POST /api/v1/verification/retry
func (h *VerificationHandler) Retry(c *gin.Context) {
	var req dto.RetryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.uc.RetryByEmail(c.Request.Context(), req.Email, req.ExtUserID); err != nil {
		if errors.Is(err, usecase.ErrVerificationActive) {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Для этой почты уже есть активная верификация. Проверьте почту и перейдите по ссылке.",
			})
			return
		}
		code, msg := h.grpcErrToHTTP(err)
		c.JSON(code, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Письмо с подтверждением отправлено. Проверьте почту."})
}

// grpcErrToHTTP маппит gRPC ошибки в HTTP коды.
func (h *VerificationHandler) grpcErrToHTTP(err error) (int, string) {
	st, ok := status.FromError(err)
	if !ok {
		h.logger.Error("VerificationHandler.CheckStatus: internal error can't get status")
		return http.StatusInternalServerError, "internal error"
	}
	switch st.Code() {
	case codes.InvalidArgument:
		return http.StatusBadRequest, "Некорректные данные"
	case codes.NotFound:
		return http.StatusNotFound, "Запись не найдена"
	case codes.AlreadyExists:
		return http.StatusConflict, "Верификация уже существует и ожидает подтверждения"
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests, "Слишком много запросов"
	case codes.FailedPrecondition:
		return http.StatusUnprocessableEntity, "Операция недоступна в текущем состоянии: " + st.Message()
	default:
		return http.StatusInternalServerError, "internal error"
	}
}
