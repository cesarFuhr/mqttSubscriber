package pid

type Repository interface {
	InsertPID(string, PID) error
}
