package configuration

type Procedure struct {
	Id    string `yaml:"id"`
	Tasks []Task `yaml:"tasks"`
}

type Task struct {
	Id      string  `yaml:"id"`
	Shell   string  `yaml:"shell"`
	Command Command `yaml:"command"`
}

type Command struct {
	Argv []string `yaml:"argv,omitempty"`
	Exec string   `yaml:"exec,omitempty"`
}
