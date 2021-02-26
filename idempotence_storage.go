package idempotence

type IdempotenceStorage interface {
	SaveIfAbsent(group IdempotenceGroup, key string) error
	Remove(group IdempotenceGroup, key string) error
}
