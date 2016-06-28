package gwaka

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	hourReg = regexp.MustCompile(`(\d+) hr`)
	minReg  = regexp.MustCompile(`(\d+) min`)
	secReg  = regexp.MustCompile(`(\d+) sec`)
)

// WakatimeWeeklyLog represents your weekly activitiy.
type WakatimeWeeklyLog struct {
	Projects  []WakatimeProjectActivity
	Languages []WakatimeLanguageActivity
}

// NewWakatimeWeeklyLog is a constuctor.
func NewWakatimeWeeklyLog() WakatimeWeeklyLog {
	return WakatimeWeeklyLog{Projects: []WakatimeProjectActivity{}, Languages: []WakatimeLanguageActivity{}}
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

// WakatimeActivity is a base activitiy.
type WakatimeActivity struct {
	Name    string
	Hours   int
	Minutes int
	Seconds int
}

func (a *WakatimeActivity) String() string {
	return fmt.Sprintf("%20s\t%2d:%2d:%2d", a.Name, a.Hours, a.Minutes, a.Seconds)
}

// WakatimeProjectActivity represents your activitiy on a project.
type WakatimeProjectActivity struct {
	*WakatimeActivity
}

func WakatimeActivityFromString(l string) WakatimeActivity {
	ret := WakatimeActivity{}
	arr := strings.Split(l, "\t")
	ret.Name = arr[0]
	m := hourReg.FindAllStringSubmatch(l, -1)
	if len(m) != 0 {
		i, err := strconv.Atoi(m[0][1])
		if err == nil {
			ret.Hours = i
		}
	}
	m = minReg.FindAllStringSubmatch(l, -1)
	if len(m) != 0 {
		i, err := strconv.Atoi(m[0][1])
		if err == nil {
			ret.Minutes = i
		}
	}
	m = secReg.FindAllStringSubmatch(l, -1)
	if len(m) != 0 {
		i, err := strconv.Atoi(m[0][1])
		if err == nil {
			ret.Seconds = i
		}
	}
	return ret
}

func WakatimeProjectActivityFromString(l string) WakatimeProjectActivity {
	a := WakatimeActivityFromString(l)
	return WakatimeProjectActivity{WakatimeActivity: &a}
}

// WakatimeLanguageActivity represents your activitiy on a language.
type WakatimeLanguageActivity struct {
	*WakatimeActivity
}

func WakatimeLanguageActivityFromString(l string) WakatimeLanguageActivity {
	a := WakatimeActivityFromString(l)
	return WakatimeLanguageActivity{WakatimeActivity: &a}
}

func (l *WakatimeLanguageActivity) String() string {
	return fmt.Sprintf("%20s\t%2d:%2d:%2d", l.Name, l.Hours, l.Minutes, l.Seconds)
}
