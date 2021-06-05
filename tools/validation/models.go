package validation

type FieldResult int

const (
	Incorrect FieldResult = 0
	Short     FieldResult = 1
	Long      FieldResult = 2
	Valid     FieldResult = 3
	Busy      FieldResult = 4
)
