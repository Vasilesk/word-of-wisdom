package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	envConfigFile = "CONFIG"
)

func NewFromEnv[T any]() (T, error) {
	var def T

	path := os.Getenv(envConfigFile)
	if path == "" {
		return def, fmt.Errorf("env %q should be set", envConfigFile)
	}

	return NewFromFile[T](path)
}

func NewFromFile[T any](path string) (T, error) {
	var def T

	const fileModeDefault = 0o644

	//nolint:nosnakecase
	file, err := os.OpenFile(path, os.O_RDONLY, fileModeDefault)
	if err != nil {
		return def, fmt.Errorf("opening file: %w", err)
	}

	defer file.Close()

	return New[T](file)
}

func New[T any](r io.Reader) (T, error) {
	var res T

	err := yaml.NewDecoder(r).Decode(&res)
	if err != nil {
		return res, fmt.Errorf("decoding yaml: %w", err)
	}

	fromenv(&res)

	return res, nil
}
