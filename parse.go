package gwaka

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func Parse() (WakatimeWeeklyLog, error) {
	return ReadFromFile("./src/20160627.log")
}

func ReadFromFile(path string) (WakatimeWeeklyLog, error) {
	ret := WakatimeWeeklyLog{Projects: []WakatimeProjectLog{}, Languages: []WakatimeLanguageLog{}}
	bf, err := ioutil.ReadFile(path)
	if err != nil {
		return ret, err
	}
	lines := strings.Split(string(bf), "\n")
	pl := findProjectLog(lines)
	ll := findLanguageLog(lines)
	hourReg := regexp.MustCompile(`(\d+) hr`)
	minReg := regexp.MustCompile(`(\d+) min`)
	secReg := regexp.MustCompile(`(\d+) sec`)
	for _, l := range pl {
		arr := strings.Split(l, "\t")
		p := WakatimeProjectLog{}
		p.Name = arr[0]
		m := hourReg.FindAllStringSubmatch(l, -1)
		if len(m) != 0 {
			i, err := strconv.Atoi(m[0][1])
			if err == nil {
				p.Hours = i
			}
		}
		m = minReg.FindAllStringSubmatch(l, -1)
		if len(m) != 0 {
			i, err := strconv.Atoi(m[0][1])
			if err == nil {
				p.Minutes = i
			}
		}

		m = secReg.FindAllStringSubmatch(l, -1)
		if len(m) != 0 {
			i, err := strconv.Atoi(m[0][1])
			if err == nil {
				p.Seconds = i
			}
		}
		ret.Projects = append(ret.Projects, p)
	}

	for _, l := range ll {
		arr := strings.Split(l, "\t")
		la := WakatimeLanguageLog{}
		la.Name = arr[0]
		m := hourReg.FindAllStringSubmatch(l, -1)
		if len(m) != 0 {
			i, err := strconv.Atoi(m[0][1])
			if err == nil {
				la.Hours = i
			}
		}
		m = minReg.FindAllStringSubmatch(l, -1)
		if len(m) != 0 {
			i, err := strconv.Atoi(m[0][1])
			if err == nil {
				la.Minutes = i
			}
		}

		m = secReg.FindAllStringSubmatch(l, -1)
		if len(m) != 0 {
			i, err := strconv.Atoi(m[0][1])
			if err == nil {
				la.Seconds = i
			}
		}
		ret.Languages = append(ret.Languages, la)
	}
	return ret, nil
}

func findProjectLog(lines []string) []string {
	f := false
	ret := []string{}
	for _, line := range lines {
		if f == true && line != "" {
			ret = append(ret, line)
		}
		if strings.Contains(line, "Projects:") {
			f = true
		}
		if strings.Contains(line, "Languages") {
			break
		}
	}
	return ret
}

func findLanguageLog(lines []string) []string {
	f := false
	ret := []string{}
	for _, line := range lines {
		if f == true && line != "" {
			ret = append(ret, line)
		}
		if strings.Contains(line, "Languages:") {
			f = true
		}
	}
	return ret
}
