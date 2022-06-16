// +build darwin,arm64,!ios

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/aarch64-apple-darwin -lbls_snark_sys -ldl -lm
*/
import "C"

