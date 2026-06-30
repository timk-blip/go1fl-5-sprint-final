package actioninfo

import "fmt"

type DataParser interface {
	Parse(dataset string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, line := range dataset {
		err := dp.Parse(line)
		if err != nil {
			fmt.Println(err)
		}
		_, errA := dp.ActionInfo()
		if errA != nil {
			fmt.Println(err)
		}

	}

}
