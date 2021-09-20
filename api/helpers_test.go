package api

import (
	// "os"
	// "path/filepath"
	// "runtime"
	// "testing"

	// "github.com/laidbackware/vmware-download-sdk/sdk"
	// "github.com/orirawlings/persistent-cookiejar"
)

// var authenticatedClient *sdk.Client

// func ensureLogin(t *testing.T) (err error) {
// 	if authenticatedClient == nil {
// 		var jar *cookiejar.Jar
// 		// Persist cookies on file system to speed up testing
// 		jar, err = cookiejar.New(&cookiejar.Options{
// 			Filename:              filepath.Join(homeDir(), ".vmware.cookies"),
// 			PersistSessionCookies: true,
// 		})
// 		if err != nil {return}
// 		user, pass := mustEnv(t, "VMW_USER"), mustEnv(t, "VMW_PASS")
// 		authenticatedClient, err = sdk.Login(user, pass, jar)
// 		if err == nil {
// 			err = jar.Save()
// 		}
// 	}
// 	return
// }

// func mustEnv(t *testing.T, k string) string {
// 	t.Helper()

// 	if v, ok := os.LookupEnv(k); ok {
// 		return v
// 	}

// 	t.Fatalf("expected environment variable %q", k)
// 	return ""
// }

// // homeDir returns the OS-specific home path as specified in the environment.
// func homeDir() string {
// 	if runtime.GOOS == "windows" {
// 		return filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"))
// 	}
// 	return os.Getenv("HOME")
// }