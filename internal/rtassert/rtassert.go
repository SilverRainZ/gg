// Copyright 2025 Bytedance Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package rtassert provides runtime assertion.
package rtassert

import (
	"fmt"

	"github.com/bytedance/gg/internal/constraints"
)

func MustNotNeg[T constraints.Number](n T) {
	if n < 0 {
		panic(fmt.Errorf("number must not be negative: %v", n))
	}
}

func ErrMustNil(err error) {
	if err != nil {
		panic(fmt.Errorf("unexpected error: %s", err))
	}
}
