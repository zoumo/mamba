package mamba

import (
	"os"
	"strings"

	"github.com/spf13/pflag"
)

const (
	underline = "_"
)

var (
	// UnderlineReplacer replace dash of underline
	UnderlineReplacer = strings.NewReplacer("-", "_")
)

var (
	automaticEnvApplied bool
	envKeyReplacer      *strings.Replacer
	envPrefix           string
)

// Flag describes a flag interface
type Flag interface {
	// IsPersistent specify whether the flag is persistent
	IsPersistent() bool
	// GetName returns the flag's name
	GetName() string
	// ApplyTo adds the flag to a given FlagSet
	ApplyTo(*pflag.FlagSet) error
}

// AutomaticEnv has Mamba check ENV variables for all.
// keys set in config, default & flags
func AutomaticEnv() {
	automaticEnvApplied = true
}

// SetEnvKeyReplacer sets the strings.Replacer on the viper object
// Useful for mapping an environmental variable to a key that does
// not match it.
func SetEnvKeyReplacer(r *strings.Replacer) {
	envKeyReplacer = r
}

// SetEnvPrefix defines a prefix that ENVIRONMENT variables will use.
// E.g. if your prefix is "spf", the env registry will look for env
// variables that start with "SPF_". Only work for automatic env
func SetEnvPrefix(in string) {
	if in != "" {
		envPrefix = in
	}
}

func mergeWithEnvPrefix(key string) string {
	if envKeyReplacer != nil {
		key = envKeyReplacer.Replace(key)
	}

	if envPrefix != "" {
		connector := underline
		if strings.HasSuffix(envPrefix, underline) {
			connector = ""
		}
		return strings.ToUpper(envPrefix + connector + key)
	}

	return strings.ToUpper(key)
}

// getEnv tries to get envKey from env. otherwise returns defValue.
// you must convert the return value to the type you want.
//
// if env key is "", and AutomaticEnv is set, mamba will try to generate
// env key by merging name with envPrefix.
// finally, if the key is "" or key is not set in env, returns the defValue.
func getEnv(name, envKey string, defValue interface{}) interface{} {

	if envKey == "" && automaticEnvApplied {
		envKey = mergeWithEnvPrefix(name)
	}

	if envKey == "" {
		return defValue
	}

	e, ok := os.LookupEnv(envKey)
	if ok {
		return e
	}

	return defValue
}
