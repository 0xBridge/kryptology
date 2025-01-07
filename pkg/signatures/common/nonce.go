//
// Copyright Coinbase, Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"github.com/0xBridge/kryptology/pkg/core/curves"
)

// Nonce is used for zero-knowledge proofs to prevent replay attacks
// and prove freshness
type Nonce = curves.Scalar
