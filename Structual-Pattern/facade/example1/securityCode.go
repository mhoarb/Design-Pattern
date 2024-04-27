package main

import "fmt"

type SecurityCode struct {
	code int
}

func NewSwcurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (s *SecurityCode) CheckCode(code int) error {
	if s.code != code {
		return fmt.Errorf("code is not correct")
	}
	fmt.Println("security code verified")
	return nil
}
