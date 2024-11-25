package services

import (
	"errors"
	"fmt"

	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/entities"
	entityErr "github.com/filipeandrade6/fiap-pedeai-clientes/domain/errors"
	"github.com/filipeandrade6/fiap-pedeai-clientes/domain/ports"
	"github.com/google/uuid"
)

type Service struct {
	repo ports.Repository
}

func New(repository ports.Repository) *Service {
	return &Service{repository}
}

// TODO: complementar os erros no repositorio que nao sao ErrNotFound

func (s *Service) Create(cliente entities.Cliente) (entities.ID, error) {
	// if err := cliente.Validate(); err != nil {
	// 	return uuid.Nil, err
	// }

	c, err := s.repo.GetClienteByCPF(cliente.CPF())
	if err != nil {
		if !errors.Is(err, entityErr.ErrNotFound) {
			return uuid.Nil, err
		}
	}
	if c != nil {
		return uuid.Nil, entityErr.ErrClienteAlreadyExistsForCPF
	}

	c, err = s.repo.GetClienteByEmail(cliente.Email())
	if err != nil {
		if !errors.Is(err, entityErr.ErrNotFound) {
			return uuid.Nil, err
		}
	}
	if c != nil {
		return uuid.Nil, entityErr.ErrClienteAlreadyExistsForEmail
	}
	// else {
	// 	cliente := c
	// 	buf := new(bytes.Buffer)
	// 	if err := json.NewEncoder(buf).Encode(someData); err != nil {
	// 		res.WriteHeader(res, "whoops", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	io.Copy(res, buf) // reads from buf, writes to res
	// }

	id := entities.NewID()

	c2, err := entities.New(id, cliente.Name(), cliente.CPF(), cliente.Email(), true)
	if err != nil {
		return uuid.Nil, fmt.Errorf("creating new cliente: %s", err)
	}

	s.repo.Create(*c2)

	return id, nil
}

func (s *Service) List() ([]*entities.Cliente, error) {
	c, err := s.repo.List()
	if err != nil {
		return nil, err // TODO: alterar isso aqui
	}

	return c, nil
}

func (s *Service) GetClienteById(id entities.ID) (*entities.Cliente, error) {
	c, err := s.repo.GetClienteById(id)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) GetClienteByCPF(cpf string) (*entities.Cliente, error) {
	c, err := s.repo.GetClienteByCPF(cpf)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) GetClienteByEmail(email string) (*entities.Cliente, error) {
	c, err := s.repo.GetClienteByEmail(email)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) Update(cliente entities.Cliente) error {
	if err := cliente.Validate(); err != nil {
		return err
	}
	// TODO: vai ter que arrumar isso aqui
	if err := s.repo.Update(cliente); err != nil {
		return err
	}

	return nil
}

func (s *Service) Remove(id entities.ID) error {
	// TODO: vai ter que arrumar isso aqui
	if err := s.repo.Remove(id); err != nil {
		return err
	}

	return nil
}
