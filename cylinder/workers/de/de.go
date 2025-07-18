package de

import (
	"github.com/bandprotocol/chain/v3/cylinder"
	"github.com/bandprotocol/chain/v3/cylinder/context"
	"github.com/bandprotocol/chain/v3/cylinder/msg"
)

var _ cylinder.Worker = &DE{}

// DE is a worker responsible for managing own nonce (DE) that being used for signing process
type DE struct {
	context *context.Context
	workers []cylinder.Worker
}

// New creates a new instance of the DE-related workers.
// It initializes the necessary components and returns the created DE instance or an error if initialization fails.
func New(ctx *context.Context) (*DE, error) {
	updateDE, err := NewUpdateDE(ctx)
	if err != nil {
		return nil, err
	}

	deleteDE, err := NewDeleteDE(ctx)
	if err != nil {
		return nil, err
	}

	return &DE{
		context: ctx,
		workers: []cylinder.Worker{updateDE, deleteDE},
	}, nil
}

// Start starts the DE-related workers.
func (de *DE) Start() {
	for _, w := range de.workers {
		go w.Start()
	}
}

// Stop stops DE-related workers.
func (de *DE) Stop() error {
	for _, w := range de.workers {
		if err := w.Stop(); err != nil {
			return err
		}
	}

	return nil
}

// GetResponseReceivers returns the message response receivers of the worker.
func (de *DE) GetResponseReceivers() []*msg.ResponseReceiver {
	receivers := []*msg.ResponseReceiver{}
	for _, w := range de.workers {
		receivers = append(receivers, w.GetResponseReceivers()...)
	}

	return receivers
}
