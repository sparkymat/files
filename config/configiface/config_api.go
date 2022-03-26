package configiface

type ConfigAPI interface {
	Username() string
	Password() string
	Port() int
	RootFolder() string
	SessionSecret() string
}
