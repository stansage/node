package config

import (
	"bytes"
	"fmt"
)

type optionStringSerializable interface {
	toFile() (string, error)
}

func (config Config) ToString() (string, error) {
	var output bytes.Buffer

	for _, item := range config.options {
		option, ok := item.(optionStringSerializable)
		if !ok {
			return "", fmt.Errorf("Unserializable option '%s': %#v", item.getName(), item)
		}

		optionValue, err := option.toFile()
		if err != nil {
			return "", err
		}
		fmt.Fprintln(&output, optionValue)
	}

	return string(output.Bytes()), nil
}
