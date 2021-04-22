package fault

type Repository interface {
	InsertFault(string, Fault) error
}
