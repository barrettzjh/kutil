package resource

import (
	"errors"
	"fmt"
	"os"
)

type Resource struct {
	Limits struct {
		CPU string `json:"cpu"`
		Memory string `json:"memory"`
	} `json:"limits"`
	Requests struct {
		CPU string `json:"cpu"`
		Memory string `json:"memory"`
	} `json:"requests"`
}

var (
	ErrEmptyName     = errors.New("name cannot be empty")
	ErrEmptyLimit    = errors.New("limit cannot be empty")
	ErrEmptyLabel    = errors.New("label cannot be empty")
	ErrInvalidType   = errors.New("invalid resource type")
	ErrInvalidFormat = errors.New("invalid resource format")
)

func handleError(err error) {
	switch err {
	case ErrEmptyName, ErrEmptyLimit, ErrEmptyLabel, ErrInvalidType, ErrInvalidFormat:
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	default:
		// handle other errors
	}
}
