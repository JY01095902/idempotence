package idempotence

import (
	"strconv"
)

type IdempotenceGroup int

func (grp IdempotenceGroup) String() string {
	switch grp {
	case YS_SAVE_SALE_ORDER:
		return "YS_SAVE_SALE_ORDER"
	default:
		return strconv.Itoa(int(grp))
	}
}

const (
	YS_SAVE_SALE_ORDER IdempotenceGroup = iota + 1
)

type Idempotence struct {
	storage IdempotenceStorage
	logger  Logger
}

func NewIdempotence(storage IdempotenceStorage) Idempotence {
	return Idempotence{
		storage: storage,
		logger:  NewLogger(),
	}
}

func (ide Idempotence) SaveIfAbsent(group IdempotenceGroup, key string) error {
	ide.logger.Infow("[Idempotence Service] save idempotence key", "key", key, "group", group, "ts", GetTimestamp())

	err := ide.storage.SaveIfAbsent(group, key)
	if err != nil {
		ide.logger.Errorw("[Idempotence Service] save idempotence key failed", "key", key, "group", group, "error", err.Error(), "ts", GetTimestamp())
	}

	return err
}

func (ide Idempotence) Remove(group IdempotenceGroup, key string) error {
	ide.logger.Infow("[Idempotence Service] remove idempotence key", "key", key, "group", group, "ts", GetTimestamp())

	err := ide.storage.Remove(group, key)
	if err != nil {
		ide.logger.Errorw("[Idempotence Service] remove idempotence key failed", "key", key, "group", group, "error", err.Error(), "ts", GetTimestamp())
	}

	return err
}
