package enum

type GenderEnum string

const (
	Male   GenderEnum = "MALE"
	Female GenderEnum = "FEMALE"
)

func Parse(gender string) (GenderEnum, bool) {
	e := GenderEnum(gender)
	switch e {
	case Male, Female:
		return e, true
	default:
		return "", false
	}
}