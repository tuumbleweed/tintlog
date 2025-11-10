package main

import (
	tl "github.com/tuumbleweed/tintlog/logger"
	"github.com/tuumbleweed/tintlog/palette"
)

func main() {
	useTid := true
	tl.InitializeConfig(&tl.Config{UseTid: &useTid, LogDir: "./log/"})
	tl.Log(tl.Info, palette.Green, "\n\n\nRegular color text %s, more regular text\nmore text: %s", "text printed with tl.Color\nstillsame tl.Color", "more color text\nnew line colored")
	tl.Log(tl.Info, palette.Red, "error: %s\n%s", "something", "went wrong")
	tl.Log(tl.Info, palette.BlueBold, "Here's config:\n'''\n%s\n'''", tl.Cfg)
	tl.Log(tl.Info, palette.RedBackground, "error: %s\n%s", "something", "went wrong")
	tl.Log(tl.Info, palette.RedBoldBackground, "error: %s\n%s", "something", "went wrong")
}
