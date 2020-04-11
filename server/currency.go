package currency

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/stan-iakov/proto/currency"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetRate(context.Context, *protos.RateRequest) (*protos.RateResponse, error) {
	return &protos.RateResponse{Rate: 0.5}, nil
}
