package database

import "log"

func (a *DatabaseAdapter) GetBankAccountByAccountNumber(acc string) (BankAccountOrm, error) {
	var bankAccountOrm BankAccountOrm

	if err := a.db.First(&bankAccountOrm, "account_number = ?", acc).Error; err != nil {
		log.Printf("Can't find bank account number %v : %v\n", acc, err)
		return bankAccountOrm, err
	}

	return bankAccountOrm, nil
}
