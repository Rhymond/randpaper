package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/reujab/wallpaper"
)

const filenamePrefix = "randpaper"

func main() {
	systray.Run(onReady, func() {})
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTooltip("RandPaper")
	mChange := systray.AddMenuItem("Change Wallpaper", "Change Wallpaper")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	for {
		select {
		case <-mChange.ClickedCh:
			change()
		case <-mQuit.ClickedCh:
			systray.Quit()
		}
	}
}

func change() {
	img, _ := wallpaper.Get()

	if _, err := os.Stat(img); !os.IsNotExist(err) {
		err := os.Remove(img)
		if err != nil {
			panic(err)
		}
	}

	di, err := downloadImage("https://source.unsplash.com/featured")
	if err != nil {
		panic(err)
	}

	err = wallpaper.SetFromFile(di)
	if err != nil {
		panic(err)
	}
}

func downloadImage(url string) (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	filename := filenamePrefix + strconv.FormatInt(time.Now().UnixNano(), 10) + ".jpg"
	filename = filepath.Join(cacheDir, filename)
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return "", errors.New("non-200 status code")
	}

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return "", err
	}

	err = file.Close()
	if err != nil {
		return "", err
	}

	return filename, nil
}