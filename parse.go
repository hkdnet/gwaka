package gwaka

import (
	"io/ioutil"
	"path"
	"strings"
	"time"
)

func ParseLatestWeek() (WakatimeWeeklyLog, error) {
	dir := "./src"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return WakatimeWeeklyLog{}, err
	}
	var maxMod time.Time
	var maxIdx int
	for i, f := range files {
		m := f.ModTime()
		if m.After(maxMod) {
			maxMod = m
			maxIdx = i
		}
	}
	latest := files[maxIdx]
	return ReadFromFile(dir + "/" + latest.Name())
}

func ReadFromFile(filePath string) (WakatimeWeeklyLog, error) {
	ret := NewWakatimeWeeklyLog()
	_, filename := path.Split(filePath)
	from, to := spanFromFilename(filename)

	ret.From = from
	ret.To = to
	bf, err := ioutil.ReadFile(filePath)
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

func spanFromFilename(filename string) (time.Time, time.Time) {
	rep := strings.Replace(filename, ".log", "", 1)
	arr := strings.Split(rep, "-")
	layout := "20060102"
	from, _ := time.Parse(layout, arr[0])
	to, _ := time.Parse(layout, arr[1])
	return from, to
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
