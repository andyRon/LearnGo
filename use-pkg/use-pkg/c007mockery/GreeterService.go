package main

type GreeterService interface {
	Greet(name string) string
}
