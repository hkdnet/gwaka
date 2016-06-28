package gwaka

type Gwaka struct {
}

func (g *Gwaka) Parse() WakatimeWeeklyLog {
	ret, err := ParseLatestWeek()
	if err != nil {
		panic(err)
	}
	return ret
}
