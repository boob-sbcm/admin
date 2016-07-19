package admin

import "github.com/qor/roles"

// Config admin config struct
type Config struct {
	Name       string
	Menu       []string
	Invisible  bool
	Permission *roles.Permission
	Themes     []ThemeInterface
	PageCount  int
	Singleton  bool
}
