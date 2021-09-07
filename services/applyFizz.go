package services

import (
	"fizzbuzz/models"
	"strconv"
)

func MakeFizzBuzz(fizz models.Fizzbuzz) (str string) {

	for i := 1; i <= fizz.Limit; i++ {
		str = conctatenateString(str, replaceCharacter(i, fizz))
	}
	return str
}

func replaceCharacter(i int, fizz models.Fizzbuzz) string {

	multipleInt1 := i % fizz.Int1
	multipleInt2 := i % fizz.Int2

	if multipleInt1 == 0 && multipleInt2 == 0 {
		return fizz.Str1 + fizz.Str2
	} else if multipleInt1 == 0 {
		return fizz.Str1
	} else if multipleInt2 == 0 {
		return fizz.Str2
	}

	return strconv.Itoa(i)
}

func conctatenateString(str string, str1 string) string {
	if str == "" {
		return str1
	}
	return str + "," + str1
}
