package main

type tomlConfig struct {
	Title          string
	Superuser      superuser
	Steam          steam
	DefaultRoles   map[string]role `toml:"roles"`
	ScoringSerivce scoringServiceConfig
}

type superuser struct {
	Name string
	Pass string `toml:"password"`
}

type steam struct {
	Port int
}

type role struct {
	Description string
	Permissions []string
}

type scoringServiceConfig struct {
	JettyPath string
	WarPath   string
	Port      int
	PortRange string
}
