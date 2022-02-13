package winapi

import (
	"unsafe"

	"github.com/lxn/win"
)

func CreateRectRgnIndirect(rect win.RECT) win.HRGN {
	return win.HRGN(createRectRgnIndirect(uintptr(unsafe.Pointer(&rect.Left))))
}
