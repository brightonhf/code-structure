package cmd

// ConsumerCommand holds the necessary data for its execution
type ConsumerCommand interface {
	Run()
}

type consumerCommand struct {
	BaseCommand
}

// NewConsumerCommand constructs ConsumerCommand
func NewConsumerCommand(baseCommand BaseCommand) ConsumerCommand {
	return &consumerCommand{baseCommand}
}

func (c consumerCommand) Run() {
	//TODO implement me
	panic("implement me")
}
