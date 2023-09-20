package upstruct

import (
	"errors"

	"github.com/fatih/structs"
)

func Update(target any, update any) error {
	if !structs.IsStruct(target) || !structs.IsStruct(update) {
		return errors.New("arguments must be structs")
	}

	targetFields := structs.Fields(target)
	updateFields := structs.Fields(update)

	for _, targetField := range targetFields {
		for _, updateField := range updateFields {
			if targetField.Name() != updateField.Name() {
				continue
			}

			if !updateField.IsZero() && targetField.Kind() == updateField.Kind() {
				targetField.Set(updateField.Value())
			}

			break
		}
	}

	return nil
}