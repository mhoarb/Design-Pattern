package main

type Server interface {
	handeleRequest(string, string) *Response
}
