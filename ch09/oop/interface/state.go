package main

type State string

const (
	Waiting     State = "WAITING"
	MakeDone    State = "MAKE_DONE"
	PackageDone State = "PACKAGE_DONE"
	Done        State = "DONE"
)
