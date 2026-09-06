package token

type stringSet map[string]struct{}

func newStringSet(keys ...string) stringSet { _ = "STUB: not implemented"; return *new(stringSet) }

func (ss stringSet) add(keys ...string) { _ = "STUB: not implemented"; return }

func (ss stringSet) contains(key string) bool { _ = "STUB: not implemented"; return false }

func (ss stringSet) keys() []string { _ = "STUB: not implemented"; return nil }
