package zomboid

import (
	"gopkg.in/ini.v1"
)

func LoadZomboidConfig(filename string) (*ZomboidConfig, error) {
	// IgnoreInlineComment is required in order to load ini values that use a semi-colon as a delimiter
	// For example, Mods=someMod;anotherMod;yetAnotherMod;
	inidata, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, filename)
	if err != nil {
		return nil, err
	}

	var zc ZomboidConfig

	err = inidata.MapTo(&zc)
	if err != nil {
		return nil, err
	}

	return &zc, nil
}

func SaveZomboidConfig(filename string, zc *ZomboidConfig) error {
	inidata := ini.Empty()
	err := ini.ReflectFrom(inidata, zc)
	if err != nil {
		return err
	}

	return inidata.SaveTo(filename)
}
