package adminweb

import "embed"

//go:embed templates/*.html static/*
var assetsFS embed.FS
