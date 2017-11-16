package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

//CheckinFN -File Names
var CheckinFN = []string{
	"dropbox.exe",
	"F1.msi",
	"Lifekids.jpg",
	"Hostteam.jpg",
}

//CheckinURL - URL's
var CheckinURL = []string{
	"https://www.dropbox.com/s/t0rvb5zql1ojbvf/DropboxInstaller.exe?dl=1",
	"https://www.dropbox.com/s/bskplm3azy0m5uw/fellowshiponecheckin2_6.msi?dl=1",
	"https://www.dropbox.com/s/8j2e4wzck2v7hyp/_NEW_LIFEKIDS.jpg?dl=1",
	"https://www.dropbox.com/s/j0qmg6ij05ckqsm/_NEW_HOST_TEAM.jpg?dl=1",
	"https://www.dropbox.com/s/ev0r0341bmjlum1/KcsSetup.exe?dl=1",
}

//UsermacFN -  OSX File Names
var UsermacFN = []string{
	"Dropbox.dmg",
	"Office.pkg",
	"Bria.dmg",
	"Firefox.dmg",
	"Chrome.dmg",
	"Vidyo.dmg",
	"Wallpaper.png",
	"Cisco.dmg",
	"Kaseya.zip",
}

//UsermacURL - OSX URL's
var UsermacURL = []string{
	"https://www.dropbox.com/s/e1yzblqltnr7440/DropboxInstaller.dmg?dl=1",
	"https://www.dropbox.com/s/3jmqji7ck7lmb7j/Microsoft_Office_2016_15.38.17090200_Installer.pkg?dl=1",
	"https://www.dropbox.com/s/2arzb1y4osqjzrr/Bria_5_5.0.0_85993.dmg?dl=1",
	"https://www.dropbox.com/s/8poqj8zqktcfzev/Firefox%2047.0.1.dmg?dl=1",
	"https://www.dropbox.com/s/82ki7ll5jh4p1ne/googlechrome.dmg?dl=1",
	"https://www.dropbox.com/s/e0vrs43i5dfnr3w/VidyoNeoInstaller-macosx-TAG_NEO_17_1_2_2575%5Bp%3Dhttps%26h%3Dlifechurch.vidyocloud.com%26f%3DRzpJUENPOklQQ0k6TW9kOlRMUzpQQzpQdWJDOkNEUjpFUDpDUDpSUEk6QkE6TkRDOkNQUjpQUjpTUjI6U1I6VFA%3D%5D.dmg?dl=1",
	"https://www.dropbox.com/s/5dx1ine0sl1mjv7/New_Hire_Background.png?dl=1",
	"https://www.dropbox.com/s/8h46z4nrrv3p542/Anyconnect-Mac.dmg?dl=1",
	"https://www.dropbox.com/s/nczqji1g1lk1wm2/KcsSetup.zip?dl=1",
}

//LightingFN - File Names
var LightingFN = []string{
	"Dropbox.dmg",
	"Vista.dmg",
	"Wallpaper",
	"Kaseya.zip",
}

//LightingURL - URL's
var LightingURL = []string{
	"https://www.dropbox.com/s/7r4sq5dtkhgw3d8/DropboxInstaller.dmg?dl=1",
	"https://www.dropbox.com/s/b7qu0aga995nuhe/Vista_Installer-2.3.17736.dmg?dl=1",
	"https://www.dropbox.com/s/bm8i0wxedejcwvx/_Worship%20Background.jpg?dl=1",
	"https://www.dropbox.com/s/nczqji1g1lk1wm2/KcsSetup.zip?dl=1",
}

//ProtoolsFN - Filenames
var ProtoolsFN = []string{
	"Protools.dmg",
	"DVS.dmg",
	"Midi.pkg",
	"Dropbox.dmg",
	"Pitch.dmg",
	"DanteController.dmg",
	"SpotifyInstaller.zip",
	"Wallpaper.jpg",
	"Kaseya.zip",
}

//ProtoolsURL - URL's
var ProtoolsURL = []string{
	"https://www.dropbox.com/s/trzchdz0jswd9vv/Pro_Tools_12_8_Mac_97877.dmg?dl=1",
	"https://www.dropbox.com/s/jiv7z7furvt7usc/DVS-3.10.1.1_macos.dmg?dl=1",
	"https://www.dropbox.com/s/rar23a8skq9ef9n/MOTU%20MIDI%20Installer.pkg?dl=1",
	"https://www.dropbox.com/s/7r4sq5dtkhgw3d8/DropboxInstaller.dmg?dl=1",
	"https://www.dropbox.com/s/z2nbhp8jii1qkuk/PitchnTime3.0.1.pkg?dl=1",
	"https://www.dropbox.com/s/anjdwq1xt5bq0jt/DanteController-3.10.0.19_osx.dmg?dl=1",
	"https://download.scdn.co/SpotifyInstaller.zip",
	"https://www.dropbox.com/s/bm8i0wxedejcwvx/_Worship%20Background.jpg?dl=1",
	"https://www.dropbox.com/s/nczqji1g1lk1wm2/KcsSetup.zip?dl=1",
}

//PropresenterFN - Filenames
var PropresenterFN = []string{
	"Propresenter.dmg",
	"Hipchat.dmg",
	"Videohub.dmg",
	"Dropbox.dmg",
	"Vlc.dmg",
	"Wallpaper",
	"Kaseya.zip",
}

