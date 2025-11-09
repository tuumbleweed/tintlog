package main

import (
	"github.com/meeeraaakiii/go-tintlog"
	"github.com/meeeraaakiii/go-tintlog/color"
)

func main() {
	useTid := true
	logger.InitializeConfig(&logger.Config{UseTid: &useTid, LogDir: "./log/"})
	logger.Log(logger.Info, color.Green, "\n\n\nRegular color text %s, more regular text\nmore text: %s", "text printed with logger.Color\nstillsame logger.Color", "more color text\nnew line colored")
	logger.Log(logger.Info, color.Red, "error: %s\n%s", "something", "went wrong")
	logger.Log(logger.Info, color.BlueBold, "Here's config:\n'''\n%s\n'''", logger.Cfg)
	logger.Log(logger.Info, color.RedBackground, "error: %s\n%s", "something", "went wrong")
	logger.Log(logger.Info, color.RedBoldBackground, "error: %s\n%s", "something", "went wrong")
}
