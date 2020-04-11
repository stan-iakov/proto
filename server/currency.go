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

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.6}, nil
}
