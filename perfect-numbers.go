package main

import (
	"errors"
	"fmt"
)

type Classification string

var ErrOnlyPositive error = errors.New("Value must be a positive integer")

const (
	ClassificationDeficient Classification = "Deficient"
	ClassificationPerfect   Classification = "Perfect"
	ClassificationAbundant  Classification = "Abundant"
)

func Classify(num int) Classification {

}

func Main() {
	fmt.Println(Classify(15))
}
