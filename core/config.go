type StorageConfig struct {
	Type    string
	Options interface{}
}

type ProfileConfig struct {
	PrimaryStorage   StorageConfig
	SecondaryStorage StorageConfig
}
