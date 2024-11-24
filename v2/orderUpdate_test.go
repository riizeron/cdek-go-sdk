package v2

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClientImpl_OrderUpdate(t *testing.T) {
	ctx := context.Background()
	timedCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	c := createTestClient()

	resp, err := c.OrderRegister(timedCtx, nil)
	require.Error(t, err)
	require.Nil(t, resp)

	registerReq := &OrderRegisterRequest{
		Type:          0,
		TariffCode:    62,
		ShipmentPoint: "OMS1",
		DeliveryPoint: "OMS2",

		Recipient: Contact{
			Name: "test",
			Phones: []Phone{
				{Number: "123"},
			},
		},
		Packages: []Package{
			{
				Number: "test",
				Weight: 1,
				Items: []PackageItem{
					{
						Name:    "test",
						WareKey: "test",
					},
				},
			},
		},
	}
	resp, err = c.OrderRegister(timedCtx, registerReq)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Greater(t, len(resp.Requests), 0)

	updateResp, err := c.OrderUpdate(ctx, &OrderUpdateRequest{
		UUID:          resp.Entity.Uuid,
		Comment:       "updated",
		DeliveryPoint: registerReq.DeliveryPoint,
		TariffCode:    registerReq.TariffCode,
		Packages:      registerReq.Packages,
	})
	require.NoError(t, err)

	statusResp, err := c.OrderStatus(ctx, updateResp.Entity.Uuid)
	require.NoError(t, err)
	require.Equal(t, statusResp.Entity.Comment, "updated")
}
