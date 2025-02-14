package linkedlist

type EducationLevel int

const (
	Diploma EducationLevel = iota + 1
	Bachelor
	Master
	PHD
)

func LevelToString(level EducationLevel) (education string) {
	switch level {
	case Diploma:
		return "diploma"
	case Bachelor:
		return "bachelor"
	case Master:
		return "master"
	case PHD:
		return "phd"
	default:
		panic("Unknown Education")
	}
}
