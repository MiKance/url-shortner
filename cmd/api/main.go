package main

import (
	"github.com/mikance/url-shortener/internal/config"
)

func main() {
	_ = config.MustLoadEnv()
}
