package tests

import (
	"testing"

	"fmt"
	"gin-api-template/utils"
)

func TestJwt(t *testing.T) {
	utils.LoadConfig()

	res, err := utils.GenerateTokenPair(1, "13800138000")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if res == nil {
		t.Errorf("Expected token pair, got nil")
	}
	fmt.Println("output------")
	fmt.Println(res.AccessToken)
	fmt.Println(res.RefreshToken)
	fmt.Println(res.ExpiresAt)
}
