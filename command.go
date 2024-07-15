package main

type Command struct {
	name        string
	description string
	callback    func() error
}
