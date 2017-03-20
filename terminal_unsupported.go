// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd appengine

package millipede

import "fmt"

func getSize() (int, error) {
	return 0, fmt.Errorf("getSize unsupported")
}
