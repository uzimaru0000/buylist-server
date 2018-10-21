package perser

import (
	"strconv"
	"strings"
)

type ingredient map[string]amount

type amount struct {
	Amount float32
	Unit   string
}

type perser struct {
	units []string
}

type Perser interface {
	Perse(str []string) ingredient
}

func NewPerser(units []string) Perser {
	return &perser{units}
}

func (perser *perser) Perse(str []string) ingredient {
	var separater [][]string
	result := make(ingredient)

	for _, s := range str {
		if s != "" {
			separater = append(separater, strings.Split(s, ","))
		}
	}

	for _, s := range separater {
		name := s[0]
		amount := newAmount(perser.units, s[1])

		val, ok := result[name]
		if ok {
			val.Amount += amount.Amount
			result[name] = val
		} else {
			result[name] = amount
		}
	}

	return result
}

func newAmount(units []string, str string) amount {
	quantity := removeBrankets(str)
	result := amount{}

	for _, x := range units {
		if strings.HasSuffix(quantity, x) {
			num, _ := strconv.ParseFloat(strings.TrimSuffix(quantity, x), 32)
			result.Amount = float32(num)
			result.Unit = x
		}
	}

	return result
}

func removeBrankets(str string) string {
	result := ""
	for _, c := range str {
		if c != '(' {
			result += string(c)
		} else {
			break
		}
	}

	return result
}
