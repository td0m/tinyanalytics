package model

import "testing"

func TestUser_Validate(t *testing.T) {
	securePassword := "SomeV4lidPassword420"
	sampleEmail := "admin@example.com"
	type fields struct {
		Email string
		Pass  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"fails on empty emails", fields{"", securePassword}, true},
		{"fails on incomplete emails", fields{"dom", securePassword}, true},
		{"fails on emails with only domain", fields{"@example.com", securePassword}, true},
		{"works with valid emails", fields{sampleEmail, securePassword}, false},
		{"fails on password too short", fields{sampleEmail, "short"}, true},
		{"fails on password too short", fields{sampleEmail, "1234567"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Email: tt.fields.Email,
				Pass:  tt.fields.Pass,
			}
			if err := u.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("User.Validate(%s) error = %v, wantErr %v", u.Email, err, tt.wantErr)
			}
		})
	}
}
