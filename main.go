package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"time"
)

var get Get

type Get struct {
	Images struct {
		Lighting struct {
			Name []string `json:"Name"`
			URL  []string `json:"URL"`
		} `json:"Lighting"`
		Propresenter struct {
			Name []string `json:"Name"`
			URL  []string `json:"URL"`
		} `json:"Propresenter"`
		Protools struct {
			Name []string `json:"Name"`
			URL  []string `json:"URL"`
		} `json:"Protools"`
		Usermac struct {
			Name []string `json:"Name"`
			URL  []string `json:"URL"`
		} `json:"Usermac"`
		Checkin struct {
			Name []string `json:"Name"`
			URL  []string `json:"URL"`
		} `json:"Checkin"`
		Smaart struct {
			Name []string `json:"Name"`
			URL  []string `json:"URL"`
		} `json:"Smaart"`
	} `json:"Images"`
}

func downloadScript(y1 []string, y2 []string) {
	start := time.Now()
	for x := range y1 {
		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		os.Chdir(user.HomeDir + "/Downloads")

		dst, err := os.Create(y1[x])
		if err != nil {
			fmt.Println(err)
		}

		resp, _ := http.Get(y2[x])

		src, err := io.Copy(dst, resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(y1[x], src, "Bytes", "Downloaded Successful!.")
	}
	elapsed := time.Since(start)
	log.Printf("Downloaded all files in %s", elapsed)
}

func main() {
	var image int
	fmt.Println("Options:\n \t1.lighting \n \t2.checkin \n \t3.protools \n \t4.propresenter \n \t5.usermac \n \t6.Smaart")
	fmt.Print("Enter the image number that you would like to download: ")
	_, err := fmt.Scan(&image)
	if err != nil {
		log.Fatal(err)
	}

	// json data
	url := "https://www.dropbox.com/s/bngfnoxe5kwv8lc/images.json?dl=1"

	dst, err := os.Create("images.json")
	if err != nil {
		fmt.Println(err)
	}

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	src, err := io.Copy(dst, res.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(src, "file created!")

	body, err := ioutil.ReadFile("./images.json")
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(body, &get)

	switch image {
	case 1:
		fmt.Println("Downloading Lighting Files...")
		downloadScript(get.Images.Lighting.Name, get.Images.Lighting.URL)
		openDropbox := exec.Command("hdiutil", "attach", "Dropbox.dmg")
		openDropbox.Run()

		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
		installDropbox.Run()

		openVista := exec.Command("hdiutil", "attach", "Vista.dmg")
		openVista.Run()

		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
		openKaseya.Run()

	case 2:
		fmt.Println("Downloading Check-In Files...")
		downloadScript(get.Images.Checkin.Name, get.Images.Checkin.URL)
		openDropbox := exec.Command("Dropbox.exe")
		openDropbox.Run()

		openF1 := exec.Command("msiexec", "/a", "F1.msi")
		openF1.Run()

		openKaseya := exec.Command("Kaseya.exe")
		openKaseya.Run()
	case 3:
		fmt.Println("Downloading Protools Files...")
		downloadScript(get.Images.Protools.Name, get.Images.Protools.URL)
		openProtools := exec.Command("open", "Protools.dmg")
		openProtools.Run()

		openDVS := exec.Command("open", "DVS.dmg")
		openDVS.Run()

		openMidi := exec.Command("open", "Midi.pkg")
		openMidi.Run()

		openDropbox := exec.Command("open", "Dropbox.dmg")
		openDropbox.Run()

		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
		installDropbox.Run()

		openPitch := exec.Command("open", "Pitchle3.pkg")
		openPitch.Run()

		openDanteController := exec.Command("open", "DanteController.dmg")
		openDanteController.Run()

		openSpotify := exec.Command("unzip", "-a", "SpotifyInstaller.zip")
		openSpotify.Run()

		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
		openKaseya.Run()
	case 4:
		fmt.Println("Downloading Propresenter Files...")
		downloadScript(get.Images.Propresenter.Name, get.Images.Propresenter.URL)
		openPropresenter := exec.Command("open", "Propresenter.dmg")
		openPropresenter.Run()

		openVideohub := exec.Command("open", "Videohub.dmg")
		openVideohub.Run()

		openVlc := exec.Command("open", "Vlc.dmg")
		openVlc.Run()

		openDropbox := exec.Command("open", "Dropbox.dmg")
		openDropbox.Run()

		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
		installDropbox.Run()

		openHipchat := exec.Command("open", "Hipchat.dmg")
		openHipchat.Run()

		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
		openKaseya.Run()
	case 5:
		fmt.Println("Downloading User Mac Files...")
		downloadScript(get.Images.Usermac.Name, get.Images.Usermac.URL)
		openOffice := exec.Command("open", "Office.pkg")
		openOffice.Run()

		deleteDock := exec.Command("defaults", "delete", "com.apple.dock", "persistent-apps")
		deleteDock.Run()

		deleteDock2 := exec.Command("killall", "Dock")
		deleteDock2.Run()

		openDropbox := exec.Command("open", "Dropbox.dmg")
		openDropbox.Run()

		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
		installDropbox.Run()

		openBria := exec.Command("open", "Bria.dmg")
		openBria.Run()

		openFirefox := exec.Command("open", "Firefox.dmg")
		openFirefox.Run()

		openChrome := exec.Command("open", "Chrome.dmg")
		openChrome.Run()

		openVidyo := exec.Command("open", "Vidyo.dmg")
		openVidyo.Run()

		openCisco := exec.Command("open", "Cisco.dmg")
		openCisco.Run()

		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
		openKaseya.Run()
	case 6:
		fmt.Println("Downloading Smaart Files...")
		downloadScript(get.Images.Smaart.Name, get.Images.Smaart.URL)
		openDropbox := exec.Command("open", "Dropbox.dmg")
		openDropbox.Run()

		installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
		installDropbox.Run()
		openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
		openKaseya.Run()
	}
}
