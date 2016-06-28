package gwaka

import (
	"io/ioutil"
	"strings"
)

func Parse() (WakatimeWeeklyLog, error) {
	return ReadFromFile("./src/20160627.log")
}

func ReadFromFile(path string) (WakatimeWeeklyLog, error) {
	ret := NewWakatimeWeeklyLog()
	bf, err := ioutil.ReadFile(path)
	if err != nil {
		return ret, err
	}
	lines := strings.Split(string(bf), "\n")
	for _, l := range findProjectLog(lines) {
		ret.Projects = append(ret.Projects, WakatimeProjectActivityFromString(l))
	}
	for _, l := range findLanguageLog(lines) {
		ret.Languages = append(ret.Languages, WakatimeLanguageActivityFromString(l))
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
