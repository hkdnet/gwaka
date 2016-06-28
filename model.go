package gwaka

import (
	"fmt"
)

// WakatimeWeeklyLog represents your weekly activitiy.
type WakatimeWeeklyLog struct {
	Projects  []WakatimeProjectLog
	Languages []WakatimeLanguageLog
}

func (l *WakatimeWeeklyLog) String() string {
	buf := []byte{}
	buf = append(buf, "Projects"...)
	for _, p := range l.Projects {
		buf = append(buf, "\n   "...)
		buf = append(buf, p.String()...)
	}
	buf = append(buf, "\nLanguages"...)
	for _, l := range l.Languages {
		buf = append(buf, "\n   "...)
		buf = append(buf, l.String()...)
	}
	return string(buf)
}

// WakatimeProjectLog represents your activitiy on a project.
type WakatimeProjectLog struct {
	Name    string
	Hours   int
	Minutes int
	Seconds int
}

func (p *WakatimeProjectLog) String() string {
	return fmt.Sprintf("%20s\t%2d:%2d:%2d", p.Name, p.Hours, p.Minutes, p.Seconds)
}

// WakatimeLanguageLog represents your activitiy on a language.
type WakatimeLanguageLog struct {
	Name    string
	Hours   int
	Minutes int
	Seconds int
}

func (l *WakatimeLanguageLog) String() string {
	return fmt.Sprintf("%20s\t%2d:%2d:%2d", l.Name, l.Hours, l.Minutes, l.Seconds)
}