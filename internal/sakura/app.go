package sakura

import (
	"fmt"
	"os"

	"github.com/s-chernyavskiy/sakura/internal/sakura/config"
)

func Start() {
	opts := config.NewOpts(os.Args)
	fmt.Println(opts)
}
