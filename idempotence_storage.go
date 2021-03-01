package idempotence

type IdempotenceStorage interface {
	SaveIfAbsent(key, group string) error
	Remove(key, group string) error
}
