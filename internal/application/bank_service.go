package application

import (
	"log"

	"github.com/SamujjalDas/grpc-go-server/internal/port"
)

type BankService struct {
	db port.BankDatabasePort
}

func NewBankService(dbPort port.BankDatabasePort) *BankService {
	return &BankService{
		db: dbPort,
	}
}

func (s *BankService) FindCurrentBalance(acc string) float64 {
	bankAccount, err := s.db.GetBankAccountByAccountNumber(acc)

	if err != nil {
		log.Println("Error finding account details", err)
	}

	return bankAccount.CurrentBalance
}
