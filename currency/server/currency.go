package server

import (
	"context"
	"currentcyproject/data"
	protos "currentcyproject/protos/currency"
	"io"
	"time"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	rates *data.ExchangeRates
	log   hclog.Logger
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{r, l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Rate: rate}, nil
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

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {
	// handle client messages
	go func() {
		for {
			rr, err := src.Recv() // Recv is a blocking method which returns on client data
			// io.EOF signals that the client has closed the connection
			if err == io.EOF {
				c.log.Info("Client has closed connection")
				break
			}

			// any other error means the transport between the server and client is unavailable
			if err != nil {
				c.log.Error("Unable to read from client", "error", err)
				break
			}

			c.log.Info("Handle client request", "request_base", rr.GetBase(), "request_dest", rr.GetDestination())
		}
	}()
	// handle server responses
	// we block here to keep the connection open
	for {
		// send a message back to the client
		err := src.Send(&protos.RateResponse{Rate: 12.1})
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
}
