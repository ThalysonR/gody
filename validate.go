package gody

import "github.com/guiferpa/gody/rule"

// DefaultValidate validates the validation subject against pre-defined rules
func DefaultValidate(validationSubject interface{}, customRules []Rule) (bool, error) {
	return RawDefaultValidate(validationSubject, DefaultTagName, customRules)
}

// Validate contains the entrypoint to validation of struct input
func Validate(validationSubject interface{}, rules []Rule) (bool, error) {
	return RawValidate(validationSubject, DefaultTagName, rules)
}

func RawDefaultValidate(validationSubject interface{}, tn string, customRules []Rule) (bool, error) {
	defaultRules := []Rule{
		rule.NotEmpty,
		rule.Required,
		rule.Enum,
		rule.Max,
		rule.Min,
		rule.MaxBound,
		rule.MinBound,
	}

	return RawValidate(validationSubject, tn, append(defaultRules, customRules...))
}

func RawValidate(validationSubject interface{}, tn string, rules []Rule) (bool, error) {
	fields, err := RawSerialize(tn, validationSubject)
	if err != nil {
		return false, err
	}

	return ValidateFields(fields, rules)
}

func ValidateFields(fields []Field, rules []Rule) (bool, error) {
	for _, field := range fields {
		for _, r := range rules {
			val, ok := field.Tags[r.Name()]
			if !ok {
				continue
			}
			if ok, err := r.Validate(field.Name, field.Value, val); err != nil {
				return ok, err
			}
		}
	}

	return true, nil
}
