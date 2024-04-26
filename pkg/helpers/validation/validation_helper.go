package validation

import (
	"github.com/gookit/validate"
)

func ValidationErrorsAsMap(validationErrors validate.Errors) map[string]any {
	eMap := make(map[string]any, len(validationErrors))

	for k, ve := range validationErrors {
		eMap[k] = ve.String()
	}

	return eMap
}
