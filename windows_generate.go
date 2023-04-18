// Code generated by 'go generate'; DO NOT EDIT.

package winapi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modGdi32    = windows.NewLazySystemDLL("Gdi32.dll")
	modMmdevapi = windows.NewLazySystemDLL("Mmdevapi.dll")
	modWtsapi32 = windows.NewLazySystemDLL("Wtsapi32.dll")
	moduser32   = windows.NewLazySystemDLL("user32.dll")

	procCreateDIBSection            = modGdi32.NewProc("CreateDIBSection")
	procCreatePen                   = modGdi32.NewProc("CreatePen")
	procCreateRectRgnIndirect       = modGdi32.NewProc("CreateRectRgnIndirect")
	procCreateSolidBrush            = modGdi32.NewProc("CreateSolidBrush")
	procExtFloodFill                = modGdi32.NewProc("ExtFloodFill")
	procPolyDraw                    = modGdi32.NewProc("PolyDraw")
	procActivateAudioInterfaceAsync = modMmdevapi.NewProc("ActivateAudioInterfaceAsync")
	procWTSCloseServer              = modWtsapi32.NewProc("WTSCloseServer")
	procWTSEnumerateSessionsExW     = modWtsapi32.NewProc("WTSEnumerateSessionsExW")
	procWTSFreeMemoryExW            = modWtsapi32.NewProc("WTSFreeMemoryExW")
	procWTSOpenServerExW            = modWtsapi32.NewProc("WTSOpenServerExW")
	procWTSVirtualChannelClose      = modWtsapi32.NewProc("WTSVirtualChannelClose")
	procWTSVirtualChannelOpen       = modWtsapi32.NewProc("WTSVirtualChannelOpen")
	procWTSVirtualChannelOpenEx     = modWtsapi32.NewProc("WTSVirtualChannelOpenEx")
	procWTSVirtualChannelRead       = modWtsapi32.NewProc("WTSVirtualChannelRead")
	procWTSVirtualChannelWrite      = modWtsapi32.NewProc("WTSVirtualChannelWrite")
	procClipCursor                  = moduser32.NewProc("ClipCursor")
	procEnumDesktopWindows          = moduser32.NewProc("EnumDesktopWindows")
	procFillRect                    = moduser32.NewProc("FillRect")
	procFindWindowExW               = moduser32.NewProc("FindWindowExW")
	procGetClassNameW               = moduser32.NewProc("GetClassNameW")
	procGetWindowTextW              = moduser32.NewProc("GetWindowTextW")
	procInvalidateRect              = moduser32.NewProc("InvalidateRect")
	procMapVirtualKeyW              = moduser32.NewProc("MapVirtualKeyW")
	procRegisterClassExW            = moduser32.NewProc("RegisterClassExW")
	procSetLayeredWindowAttributes  = moduser32.NewProc("SetLayeredWindowAttributes")
	procSetWindowRgn                = moduser32.NewProc("SetWindowRgn")
	procSetWindowTextW              = moduser32.NewProc("SetWindowTextW")
	procShowCursor                  = moduser32.NewProc("ShowCursor")
	procUpdateLayeredWindow         = moduser32.NewProc("UpdateLayeredWindow")
)

func createDIBSection(hdc uintptr, pbmi uintptr, usage uint, ppvBits uintptr, hSection uintptr, offset uint32) (hBitMap uintptr) {
	r0, _, _ := syscall.Syscall6(procCreateDIBSection.Addr(), 6, uintptr(hdc), uintptr(pbmi), uintptr(usage), uintptr(ppvBits), uintptr(hSection), uintptr(offset))
	hBitMap = uintptr(r0)
	return
}

func createPen(iStyle int, cWidth int, color uint32) (hpen uintptr) {
	r0, _, _ := syscall.Syscall(procCreatePen.Addr(), 3, uintptr(iStyle), uintptr(cWidth), uintptr(color))
	hpen = uintptr(r0)
	return
}

