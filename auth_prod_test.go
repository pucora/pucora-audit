package audit

import (
	"testing"

	apikeys "github.com/pucora/pucora-apikeys/v2"
	jose "github.com/pucora/pucora-jose/v2"
	"github.com/pucora/lura/v2/config"
)

func TestParseAuthProdJWTFlags(t *testing.T) {
	cfg := &config.ServiceConfig{
		Endpoints: []*config.EndpointConfig{
			{
				ExtraConfig: config.ExtraConfig{
					jose.ValidatorNamespace: map[string]interface{}{
						"disable_jwk_security": true,
						"operation_debug":      true,
					},
				},
			},
		},
	}
	flags := parseAuthProd(cfg)
	if !flags.JWTDisableJWKSecurity || !flags.JWTOperationDebug {
		t.Fatalf("expected jwt prod flags, got %+v", flags)
	}
}

func TestParseAuthProdRevokerStandalone(t *testing.T) {
	cfg := &config.ServiceConfig{
		ExtraConfig: config.ExtraConfig{
			revokerNamespace: map[string]interface{}{
				"N": 1000,
			},
		},
	}
	if !parseAuthProd(cfg).RevokerMissingAPIKey {
		t.Fatal("expected missing revoke server api key")
	}
}

func TestParseAuthProdRevokerGatewayClient(t *testing.T) {
	cfg := &config.ServiceConfig{
		ExtraConfig: config.ExtraConfig{
			revokerNamespace: map[string]interface{}{
				"revoke_server_ping_url": "http://revoker:8080",
			},
		},
		Endpoints: []*config.EndpointConfig{{Endpoint: "/"}},
	}
	if !parseAuthProd(cfg).RevokerMissingAPIKey {
		t.Fatal("expected missing api key for gateway client")
	}
}

func TestParseAuthProdRevokerDIYGateway(t *testing.T) {
	cfg := &config.ServiceConfig{
		ExtraConfig: config.ExtraConfig{
			revokerNamespace: map[string]interface{}{
				"port":        1234,
				"token_keys":  []interface{}{"jti"},
				"revoke_server_ping_url": "",
			},
		},
		Endpoints: []*config.EndpointConfig{{Endpoint: "/"}},
	}
	if parseAuthProd(cfg).RevokerMissingAPIKey {
		t.Fatal("diy gateway bloomfilter should not require revoke_server_api_key")
	}
}

func TestParseAuthProdAPIKeys(t *testing.T) {
	cfg := &config.ServiceConfig{
		ExtraConfig: config.ExtraConfig{
			apikeys.Namespace: map[string]interface{}{
				"hash":     "plain",
				"strategy": "query_string",
				"keys":     []interface{}{map[string]interface{}{"key": "x", "roles": []interface{}{"user"}}},
			},
		},
	}
	flags := parseAuthProd(cfg)
	if !flags.APIKeysPlainHash || !flags.APIKeysQueryStrategy {
		t.Fatalf("unexpected flags: %+v", flags)
	}
}

func TestAuthProdAuditRules(t *testing.T) {
	cfg := &config.ServiceConfig{
		Endpoints: []*config.EndpointConfig{
			{
				ExtraConfig: config.ExtraConfig{
					jose.ValidatorNamespace: map[string]interface{}{
						"disable_jwk_security": true,
					},
				},
			},
		},
	}
	cfg.Normalize()
	result, err := Audit(cfg, nil, []string{SeverityHigh})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Recommendations) == 0 {
		t.Fatal("expected auth prod recommendations")
	}
	found := false
	for _, rec := range result.Recommendations {
		if rec.Rule == "1.3.1" {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected rule 1.3.1, got %+v", result.Recommendations)
	}
}
