package backend

import (
	"io"

	"golang.org/x/sys/windows/registry"
)

func DeleteRegistryHistory() error {
	key, _ := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\AppCompatFlags\Compatibility Assistant\Store`, registry.ALL_ACCESS)
	DeleteSingle(key)
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
