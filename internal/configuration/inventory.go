package configuration

type InventoryItem struct {
	Id   string     `yaml:"id"`
	Host Host       `yaml:"host"`
	Auth AuthMethod `yaml:"auth"`
}

type Host struct {
	Address string `yaml:"address"`
	User    string `yaml:"user"`
	Port    uint16 `yaml:"port"`
}

type AuthMethod struct {
	Password PasswordAuthConfig `yaml:"password,omitempty"`
	Key      SSHKeyAuthConfig   `yaml:"key,omitempty"`
}
