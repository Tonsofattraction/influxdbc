package influxdbc

type Series struct {
	Name    string
	Columns []string
	Points  [][]string
}

func NewSeries(name string, cols ...string) *Series {
	s := new(Series)
	s.Name = name
	s.Columns = cols
	s.Points = make([][]string, 0)
	return s
}

func (s *Series) AddPoint(point ...string) {
	s.Points = append(s.Points, point)
}
