// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package database

import (
	"testing"
)

func TestDatabase_decrypt(t *testing.T) {
	// setup types
	key := "C639A572E14D5075C526FDDD43E4ECF6"
	value := []byte("abc")

	encrypted, err := encrypt(key, value)
	if err != nil {
		t.Errorf("unable to encrypt value: %v", err)
	}

	// setup tests
	tests := []struct {
		failure bool
		key     string
		value   []byte
	}{
		{
			failure: false,
			key:     key,
			value:   encrypted,
		},
		{
			failure: true,
			key:     "",
			value:   encrypted,
		},
		{
			failure: true,
			key:     key,
			value:   value,
		},
	}

	// run tests
	for _, test := range tests {
		_, err := decrypt(test.key, test.value)

		if test.failure {
			if err == nil {
				t.Errorf("decrypt should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("decrypt returned err: %v", err)
		}
	}
}

func TestDatabase_encrypt(t *testing.T) {
	// setup types
	key := "C639A572E14D5075C526FDDD43E4ECF6"
	value := []byte("abc")

	// setup tests
	tests := []struct {
		failure bool
		key     string
		value   []byte
	}{
		{
			failure: false,
			key:     key,
			value:   value,
		},
		{
			failure: true,
			key:     "",
			value:   value,
		},
	}

	// run tests
	for _, test := range tests {
		_, err := encrypt(test.key, test.value)

		if test.failure {
			if err == nil {
				t.Errorf("encrypt should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("encrypt returned err: %v", err)
		}
	}
}