//PropresenterURL - URL's
var PropresenterURL = []string{
	"https://www.dropbox.com/s/z7pogswlir1qeob/ProPresenter6_6.2.7_b16049.dmg?dl=1",
	"https://www.dropbox.com/s/b84rh0f29f9z0by/HipChat-4.30.0-742.dmg?dl=1",
	"https://www.dropbox.com/s/drodolosuiv5jg3/Blackmagic_Videohub_6.4.dmg?dl=1",
	"https://www.dropbox.com/s/7r4sq5dtkhgw3d8/DropboxInstaller.dmg?dl=1",
	"https://www.dropbox.com/s/29txln6sxp68nez/vlc-2.2.6.dmg?dl=1",
	"https://www.dropbox.com/s/bm8i0wxedejcwvx/_Worship%20Background.jpg?dl=1",
	"https://www.dropbox.com/s/nczqji1g1lk1wm2/KcsSetup.zip?dl=1",
}

func downloadScript(y1 []string, y2 []string) {
	start := time.Now()
	for x := range y1 {
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

func commands() {
	var cmdProtools = &cobra.Command{
		Use:   "protools [download full image]",
		Short: "Protools image download",
		Long:  `protools is used for downloading and executeing protools specific commands for mac osx machines`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Downloading Protools Files...")
			downloadScript(ProtoolsFN, ProtoolsURL)
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
		},
	}

	var cmdPropresenter = &cobra.Command{
		Use:   "propresenter [download full image]",
		Short: "Propresenter image download",
		Long:  `propresenter is used for downloading and executeing propresenter specific commands for that given machine.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Downloading Propresenter Files...")
			downloadScript(PropresenterFN, PropresenterURL)
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
		},
	}

	var cmdLighting = &cobra.Command{
		Use:   "lighting [download full image]",
		Short: "Lighting image download",
		Long:  `Lighting is used for downloading and executeing lighting specific commands for that given machine.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Downloading Lighting Files...")
			downloadScript(LightingFN, LightingURL)
			openDropbox := exec.Command("hdiutil", "attach", "Dropbox.dmg")
			openDropbox.Run()

			installDropbox := exec.Command("hdiutil", "attach", "/Volumes/Dropbox Installer/Dropbox.app")
			installDropbox.Run()

			openVista := exec.Command("hdiutil", "attach", "Vista.dmg")
			openVista.Run()

			openKaseya := exec.Command("unzip", "-a", "Kaseya.zip")
			openKaseya.Run()
		},
	}

	var cmdUserMac = &cobra.Command{
		Use:   "usermac [download full image]",
		Short: "Usermac image download",
		Long:  `Usermac is used for downloading and executeing User MacOSX specific commands for that given machine.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Downloading User Mac Files...")
			downloadScript(UsermacFN, UsermacURL)
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
		},
	}

	var cmdCheckin = &cobra.Command{
		Use:   "checkin [download full image]",
		Short: "Checkin image download",
		Long:  `checkin is used for downloading and executeing checkin specific commands for that given machine.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Downloading Check-In Files...")
			downloadScript(CheckinFN, CheckinURL)
			openDropbox := exec.Command("Dropbox.exe")
			openDropbox.Run()

			openF1 := exec.Command("msiexec", "/a", "F1.msi")
			openF1.Run()

			openKaseya := exec.Command("Kaseya.exe")
			openKaseya.Run()
		},
	}

	var rootCmd = &cobra.Command{Use: "get"}
	rootCmd.AddCommand(cmdProtools, cmdPropresenter, cmdLighting, cmdUserMac, cmdCheckin)
	rootCmd.Execute()

}

// NewMainWindow - Created the Main Window for the Desktop Application.
// func NewMainWindow() {
// 	seperate := ui.NewHorizontalSeparator()
// 	title := ui.NewLabel("Choose the image you wish to download.")
// 	protools := ui.NewButton("Protools")
// 	propresenter := ui.NewButton("Propresenter")
// 	lighting := ui.NewButton("Lighting")
// 	usermac := ui.NewButton("User Mac OSX")
// 	checkin := ui.NewButton("Check-In")

// 	downloading := ui.NewLabel("Downloading... Please Wait.")

// 	mbox := ui.NewVerticalBox()

// 	mbox.SetPadded(true)
// 	mbox.Append(title, false)
// 	mbox.Append(seperate, false)
// 	mbox.Append(protools, false)
// 	mbox.Append(propresenter, false)
// 	mbox.Append(lighting, false)
// 	mbox.Append(usermac, false)
// 	mbox.Append(checkin, false)

// 	window := ui.NewWindow("Life.Church IT", 800, 800, true)
// 	window.Show()
// 	window.SetMargined(true)
// 	window.SetChild(mbox)

// 	protools.OnClicked(func(*ui.Button) {
// 		mbox.Append(downloading, false)
// 		downloadScript(ProtoolsFN, ProtoolsURL)
// 	})
// 	propresenter.OnClicked(func(*ui.Button) {
// 		mbox.Append(downloading, false)
// 		downloadScript(PropresenterFN, PropresenterURL)
// 	})
// 	lighting.OnClicked(func(*ui.Button) {
// 		mbox.Append(downloading, false)
// 		downloadScript(LightingFN, LightingURL)
// 	})
// 	checkin.OnClicked(func(*ui.Button) {
// 		mbox.Append(downloading, false)
// 		downloadScript(CheckinFN, CheckinURL)
// 	})
// 	usermac.OnClicked(func(*ui.Button) {
// 		mbox.Append(downloading, false)
// 		downloadScript(UsermacFN, UsermacURL)
// 	})

// 	window.OnClosing(func(*ui.Window) bool {
// 		ui.Quit()
// 		return true
// 	})
// 	window.Show()
// }

func main() {
	commands()

	// err := ui.Main(func() { // Start
	// 	NewMainWindow()
	// })
	// if err != nil {
	// 	panic(fmt.Errorf("error initializing UI library: %v", err))
	// }
}
