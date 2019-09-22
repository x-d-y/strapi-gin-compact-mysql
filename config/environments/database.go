package environments

import "github.com/xdy/gin/config"

type database struct{}

func init() {
	config.Register()
}
