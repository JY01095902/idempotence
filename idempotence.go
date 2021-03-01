package idempotence

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

func (ide Idempotence) SaveIfAbsent(key, group string) error {
	ide.logger.Infow("[Idempotence Service] save idempotence key", "key", key, "group", group, "ts", GetTimestamp())

	err := ide.storage.SaveIfAbsent(key, group)
	if err != nil {
		ide.logger.Errorw("[Idempotence Service] save idempotence key failed", "key", key, "group", group, "error", err.Error(), "ts", GetTimestamp())
	}

	return err
}

func (ide Idempotence) Remove(key, group string) error {
	ide.logger.Infow("[Idempotence Service] remove idempotence key", "key", key, "group", group, "ts", GetTimestamp())

	err := ide.storage.Remove(key, group)
	if err != nil {
		ide.logger.Errorw("[Idempotence Service] remove idempotence key failed", "key", key, "group", group, "error", err.Error(), "ts", GetTimestamp())
	}

	return err
}
