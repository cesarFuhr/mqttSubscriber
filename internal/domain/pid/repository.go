package pid

type Repository interface {
	Store(string, PID) error
}
