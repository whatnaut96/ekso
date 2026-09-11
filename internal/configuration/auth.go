package configuration

type PasswordAuthConfig struct {
	VaultKeyName string `yaml:"vault"`
}

type SSHKeyAuthConfig struct {
	Path          string `yaml:"path"`
	PassphraseEnv string `yaml:"passphrase_env"`
}
