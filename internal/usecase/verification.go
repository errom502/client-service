package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/errom502/client-service/internal/client"
	"github.com/errom502/client-service/internal/dto"
)

var (
	// ErrVerificationActive — для почты уже есть активная верификация (pending).
	// Нельзя сделать retry пока не истечёт текущая.
	ErrVerificationActive = errors.New("active verification already exists for this email")
)

type VerificationUsecase struct {
	client *client.VerificationClient
}

func NewVerificationUsecase(c *client.VerificationClient) *VerificationUsecase {
	return &VerificationUsecase{client: c}
}

// CheckStatus возвращает последнюю верификацию по email и сообщение для пользователя.
// Сортировка на стороне сервиса created_at asc — берём последний элемент.
func (u *VerificationUsecase) CheckStatus(ctx context.Context, email string) (*dto.StatusResponse, error) {
	verifications, err := u.client.FilterByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("VerificationUsecase.CheckStatus: %w", err)
	}

	if len(verifications) == 0 {
		return &dto.StatusResponse{
			Found:   false,
			Message: "Верификация для данной почты не найдена",
		}, nil
	}

	latest := verifications[len(verifications)-1]

	resp := &dto.StatusResponse{
		Found:        true,
		Status:       latest.Status,
		StatusText:   latest.StatusText,
		Verification: latest,
	}

	switch latest.Status {
	case dto.VerificationStatusPending:
		resp.Message = "Письмо с подтверждением отправлено. Проверьте почту и перейдите по ссылке."
	case dto.VerificationStatusVerified:
		resp.Message = "Почта успешно верифицирована."
	case dto.VerificationStatusFailed:
		resp.Message = "Верификация не удалась. Вы можете запустить повторную верификацию."
	case dto.VerificationStatusExpired:
		resp.Message = "Срок верификации истёк. Вы можете запустить повторную верификацию."
	default:
		resp.Message = "Неизвестный статус верификации."
	}

	return resp, nil
}

// Create создаёт новую верификацию для email.
func (u *VerificationUsecase) Create(ctx context.Context, extUserID, email string) (string, error) {
	id, err := u.client.Create(ctx, extUserID, email)
	if err != nil {
		return "", fmt.Errorf("VerificationUsecase.Create: %w", err)
	}
	return id, nil
}

// RetryByEmail реализует логику повторной верификации:
//   - нет верификаций → создать новую
//   - статус pending → ошибка (активная верификация)
//   - статус verified → создать новую
//   - статус expired/failed → retry существующей
func (u *VerificationUsecase) RetryByEmail(ctx context.Context, email, extUserID string) error {
	verifications, err := u.client.FilterByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("VerificationUsecase.RetryByEmail: filter: %w", err)
	}

	// Нет верификаций — создаём первую
	if len(verifications) == 0 {
		if _, err := u.client.Create(ctx, extUserID, email); err != nil {
			return fmt.Errorf("VerificationUsecase.RetryByEmail: create: %w", err)
		}
		return nil
	}

	latest := verifications[len(verifications)-1]

	switch latest.Status {
	case dto.VerificationStatusPending:
		// Активная верификация — нельзя создать новую или сделать retry
		return ErrVerificationActive

	case dto.VerificationStatusVerified:
		// Уже верифицирована — создаём новую верификацию
		if _, err := u.client.Create(ctx, extUserID, email); err != nil {
			return fmt.Errorf("VerificationUsecase.RetryByEmail: create after verified: %w", err)
		}
		return nil

	case dto.VerificationStatusExpired, dto.VerificationStatusFailed:
		// Retry существующей верификации
		if err := u.client.Retry(ctx, latest.ID); err != nil {
			return fmt.Errorf("VerificationUsecase.RetryByEmail: retry: %w", err)
		}
		return nil

	default:
		return fmt.Errorf("VerificationUsecase.RetryByEmail: unknown status: %s", latest.StatusText)
	}
}
