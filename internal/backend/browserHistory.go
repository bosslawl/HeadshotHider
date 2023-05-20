package backend

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

/*
	Deletes all browser history
	I may change this to filter in the history db for only headshot related stuff
*/

func ChromeHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "chrome.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Google\\Chrome\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func OperaHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "opera.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Roaming\\Opera Software\\Opera Stable"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func OperaGXHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "opera.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Roaming\\Opera Software\\Opera GX Stable"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func EdgeHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "msedge.exe", "/T")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func FirefoxHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "firefox.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Roaming\\Mozilla\\Firefox\\Profiles"

	filepath.Walk(history, func(path string, info os.FileInfo, err error) error {
		if info.Name() == "places.sqlite" {
			err = os.Remove(path)
			if err != nil {
				fmt.Println("Error deleting file:", err)
				return err
			}
		}
		return nil
	})
}

func BraveHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "brave.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\BraveSoftware\\Brave-Browser\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func VivaldiHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "vivaldi.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Vivaldi\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func YandexHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "yandex.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Yandex\\YandexBrowser\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func TorHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "tor.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Tor Browser\\Browser\\TorBrowser\\Data\\Browser\\profile.default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func ChromiumHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "chromium.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Chromium\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func IridiumHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "iridium.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Iridium\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func EpicHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "epicbrowser.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Epic Privacy Browser\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func ComodoHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "comodo.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Comodo\\Dragon\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func SRWareHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "srware.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\SRWare Iron\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func TorchHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "torch.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Torch\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func DuckDuckHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "duckduckgo.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\DuckDuckGo\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func CentBrowserHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "centbrowser.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\CentBrowser\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func SlimjetHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "slimjet.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Slimjet\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func MaxthonHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "maxthon.exe")
	cmd.Run()

	time.Sleep(5*time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Maxthon\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		fmt.Println(err)
	}
}

func BrowserHistory() {
	ChromeLocation := UserHomeDir() + "\\AppData\\Local\\Google\\Chrome\\User Data\\Default"
	if _, err := os.Stat(ChromeLocation); err == nil {
		ChromeHistory()
	}

	OperaLocation := UserHomeDir() + "\\AppData\\Roaming\\Opera Software\\Opera Stable"
	if _, err := os.Stat(OperaLocation); err == nil {
		OperaHistory()
	}

	OperaGXLocation := UserHomeDir() + "\\AppData\\Roaming\\Opera Software\\Opera GX Stable"
	if _, err := os.Stat(OperaGXLocation); err == nil {
		OperaGXHistory()
	}

	FirefoxLocation := UserHomeDir() + "\\AppData\\Roaming\\Mozilla\\Firefox\\Profiles\\"
	if _, err := os.Stat(FirefoxLocation); err == nil {
		FirefoxHistory()
	}

	EdgeLocation := UserHomeDir() + "\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default"
	if _, err := os.Stat(EdgeLocation); err == nil {
		EdgeHistory()
	}

	BraveLocation := UserHomeDir() + "\\AppData\\Local\\BraveSoftware\\Brave-Browser\\User Data\\Default"
	if _, err := os.Stat(BraveLocation); err == nil {
		BraveHistory()
	}

	VivaldiLocation := UserHomeDir() + "\\AppData\\Local\\Vivaldi\\User Data\\Default"
	if _, err := os.Stat(VivaldiLocation); err == nil {
		VivaldiHistory()
	}

	YandexLocation := UserHomeDir() + "\\AppData\\Local\\Yandex\\YandexBrowser\\User Data\\Default"
	if _, err := os.Stat(YandexLocation); err == nil {
		YandexHistory()
	}

	TorLocation := UserHomeDir() + "\\AppData\\Local\\Tor Browser\\Browser\\TorBrowser\\Data\\Browser\\profile.default"
	if _, err := os.Stat(TorLocation); err == nil {
		TorHistory()
	}

	ChromiumLocation := UserHomeDir() + "\\AppData\\Local\\Chromium\\User Data\\Default"
	if _, err := os.Stat(ChromiumLocation); err == nil {
		ChromiumHistory()
	}

	IridiumLocation := UserHomeDir() + "\\AppData\\Local\\Iridium\\User Data\\Default"
	if _, err := os.Stat(IridiumLocation); err == nil {
		IridiumHistory()
	}

	EpicLocation := UserHomeDir() + "\\AppData\\Local\\Epic Privacy Browser\\User Data\\Default"
	if _, err := os.Stat(EpicLocation); err == nil {
		EpicHistory()
	}

	ComodoLocation := UserHomeDir() + "\\AppData\\Local\\Comodo\\Dragon\\User Data\\Default"
	if _, err := os.Stat(ComodoLocation); err == nil {
		ComodoHistory()
	}

	SRWareLocation := UserHomeDir() + "\\AppData\\Local\\SRWare Iron\\User Data\\Default"
	if _, err := os.Stat(SRWareLocation); err == nil {
		SRWareHistory()
	}

	TorchLocation := UserHomeDir() + "\\AppData\\Local\\Torch\\User Data\\Default"
	if _, err := os.Stat(TorchLocation); err == nil {
		TorchHistory()
	}

	DuckDuckLocation := UserHomeDir() + "\\AppData\\Local\\DuckDuckGo\\User Data\\Default"
	if _, err := os.Stat(DuckDuckLocation); err == nil {
		DuckDuckHistory()
	}

	CentBrowserLocation := UserHomeDir() + "\\AppData\\Local\\CentBrowser\\User Data\\Default"
	if _, err := os.Stat(CentBrowserLocation); err == nil {
		CentBrowserHistory()
	}

	SlimjetLocation := UserHomeDir() + "\\AppData\\Local\\Slimjet\\User Data\\Default"
	if _, err := os.Stat(SlimjetLocation); err == nil {
		SlimjetHistory()
	}

	MaxthonLocation := UserHomeDir() + "\\AppData\\Local\\Maxthon\\User Data\\Default"
	if _, err := os.Stat(MaxthonLocation); err == nil {
		MaxthonHistory()
	}
}