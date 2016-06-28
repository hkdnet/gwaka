package gwaka

import "fmt"

type Gwaka struct {
}

func (g *Gwaka) Parse() {
	ret, err := ParseAll()
	if err != nil {
		panic(err)
	}
	for _, log := range ret {
		fmt.Println("-------------------")
		fmt.Println(log.String())
	}
}
