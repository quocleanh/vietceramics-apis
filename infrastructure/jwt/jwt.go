package jwt

import (
    "errors"
    "time"

    jwt "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
    "vietceramics-apis/config"
)

// Service encapsulates JWT token generation and validation logic.
type Service struct {
    secret     []byte
    expiration time.Duration
}

// NewService creates a new JWT service using the provided configuration.
func NewService(cfg *config.Config) *Service {
    return &Service{
        secret:     []byte(cfg.JWTSecret),
        expiration: time.Duration(cfg.JWTExpirationSeconds) * time.Second,
    }
}

// Claims represents the payload stored in JWT tokens.
type Claims struct {
    UserID   uuid.UUID `json:"user_id"`
    Username string    `json:"username"`
    Role     string    `json:"role"`
    jwt.RegisteredClaims
}

// GenerateToken issues a JWT for the given user information.
func (s *Service) GenerateToken(userID uuid.UUID, username, role string) (string, error) {
    now := time.Now()
    claims := &Claims{
        UserID:   userID,
        Username: username,
        Role:     role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(now.Add(s.expiration)),
            IssuedAt:  jwt.NewNumericDate(now),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signed, err := token.SignedString(s.secret)
    if err != nil {
        return "", err
    }
    return signed, nil
}

// ValidateToken parses and validates a JWT and returns the embedded claims.
func (s *Service) ValidateToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        // Ensure the token uses the expected signing method.
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return s.secret, nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    return nil, errors.New("invalid token claims")
}
