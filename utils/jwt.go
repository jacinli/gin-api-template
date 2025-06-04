package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWTClaims JWT 声明结构
type JWTClaims struct {
	UserID  uint   `json:"user_id"`
	Phone   string `json:"phone"`
	TokenID string `json:"token_id"`
	jwt.RegisteredClaims
}

// TokenPair 令牌对
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

// GenerateTokenPair 生成访问令牌和刷新令牌
func GenerateTokenPair(userID uint, phone string) (*TokenPair, error) {
	if AppConfig == nil {
		return nil, errors.New("配置未初始化")
	}

	// 生成访问令牌
	accessTokenID := uuid.New().String()
	accessTokenExpire := time.Now().Add(time.Duration(AppConfig.JWTAccessTokenExpireMinutes) * time.Minute)

	accessClaims := &JWTClaims{
		UserID:  userID,
		Phone:   phone,
		TokenID: accessTokenID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpire),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-api-template",
			Subject:   phone,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(AppConfig.JWTSecret))
	if err != nil {
		LogError("生成访问令牌失败: " + err.Error())
		return nil, err
	}

	// 生成刷新令牌
	refreshTokenID := uuid.New().String()
	refreshTokenExpire := time.Now().Add(time.Duration(AppConfig.JWTRefreshTokenExpireHours) * time.Hour)

	refreshClaims := &JWTClaims{
		UserID:  userID,
		Phone:   phone,
		TokenID: refreshTokenID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExpire),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-api-template",
			Subject:   phone,
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(AppConfig.JWTSecret))
	if err != nil {
		LogError("生成刷新令牌失败: " + err.Error())
		return nil, err
	}

	LogInfo("成功生成令牌对，用户ID: " + string(rune(userID)))

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresAt:    accessTokenExpire.Unix(),
	}, nil
}

// ValidateToken 验证令牌
func ValidateToken(tokenString string) (*JWTClaims, error) {
	if AppConfig == nil {
		return nil, errors.New("配置未初始化")
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效的签名方法")
		}
		return []byte(AppConfig.JWTSecret), nil
	})

	if err != nil {
		LogError("令牌解析失败: " + err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("无效的令牌")
	}

	LogInfo("令牌验证成功，用户ID: " + string(rune(claims.UserID)))
	return claims, nil
}

// ExtractTokenFromHeader 从请求头中提取令牌
func ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("授权头为空")
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return "", errors.New("无效的授权头格式")
	}

	return authHeader[len(bearerPrefix):], nil
}
