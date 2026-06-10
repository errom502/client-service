package dto

// VerificationStatus статус верификации совпадает с proto enum.
type VerificationStatus int32

const (
	VerificationStatusInvalid  VerificationStatus = 0
	VerificationStatusPending  VerificationStatus = 1
	VerificationStatusVerified VerificationStatus = 2
	VerificationStatusFailed   VerificationStatus = 3
	VerificationStatusExpired  VerificationStatus = 4
)

func (s VerificationStatus) String() string {
	switch s {
	case VerificationStatusPending:
		return "pending"
	case VerificationStatusVerified:
		return "verified"
	case VerificationStatusFailed:
		return "failed"
	case VerificationStatusExpired:
		return "expired"
	default:
		return "unknown"
	}
}

type Verification struct {
	ID         string             `json:"id"`
	ExtID      string             `json:"ext_id"`
	Target     string             `json:"target"`
	Status     VerificationStatus `json:"status"`
	StatusText string             `json:"status_text"`
	ExpiresAt  *string            `json:"expires_at,omitempty"`
	VerifiedAt *string            `json:"verified_at,omitempty"`
	CreatedAt  string             `json:"created_at"`
}

// CreateRequest запрос на создание верификации от фронтенда.
type CreateRequest struct {
	ExtUserID string `json:"ext_user_id" binding:"required"`
	Email     string `json:"email"       binding:"required"`
}

// RetryRequest запрос на повторную верификацию от фронтенда.
type RetryRequest struct {
	Email     string `json:"email"       binding:"required"`
	ExtUserID string `json:"ext_user_id" binding:"required"`
}

// StatusResponse ответ для страницы проверки статуса.
type StatusResponse struct {
	Found        bool               `json:"found"`
	Status       VerificationStatus `json:"status,omitempty"`
	StatusText   string             `json:"status_text,omitempty"`
	Message      string             `json:"message"`
	Verification *Verification      `json:"verification,omitempty"`
}
