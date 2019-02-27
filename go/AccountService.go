package main

func NewAccountService(repo AccountRepository) *AccountServiceRepo {
	return &AccountServiceRepo{
		Repo: repo,
	}
}

type AccountServiceRepo struct {
	Repo AccountRepository
}

func (s *AccountServiceRepo) GetById(id int) Account {
	return s.Repo.GetById(id)
}

func (s *AccountServiceRepo) GetAll() []Account {
	return s.Repo.GetAll()
}

func (s *AccountServiceRepo) Delete(id int) (deleted bool) {
	return s.Repo.Delete(id)
}

func (s *AccountServiceRepo) Insert(item Account) (Account, error) {
	return s.Repo.Insert(item)
}

func (s *AccountServiceRepo) Update(id int, item Account) (Account, error) {
	return s.Repo.Update(id, item)
}

func (s *AccountServiceRepo) Patch(id int, item Account) (Account, error) {
	return s.Repo.Patch(id, item)
}
