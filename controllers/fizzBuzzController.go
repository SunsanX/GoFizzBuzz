package controllers

import (
	"fizzbuzz/models"
	"fizzbuzz/services"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gofiber/fiber"
)

type FizzControl struct {
	Fizz  models.Fizzbuzz
	Stats []models.Stat
}

func (f *FizzControl) HandleFizzBuzz(c *fiber.Ctx) {

	if err := c.BodyParser(&f.Fizz); err != nil {
		c.SendStatus(404)
	}
	if err := validateFizz(f.Fizz); err != nil {
		c.SendString(err.Error())
		return
	}

	f.addStats()

	str := services.MakeFizzBuzz(f.Fizz)

	c.SendString(str)
}

func (f *FizzControl) GetStats(c *fiber.Ctx) {

	if len(f.Stats) == 0 {
		c.SendString("Stats aren't available !")
	}

	sort.SliceStable(f.Stats, func(i, j int) bool {
		return f.Stats[i].NbrReq > f.Stats[j].NbrReq
	})

	str := ""
	for i, j := range f.Stats {
		str = str + "\n" + strconv.Itoa(i+1) + " - with " + strconv.Itoa(j.NbrReq) + " request ---> \"" + j.Req + "\"."
	}
	c.SendString(str)
}

func (f *FizzControl) addStats() {

	str := strconv.Itoa(f.Fizz.Int1) + "/" + strconv.Itoa(f.Fizz.Int2) + "/" + strconv.Itoa(f.Fizz.Limit) + "/" + f.Fizz.Str1 + "/" + f.Fizz.Str2
	if len(f.Stats) == 0 {
		f.Stats = append(f.Stats, models.Stat{Req: str, NbrReq: 1})
		return
	}

	for i, j := range f.Stats {
		if str == j.Req {
			f.Stats[i].NbrReq = f.Stats[i].NbrReq + 1
			return
		}
	}

	f.Stats = append(f.Stats, models.Stat{Req: str, NbrReq: 1})
}

func validateFizz(fizz models.Fizzbuzz) error {

	var errTab []string

	if fizz.Int1 == 0 || fizz.Int2 == 0 {
		err1 := fmt.Errorf("Int1, Int2 Should be > 0.")
		errTab = append(errTab, err1.Error())
	}

	if fizz.Limit == 0 {
		fizz.Limit = 100
	} else if fizz.Limit > 100 {
		fizz.Limit = 100
		err2 := fmt.Errorf("Limit shouldn't exceed 100.")
		errTab = append(errTab, err2.Error())
	}

	if fizz.Str1 == "" || fizz.Str2 == "" {
		err2 := fmt.Errorf("Str1, Str2 Shouldn't be \"\".")
		errTab = append(errTab, err2.Error())
	}

	if len(errTab) == 0 {
		return nil
	}

	return fmt.Errorf(strings.Join(errTab, "\n"))
}
