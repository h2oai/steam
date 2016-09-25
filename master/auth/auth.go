/*
  Copyright (C) 2016 H2O.ai, Inc. <http://h2o.ai/>

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU Affero General Public License as
  published by the Free Software Foundation, either version 3 of the
  License, or (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package auth

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	_ "net/http/pprof"
	"regexp"
)

const SystemIdentityName = "system"

var usernameRegexp = regexp.MustCompile(`^\S+$`) // no whitespace
var passwordRegexp = regexp.MustCompile(`^\S+$`) // non-empty

func ValidateUsername(name string) error {
	if name == SystemIdentityName {
		return fmt.Errorf("\"%s\" is a reserved username", SystemIdentityName)
	}
	if !usernameRegexp.MatchString(name) {
		return fmt.Errorf("Username cannot contain whitespace characters")
	}
	if len(name) < 3 {
		return fmt.Errorf("Username must be at least 3 characters long")
	}
	return nil
}

func ValidatePassword(password string) error {
	if !passwordRegexp.MatchString(password) {
		return fmt.Errorf("Password cannot contain whitespace characters")
	}
	if len(password) < 8 {
		return fmt.Errorf("Password must be at least 8 characters long")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Password encryption failed: %s", err)
	}
	return string(hash), nil
}

func VerifyPassword(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}
