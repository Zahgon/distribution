package token

import (
	"crypto"
)

type actionSet struct {
	stringSet
}

func newActionSet(actions ...string) actionSet { _ = "STUB: not implemented"; return *new(actionSet) }

func (s actionSet) contains(action string) bool { _ = "STUB: not implemented"; return false }

func contains(ss []string, q string) bool { _ = "STUB: not implemented"; return false }

func containsAny(ss []string, q []string) bool { _ = "STUB: not implemented"; return false }

func hashAndEncode(payload string) string { _ = "STUB: not implemented"; return "" }

func GetRFC7638Thumbprint(publickey crypto.PublicKey) string { _ = "STUB: not implemented"; return "" }

func GetJWKThumbprint(publickey crypto.PublicKey) string { _ = "STUB: not implemented"; return "" }

func getJWKThumbprint(publickey crypto.PublicKey, skipED25519 bool) string {
	_ = "STUB: not implemented"
	return ""
}
