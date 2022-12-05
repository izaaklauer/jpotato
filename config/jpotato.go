package config

type Jpotato struct {
HelloWorldMessage string `hcl:"hello_world_message,attr"`

// ... your config here
}

// DefaultJpotatoConfig returns default config values
func DefaultJpotatoConfig() Jpotato {
	return Jpotato{
	HelloWorldMessage:
		"hello from the default config",
	}
}
