package config

import (
	"path/filepath"
	"runtime"
)

// membuat config untuk mengakses static asset folder agar tidak erorr saat di deploy

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "..")
)