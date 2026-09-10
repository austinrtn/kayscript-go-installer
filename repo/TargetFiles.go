package repo

import "path/filepath"

type Targets struct {
	UdevRule File
	SudoersRule File
	RootService File
	UserService File
	Kayscript   File
}

func NewTargets(workDir string) (Targets, error) {
	home, err := GetHomeDir()
	if err != nil {
		return Targets{}, err
	}

	return Targets{
		UdevRule: NewFile(
			"udev_rule",
			"/etc/udev/rules.d/99-usb-connected.rules",
			0o644,
			true,
			workDir,
		),
		SudoersRule: NewFile(
			"sudoers_rule",
			filepath.Join("/etc", "sudoers.d", "kayscript-bypass"),
			0o644,
			true,
			workDir,
		),
		RootService: NewFile(
			"root_service",
			filepath.Join("/etc", "systemd", "system", "kayscript-usb.service"),
			0o644,
			true,
			workDir,
		),
		UserService: NewFile(
			"user_service",
			filepath.Join(home, ".config", "systemd", "user", "kayscript.service"),
			0o644,
			false,
			workDir,
		),
		Kayscript: NewFile("kayscript", filepath.Join(home, "KayScript"), 0o755, true, workDir,),
	}, nil
}

func (targets Targets) All() []File {
	return []File{
		targets.SudoersRule,
		targets.RootService,
		targets.UserService,
		targets.Kayscript,
	}
}
