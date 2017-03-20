// +build darwin dragonfly freebsd linux netbsd openbsd
// +build !appengine

package millipede

import "golang.org/x/crypto/ssh/terminal"

func getSize() (int, error) {
	w, _, err := terminal.GetSize(0)
	return w, err
}
