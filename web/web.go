package web

import "embed"

type Web struct {
	Statics *embed.FS
	Pages   *embed.FS
}
