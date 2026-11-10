// Copyright IBM Corp. 2018, 2025
// SPDX-License-Identifier: MPL-2.0

package kerberos

import (
	"testing"

	"github.com/jcmturner/gokrb5/v8/credentials"
	"github.com/jcmturner/gokrb5/v8/iana/nametype"
	"github.com/jcmturner/gokrb5/v8/types"
)

// A FreeIPA host ticket: the PAC effective name is the FQDN, the principal is host/<fqdn>.
func TestLoginName_PACEffectiveNameVsPrincipal(t *testing.T) {
	cname := types.NewPrincipalName(nametype.KRB_NT_PRINCIPAL, "host/obs.e-bridge.internal")
	creds := credentials.NewFromPrincipalName(cname, "E-BRIDGE.INTERNAL")
	creds.SetADCredentials(credentials.ADCredentials{EffectiveName: "obs.e-bridge.internal"})

	got, err := loginName(creds, false)
	if err != nil || got != "obs.e-bridge.internal" {
		t.Fatalf("default: got %q, %v; want the PAC effective name", got, err)
	}
	got, err = loginName(creds, true)
	if err != nil || got != "host/obs.e-bridge.internal" {
		t.Fatalf("use_principal_name: got %q, %v; want the client principal", got, err)
	}
}
