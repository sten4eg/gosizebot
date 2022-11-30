package features

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	R = rand.New(rand.NewSource(time.Now().Unix()))

	Words = map[byte][]string{
		'M': {
			"канал", "слайс", "интерфейс",
		},
		'F': {
			"горутина", "мапа", "паника",
		},
	}

	Special1 = "Я Даниэль Подольский 😏"
	Special2 = "Я питонист..."
	Special3 = "Я смотрю Гошу Дударя..."
	Special4 = "Я смотрю Хауди Хо..."
	Special5 = "Так, коллеги. Сейчас начну банить."
)

func GoSize() string {
	ticket := R.Int() % 101

	switch ticket {
	case 100:
		return Special1
	case 99:
		return Special2
	case 98:
		return Special3
	case 97:
		return Special4
	case 96:
		return Special5
	}

	var result string

	if R.Int()%2 == 0 {
		if R.Int()%2 == 0 {
			result = "Мой " + Words['M'][R.Int()%len(Words['M'])] + " "
		} else {
			result = "Моя " + Words['F'][R.Int()%len(Words['F'])] + " "
		}
	} else {
		if R.Int()%2 == 0 {
			word := Words['M'][R.Int()%len(Words['M'])]
			word = strings.ToUpper(string([]rune(word)[0])) + string([]rune(word)[1:])
			result = word + " у меня "
		} else {
			word := Words['F'][R.Int()%len(Words['F'])]
			word = strings.ToUpper(string([]rune(word)[0])) + string([]rune(word)[1:])
			result = word + " у меня "
		}
	}

	size := R.Int63() % 30

	if size == 0 {
		result += "nil "
	} else {
		result += strconv.FormatInt(size, 10) + " "
	}

	result += "см "

	switch true {
	case size < 10:
		result += "😥"
	case size < 20:
		result += "😕"
	default:
		result += "😄"
	}

	return result
}
