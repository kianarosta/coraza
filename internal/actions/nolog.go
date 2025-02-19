// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package actions

import (
	"github.com/corazawaf/coraza/v3/experimental/plugins/plugintypes"
	"github.com/corazawaf/coraza/v3/internal/corazawaf"
)

type nologFn struct{}

func (a *nologFn) Init(r plugintypes.RuleMetadata, data string) error {
	if len(data) > 0 {
		return ErrUnexpectedArguments
	}

	r.(*corazawaf.Rule).Log = false
	r.(*corazawaf.Rule).Audit = false
	return nil
}

func (a *nologFn) Evaluate(_ plugintypes.RuleMetadata, _ plugintypes.TransactionState) {}

func (a *nologFn) Type() plugintypes.ActionType {
	return plugintypes.ActionTypeNondisruptive
}

func nolog() plugintypes.Action {
	return &nologFn{}
}

var (
	_ plugintypes.Action = &nologFn{}
	_ ruleActionWrapper  = nolog
)
