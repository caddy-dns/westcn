package westcn

import (
	"fmt"
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/westcn"
)

func TestUnmarshalCaddyFile(t *testing.T) {
	tests := []string{
		`westcn {
			username theusername
			api_password itsaapipassword
		}`}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			// given
			dispenser := caddyfile.NewTestDispenser(tc)
			p := Provider{&westcn.Provider{}}
			// when
			err := p.UnmarshalCaddyfile(dispenser)
			// then
			if err != nil {
				t.Errorf("UnmarshalCaddyfile failed with %v", err)
				return
			}

			expectedUsername := "theusername"
			actualUsername := p.Provider.Username
			if expectedUsername != actualUsername {
				t.Errorf("Expected Username to be '%s' but got '%s'", expectedUsername, actualUsername)
			}

			expectedAPIPassword := "itsaapipassword"
			actualAPIPassword := p.Provider.APIPassword
			if expectedAPIPassword != actualAPIPassword {
				t.Errorf("Expected APIPassword to be '%s' but got '%s'", expectedAPIPassword, actualAPIPassword)
			}
		})
	}
}
