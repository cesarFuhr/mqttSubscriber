package dtc

type Repository interface {
	InsertDTC(string, DTC) error
}
