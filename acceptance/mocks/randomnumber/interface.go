package randomnumber

//go:generate mockery --case underscore --inpackage --name Interface
type Interface interface {
	GetRandomNumber() (int, error)
}
