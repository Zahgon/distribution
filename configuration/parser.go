package configuration

import (
	"reflect"
)

type Version string

func MajorMinorVersion(major, minor uint) Version { _ = "STUB: not implemented"; return *new(Version) }

func (version Version) major() (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func (version Version) Major() uint { _ = "STUB: not implemented"; return 0 }

func (version Version) minor() (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func (version Version) Minor() uint { _ = "STUB: not implemented"; return 0 }

type VersionedParseInfo struct {
	Version Version

	ParseAs reflect.Type

	ConversionFunc func(any) (any, error)
}

type envVar struct {
	name  string
	value string
}

type envVars []envVar

func (a envVars) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a envVars) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (a envVars) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

type Parser struct {
	prefix  string
	mapping map[Version]VersionedParseInfo
	env     envVars
}

func NewParser(prefix string, parseInfos []VersionedParseInfo) *Parser {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) Parse(in []byte, v any) error { _ = "STUB: not implemented"; return nil }

func (p *Parser) overwriteFields(v reflect.Value, fullpath string, path []string, payload string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) overwriteStruct(v reflect.Value, fullpath string, path []string, payload string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) overwriteMap(m reflect.Value, fullpath string, path []string, payload string) error {
	_ = "STUB: not implemented"
	return nil
}