func createRectRgnIndirect(rect uintptr) (rgn uintptr) {
	r0, _, _ := syscall.Syscall(procCreateRectRgnIndirect.Addr(), 1, uintptr(rect), 0, 0)
	rgn = uintptr(r0)
	return
}

func createSolidBrush(color uint32) (hbrush uintptr) {
	r0, _, _ := syscall.Syscall(procCreateSolidBrush.Addr(), 1, uintptr(color), 0, 0)
	hbrush = uintptr(r0)
	return
}

func extFloodFill(hdc uintptr, x int, y int, color uint32, opType uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procExtFloodFill.Addr(), 5, uintptr(hdc), uintptr(x), uintptr(y), uintptr(color), uintptr(opType), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func polyDraw(hdc uintptr, apt uintptr, aj uintptr, cpt int) (err error) {
	r1, _, e1 := syscall.Syscall6(procPolyDraw.Addr(), 4, uintptr(hdc), uintptr(apt), uintptr(aj), uintptr(cpt), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func activateAudioInterfaceAsync(deviceInterfacePath *uint16, riid uintptr, activationParams uintptr, completionHandler uintptr, createAsync uintptr) (hresult int32) {
	r0, _, _ := syscall.Syscall6(procActivateAudioInterfaceAsync.Addr(), 5, uintptr(unsafe.Pointer(deviceInterfacePath)), uintptr(riid), uintptr(activationParams), uintptr(completionHandler), uintptr(createAsync), 0)
	hresult = int32(r0)
	return
}

func wtsCloseServerExW(hServer uintptr) {
	syscall.Syscall(procWTSCloseServer.Addr(), 1, uintptr(hServer), 0, 0)
	return
}

func wtsEnumerateSessionsEx(hServer uintptr, pLevel *uint32, Filter uint32, ppSessionInfo uintptr, pCount *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procWTSEnumerateSessionsExW.Addr(), 5, uintptr(hServer), uintptr(unsafe.Pointer(pLevel)), uintptr(Filter), uintptr(ppSessionInfo), uintptr(unsafe.Pointer(pCount)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func wtsFreeMemoryEx(wtsTypeClass uintptr, pMemory uintptr, NumberOfEntries uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procWTSFreeMemoryExW.Addr(), 3, uintptr(wtsTypeClass), uintptr(pMemory), uintptr(NumberOfEntries))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func wtsOpenServerExW(pServerName *uint16) (handle uintptr) {
	r0, _, _ := syscall.Syscall(procWTSOpenServerExW.Addr(), 1, uintptr(unsafe.Pointer(pServerName)), 0, 0)
	handle = uintptr(r0)
	return
}

func wtsVirtualChannelClose(hChannelHandle uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procWTSVirtualChannelClose.Addr(), 1, uintptr(hChannelHandle), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func wtsVirtualChannelOpen(hServer uintptr, SessionId uint32, pVirtualName *uint16) (handle uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procWTSVirtualChannelOpen.Addr(), 3, uintptr(hServer), uintptr(SessionId), uintptr(unsafe.Pointer(pVirtualName)))
	handle = uintptr(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func wtsVirtualChannelOpenEx(SessionId uint32, pVirtualName *byte, flags uint32) (handle uintptr, err error) {
	r0, _, e1 := syscall.Syscall(procWTSVirtualChannelOpenEx.Addr(), 3, uintptr(SessionId), uintptr(unsafe.Pointer(pVirtualName)), uintptr(flags))
	handle = uintptr(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func wtsVirtualChannelRead(hChannelHandle uintptr, TimeOut uint32, Buffer uintptr, BufferSize uint32, pBytesRead *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procWTSVirtualChannelRead.Addr(), 5, uintptr(hChannelHandle), uintptr(TimeOut), uintptr(Buffer), uintptr(BufferSize), uintptr(unsafe.Pointer(pBytesRead)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func wtsVirtualChannelWrite(hChannelHandle uintptr, Buffer uintptr, Length uint32, pBytesWritten *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procWTSVirtualChannelWrite.Addr(), 4, uintptr(hChannelHandle), uintptr(Buffer), uintptr(Length), uintptr(unsafe.Pointer(pBytesWritten)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func clipCursor(rect uintptr) (ok int, err error) {
	r0, _, e1 := syscall.Syscall(procClipCursor.Addr(), 1, uintptr(rect), 0, 0)
	ok = int(r0)
	if ok == 0 {
		err = errnoErr(e1)
	}
	return
}

func enumDesktopWindows(hDesktop uintptr, lpEnumFunc uintptr, lParam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procEnumDesktopWindows.Addr(), 3, uintptr(hDesktop), uintptr(lpEnumFunc), uintptr(lParam))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func fillRect(hdc uintptr, rect uintptr, hbr uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procFillRect.Addr(), 3, uintptr(hdc), uintptr(rect), uintptr(hbr))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func findWindowEx(hwndParent uintptr, hwndChildAfter uintptr, lpszClass *uint16, lpszWindow *uint16) (hwnd uintptr) {
	r0, _, _ := syscall.Syscall6(procFindWindowExW.Addr(), 4, uintptr(hwndParent), uintptr(hwndChildAfter), uintptr(unsafe.Pointer(lpszClass)), uintptr(unsafe.Pointer(lpszWindow)), 0, 0)
	hwnd = uintptr(r0)
	return
}

func getClassName(hwnd uintptr, lpClassName uintptr, nMax int) (length int) {
	r0, _, _ := syscall.Syscall(procGetClassNameW.Addr(), 3, uintptr(hwnd), uintptr(lpClassName), uintptr(nMax))
	length = int(r0)
	return
}

func getWindowText(hwnd uintptr, lpString uintptr, nMax int) (length int) {
	r0, _, _ := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(lpString), uintptr(nMax))
	length = int(r0)
	return
}

func invalidateRect(hwnd uintptr, rect uintptr, bErase bool) (err error) {
	var _p0 uint32
	if bErase {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall(procInvalidateRect.Addr(), 3, uintptr(hwnd), uintptr(rect), uintptr(_p0))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func mapVirtualKey(uCode uint32, uMapType uint32) (code uint32) {
	r0, _, _ := syscall.Syscall(procMapVirtualKeyW.Addr(), 2, uintptr(uCode), uintptr(uMapType), 0)
	code = uint32(r0)
	return
}

func registerClassEx(windowClass uintptr) (atom uint16, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterClassExW.Addr(), 1, uintptr(windowClass), 0, 0)
	atom = uint16(r0)
	if atom == 0 {
		err = errnoErr(e1)
	}
	return
}

func setLayeredWindowAttributes(hwnd uintptr, color uint32, bAlpha byte, dwFlags uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetLayeredWindowAttributes.Addr(), 4, uintptr(hwnd), uintptr(color), uintptr(bAlpha), uintptr(dwFlags), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setWindowRgn(hwnd uintptr, hRgn uintptr, bRedraw bool) (err error) {
	var _p0 uint32
	if bRedraw {
		_p0 = 1
	}
	r1, _, e1 := syscall.Syscall(procSetWindowRgn.Addr(), 3, uintptr(hwnd), uintptr(hRgn), uintptr(_p0))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func setWindowText(hwnd uintptr, lpString *uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procSetWindowTextW.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(lpString)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func showCursor(state bool) (counter int) {
	var _p0 uint32
	if state {
		_p0 = 1
	}
	r0, _, _ := syscall.Syscall(procShowCursor.Addr(), 1, uintptr(_p0), 0, 0)
	counter = int(r0)
	return
}

func updateLayeredWindow(hwnd uintptr, hdcDst uintptr, pptDst uintptr, psize uintptr, hdcSrc uintptr, pptSrc uintptr, crKey uint32, pblend uintptr, dwFlags uint32) (ok bool) {
	r0, _, _ := syscall.Syscall9(procUpdateLayeredWindow.Addr(), 9, uintptr(hwnd), uintptr(hdcDst), uintptr(pptDst), uintptr(psize), uintptr(hdcSrc), uintptr(pptSrc), uintptr(crKey), uintptr(pblend), uintptr(dwFlags))
	ok = r0 != 0
	return
}
