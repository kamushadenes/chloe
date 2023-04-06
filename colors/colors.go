package colors

import "github.com/fatih/color"

var Yellow = color.New(color.FgYellow).SprintFunc()
var Red = color.New(color.FgRed).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Cyan = color.New(color.FgCyan).SprintFunc()

var BoldYellow = color.New(color.FgYellow).Add(color.Bold).SprintFunc()
var BoldRed = color.New(color.FgRed).Add(color.Bold).SprintFunc()
var BoldGreen = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
var BoldCyan = color.New(color.FgCyan).Add(color.Bold).SprintFunc()
var BoldWhite = color.New(color.FgWhite).Add(color.Bold).SprintFunc()
var BoldPurple = color.New(color.FgMagenta).Add(color.Bold).SprintFunc()

var BgBlue = color.New(color.BgBlue).Add(color.FgWhite).Add(color.Bold).SprintFunc()
var BgRed = color.New(color.BgRed).Add(color.FgWhite).Add(color.Bold).SprintFunc()
var BgGreen = color.New(color.BgGreen).Add(color.FgWhite).Add(color.Bold).SprintFunc()
var BgYellow = color.New(color.BgYellow).Add(color.FgWhite).Add(color.Bold).SprintFunc()
var BgBlack = color.New(color.BgBlack).Add(color.FgWhite).Add(color.Bold).SprintFunc()

var Bold = color.New(color.Bold).SprintFunc()
var Italic = color.New(color.Italic).SprintFunc()
