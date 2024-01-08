package main

import (
	"fmt"
	"reflect"
)

type ResistorValues struct {
	Black  int
	Brown  int
	Red    int
	Orange int
	Yellow int
	Green  int
	Blue   int
	Violet int
	Grey   int
	White  int
}

func SetResistorValues() ResistorValues {
	return ResistorValues{
		Black:  0,
		Brown:  1,
		Red:    2,
		Orange: 3,
		Yellow: 4,
		Green:  5,
		Blue:   6,
		Violet: 7,
		Grey:   8,
		White:  9,
	}
}

func DetermineBandValues(resistorValues ResistorValues, bands []string) string {
	resistance := ""
	r := reflect.ValueOf(resistorValues)
	for idx, band := range bands {
		f := reflect.Indirect(r).FieldByName(band)
		fmt.Println(f)
		value := f.Interface()
		if idx == 1 || idx == 2 {
			resistance = resistance + value.(string)
		} else if idx == 3 {
			for i := 1; i <= value.(int); i++ {
				resistance = resistance + "0"
			}
		}
	}
	return resistance
}

func main() {
	resistorValues := SetResistorValues()

	bands := []string{"orange", "orange", "black"}
	resistanceString := DetermineBandValues(resistorValues, bands)
	fmt.Println(resistanceString)
}
