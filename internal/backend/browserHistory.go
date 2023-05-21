package backend

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/bosslawl/HeadshotHider/v2/internal/util"
)

/*
	Deletes all browser history
	I may change this to filter in the history db for only headshot related stuff
*/

func ChromeHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "chrome.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Google\\Chrome\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func OperaHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "opera.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Roaming\\Opera Software\\Opera Stable"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func OperaGXHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "opera.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Roaming\\Opera Software\\Opera GX Stable"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func EdgeHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "msedge.exe", "/T")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func FirefoxHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "firefox.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

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

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\BraveSoftware\\Brave-Browser\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func VivaldiHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "vivaldi.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Vivaldi\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func YandexHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "yandex.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Yandex\\YandexBrowser\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func TorHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "tor.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Tor Browser\\Browser\\TorBrowser\\Data\\Browser\\profile.default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func ChromiumHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "chromium.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Chromium\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func IridiumHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "iridium.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Iridium\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func EpicHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "epicbrowser.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Epic Privacy Browser\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func ComodoHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "comodo.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Comodo\\Dragon\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func SRWareHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "srware.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\SRWare Iron\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func TorchHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "torch.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Torch\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func DuckDuckHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "duckduckgo.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\DuckDuckGo\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func CentBrowserHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "centbrowser.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\CentBrowser\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func SlimjetHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "slimjet.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Slimjet\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func MaxthonHistory() {
	cmd := exec.Command("taskkill", "/F", "/IM", "maxthon.exe")
	cmd.Run()

	time.Sleep(5 * time.Second)

	history := UserHomeDir() + "\\AppData\\Local\\Maxthon\\User Data\\Default"

	err := os.Remove(history + "\\History")
	if err != nil {
		util.Logger.Println("Error deleting browser history:", err)
	}
}

func BrowserHistory() error {
	ChromeLocation := UserHomeDir() + "\\AppData\\Local\\Google\\Chrome\\User Data\\Default"
	if _, err := os.Stat(ChromeLocation); err == nil {
		ChromeHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	OperaLocation := UserHomeDir() + "\\AppData\\Roaming\\Opera Software\\Opera Stable"
	if _, err := os.Stat(OperaLocation); err == nil {
		OperaHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	OperaGXLocation := UserHomeDir() + "\\AppData\\Roaming\\Opera Software\\Opera GX Stable"
	if _, err := os.Stat(OperaGXLocation); err == nil {
		OperaGXHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	FirefoxLocation := UserHomeDir() + "\\AppData\\Roaming\\Mozilla\\Firefox\\Profiles\\"
	if _, err := os.Stat(FirefoxLocation); err == nil {
		FirefoxHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	EdgeLocation := UserHomeDir() + "\\AppData\\Local\\Microsoft\\Edge\\User Data\\Default"
	if _, err := os.Stat(EdgeLocation); err == nil {
		EdgeHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	BraveLocation := UserHomeDir() + "\\AppData\\Local\\BraveSoftware\\Brave-Browser\\User Data\\Default"
	if _, err := os.Stat(BraveLocation); err == nil {
		BraveHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	VivaldiLocation := UserHomeDir() + "\\AppData\\Local\\Vivaldi\\User Data\\Default"
	if _, err := os.Stat(VivaldiLocation); err == nil {
		VivaldiHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	YandexLocation := UserHomeDir() + "\\AppData\\Local\\Yandex\\YandexBrowser\\User Data\\Default"
	if _, err := os.Stat(YandexLocation); err == nil {
		YandexHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	TorLocation := UserHomeDir() + "\\AppData\\Local\\Tor Browser\\Browser\\TorBrowser\\Data\\Browser\\profile.default"
	if _, err := os.Stat(TorLocation); err == nil {
		TorHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	ChromiumLocation := UserHomeDir() + "\\AppData\\Local\\Chromium\\User Data\\Default"
	if _, err := os.Stat(ChromiumLocation); err == nil {
		ChromiumHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	IridiumLocation := UserHomeDir() + "\\AppData\\Local\\Iridium\\User Data\\Default"
	if _, err := os.Stat(IridiumLocation); err == nil {
		IridiumHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	EpicLocation := UserHomeDir() + "\\AppData\\Local\\Epic Privacy Browser\\User Data\\Default"
	if _, err := os.Stat(EpicLocation); err == nil {
		EpicHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	ComodoLocation := UserHomeDir() + "\\AppData\\Local\\Comodo\\Dragon\\User Data\\Default"
	if _, err := os.Stat(ComodoLocation); err == nil {
		ComodoHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	SRWareLocation := UserHomeDir() + "\\AppData\\Local\\SRWare Iron\\User Data\\Default"
	if _, err := os.Stat(SRWareLocation); err == nil {
		SRWareHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	TorchLocation := UserHomeDir() + "\\AppData\\Local\\Torch\\User Data\\Default"
	if _, err := os.Stat(TorchLocation); err == nil {
		TorchHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	DuckDuckLocation := UserHomeDir() + "\\AppData\\Local\\DuckDuckGo\\User Data\\Default"
	if _, err := os.Stat(DuckDuckLocation); err == nil {
		DuckDuckHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	CentBrowserLocation := UserHomeDir() + "\\AppData\\Local\\CentBrowser\\User Data\\Default"
	if _, err := os.Stat(CentBrowserLocation); err == nil {
		CentBrowserHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	SlimjetLocation := UserHomeDir() + "\\AppData\\Local\\Slimjet\\User Data\\Default"
	if _, err := os.Stat(SlimjetLocation); err == nil {
		SlimjetHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}

	MaxthonLocation := UserHomeDir() + "\\AppData\\Local\\Maxthon\\User Data\\Default"
	if _, err := os.Stat(MaxthonLocation); err == nil {
		MaxthonHistory()
	} else {
		util.Logger.Println("Error deleting browser history:", err)
	}
	return nil
}
