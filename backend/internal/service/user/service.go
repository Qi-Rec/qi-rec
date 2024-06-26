package user

import (
	"context"
	"fmt"
	"net/mail"

	"qi-rec/internal/domain"
	"qi-rec/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserExists        = fmt.Errorf("user with such email already exists")
	ErrInvalidEmail      = fmt.Errorf("invalid email")
	ErrEmptyPassword     = fmt.Errorf("password is too short")
	ErrUserNotFound      = fmt.Errorf("user not found")
	ErrorInvalidPassword = fmt.Errorf("wrong password")
)

type Service struct {
	user    repository.UserRepository
	history repository.HistoryRepository
}

func NewService(repo repository.UserRepository, history repository.HistoryRepository) *Service {
	return &Service{user: repo, history: history}
}

func (s *Service) SignUp(ctx context.Context, email string, password string) (*domain.User, error) {
	ok, err := s.user.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to check if user exists by email: %w", err)
	}
	if ok {
		return nil, ErrUserExists
	}

	if err := validate(email, password); err != nil {
		return nil, fmt.Errorf("invalid user info: %w", err)
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user, err := s.user.CreateUser(ctx, email, hashedPassword)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *Service) SignIn(ctx context.Context, email string, password string) (*domain.User, error) {
	ok, err := s.user.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to check if user exists by email: %w", err)
	}

	if !ok {
		return nil, ErrUserNotFound
	}

	user, err := s.user.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrorInvalidPassword
	}

	return user, nil
}

func (s *Service) GetUserHistory(ctx context.Context, userID int) ([]*domain.Track, error) {
	ok, err := s.user.ExistsByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to check if user exists by id: %w", err)
	}

	if !ok {
		return nil, ErrUserNotFound
	}

	history, err := s.history.GetHistoryByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user history: %w", err)
	}

	return history, nil
}

func (s *Service) AddTrackToHistory(ctx context.Context, userID int, track *domain.Track) error {
	ok, err := s.user.ExistsByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to check if user exists by id: %w", err)
	}

	if !ok {
		return ErrUserNotFound
	}

	if err := s.history.AddTrack(ctx, userID, track); err != nil {
		return fmt.Errorf("failed to add track to history: %w", err)
	}

	return nil
}

func validate(email string, password string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return ErrInvalidEmail
	}

	if len(password) == 0 {
		return ErrEmptyPassword
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
