package hide

import (
	"HeadshotHider/core/util"
	"io"

	"golang.org/x/sys/windows/registry"
)

func DeleteRegistryKey(key registry.Key) error {
	// Find all the value names inside the key and delete them.
	valueNames, err := key.ReadValueNames(1000)
	if err != nil && err != io.EOF {
		util.Logger.Printf("Error reading value names from registry key: %v", err)
		return err
	}

	for _, valueName := range valueNames {
		err = key.DeleteValue(valueName)
		if err != nil && err != io.EOF {
			util.Logger.Printf("Error deleting value %s from registry key: %v", valueName, err)
			return err
		}
	}

	// Find the subkeys and delete those recursively.
	subkeyNames, err := key.ReadSubKeyNames(1000)
	if err != nil && err != io.EOF {
		util.Logger.Printf("Error reading subkey names from registry key: %v", err)
		return err
	}

	for _, subkeyName := range subkeyNames {
		subKey, err := registry.OpenKey(key, subkeyName, registry.ALL_ACCESS)
		if err != nil {
			util.Logger.Printf("Error opening subkey %s: %v", subkeyName, err)
			return err
		}

		err = DeleteRegistryKey(subKey)
		if err != nil {
			util.Logger.Printf("Error deleting subkey %s: %v", subkeyName, err)
			return err
		}

		err = subKey.Close()
		if err != nil {
			util.Logger.Printf("Error closing subkey %s: %v", subkeyName, err)
			return err
		}
	}

	return nil
}

func DeleteSingle(key registry.Key) error {
	// Find all the value names inside the key and delete them.
	valueNames, err := key.ReadValueNames(1000)
	if err != nil && err != io.EOF {
		return err
	}

	for _, valueName := range valueNames {
		err = key.DeleteValue(valueName)
		if err != nil && err != io.EOF {
			return err
		}
	}

	// Find the subkeys and delete those recursively.
	subkeyNames, err := key.ReadSubKeyNames(1000)
	if err != nil && err != io.EOF {
		return err
	}

	for _, subkeyName := range subkeyNames {
		subKey, err := registry.OpenKey(key, subkeyName, registry.ALL_ACCESS)
		if err != nil {
			return err
		}

		err = DeleteRegistryKey(subKey)
		if err != nil {
			return err
		}
	}
	return nil
}
