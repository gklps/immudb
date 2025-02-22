/*
Copyright 2022 CodeNotary, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"testing"

	"github.com/codenotary/immudb/pkg/api/schema"
	"github.com/codenotary/immudb/pkg/signer"
	"github.com/gklps/immudb/embedded/store"
	"github.com/stretchr/testify/assert"
)

func TestNewStateSigner(t *testing.T) {
	s, _ := signer.NewSigner("./../../test/signer/ec3.key")
	rs := NewStateSigner(s)
	assert.IsType(t, &stateSigner{}, rs)
}

func TestStateSigner_Sign(t *testing.T) {
	s, _ := signer.NewSigner("./../../test/signer/ec3.key")
	stSigner := NewStateSigner(s)
	state := &schema.ImmutableState{}
	err := stSigner.Sign(state)
	assert.NoError(t, err)
	assert.IsType(t, &schema.ImmutableState{}, state)
}

func TestStateSigner_Err(t *testing.T) {
	s, _ := signer.NewSigner("./../../test/signer/ec3.key")
	stSigner := NewStateSigner(s)
	err := stSigner.Sign(nil)
	assert.Error(t, store.ErrIllegalArguments, err)
}
