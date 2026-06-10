package client

import (
	"context"
	"fmt"

	"github.com/errom502/client-service/internal/dto"
	gatewayv1 "github.com/errom502/protolib/gen/gateway-service/v1"
	verificationv1 "github.com/errom502/protolib/gen/verification-service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type VerificationClient struct {
	conn   *grpc.ClientConn
	client gatewayv1.GatewayServiceClient
}

func NewVerificationClient(addr string) (*VerificationClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("NewVerificationClient: dial %s: %w", addr, err)
	}
	return &VerificationClient{
		conn:   conn,
		client: gatewayv1.NewGatewayServiceClient(conn),
	}, nil
}

func (c *VerificationClient) Close() error {
	return c.conn.Close()
}

func (c *VerificationClient) Create(ctx context.Context, extUserID, email string) (string, error) {
	resp, err := c.client.Create(ctx, &gatewayv1.CreateRequest{
		ExtId:        extUserID,
		ProviderType: verificationv1.ProviderType_PROVIDER_TYPE_EMAIL,
		Target:       email,
	})
	if err != nil {
		return "", err
	}
	return resp.GetId(), nil
}

func (c *VerificationClient) Retry(ctx context.Context, verificationID string) error {
	_, err := c.client.Retry(ctx, &gatewayv1.RetryRequest{
		Id: verificationID,
	})
	return err
}

// Filter возвращает верификации по email, отсортированные created_at asc.
func (c *VerificationClient) FilterByEmail(ctx context.Context, email string) ([]*dto.Verification, error) {
	resp, err := c.client.Filter(ctx, &gatewayv1.FilterRequest{
		Targets: []string{email},
		Limit:   100,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*dto.Verification, 0, len(resp.GetVerifications()))
	for _, v := range resp.GetVerifications() {
		result = append(result, protoToDTO(v))
	}
	return result, nil
}

func protoToDTO(v *verificationv1.Verification) *dto.Verification {
	d := &dto.Verification{
		ID:         v.GetId(),
		ExtID:      v.GetExtId(),
		Target:     v.GetTarget(),
		Status:     dto.VerificationStatus(v.GetStatus()),
		StatusText: dto.VerificationStatus(v.GetStatus()).String(),
		CreatedAt:  formatTS(v.GetCreatedAt()),
	}
	if v.GetExpiresAt() != nil {
		s := formatTS(v.GetExpiresAt())
		d.ExpiresAt = &s
	}
	if v.GetVerifiedAt() != nil {
		s := formatTS(v.GetVerifiedAt())
		d.VerifiedAt = &s
	}
	return d
}

func formatTS(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return ""
	}
	return ts.AsTime().Format("2006-01-02 15:04:05")
}
