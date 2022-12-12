package utils

import (
	"hole/src/pkg/utils"
	"testing"
)

const password = "123456"

const hashPassword = "$2a$10$9oEPoU.mIwXZWJ3a2OCYzeGB2Wk8/L97UDg0yo.u2EBTF3zo/0uze"

func TestGeneratePassword(t *testing.T) {
	password, err := utils.GeneratePassword(password)
	if err != nil {
		t.Error(err)
	}

	t.Logf("password: %s", password)
}

func TestVerifyPassword(t *testing.T) {
	type args struct {
		password     string
		hashPassword string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "PASS",
			args: struct {
				password     string
				hashPassword string
			}{password: password, hashPassword: hashPassword},
			wantErr: false,
		},
		{
			name: "ERROR",
			args: struct {
				password     string
				hashPassword string
			}{password: password, hashPassword: hashPassword[1:]},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := utils.VerifyPassword(tt.args.password, tt.args.hashPassword); (err != nil) != tt.wantErr {
				t.Errorf("VerifyPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
