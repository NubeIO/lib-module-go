package nmodule

import (
	"errors"
	"strings"
)

func ExtractRPCErrorMessage(err error) error {
	if err == nil {
		return nil
	}
	if strings.Contains(err.Error(), "desc = ") {
		parts := strings.Split(err.Error(), "desc = ")
		if len(parts) == 2 {
			return errors.New(strings.TrimSpace(parts[1]))
		}
	}
	return err
}
