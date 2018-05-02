package config

import (
	"fmt"
	"strings"
)

type optionCliSerializable interface {
	toCli() (string, error)
}

func (config Config) ToArguments() ([]string, error) {
	arguments := make([]string, 0)

	for _, item := range config.options {
		option, ok := item.(optionCliSerializable)
		if !ok {
			return nil, fmt.Errorf("Unserializable option '%s': %#v", item.getName(), item)
		}

		optionValue, err := option.toCli()
		if err != nil {
			return nil, err
		}

		optionArguments := strings.Split(optionValue, " ")
		arguments = append(arguments, optionArguments...)
	}

	return arguments, nil
}
