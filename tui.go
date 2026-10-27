package main

import (
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	client 	*Client
	options []string
	cursor	int
}

func NewModel() Model {
	return Model{

	}
}
