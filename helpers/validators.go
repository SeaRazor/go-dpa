package helpers

import "strconv"

func ValidateId(id string) error {
	if _, err := strconv.Atoi(id); err != nil {
		return err
	}
	return nil
}
