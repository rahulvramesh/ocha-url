package validators

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/rahulvramesh/ocha-url/internal/app/ocha/models"
	"regexp"
)

func CreateStructLevelValidation(sl validator.StructLevel) {

	item := sl.Current().Interface().(models.CreateRequest)

	// Case 1 - URL shouldn't be null & Case 2 - URL Format Validation will be handled by Validator
	// Case 3 - if the key is not empty, check the regex and length
	myRegex, _ := regexp.Compile("^[0-9a-zA-Z_]{6}$")
	if item.Key != "" && len(item.Key) != 6 && myRegex.MatchString("item.Key") {
		// Check Length & Regex
		sl.ReportError(item.Key, "key", "Key", "", "")
	}
}
