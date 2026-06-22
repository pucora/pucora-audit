package audit

import (
	"strings"

	apikeys "github.com/pucora/pucora-apikeys/v2"
	jose "github.com/pucora/pucora-jose/v2"
	"github.com/pucora/lura/v2/config"
)

const revokerNamespace = "github_com/devopsfaith/bloomfilter"

// AuthProd captures production auth misconfigurations detected during parse.
type AuthProd struct {
	JWTDisableJWKSecurity bool `json:"jwt_disable_jwk_security,omitempty"`
	JWTOperationDebug     bool `json:"jwt_operation_debug,omitempty"`
	RevokerMissingAPIKey  bool `json:"revoker_missing_api_key,omitempty"`
	APIKeysPlainHash      bool `json:"api_keys_plain_hash,omitempty"`
	APIKeysQueryStrategy  bool `json:"api_keys_query_strategy,omitempty"`
}

func parseAuthProd(cfg *config.ServiceConfig) AuthProd {
	if cfg == nil {
		return AuthProd{}
	}
	flags := AuthProd{}
	inspectRevokerConfig(cfg, &flags)
	inspectAPIKeysConfig(cfg.ExtraConfig, &flags)
	for _, ep := range cfg.Endpoints {
		if ep == nil {
			continue
		}
		inspectJWTValidator(ep.ExtraConfig, &flags)
		inspectAPIKeysConfig(ep.ExtraConfig, &flags)
	}
	return flags
}

func inspectJWTValidator(extra config.ExtraConfig, flags *AuthProd) {
	raw, ok := extra[jose.ValidatorNamespace].(map[string]interface{})
	if !ok {
		return
	}
	if asBool(raw["disable_jwk_security"]) {
		flags.JWTDisableJWKSecurity = true
	}
	if asBool(raw["operation_debug"]) {
		flags.JWTOperationDebug = true
	}
}

func inspectRevokerConfig(cfg *config.ServiceConfig, flags *AuthProd) {
	raw, ok := cfg.ExtraConfig[revokerNamespace].(map[string]interface{})
	if !ok {
		return
	}
	key, _ := raw["revoke_server_api_key"].(string)
	if strings.TrimSpace(key) != "" {
		return
	}
	pingURL, _ := raw["revoke_server_ping_url"].(string)
	if strings.TrimSpace(pingURL) != "" {
		flags.RevokerMissingAPIKey = true
		return
	}
	if len(cfg.Endpoints) == 0 {
		flags.RevokerMissingAPIKey = true
	}
}

func inspectAPIKeysConfig(extra config.ExtraConfig, flags *AuthProd) {
	raw, ok := extra[apikeys.Namespace].(map[string]interface{})
	if !ok {
		return
	}
	hash, _ := raw["hash"].(string)
	if hash == "" || hash == "plain" {
		flags.APIKeysPlainHash = true
	}
	if strategy, _ := raw["strategy"].(string); strategy == "query_string" {
		flags.APIKeysQueryStrategy = true
	}
}

func asBool(v interface{}) bool {
	b, ok := v.(bool)
	return ok && b
}
