package v2

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
	"strings"
)

type Client interface {
	Auth(ctx context.Context) (*AuthResponse, error)
	DeliveryPoints(ctx context.Context, input *DeliveryPointsRequest) (*DeliveryPointsResponse, error)
	Regions(ctx context.Context, input *RegionsRequest) (*RegionsResponse, error)
	Cities(ctx context.Context, input *CitiesRequest) (*CitiesResponse, error)
	CalculatorTariffList(ctx context.Context, input *CalculatorTariffListRequest) (*CalculatorTrafiffListResponse, error)
	CalculatorTariff(ctx context.Context, input *CalculatorTariffRequest) (*Tariff, error)
	OrderRegister(ctx context.Context, input *OrderRegisterRequest) (*Response, error)
	OrderDelete(ctx context.Context, uuid string) (*Response, error)
	OrderUpdate(ctx context.Context, input *OrderUpdateRequest) (*OrderUpdateResponse, error)
	OrderStatus(ctx context.Context, uuid string) (*Response, error)
}

type Options struct {
	Endpoint    string
	Credentials *Credentials
}

//func NewClient(opts *Options) Client {
//	return &clientImpl{opts: opts}
//}

func NewClient(url string, clientId string, clientSecretId string) Client {
	opts := &Options{
		Endpoint: url,
		Credentials: &Credentials{
			ClientID:     clientId,
			ClientSecret: clientSecretId,
		},
	}
	return &clientImpl{opts: opts}
}

type clientImpl struct {
	opts        *Options
	accessToken string
	expireIn    int
}

func (c *clientImpl) buildUri(p string, values interface{}) string {
	v, _ := query.Values(values)
	return strings.TrimRight(fmt.Sprintf(
		"%s/%s?%s",
		strings.TrimRight(c.opts.Endpoint, "/"),
		strings.TrimLeft(p, "/"),
		v.Encode(),
	), "?")
}
