package test

import (
	"gin-framework/app/utils"
	"testing"
)

func TestJwt(t *testing.T) {
	var secret = "test"
	// token, _ := utils.GenerateToken(222, secret)
	// t.Log(token)
	claim, err := utils.ParseToken("", secret)
	if err != nil {
		t.Log(claim)
		t.Error("test fail...", err.Error())
	} else {
		t.Log("test log", claim)
	}

}
