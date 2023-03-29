package http

import "embed"

// WebFS is the embedded filesystem containing the web folder.
//
//go:embed web/*
var WebFS embed.FS
