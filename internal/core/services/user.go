package services

import (
	"context"
	"github.com/lucasvavon/slipx-api/internal/core/domain"
	"github.com/lucasvavon/slipx-api/internal/core/ports"
	"github.com/lucasvavon/slipx-api/internal/core/utils"
)

/**
 * UserService implements ports.UserService interface
 * and provides an access to the user repository
 * and cache service
 */

type UserService struct {
	repo  ports.UserRepository
	cache ports.CacheRepository
}

// NewUserService creates a new user service instance
func NewUserService(repo ports.UserRepository, cache ports.CacheRepository) *UserService {
	return &UserService{
		repo,
		cache,
	}
}

// Register creates a new user
func (us *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	user, err = us.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	cacheKey := utils.GenerateCacheKey("user", user.ID)
	userSerialized, err := utils.Serialize(user)
	if err != nil {
		return nil, err
	}

	err = us.cache.Set(ctx, cacheKey, userSerialized, 0)
	if err != nil {
		return nil, err
	}

	err = us.cache.DeleteByPrefix(ctx, "users:*")
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser gets a user by ID
func (us *UserService) GetUser(ctx context.Context, id uint) (*domain.User, error) {
	var user *domain.User

	cacheKey := utils.GenerateCacheKey("user", id)
	cachedUser, err := us.cache.Get(ctx, cacheKey)
	if err == nil {
		err := utils.Deserialize(cachedUser, &user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	user, err = us.repo.GetUserByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	userSerialized, err := utils.Serialize(user)
	if err != nil {
		return nil, err
	}

	err = us.cache.Set(ctx, cacheKey, userSerialized, 0)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUsers lists all users
func (us *UserService) GetUsers(ctx context.Context, skip, limit uint) ([]domain.User, error) {
	var users []domain.User

	params := utils.GenerateCacheKeyParams(skip, limit)
	cacheKey := utils.GenerateCacheKey("users", params)

	cachedUsers, err := us.cache.Get(ctx, cacheKey)
	if err == nil {
		err := utils.Deserialize(cachedUsers, &users)
		if err != nil {
			return nil, err
		}
		return users, nil
	}

	users, err = us.repo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	usersSerialized, err := utils.Serialize(users)
	if err != nil {
		return nil, err
	}

	err = us.cache.Set(ctx, cacheKey, usersSerialized, 0)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser updates a user's name, email, and password
func (us *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	existingUser, err := us.repo.GetUserByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	emptyData := user.Name == "" &&
		user.Email == "" &&
		user.Password == ""
	sameData := existingUser.Name == user.Name &&
		existingUser.Email == user.Email
	if emptyData || sameData {
		return nil, err
	}

	var hashedPassword string

	if user.Password != "" {
		hashedPassword, err = utils.HashPassword(user.Password)
		if err != nil {
			return nil, err
		}
	}

	user.Password = hashedPassword

	_, err = us.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	cacheKey := utils.GenerateCacheKey("user", user.ID)

	err = us.cache.Delete(ctx, cacheKey)
	if err != nil {
		return nil, err
	}

	userSerialized, err := utils.Serialize(user)
	if err != nil {
		return nil, err
	}

	err = us.cache.Set(ctx, cacheKey, userSerialized, 0)
	if err != nil {
		return nil, err
	}

	err = us.cache.DeleteByPrefix(ctx, "users:*")
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func (us *UserService) DeleteUser(ctx context.Context, id uint) error {
	_, err := us.repo.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	cacheKey := utils.GenerateCacheKey("user", id)

	err = us.cache.Delete(ctx, cacheKey)
	if err != nil {
		return err
	}

	err = us.cache.DeleteByPrefix(ctx, "users:*")
	if err != nil {
		return err
	}

	return us.repo.DeleteUser(ctx, id)
}
