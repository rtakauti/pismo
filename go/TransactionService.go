package main

func NewTransactionService(repo TransactionRepository) *TransactionServiceRepo {
	return &TransactionServiceRepo{
		Repo: repo,
	}
}

type TransactionServiceRepo struct {
	Repo TransactionRepository
}

func (s *TransactionServiceRepo) GetById(id int64) Transaction {
	return s.Repo.GetById(id)
}

func (s *TransactionServiceRepo) GetAll() []Transaction {
	return s.Repo.GetAll()
}

func (s *TransactionServiceRepo) Delete(id int64) (deleted bool) {
	return s.Repo.Delete(id)
}

func (s *TransactionServiceRepo) Insert(item Transaction) (Transaction, error) {
	return s.Repo.Insert(item)
}

func (s *TransactionServiceRepo) Update(id int64, item Transaction) (Transaction, error) {
	return s.Repo.Update(id, item)
}
