package gwaka

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

var (
	dir = "./src"
)

func ParseAll() ([]WakatimeWeeklyLog, error) {
	ret := []WakatimeWeeklyLog{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return ret, err
	}
	for _, file := range files {
		tmp, err := ReadFromFile(dir + "/" + file.Name())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error:\n%s\n", err)
			continue
		}
		ret = append(ret, tmp)
	}
	return ret, nil
}

func ParseLatestWeek() (WakatimeWeeklyLog, error) {
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

	if from == to {
		return ret, errors.New("Invalid filename: " + filename)
	}

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
	if len(arr) != 2 {
		return time.Time{}, time.Time{}
	}
	layout := "20060102"
	from, err := time.Parse(layout, arr[0])
	if err != nil {
		return time.Time{}, time.Time{}
	}
	to, err := time.Parse(layout, arr[1])
	if err != nil {
		return time.Time{}, time.Time{}
	}
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
