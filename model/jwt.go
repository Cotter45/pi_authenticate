package model

import (
  "fmt"
  "time"

  "github.com/golang-jwt/jwt"
)

type Claims struct {
  Email  string `json:"email"`
  UserID uint   `json:"id"`
  AppID  uint   `json:"appId"`
  ApiKey string `json:"apiKey"`
  jwt.StandardClaims
}

type JWTManager struct {
  secretKey     string
  tokenDuration time.Duration
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
  return &JWTManager{secretKey, tokenDuration}
}

func (manager *JWTManager) Generate(user *SafeUser) (string, error) {
  claims := Claims{
    StandardClaims: jwt.StandardClaims{
      ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
    },
    UserID: user.ID,
    Email:  user.Email,
    AppID:  user.AppId,
    ApiKey: user.ApiKey,
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString([]byte(manager.secretKey))
}

func (manager *JWTManager) Verify(accessToken []byte) (*Claims, error) {
  token, err := jwt.ParseWithClaims(
    string(accessToken),
    &Claims{},
    func(token *jwt.Token) (interface{}, error) {
      _, ok := token.Method.(*jwt.SigningMethodHMAC)
      if !ok {
        return nil, fmt.Errorf("unexpected token signing method")
      }

      return []byte(manager.secretKey), nil
    },
  )

  if err != nil {
    return nil, fmt.Errorf("invalid token: %w", err)
  }

  claims, ok := token.Claims.(*Claims)
  if !ok {
    return nil, fmt.Errorf("invalid token claims")
  }

  return claims, nil
}
