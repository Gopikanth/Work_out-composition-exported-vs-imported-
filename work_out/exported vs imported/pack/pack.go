package pack

type Planet struct {
	Name                   string
	Size                   int
	Distance_between_earth float64
}

type Solar_system struct {
	All_planents []Planet
}

func (s *Solar_system) All() []Planet {
	return s.All_planents
}
func (s *Solar_system) Farest() []Planet {
	var printout []Planet
	for _, ok := range s.All_planents {
		if ok.Distance_between_earth < 10 {
			printout = append(printout, ok)
		}

	}
	return printout
}
