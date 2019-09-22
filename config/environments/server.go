package environments

import "github.com/xdy/gin/config"

type server struct{}

func init() {
	config.Register()
}
