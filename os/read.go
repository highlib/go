package os

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

// ErrorDataConversion is a custom error for wrong input data type (in relation to variable pass per param).
var ErrorDataConversion = errors.New("error on data conversion")

// Read reads an input from terminal putting her inside variable passed per reference with an optional message (return ErrorDataConversion case error).
func Read[Type string | int | float64 | bool](variable *Type, message ...string) error {

	var (
		// Like consts.
		typeString = reflect.String.String()
		typeInt    = reflect.Int.String()
		typeFloat  = reflect.Float64.String()
		typeBool   = reflect.Bool.String()
		// Complex variables.
		reader = bufio.NewReader(os.Stdin)
		// Simple variables.
		parsed any
	)

	// Print input prompt.
	if len(message) > 0 {
		for index, phrase := range message {
			if index == (len(message) - 1) {
				print(phrase)
			} else {
				println(phrase)
			}
		}
	} else {
		print("Input: ")
	}

	// Get and validade input errors.
	answer, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("error on input data")
	}
	answer = answer[:len(answer)-1]

	// Check and assign the input inside reference variable.
	switch reflect.TypeOf(*variable).String() {
	case typeString:
		*variable, _ = any(answer).(Type)
	case typeInt:
		// "parsed" and "err" variable are local.
		parsed, err := strconv.ParseInt(answer, 10, 64)
		if err != nil {
			return fmt.Errorf("%w (%s)", ErrorDataConversion, typeInt)
		}
		*variable = any(int(parsed)).(Type)
	case typeFloat:
		if parsed, err = strconv.ParseFloat(answer, 64); err != nil {
			return fmt.Errorf("%w (%s)", ErrorDataConversion, typeFloat)
		}
		*variable = parsed.(Type)
	case typeBool:
		if parsed, err = strconv.ParseBool(answer); err != nil {
			return fmt.Errorf("%w (%s)", ErrorDataConversion, typeBool)
		}
		*variable = parsed.(Type)
	}
	return nil

}
