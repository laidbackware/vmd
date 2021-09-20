package api

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/laidbackware/vmware-download-sdk/sdk"
	"github.com/orirawlings/persistent-cookiejar"
)

var authenticatedClient *sdk.Client

func ensureLogin() (err error) {
	if authenticatedClient == nil {
		var jar *cookiejar.Jar
		// Store cookies under the user profile
		jar, err = cookiejar.New(&cookiejar.Options{
			Filename:              filepath.Join(homeDir(), ".vmware.cookies"),
			PersistSessionCookies: true,
		})
		if err != nil {return}
		user, pass := mustEnv("VMW_USER"), mustEnv("VMW_PASS")
		fmt.Println("Logging in...")
		authenticatedClient, err = sdk.Login(user, pass, jar)
		if err == nil {
			err = jar.Save()
		}
	}
	return
}

func mustEnv(k string) string {
	if v, ok := os.LookupEnv(k); ok {
		return v
	}
	return ""
}

// homeDir returns the OS-specific home path as specified in the environment.
func homeDir() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"))
	}
	return os.Getenv("HOME")
}