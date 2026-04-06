package authService

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"libro/apiSchema/authSchema"
	"libro/apiSchema/userSchema"
	"libro/models/user"
	"libro/repositories"
	"libro/statics/configs"
	"libro/statics/customErr"
)

type Service struct {
	repos *repositories.InitialRepositories
	cfg   *configs.Config
}

func New(repos *repositories.InitialRepositories, cfg *configs.Config) *Service {
	return &Service{repos: repos, cfg: cfg}
}

func (s *Service) Register(req authSchema.RegisterRequest) (*authSchema.RegisterResponse, error) {
	if req.Email == "" || req.Password == "" || req.Name == "" {
		return nil, customErr.ErrInvalidInput
	}
	if _, err := s.repos.UserRepo.FindByEmail(req.Email); err == nil {
		return nil, customErr.ErrEmailAlreadyExists
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &user.User{Name: req.Name, Email: req.Email, PasswordHash: string(passwordHash)}
	if err := s.repos.UserRepo.Create(u); err != nil {
		return nil, err
	}
	tokens, err := s.issueTokens(u.ID, u.Email)
	if err != nil {
		return nil, err
	}
	return &authSchema.RegisterResponse{User: authSchema.AuthUserResponse{ID: u.ID, Name: u.Name, Email: u.Email}, AccessToken: tokens.AccessToken, RefreshToken: tokens.RefreshToken}, nil
}

func (s *Service) Login(req authSchema.LoginRequest) (*authSchema.LoginResponse, error) {
	u, err := s.repos.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, customErr.ErrInvalidCredentials
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)) != nil {
		return nil, customErr.ErrInvalidCredentials
	}
	tokens, err := s.issueTokens(u.ID, u.Email)
	if err != nil {
		return nil, err
	}
	resp := authSchema.LoginResponse{User: authSchema.AuthUserResponse{ID: u.ID, Name: u.Name, Email: u.Email}, AccessToken: tokens.AccessToken, RefreshToken: tokens.RefreshToken}
	return &resp, nil
}

func (s *Service) RefreshToken(token string) (*authSchema.RefreshTokenResponse, error) {
	claims, err := s.parseToken(token, s.cfg.JWTRefreshSecret)
	if err != nil {
		return nil, customErr.ErrUnauthorized
	}
	uid := uint(claims["userId"].(float64))
	stored, err := s.repos.AuthRepo.GetRefreshToken(uid)
	if err != nil || stored != token {
		return nil, customErr.ErrUnauthorized
	}
	email, _ := claims["email"].(string)
	pair, err := s.issueTokens(uid, email)
	if err != nil {
		return nil, err
	}
	return &authSchema.RefreshTokenResponse{AccessToken: pair.AccessToken, RefreshToken: pair.RefreshToken}, nil
}

func (s *Service) Logout(userID uint) error { return s.repos.AuthRepo.DeleteRefreshToken(userID) }

func (s *Service) Me(userID uint) (*authSchema.MeResponse, error) {
	u, err := s.repos.UserRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	return &authSchema.MeResponse{ID: u.ID, Name: u.Name, Email: u.Email, CreatedAt: u.CreatedAt.Format(time.RFC3339)}, nil
}

func (s *Service) GetProfile(userID uint) (*userSchema.ProfileResponse, error) {
	u, err := s.repos.UserRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	return &userSchema.ProfileResponse{ID: u.ID, Name: u.Name, Email: u.Email, CreatedAt: u.CreatedAt.Format(time.RFC3339), UpdatedAt: u.UpdatedAt.Format(time.RFC3339)}, nil
}

func (s *Service) UpdateProfile(userID uint, name string) (*userSchema.ProfileResponse, error) {
	u, err := s.repos.UserRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}
	u.Name = name
	if err := s.repos.UserRepo.Update(u); err != nil {
		return nil, err
	}
	return s.GetProfile(userID)
}

func (s *Service) ChangePassword(userID uint, currentPassword, newPassword string) error {
	u, err := s.repos.UserRepo.FindByID(userID)
	if err != nil {
		return err
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(currentPassword)) != nil {
		return customErr.ErrInvalidCredentials
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return s.repos.UserRepo.Update(u)
}

func (s *Service) issueTokens(userID uint, email string) (*struct{ AccessToken, RefreshToken string }, error) {
	accessToken, err := s.buildToken(userID, email, s.cfg.JWTSecret, time.Minute*time.Duration(s.cfg.AccessTokenTTLMinutes))
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.buildToken(userID, email, s.cfg.JWTRefreshSecret, time.Hour*time.Duration(s.cfg.RefreshTokenTTLHours))
	if err != nil {
		return nil, err
	}
	if err := s.repos.AuthRepo.StoreRefreshToken(userID, refreshToken, time.Hour*time.Duration(s.cfg.RefreshTokenTTLHours)); err != nil {
		return nil, err
	}
	return &struct{ AccessToken, RefreshToken string }{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *Service) buildToken(userID uint, email, secret string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{"userId": userID, "email": email, "exp": time.Now().Add(ttl).Unix()}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (s *Service) parseToken(tokenString, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}
	return claims, nil
}

func IsUniqueErr(err error) bool { return errors.Is(err, gorm.ErrDuplicatedKey) }
