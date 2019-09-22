package config

//CustomCfg is the object for custom config
type CustomCfg struct {
	Filename string
}

var cfgFile CustomCfg

//ParseCfg interface
func (c CustomCfg) ParseCfg(env string) {
	// workPath, err := os.Getwd()
	// if err != nil {
	// 	panic("error")
	// }
	// cfgFile = path.Join(workPath, "config", "environments", env, "custom.json")
	// fmt.Println(cfgFile)
}

func init() {
	cfgFile.Filename = "custom.json"
	Register(cfgFile)
}
