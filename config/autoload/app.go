package autoload

type App struct {
	Mode       string   `mapstructure:"mode" json:"mode" yaml:"mode"`
	Address    string   `mapstructure:"address" json:"address" yaml:"address"`
	Port       int      `mapstructure:"port" json:"port" yaml:"port"`
	BaseURL    []string `mapstructure:"baseURL" json:"baseURL" yaml:"baseURL"`
	AppVersion string   `mapstructure:"appVersion" json:"appVersion" yaml:"appVersion"`
	PlayAesKey string   `mapstructure:"play_aes_key" json:"play_aes_key" yaml:"play_aes_key"`
	PlayAesIv  string   `mapstructure:"play_aes_iv" json:"play_aes_iv" yaml:"play_aes_iv"`
}
