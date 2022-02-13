package winapi

import (
	"github.com/lxn/win"
)

func EnumChildWindows(hwnd win.HWND) []win.HWND {
	var res = []win.HWND{}

	var chwnd win.HWND
	for chwnd = FindWindowEx(hwnd, chwnd, nil, nil); chwnd != win.HWND(uintptr(0)); chwnd = FindWindowEx(hwnd, chwnd, nil, nil) {
		res = append(res, chwnd)
	}

	return res
}

func SendMessage(hwnd win.HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {
	return win.SendMessage(hwnd, msg, wParam, lParam)
}
