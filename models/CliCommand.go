package models

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}
