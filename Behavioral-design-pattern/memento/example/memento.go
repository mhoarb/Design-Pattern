package main

type Memento struct {
	state string
}

func (m *Memento) GetSavedData() string {
	return m.state
}
