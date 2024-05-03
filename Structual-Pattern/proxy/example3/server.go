package main

type Server interface {
	handleRequest(string, string) *Responses
}
