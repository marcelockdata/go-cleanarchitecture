package repository_test

import (
	"cleanarchitecture/internal/domain"
	"cleanarchitecture/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByUsername(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)

	t.Run("Usuário encontrado", func(t *testing.T) {
		// Configurando o mock para retornar um usuário válido
		expectedUser := &domain.User{
			ID:       1,
			Username: "testuser",
			Password: "hashedpassword",
		}
		mockRepo.On("FindByUsername", "testuser").Return(expectedUser, nil)

		// Chamando o método do mock
		user, err := mockRepo.FindByUsername("testuser")

		// Validações
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser, user)

		// Verificando se o método foi chamado com os argumentos corretos
		mockRepo.AssertCalled(t, "FindByUsername", "testuser")
	})

}

func TestFindByUsername_ValidCredentials(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)

	// Configurando o mock para retornar um usuário válido
	expectedUser := &domain.User{
		ID:       1,
		Username: "testuser",
		Password: "hashedpassword",
	}
	mockRepo.On("FindByUsername", "testuser").Return(expectedUser, nil)

	// Chamando o método do mock
	user, err := mockRepo.FindByUsername("testuser")

	// Validações
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)

	// Verificando se o método foi chamado com os argumentos corretos
	mockRepo.AssertCalled(t, "FindByUsername", "testuser")
}

func TestFindByUsername_InvalidCredentials(t *testing.T) {

	mockRepo := new(repository.MockUserRepository)

	// Configurando o mock para retornar erro de credenciais inválidas
	mockRepo.On("FindByUsername", "unknownuser").Return(nil, domain.ErrInvalidCredentials)

	// Chamando o método do mock
	user, err := mockRepo.FindByUsername("unknownuser")

	// Validações
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, domain.ErrInvalidCredentials, err)

	// Verificando se o método foi chamado com os argumentos corretos
	mockRepo.AssertCalled(t, "FindByUsername", "unknownuser")

}
