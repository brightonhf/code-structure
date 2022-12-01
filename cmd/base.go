package cmd

import (
	"code_structure/internals/config"
	"context"
)

type (
	// BaseCommand hold common command properties
	BaseCommand struct {
		ctx            context.Context
		config         *config.Specification
		notifyOnListen chan string
	}
)

func NewBaseCommand(ctx context.Context, config *config.Specification) BaseCommand {

	return BaseCommand{
		ctx:    ctx,
		config: config,
	}

}
