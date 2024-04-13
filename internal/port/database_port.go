package port

import (
	db "github.com/SamujjalDas/grpc-go-server/internal/adapter/database"
	"github.com/google/uuid"
)

type DummyDatabasePort interface {
	Save(data *db.DummyOrm) (uuid.UUID, error)
	GetByUuid(uuid *uuid.UUID) (db.DummyOrm, error)
}

type BankDatabasePort interface {
	GetBankAccountByAccountNumber(acc string) (db.BankAccountOrm, error)
}
