package service

import (
	"example.com/bank-account/domain"
	"example.com/bank-account/repository"
)

func FetchAccount(id string) domain.Account {
	return repository.Find(id)
}
