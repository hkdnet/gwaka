package gwaka

type Gwaka struct {
}

func (g *Gwaka) Parse() WakatimeWeeklyLog {
	ret, err := Parse()
	if err != nil {
		panic(err)
	}
	return ret
}
