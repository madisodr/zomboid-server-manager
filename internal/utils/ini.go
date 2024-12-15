package utils

import (
	"errors"
	"reflect"

	"gopkg.in/ini.v1"
)

func LoadIni(filename string, v interface{}) error {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return errors.New("v must be a pointer to a struct")
	}

	// IgnoreInlineComment is required in order to load ini values that use a semi-colon as a delimiter
	// For example, Mods=someMod;anotherMod;yetAnotherMod;
	inidata, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, filename)
	if err != nil {
		return err
	}

	err = inidata.MapTo(v)
	if err != nil {
		return err
	}

	return nil
}

func SaveIni(filename string, v interface{}) error {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return errors.New("v must be a pointer to a struct")
	}

	inidata := ini.Empty()
	err := ini.ReflectFrom(inidata, v)
	if err != nil {
		return err
	}

	return inidata.SaveTo(filename)
}
