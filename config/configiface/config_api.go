package configiface

type ConfigAPI interface {
	AuthDisabled() bool
	Username() string
	Password() string
	Port() int
	RootFolder() string
	SessionSecret() string
}
