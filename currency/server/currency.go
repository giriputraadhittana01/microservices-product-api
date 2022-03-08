package server

import (
	"context"
	protos "currentcyproject/protos/currency"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}

func (c *Currency) Calculate(ctx context.Context, rr *protos.CalculateRequest) (*protos.CalculateResponse, error) {
	c.log.Info("Handle Calculate", "X", rr.GetX(), "Y", rr.GetY())
	res := rr.GetX() * rr.GetY()
	return &protos.CalculateResponse{Result: res}, nil
}

func (c *Currency) Greeting(ctx context.Context, rr *protos.GreetingRequest) (*protos.GreetingResponse, error) {
	c.log.Info("Handle Greeting", "Name", rr.Name)
	res := "Hello " + rr.GetName()
	return &protos.GreetingResponse{Result: res}, nil
}
