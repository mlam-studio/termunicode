package termunicode

import (
	"os"
	"runtime"
)

func IsUnicodeSupported() bool {
	if runtime.GOOS != "windows" {
		return os.Getenv("TERM") != "linux"
	}

	return os.Getenv("WT_SESSION") != "" ||
		os.Getenv("TERMINUS_SUBLIME") != "" ||
		os.Getenv("ConEmuTask") == "{cmd::Cmder}" ||
		os.Getenv("TERM_PROGRAM") == "Terminus-Sublime" ||
		os.Getenv("TERM_PROGRAM") == "vscode" ||
		os.Getenv("TERM") == "xterm-256color" ||
		os.Getenv("TERM") == "alacritty" ||
		os.Getenv("TERM") == "rxvt-unicode" ||
		os.Getenv("TERM") == "rxvt-unicode-256color" ||
		os.Getenv("TERMINAL_EMULATOR") == "JetBrains-JediTerm"
}
