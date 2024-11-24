package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type OrderRegisterRequest struct {
	Type int `json:"type,omitempty"`
	// 136 - не только наземная, 234 - только наземная
	TariffCode    int    `json:"tariff_code"`
	ShipmentPoint string `json:"shipment_point"`
	DeliveryPoint string `json:"delivery_point"`

	// TODO: Че это за хуйня вообще?
	DeliveryRecipientCost    Payment `json:"delivery_recipient_cost"`
	DeliveryRecipientCostAdv Cost    `json:"delivery_recipient_cost_adv"`

	Recipient Contact `json:"recipient"`

	Packages []Package `json:"packages"`

	// TODO: Maybe будет полезным
	WidgetToken string `json:"widget_token,omitempty"`
}

func (c *clientImpl) OrderRegister(ctx context.Context, input *OrderRegisterRequest) (*Response, error) {
	payload, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.buildUri("/v2/orders", nil),
		bytes.NewReader(payload),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	accessToken, err := c.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := jsonReq[Response](req)
	if err != nil {
		return nil, err
	}

	if err := validateResponse(resp.Requests); err != nil {
		return nil, err
	}

	return resp, nil
}
