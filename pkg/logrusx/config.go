package logrusx

type Config struct {
	Dir       string `mapstructure:"dir"`
	FileName  string `mapstructure:"file_name"`
	MaxSize   int    `mapstructure:"max_size"`
	LocalTime bool   `mapstructure:"local_time"`
	Compress  bool   `mapstructure:"compress"`
}
