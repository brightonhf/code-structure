package cmd

// MigrateCommand ...
type MigrateCommand interface {
	Up()
	Down()
	Run()
}

type migrateCommand struct {
	BaseCommand
}

// NewMigrateCommand constructs MigrateCommand
func NewMigrateCommand(baseCommand BaseCommand) MigrateCommand {
	return &migrateCommand{
		BaseCommand: baseCommand,
	}
}

func (m migrateCommand) Up() {
	//TODO implement me
	panic("implement me")
}

func (m migrateCommand) Down() {
	//TODO implement me
	panic("implement me")
}

func (m migrateCommand) Run() {
	//TODO implement me
	panic("implement me")
}
