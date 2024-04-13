package main

type SpreadSheet struct {
	Cells map[string]string
}

func (s *SpreadSheet) GetFormat() string {
	return "xlsx"
}
func (s *SpreadSheet) Clone() Document {
	newDoc := SpreadSheet{Cells: make(map[string]string)}
	for k, v := range s.Cells {
		newDoc.Cells[k] = v
	}
	return &newDoc
}
