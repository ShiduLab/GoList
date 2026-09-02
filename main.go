package main

import (
	_ "embed"
	"encoding/base64"
	"encoding/binary"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"
)

const (
	appTitle = "GoList!"

	WS_OVERLAPPED        = 0x00000000
	WS_CAPTION           = 0x00C00000
	WS_SYSMENU           = 0x00080000
	WS_THICKFRAME        = 0x00040000
	WS_MINIMIZEBOX       = 0x00020000
	WS_MAXIMIZEBOX       = 0x00010000
	WS_VISIBLE           = 0x10000000
	WS_CHILD             = 0x40000000
	WS_BORDER            = 0x00800000
	WS_VSCROLL           = 0x00200000
	WS_HSCROLL           = 0x00100000
	WS_TABSTOP           = 0x00010000
	ES_AUTOHSCROLL       = 0x0080
	BS_PUSHBUTTON        = 0x00000000
	BS_AUTOCHECKBOX      = 0x00000003
	CW_USEDEFAULT        = ^uintptr(0x7fffffff)
	SW_HIDE              = 0
	SW_SHOW              = 5
	SS_CENTER            = 0x00000001
	SS_RIGHT             = 0x00000002
	WM_CREATE            = 0x0001
	WM_DESTROY           = 0x0002
	WM_SIZE              = 0x0005
	WM_COMMAND           = 0x0111
	WM_CLOSE             = 0x0010
	WM_CONTEXTMENU       = 0x007B
	WM_RBUTTONUP         = 0x0205
	WM_NOTIFY            = 0x004E
	WM_DROPFILES         = 0x0233
	WM_SETFONT           = 0x0030
	WM_SETICON           = 0x0080
	BM_GETCHECK          = 0x00F0
	BM_SETCHECK          = 0x00F1
	BST_CHECKED          = 1
	COLOR_WINDOW         = 5
	IDC_ARROW            = 32512
	DEFAULT_GUI_FONT     = 17
	MB_OK                = 0x00000000
	MB_ICONINFORMATION   = 0x00000040
	MB_ICONERROR         = 0x00000010
	BIF_RETURNONLYFSDIRS = 0x0001
	BIF_NEWDIALOGSTYLE   = 0x0040
	OFN_OVERWRITEPROMPT  = 0x00000002
	OFN_PATHMUSTEXIST    = 0x00000800
	GMEM_MOVEABLE        = 0x0002
	CF_UNICODETEXT       = 13
	MF_STRING            = 0x00000000
	MF_CHECKED           = 0x00000008
	MF_POPUP             = 0x00000010
	MF_SEPARATOR         = 0x00000800
	TPM_RIGHTBUTTON      = 0x0002
	TPM_RETURNCMD        = 0x0100
	GWLP_WNDPROC         = -4
	ICON_SMALL           = 0
	ICON_BIG             = 1
	LR_DEFAULTCOLOR      = 0x0000

	LVS_REPORT            = 0x0001
	LVS_SHOWSELALWAYS     = 0x0008
	LVS_EX_GRIDLINES      = 0x00000001
	LVS_EX_HEADERDRAGDROP = 0x00000010
	LVS_EX_FULLROWSELECT  = 0x00000020
	LVS_EX_DOUBLEBUFFER   = 0x00010000

	LVM_FIRST                    = 0x1000
	LVM_DELETEALLITEMS           = LVM_FIRST + 9
	LVM_GETCOLUMNWIDTH           = LVM_FIRST + 29
	LVM_GETHEADER                = LVM_FIRST + 31
	LVM_SETCOLUMNWIDTH           = LVM_FIRST + 30
	LVM_SETEXTENDEDLISTVIEWSTYLE = LVM_FIRST + 54
	LVM_SETCOLUMNORDERARRAY      = LVM_FIRST + 58
	LVM_GETCOLUMNORDERARRAY      = LVM_FIRST + 59
	LVM_SETITEMW                 = LVM_FIRST + 76
	LVM_INSERTITEMW              = LVM_FIRST + 77
	LVM_INSERTCOLUMNW            = LVM_FIRST + 97

	LVIF_TEXT       = 0x0001
	LVCF_FMT        = 0x0001
	LVCF_WIDTH      = 0x0002
	LVCF_TEXT       = 0x0004
	LVCF_SUBITEM    = 0x0008
	LVCFMT_LEFT     = 0x0000
	LVCFMT_RIGHT    = 0x0001
	LVN_COLUMNCLICK = -108

	ICC_LISTVIEW_CLASSES = 0x00000001
)

const (
	IDPath           = 101
	IDBrowse         = 102
	IDList           = 103
	IDCopy           = 104
	IDSave           = 105
	IDRecursive      = 106
	IDAddContext     = 107
	IDRemoveContext  = 108
	IDDeleteList     = 109
	IDTable          = 900
	IDStatus         = 901
	IDHint           = 902
	IDBrand          = 903
	IDColumnBase     = 1200
	IDHTMLColumnBase = 3000
	IDHTMLAll        = 3100
	IDHTMLNone       = 3101
	IDHTMLOK         = 3102
	IDHTMLCancel     = 3103
	IDExportHTML     = 3200
	IDExportTXT      = 3201
	IDExportTSV      = 3202
	IDExportCSV      = 3203
	IDExportJSON     = 3204
)

type WNDCLASSEX struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     uintptr
	HIcon         uintptr
	HCursor       uintptr
	HbrBackground uintptr
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       uintptr
}

type MSG struct {
	Hwnd     uintptr
	Message  uint32
	WParam   uintptr
	LParam   uintptr
	Time     uint32
	Pt       POINT
	LPrivate uint32
}

type POINT struct{ X, Y int32 }
type RECT struct{ Left, Top, Right, Bottom int32 }

type BROWSEINFO struct {
	HwndOwner      uintptr
	PidlRoot       uintptr
	PszDisplayName *uint16
	LpszTitle      *uint16
	UlFlags        uint32
	Lpfn           uintptr
	LParam         uintptr
	IImage         int32
}

type OPENFILENAME struct {
	LStructSize       uint32
	HwndOwner         uintptr
	HInstance         uintptr
	LpstrFilter       *uint16
	LpstrCustomFilter *uint16
	NMaxCustFilter    uint32
	NFilterIndex      uint32
	LpstrFile         *uint16
	NMaxFile          uint32
	LpstrFileTitle    *uint16
	NMaxFileTitle     uint32
	LpstrInitialDir   *uint16
	LpstrTitle        *uint16
	Flags             uint32
	NFileOffset       uint16
	NFileExtension    uint16
	LpstrDefExt       *uint16
	LCustData         uintptr
	LpfnHook          uintptr
	LpTemplateName    *uint16
	PvReserved        uintptr
	DwReserved        uint32
	FlagsEx           uint32
}

type INITCOMMONCONTROLSEX struct {
	DwSize uint32
	DwICC  uint32
}

type LVCOLUMN struct {
	Mask       uint32
	Fmt        int32
	Cx         int32
	PszText    *uint16
	CchTextMax int32
	ISubItem   int32
	IImage     int32
	IOrder     int32
}

type LVITEM struct {
	Mask       uint32
	IItem      int32
	ISubItem   int32
	State      uint32
	StateMask  uint32
	PszText    *uint16
	CchTextMax int32
	IImage     int32
	LParam     uintptr
}

type NMHDR struct {
	HwndFrom uintptr
	IdFrom   uintptr
	Code     uint32
}

type NMLISTVIEW struct {
	Hdr       NMHDR
	IItem     int32
	ISubItem  int32
	UNewState uint32
	UOldState uint32
	UChanged  uint32
	PtAction  POINT
	LParam    uintptr
}

type Entry struct {
	Name            string `json:"name"`
	RelativePath    string `json:"relative_path"`
	Kind            string `json:"type"`
	Size            int64  `json:"size_bytes"`
	Title           string `json:"title,omitempty"`
	Artist          string `json:"artist,omitempty"`
	Album           string `json:"album,omitempty"`
	DurationSeconds int64  `json:"duration_seconds,omitempty"`
	SampledBy       string `json:"sampled_by,omitempty"`
	Year            string `json:"year,omitempty"`
	Track           string `json:"track,omitempty"`
	Genre           string `json:"genre,omitempty"`
	Comment         string `json:"comment,omitempty"`
	AlbumArtist     string `json:"album_artist,omitempty"`
	Composer        string `json:"composer,omitempty"`
	DiscNumber      string `json:"disc_number,omitempty"`
	Publisher       string `json:"publisher,omitempty"`
	Copyright       string `json:"copyright,omitempty"`
	ISRC            string `json:"isrc,omitempty"`
	TagVersion      string `json:"id3_version,omitempty"`
	ExtraID3        string `json:"extra_id3,omitempty"`
	Created         string `json:"created"`
	Modified        string `json:"modified"`
}

type TableLayout struct {
	Widths  []int32 `json:"widths"`
	Order   []int32 `json:"order"`
	Visible []bool  `json:"visible,omitempty"`
}

type ColumnDef struct {
	Title          string
	Width          int32
	Align          int32
	DefaultVisible bool
}

var columns = []ColumnDef{
	// Vista di default: attributi universali del filesystem.
	// Le colonne specialistiche (ID3/audio) restano disponibili dal menu
	// contestuale delle intestazioni e vengono memorizzate se abilitate.
	{"Nome", 220, LVCFMT_LEFT, true},
	{"Titolo", 220, LVCFMT_LEFT, false},
	{"Artista", 160, LVCFMT_LEFT, false},
	{"Album", 180, LVCFMT_LEFT, false},
	{"Durata", 75, LVCFMT_RIGHT, false},
	{"Campionato da", 260, LVCFMT_LEFT, false},
	{"Anno", 70, LVCFMT_LEFT, false},
	{"Traccia", 70, LVCFMT_LEFT, false},
	{"Genere", 120, LVCFMT_LEFT, false},
	{"Commento", 260, LVCFMT_LEFT, false},
	{"Artista album", 160, LVCFMT_LEFT, false},
	{"Compositore", 160, LVCFMT_LEFT, false},
	{"Numero disco", 90, LVCFMT_LEFT, false},
	{"Editore", 160, LVCFMT_LEFT, false},
	{"Copyright", 180, LVCFMT_LEFT, false},
	{"ISRC", 120, LVCFMT_LEFT, false},
	{"ID3", 105, LVCFMT_LEFT, false},
	{"Extra ID3", 280, LVCFMT_LEFT, false},
	{"Tipo", 70, LVCFMT_LEFT, true},
	{"Dimensione", 105, LVCFMT_RIGHT, true},
	{"Creato", 145, LVCFMT_LEFT, true},
	{"Modificato", 145, LVCFMT_LEFT, true},
	{"Percorso", 330, LVCFMT_LEFT, true},
}

type AudioMeta struct {
	Title, Artist, Album, Year, Track, Genre, Comment string
	AlbumArtist, Composer, DiscNumber, Publisher      string
	Copyright, ISRC, TagVersion, ExtraID3             string
	Duration                                          int64
}

//go:embed Botolo.ico
var botoloICO []byte

//go:embed Botolo.png
var botoloPNG []byte

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	advapi32 = syscall.NewLazyDLL("advapi32.dll")
	shell32  = syscall.NewLazyDLL("shell32.dll")
	ole32    = syscall.NewLazyDLL("ole32.dll")
	comdlg32 = syscall.NewLazyDLL("comdlg32.dll")
	comctl32 = syscall.NewLazyDLL("comctl32.dll")
	gdi32    = syscall.NewLazyDLL("gdi32.dll")

	pRegisterClassExW         = user32.NewProc("RegisterClassExW")
	pCreateWindowExW          = user32.NewProc("CreateWindowExW")
	pDefWindowProcW           = user32.NewProc("DefWindowProcW")
	pShowWindow               = user32.NewProc("ShowWindow")
	pUpdateWindow             = user32.NewProc("UpdateWindow")
	pGetMessageW              = user32.NewProc("GetMessageW")
	pTranslateMessage         = user32.NewProc("TranslateMessage")
	pDispatchMessageW         = user32.NewProc("DispatchMessageW")
	pPostQuitMessage          = user32.NewProc("PostQuitMessage")
	pSetWindowTextW           = user32.NewProc("SetWindowTextW")
	pGetWindowTextLengthW     = user32.NewProc("GetWindowTextLengthW")
	pGetWindowTextW           = user32.NewProc("GetWindowTextW")
	pSendMessageW             = user32.NewProc("SendMessageW")
	pMessageBoxW              = user32.NewProc("MessageBoxW")
	pMoveWindow               = user32.NewProc("MoveWindow")
	pGetClientRect            = user32.NewProc("GetClientRect")
	pLoadCursorW              = user32.NewProc("LoadCursorW")
	pOpenClipboard            = user32.NewProc("OpenClipboard")
	pCloseClipboard           = user32.NewProc("CloseClipboard")
	pEmptyClipboard           = user32.NewProc("EmptyClipboard")
	pSetClipboardData         = user32.NewProc("SetClipboardData")
	pCreatePopupMenu          = user32.NewProc("CreatePopupMenu")
	pAppendMenuW              = user32.NewProc("AppendMenuW")
	pTrackPopupMenu           = user32.NewProc("TrackPopupMenu")
	pDestroyMenu              = user32.NewProc("DestroyMenu")
	pGetCursorPos             = user32.NewProc("GetCursorPos")
	pSetWindowLongPtrW        = user32.NewProc("SetWindowLongPtrW")
	pCallWindowProcW          = user32.NewProc("CallWindowProcW")
	pCreateIconFromResourceEx = user32.NewProc("CreateIconFromResourceEx")
	pDestroyIcon              = user32.NewProc("DestroyIcon")
	pDestroyWindow            = user32.NewProc("DestroyWindow")
	pEnableWindow             = user32.NewProc("EnableWindow")
	pSetForegroundWindow      = user32.NewProc("SetForegroundWindow")

	pGetModuleHandleW = kernel32.NewProc("GetModuleHandleW")
	pGlobalAlloc      = kernel32.NewProc("GlobalAlloc")
	pGlobalLock       = kernel32.NewProc("GlobalLock")
	pGlobalUnlock     = kernel32.NewProc("GlobalUnlock")

	pRegCreateKeyExW = advapi32.NewProc("RegCreateKeyExW")
	pRegSetValueExW  = advapi32.NewProc("RegSetValueExW")
	pRegDeleteTreeW  = advapi32.NewProc("RegDeleteTreeW")
	pRegCloseKey     = advapi32.NewProc("RegCloseKey")

	pDragAcceptFiles                         = shell32.NewProc("DragAcceptFiles")
	pDragQueryFileW                          = shell32.NewProc("DragQueryFileW")
	pDragFinish                              = shell32.NewProc("DragFinish")
	pSHBrowseForFolderW                      = shell32.NewProc("SHBrowseForFolderW")
	pSHGetPathFromIDListW                    = shell32.NewProc("SHGetPathFromIDListW")
	pSetCurrentProcessExplicitAppUserModelID = shell32.NewProc("SetCurrentProcessExplicitAppUserModelID")

	pCoInitializeEx = ole32.NewProc("CoInitializeEx")
	pCoTaskMemFree  = ole32.NewProc("CoTaskMemFree")

	pGetSaveFileNameW     = comdlg32.NewProc("GetSaveFileNameW")
	pInitCommonControlsEx = comctl32.NewProc("InitCommonControlsEx")
	pGetStockObject       = gdi32.NewProc("GetStockObject")

	contextMenuBusy                                                                           bool
	hwndMain, hwndPath, hwndTable, hwndHeader, hwndRecursive, hwndStatus, hwndBrand, hwndHint uintptr
	oldHeaderWndProc                                                                          uintptr
	headerWndProcCallback                                                                     uintptr
	currentFolder                                                                             string
	currentEntries                                                                            []Entry
	currentText                                                                               string
	currentRecursive                                                                          bool
	currentSortColumn                                                                         = -1
	currentSortAscending                                                                      = true
	columnVisible                                                                             []bool
	columnSavedWidths                                                                         []int32
	hIconBig, hIconSmall                                                                      uintptr
	htmlDlgClassRegistered                                                                    bool
	htmlDlgHwnd                                                                               uintptr
	htmlDlgDone                                                                               bool
	htmlDlgAccepted                                                                           bool
	htmlDlgSelection                                                                          []bool
	exportDlgFormatLabel                                                                      = "HTML"
	startupExportFormat                                                                       string
	startupExportPreset                                                                       string
)

func wptr(s string) *uint16 {
	p, _ := syscall.UTF16PtrFromString(s)
	return p
}

func utf16z(s string) []uint16 { return append(syscall.StringToUTF16(s), 0) }
func loword(v uintptr) uint16  { return uint16(v & 0xffff) }

func createControl(class, text string, style uint32, x, y, w, h int32, parent uintptr, id int) uintptr {
	r, _, _ := pCreateWindowExW.Call(
		0,
		uintptr(unsafe.Pointer(wptr(class))),
		uintptr(unsafe.Pointer(wptr(text))),
		uintptr(style),
		uintptr(x), uintptr(y), uintptr(w), uintptr(h),
		parent, uintptr(id), 0, 0,
	)
	if r != 0 {
		font, _, _ := pGetStockObject.Call(DEFAULT_GUI_FONT)
		pSendMessageW.Call(r, WM_SETFONT, font, 1)
	}
	return r
}

func setText(hwnd uintptr, s string) { pSetWindowTextW.Call(hwnd, uintptr(unsafe.Pointer(wptr(s)))) }

func getText(hwnd uintptr) string {
	n, _, _ := pGetWindowTextLengthW.Call(hwnd)
	buf := make([]uint16, n+1)
	pGetWindowTextW.Call(hwnd, uintptr(unsafe.Pointer(&buf[0])), n+1)
	return syscall.UTF16ToString(buf)
}

func msgBox(text, title string, flags uintptr) int {
	r, _, _ := pMessageBoxW.Call(hwndMain, uintptr(unsafe.Pointer(wptr(text))), uintptr(unsafe.Pointer(wptr(title))), flags)
	return int(r)
}

func createIconFromICO(data []byte, desired int) uintptr {
	if len(data) < 6 || binary.LittleEndian.Uint16(data[0:2]) != 0 || binary.LittleEndian.Uint16(data[2:4]) != 1 {
		return 0
	}
	count := int(binary.LittleEndian.Uint16(data[4:6]))
	if count <= 0 || len(data) < 6+count*16 {
		return 0
	}
	best := -1
	bestDiff := int(^uint(0) >> 1)
	for i := 0; i < count; i++ {
		p := 6 + i*16
		w := int(data[p])
		h := int(data[p+1])
		if w == 0 {
			w = 256
		}
		if h == 0 {
			h = 256
		}
		d := w - desired
		if d < 0 {
			d = -d
		}
		d2 := h - desired
		if d2 < 0 {
			d2 = -d2
		}
		d += d2
		if d < bestDiff {
			bestDiff = d
			best = p
		}
	}
	if best < 0 {
		return 0
	}
	sz := int(binary.LittleEndian.Uint32(data[best+8 : best+12]))
	off := int(binary.LittleEndian.Uint32(data[best+12 : best+16]))
	if sz <= 0 || off < 0 || off+sz > len(data) {
		return 0
	}
	r, _, _ := pCreateIconFromResourceEx.Call(uintptr(unsafe.Pointer(&data[off])), uintptr(sz), 1, 0x00030000, uintptr(desired), uintptr(desired), LR_DEFAULTCOLOR)
	return r
}

func ensureContextIcon() string {
	exe, err := os.Executable()
	if err == nil {
		exe, _ = filepath.Abs(exe)
		p := filepath.Join(filepath.Dir(exe), "GoList.Botolo.ico")
		if os.WriteFile(p, botoloICO, 0644) == nil {
			return p
		}
	}
	base := os.Getenv("LOCALAPPDATA")
	if base == "" {
		base = os.TempDir()
	}
	dir := filepath.Join(base, "ShiduLab", "GoList")
	if os.MkdirAll(dir, 0755) == nil {
		p := filepath.Join(dir, "GoList.Botolo.ico")
		if os.WriteFile(p, botoloICO, 0644) == nil {
			return p
		}
	}
	return ""
}

func browseFolder() string {
	display := make([]uint16, 260)
	bi := BROWSEINFO{
		HwndOwner: hwndMain, PszDisplayName: &display[0],
		LpszTitle: wptr("Scegli la cartella da listare"),
		UlFlags:   BIF_RETURNONLYFSDIRS | BIF_NEWDIALOGSTYLE,
	}
	pidl, _, _ := pSHBrowseForFolderW.Call(uintptr(unsafe.Pointer(&bi)))
	if pidl == 0 {
		return ""
	}
	defer pCoTaskMemFree.Call(pidl)
	path := make([]uint16, 32768)
	ok, _, _ := pSHGetPathFromIDListW.Call(pidl, uintptr(unsafe.Pointer(&path[0])))
	if ok == 0 {
		return ""
	}
	return syscall.UTF16ToString(path)
}

func humanSize(n int64) string {
	const unit = 1024
	if n < unit {
		return fmt.Sprintf("%d B", n)
	}
	div, exp := int64(unit), 0
	for x := n / unit; x >= unit && exp < 4; x /= unit {
		div *= unit
		exp++
	}
	units := []string{"KiB", "MiB", "GiB", "TiB", "PiB"}
	return fmt.Sprintf("%.2f %s", float64(n)/float64(div), units[exp])
}

func fileTimes(info os.FileInfo) (created, modified string) {
	modified = info.ModTime().Format("2006-01-02 15:04:05")
	if d, ok := info.Sys().(*syscall.Win32FileAttributeData); ok {
		created = time.Unix(0, d.CreationTime.Nanoseconds()).Local().Format("2006-01-02 15:04:05")
	}
	return
}

func formatDuration(sec int64) string {
	if sec <= 0 {
		return ""
	}
	h := sec / 3600
	m := (sec % 3600) / 60
	ss := sec % 60
	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, ss)
	}
	return fmt.Sprintf("%d:%02d", m, ss)
}

func syncSafe32(b []byte) int {
	if len(b) < 4 {
		return 0
	}
	return int(b[0]&0x7f)<<21 | int(b[1]&0x7f)<<14 | int(b[2]&0x7f)<<7 | int(b[3]&0x7f)
}

func latin1String(b []byte) string {
	for i, c := range b {
		if c == 0 {
			b = b[:i]
			break
		}
	}
	r := make([]rune, len(b))
	for i, c := range b {
		r[i] = rune(c)
	}
	return strings.TrimSpace(string(r))
}

func decodeID3Text(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	enc := b[0]
	data := b[1:]
	switch enc {
	case 0:
		return latin1String(data)
	case 3:
		if i := strings.IndexByte(string(data), 0); i >= 0 {
			data = data[:i]
		}
		return strings.TrimSpace(string(data))
	case 1, 2:
		if len(data) < 2 {
			return ""
		}
		var order binary.ByteOrder = binary.BigEndian
		if enc == 1 {
			if data[0] == 0xff && data[1] == 0xfe {
				order = binary.LittleEndian
				data = data[2:]
			} else if data[0] == 0xfe && data[1] == 0xff {
				data = data[2:]
			}
		}
		if len(data)%2 != 0 {
			data = data[:len(data)-1]
		}
		u := make([]uint16, 0, len(data)/2)
		for i := 0; i+1 < len(data); i += 2 {
			v := order.Uint16(data[i : i+2])
			if v == 0 {
				break
			}
			u = append(u, v)
		}
		return strings.TrimSpace(string(utf16.Decode(u)))
	}
	return ""
}

func parseMPEGHeader(h uint32) (bitrateKbps, sampleRate, samplesPerFrame int, layer int, version int, ok bool) {
	if h&0xffe00000 != 0xffe00000 {
		return
	}
	versionBits := int((h >> 19) & 0x3)
	layerBits := int((h >> 17) & 0x3)
	bitrateIdx := int((h >> 12) & 0xf)
	sampleIdx := int((h >> 10) & 0x3)
	if versionBits == 1 || layerBits == 0 || bitrateIdx == 0 || bitrateIdx == 15 || sampleIdx == 3 {
		return
	}
	switch versionBits {
	case 3:
		version = 1
	case 2:
		version = 2
	case 0:
		version = 25
	default:
		return
	}
	switch layerBits {
	case 3:
		layer = 1
	case 2:
		layer = 2
	case 1:
		layer = 3
	default:
		return
	}
	brMpeg1L1 := []int{0, 32, 64, 96, 128, 160, 192, 224, 256, 288, 320, 352, 384, 416, 448, 0}
	brMpeg1L2 := []int{0, 32, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320, 384, 0}
	brMpeg1L3 := []int{0, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320, 0}
	brMpeg2L1 := []int{0, 32, 48, 56, 64, 80, 96, 112, 128, 144, 160, 176, 192, 224, 256, 0}
	brMpeg2L23 := []int{0, 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160, 0}
	if version == 1 {
		switch layer {
		case 1:
			bitrateKbps = brMpeg1L1[bitrateIdx]
		case 2:
			bitrateKbps = brMpeg1L2[bitrateIdx]
		case 3:
			bitrateKbps = brMpeg1L3[bitrateIdx]
		}
	} else {
		if layer == 1 {
			bitrateKbps = brMpeg2L1[bitrateIdx]
		} else {
			bitrateKbps = brMpeg2L23[bitrateIdx]
		}
	}
	baseRates := []int{44100, 48000, 32000}
	sampleRate = baseRates[sampleIdx]
	if version == 2 {
		sampleRate /= 2
	} else if version == 25 {
		sampleRate /= 4
	}
	switch layer {
	case 1:
		samplesPerFrame = 384
	case 2:
		samplesPerFrame = 1152
	case 3:
		if version == 1 {
			samplesPerFrame = 1152
		} else {
			samplesPerFrame = 576
		}
	}
	ok = bitrateKbps > 0 && sampleRate > 0
	return
}

func decodeID3Payload(enc byte, data []byte) string {
	buf := make([]byte, 1+len(data))
	buf[0] = enc
	copy(buf[1:], data)
	return decodeID3Text(buf)
}

func splitEncodedField(enc byte, data []byte) (string, []byte) {
	if enc == 1 || enc == 2 {
		for i := 0; i+1 < len(data); i += 2 {
			if data[i] == 0 && data[i+1] == 0 {
				return decodeID3Payload(enc, data[:i]), data[i+2:]
			}
		}
		return decodeID3Payload(enc, data), nil
	}
	for i, b := range data {
		if b == 0 {
			return decodeID3Payload(enc, data[:i]), data[i+1:]
		}
	}
	return decodeID3Payload(enc, data), nil
}

func decodeID3Comment(data []byte) string {
	if len(data) < 4 {
		return ""
	}
	enc := data[0]
	_, rest := splitEncodedField(enc, data[4:]) // encoding + 3-byte language + description
	if len(rest) == 0 {
		return ""
	}
	return decodeID3Payload(enc, rest)
}

func decodeTXXX(data []byte) (string, string) {
	if len(data) < 2 {
		return "", ""
	}
	enc := data[0]
	desc, rest := splitEncodedField(enc, data[1:])
	if len(rest) == 0 {
		return strings.TrimSpace(desc), ""
	}
	return strings.TrimSpace(desc), strings.TrimSpace(decodeID3Payload(enc, rest))
}

func applyID3Frame(meta *AudioMeta, id string, data []byte, extras *[]string) {
	text := func() string { return decodeID3Text(data) }
	setIfEmpty := func(dst *string, v string) {
		if *dst == "" {
			*dst = strings.TrimSpace(v)
		}
	}
	switch id {
	case "TIT2", "TT2":
		setIfEmpty(&meta.Title, text())
	case "TPE1", "TP1":
		setIfEmpty(&meta.Artist, text())
	case "TALB", "TAL":
		setIfEmpty(&meta.Album, text())
	case "TYER", "TDRC", "TYE":
		setIfEmpty(&meta.Year, text())
	case "TRCK", "TRK":
		setIfEmpty(&meta.Track, text())
	case "TCON", "TCO":
		setIfEmpty(&meta.Genre, text())
	case "TPE2", "TP2":
		setIfEmpty(&meta.AlbumArtist, text())
	case "TCOM", "TCM":
		setIfEmpty(&meta.Composer, text())
	case "TPOS", "TPA":
		setIfEmpty(&meta.DiscNumber, text())
	case "TPUB", "TPB":
		setIfEmpty(&meta.Publisher, text())
	case "TCOP", "TCR":
		setIfEmpty(&meta.Copyright, text())
	case "TSRC", "TRC":
		setIfEmpty(&meta.ISRC, text())
	case "COMM", "COM":
		setIfEmpty(&meta.Comment, decodeID3Comment(data))
	case "TXXX", "TXX":
		k, v := decodeTXXX(data)
		if k != "" || v != "" {
			if k == "" {
				*extras = append(*extras, v)
			} else {
				*extras = append(*extras, k+"="+v)
			}
		}
	}
}

func readMP3Metadata(path string, fileSize int64) AudioMeta {
	var meta AudioMeta
	f, err := os.Open(path)
	if err != nil {
		return meta
	}
	defer f.Close()

	audioStart := int64(0)
	hasID3v1 := false
	hasID3v2 := false
	var id3v2Label string
	extra := make([]string, 0, 4)

	head := make([]byte, 10)
	if _, err = f.ReadAt(head, 0); err == nil && string(head[:3]) == "ID3" {
		hasID3v2 = true
		ver := int(head[3])
		id3v2Label = fmt.Sprintf("ID3v2.%d", ver)
		tagSize := syncSafe32(head[6:10])
		audioStart = int64(10 + tagSize)
		if tagSize > 0 && tagSize <= 32*1024*1024 {
			tag := make([]byte, tagSize)
			if _, err := f.ReadAt(tag, 10); err == nil {
				pos := 0
				for pos < len(tag) {
					if ver == 2 {
						if pos+6 > len(tag) {
							break
						}
						id := string(tag[pos : pos+3])
						if id == "\x00\x00\x00" {
							break
						}
						sz := int(tag[pos+3])<<16 | int(tag[pos+4])<<8 | int(tag[pos+5])
						if sz <= 0 || pos+6+sz > len(tag) {
							break
						}
						applyID3Frame(&meta, id, tag[pos+6:pos+6+sz], &extra)
						pos += 6 + sz
						continue
					}
					if pos+10 > len(tag) {
						break
					}
					id := string(tag[pos : pos+4])
					if id == "\x00\x00\x00\x00" {
						break
					}
					var sz int
					if ver == 4 {
						sz = syncSafe32(tag[pos+4 : pos+8])
					} else {
						sz = int(binary.BigEndian.Uint32(tag[pos+4 : pos+8]))
					}
					if sz <= 0 || pos+10+sz > len(tag) {
						break
					}
					applyID3Frame(&meta, id, tag[pos+10:pos+10+sz], &extra)
					pos += 10 + sz
				}
			}
		}
	}

	if fileSize >= 128 {
		tail := make([]byte, 128)
		if _, err := f.ReadAt(tail, fileSize-128); err == nil && string(tail[:3]) == "TAG" {
			hasID3v1 = true
			if meta.Title == "" {
				meta.Title = latin1String(tail[3:33])
			}
			if meta.Artist == "" {
				meta.Artist = latin1String(tail[33:63])
			}
			if meta.Album == "" {
				meta.Album = latin1String(tail[63:93])
			}
			if meta.Year == "" {
				meta.Year = latin1String(tail[93:97])
			}
			if meta.Comment == "" {
				meta.Comment = latin1String(tail[97:127])
			}
			if meta.Track == "" && tail[125] == 0 && tail[126] != 0 {
				meta.Track = strconv.Itoa(int(tail[126]))
			}
		}
	}

	parts := make([]string, 0, 2)
	if hasID3v2 {
		parts = append(parts, id3v2Label)
	}
	if hasID3v1 {
		parts = append(parts, "ID3v1")
	}
	meta.TagVersion = strings.Join(parts, " + ")
	meta.ExtraID3 = strings.Join(extra, " | ")

	searchLen := int64(128 * 1024)
	if fileSize-audioStart < searchLen {
		searchLen = fileSize - audioStart
	}
	if searchLen <= 4 {
		return meta
	}
	buf := make([]byte, int(searchLen))
	n, _ := f.ReadAt(buf, audioStart)
	buf = buf[:n]
	frameAt := -1
	var hdr uint32
	var bitrate, rate, spf, layer, version int
	for i := 0; i+4 <= len(buf); i++ {
		h := binary.BigEndian.Uint32(buf[i : i+4])
		br, sr, sf, ly, v, ok := parseMPEGHeader(h)
		if ok {
			frameAt = i
			hdr = h
			bitrate, rate, spf, layer, version = br, sr, sf, ly, v
			break
		}
	}
	if frameAt < 0 {
		return meta
	}
	firstOffset := audioStart + int64(frameAt)

	probe := make([]byte, 256)
	n, _ = f.ReadAt(probe, firstOffset)
	probe = probe[:n]
	if layer == 3 && len(probe) >= 16 {
		mode := int((hdr >> 6) & 0x3)
		mono := mode == 3
		crcBytes := 0
		if ((hdr >> 16) & 0x1) == 0 {
			crcBytes = 2
		}
		sideInfo := 0
		if version == 1 {
			if mono {
				sideInfo = 17
			} else {
				sideInfo = 32
			}
		} else {
			if mono {
				sideInfo = 9
			} else {
				sideInfo = 17
			}
		}
		x := 4 + crcBytes + sideInfo
		if x+12 <= len(probe) && (string(probe[x:x+4]) == "Xing" || string(probe[x:x+4]) == "Info") {
			flags := binary.BigEndian.Uint32(probe[x+4 : x+8])
			if flags&1 != 0 {
				frames := binary.BigEndian.Uint32(probe[x+8 : x+12])
				if frames > 0 && rate > 0 {
					meta.Duration = int64((float64(frames)*float64(spf))/float64(rate) + 0.5)
					return meta
				}
			}
		}
		if vi := strings.Index(string(probe), "VBRI"); vi >= 0 && vi+18 <= len(probe) {
			frames := binary.BigEndian.Uint32(probe[vi+14 : vi+18])
			if frames > 0 && rate > 0 {
				meta.Duration = int64((float64(frames)*float64(spf))/float64(rate) + 0.5)
				return meta
			}
		}
	}
	audioBytes := fileSize - firstOffset
	if hasID3v1 {
		audioBytes -= 128
	}
	if audioBytes > 0 && bitrate > 0 {
		meta.Duration = int64((float64(audioBytes)*8.0)/(float64(bitrate)*1000.0) + 0.5)
	}
	return meta
}

func readFLACMetadata(path string) AudioMeta {
	var meta AudioMeta
	f, err := os.Open(path)
	if err != nil {
		return meta
	}
	defer f.Close()
	magic := make([]byte, 4)
	if _, err = f.Read(magic); err != nil || string(magic) != "fLaC" {
		return meta
	}
	last := false
	extra := make([]string, 0, 4)
	for !last {
		h := make([]byte, 4)
		if _, err = f.Read(h); err != nil {
			break
		}
		last = h[0]&0x80 != 0
		typ := h[0] & 0x7f
		l := int(h[1])<<16 | int(h[2])<<8 | int(h[3])
		if l < 0 || l > 64*1024*1024 {
			break
		}
		if typ != 0 && typ != 4 {
			_, _ = f.Seek(int64(l), 1)
			continue
		}
		data := make([]byte, l)
		if _, err = f.Read(data); err != nil {
			break
		}
		if typ == 0 && len(data) >= 18 {
			x := binary.BigEndian.Uint64(data[10:18])
			sampleRate := int64((x >> 44) & 0xfffff)
			totalSamples := int64(x & 0xfffffffff)
			if sampleRate > 0 && totalSamples > 0 {
				meta.Duration = int64(float64(totalSamples)/float64(sampleRate) + 0.5)
			}
		} else if typ == 4 && len(data) >= 8 {
			pos := 0
			vendorLen := int(binary.LittleEndian.Uint32(data[pos : pos+4]))
			pos += 4 + vendorLen
			if pos+4 > len(data) {
				continue
			}
			count := int(binary.LittleEndian.Uint32(data[pos : pos+4]))
			pos += 4
			for i := 0; i < count && pos+4 <= len(data); i++ {
				cl := int(binary.LittleEndian.Uint32(data[pos : pos+4]))
				pos += 4
				if cl < 0 || pos+cl > len(data) {
					break
				}
				c := string(data[pos : pos+cl])
				pos += cl
				if eq := strings.IndexByte(c, '='); eq > 0 {
					key := strings.ToUpper(c[:eq])
					val := strings.TrimSpace(c[eq+1:])
					switch key {
					case "TITLE":
						if meta.Title == "" {
							meta.Title = val
						}
					case "ARTIST":
						if meta.Artist == "" {
							meta.Artist = val
						}
					case "ALBUM":
						if meta.Album == "" {
							meta.Album = val
						}
					case "DATE", "YEAR":
						if meta.Year == "" {
							meta.Year = val
						}
					case "TRACKNUMBER":
						if meta.Track == "" {
							meta.Track = val
						}
					case "GENRE":
						if meta.Genre == "" {
							meta.Genre = val
						}
					case "COMMENT":
						if meta.Comment == "" {
							meta.Comment = val
						}
					case "ALBUMARTIST", "ALBUM ARTIST":
						if meta.AlbumArtist == "" {
							meta.AlbumArtist = val
						}
					case "COMPOSER":
						if meta.Composer == "" {
							meta.Composer = val
						}
					case "DISCNUMBER":
						if meta.DiscNumber == "" {
							meta.DiscNumber = val
						}
					case "PUBLISHER", "LABEL":
						if meta.Publisher == "" {
							meta.Publisher = val
						}
					case "COPYRIGHT":
						if meta.Copyright == "" {
							meta.Copyright = val
						}
					case "ISRC":
						if meta.ISRC == "" {
							meta.ISRC = val
						}
					default:
						extra = append(extra, key+"="+val)
					}
				}
			}
		}
	}
	meta.ExtraID3 = strings.Join(extra, " | ")
	return meta
}

func readWAVMetadata(path string) AudioMeta {
	var meta AudioMeta
	f, err := os.Open(path)
	if err != nil {
		return meta
	}
	defer f.Close()
	h := make([]byte, 12)
	if _, err = f.Read(h); err != nil || string(h[:4]) != "RIFF" || string(h[8:12]) != "WAVE" {
		return meta
	}
	var byteRate uint32
	var dataSize uint32
	for {
		ch := make([]byte, 8)
		if _, err = f.Read(ch); err != nil {
			break
		}
		id := string(ch[:4])
		sz := binary.LittleEndian.Uint32(ch[4:8])
		if sz > 512*1024*1024 {
			break
		}
		switch id {
		case "fmt ":
			data := make([]byte, sz)
			if _, err = f.Read(data); err != nil {
				return meta
			}
			if len(data) >= 12 {
				byteRate = binary.LittleEndian.Uint32(data[8:12])
			}
		case "data":
			dataSize = sz
			_, _ = f.Seek(int64(sz), 1)
		case "LIST":
			if sz < 4 {
				_, _ = f.Seek(int64(sz), 1)
				break
			}
			data := make([]byte, sz)
			if _, err = f.Read(data); err != nil {
				return meta
			}
			if string(data[:4]) == "INFO" {
				pos := 4
				for pos+8 <= len(data) {
					sid := string(data[pos : pos+4])
					sl := int(binary.LittleEndian.Uint32(data[pos+4 : pos+8]))
					pos += 8
					if sl < 0 || pos+sl > len(data) {
						break
					}
					val := strings.TrimSpace(strings.TrimRight(string(data[pos:pos+sl]), "\x00"))
					switch sid {
					case "INAM":
						if meta.Title == "" {
							meta.Title = val
						}
					case "IART":
						if meta.Artist == "" {
							meta.Artist = val
						}
					case "IPRD":
						if meta.Album == "" {
							meta.Album = val
						}
					case "ICRD":
						if meta.Year == "" {
							meta.Year = val
						}
					case "IGNR":
						if meta.Genre == "" {
							meta.Genre = val
						}
					case "ICMT":
						if meta.Comment == "" {
							meta.Comment = val
						}
					}
					pos += sl
					if sl%2 == 1 {
						pos++
					}
				}
			}
		default:
			_, _ = f.Seek(int64(sz), 1)
		}
		if sz%2 == 1 {
			_, _ = f.Seek(1, 1)
		}
	}
	if byteRate > 0 && dataSize > 0 {
		meta.Duration = int64(float64(dataSize)/float64(byteRate) + 0.5)
	}
	return meta
}

func readAudioMetadata(path string, fileSize int64) AudioMeta {
	switch strings.ToLower(filepath.Ext(path)) {
	case ".mp3":
		return readMP3Metadata(path, fileSize)
	case ".flac":
		return readFLACMetadata(path)
	case ".wav", ".wave":
		return readWAVMetadata(path)
	}
	return AudioMeta{}
}

func parseSampledBy(text string) string {
	// Parser volutamente permissivo per l'archivio storico.
	// Non pretende piu' che il marcatore sia scritto ESATTAMENTE [S] o [St]:
	// basta un gruppo tra parentesi seguito da un marcatore tra [] oppure,
	// come ulteriore sicurezza, un anno a 4 cifre subito prima del gruppo.
	// Esempi:
	//   ... 1969 (GangStarr) [St].mp3
	//   ... 1972 (Wu-Tang and other)[St].mp3
	//   ... 1971 (Black Sheep, Mos Def) [S].mp3
	normalized := strings.NewReplacer("［", "[", "］", "]", "（", "(", "）", ")").Replace(text)
	base := strings.TrimSpace(normalized)
	if ext := filepath.Ext(base); ext != "" {
		base = strings.TrimSuffix(base, ext)
	}

	// Cerca da destra una coppia [...] che venga dopo l'ultima parentesi (...).
	closeBracket := strings.LastIndex(base, "]")
	openBracket := -1
	if closeBracket >= 0 {
		openBracket = strings.LastIndex(base[:closeBracket], "[")
	}

	searchEnd := len(base)
	hasBracketMarker := false
	if openBracket >= 0 && closeBracket > openBracket {
		// Accettiamo il marcatore storico anche se contiene spazi o varianti
		// non previste: il fatto che sia subito dopo (...) e' il vero segnale.
		hasBracketMarker = true
		searchEnd = openBracket
	}

	prefix := strings.TrimRight(strings.TrimSpace(base[:searchEnd]), " -_–—·")
	close := strings.LastIndex(prefix, ")")
	if close < 0 {
		return ""
	}
	if strings.Trim(prefix[close+1:], " \t-_–—·") != "" {
		return ""
	}

	// Trova la parentesi aperta corrispondente all'ultima chiusa.
	depth := 0
	open := -1
	for i := close; i >= 0; i-- {
		switch prefix[i] {
		case ')':
			depth++
		case '(':
			depth--
			if depth == 0 {
				open = i
				i = -1
			}
		}
	}
	if open < 0 || open+1 >= close {
		return ""
	}

	// Se c'e' [...] subito dopo, e' sufficiente. Se per qualche motivo il
	// marcatore e' assente/non standard, accettiamo solo il formato archivio
	// con anno a 4 cifre prima della parentesi, evitando casi tipo A Proi (lassa).
	if !hasBracketMarker {
		before := strings.TrimSpace(prefix[:open])
		hasYear := false
		for i := 0; i+4 <= len(before); i++ {
			chunk := before[i : i+4]
			if chunk[0] >= '1' && chunk[0] <= '2' &&
				chunk[1] >= '0' && chunk[1] <= '9' &&
				chunk[2] >= '0' && chunk[2] <= '9' &&
				chunk[3] >= '0' && chunk[3] <= '9' {
				hasYear = true
			}
		}
		if !hasYear {
			return ""
		}
	}

	return strings.TrimSpace(prefix[open+1 : close])
}

func sampledByFor(filename string, meta AudioMeta) string {
	// Prima fonte: il nome file, che è dove l'archivio storico annota i sample.
	// Fallback: eventuale Titolo/Commento ID3 con la stessa notazione [S]/[St].
	for _, source := range []string{filename, meta.Title, meta.Comment} {
		if v := parseSampledBy(source); v != "" {
			return v
		}
	}
	return ""
}

func folderTreeSize(path string) (int64, bool) {
	var total int64
	complete := true
	err := filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			complete = false
			return nil
		}
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil {
			complete = false
			return nil
		}
		total += info.Size()
		return nil
	})
	if err != nil {
		complete = false
	}
	return total, complete
}

func scanFolder(root string, recursive bool) ([]Entry, error) {
	root = filepath.Clean(root)
	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("non è una cartella")
	}
	out := make([]Entry, 0, 128)

	add := func(path string, d os.DirEntry) error {
		info, err := d.Info()
		if err != nil {
			return nil
		}
		rel, _ := filepath.Rel(root, path)
		kind := strings.TrimPrefix(strings.ToLower(filepath.Ext(d.Name())), ".")
		size := info.Size()
		if d.IsDir() {
			kind, size = "DIR", 0
		}
		if kind == "" {
			kind = "FILE"
		}
		created, modified := fileTimes(info)
		meta := AudioMeta{}
		if !d.IsDir() {
			meta = readAudioMetadata(path, size)
		}
		out = append(out, Entry{
			Name: d.Name(), RelativePath: rel, Kind: kind, Size: size,
			Title: meta.Title, Artist: meta.Artist, Album: meta.Album, DurationSeconds: meta.Duration,
			SampledBy: sampledByFor(d.Name(), meta), Year: meta.Year, Track: meta.Track, Genre: meta.Genre,
			Comment: meta.Comment, AlbumArtist: meta.AlbumArtist, Composer: meta.Composer, DiscNumber: meta.DiscNumber,
			Publisher: meta.Publisher, Copyright: meta.Copyright, ISRC: meta.ISRC, TagVersion: meta.TagVersion, ExtraID3: meta.ExtraID3,
			Created: created, Modified: modified,
		})
		return nil
	}

	if recursive {
		err = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if path == root {
				return nil
			}
			return add(path, d)
		})
	} else {
		var ds []os.DirEntry
		ds, err = os.ReadDir(root)
		if err == nil {
			for _, d := range ds {
				_ = add(filepath.Join(root, d.Name()), d)
			}
		}
	}
	if err != nil {
		return nil, err
	}

	// Le cartelle hanno un peso reale: in modalità ricorsiva lo ricaviamo
	// dagli elementi già letti (senza una seconda scansione); in modalità
	// normale calcoliamo il contenuto di ciascuna sottocartella.
	if recursive {
		sep := string(os.PathSeparator)
		for i := range out {
			if out[i].Kind != "DIR" {
				continue
			}
			prefix := out[i].RelativePath + sep
			var total int64
			for _, e := range out {
				if e.Kind != "DIR" && strings.HasPrefix(e.RelativePath, prefix) {
					total += e.Size
				}
			}
			out[i].Size = total
		}
	} else {
		for i := range out {
			if out[i].Kind != "DIR" {
				continue
			}
			sz, complete := folderTreeSize(filepath.Join(root, out[i].RelativePath))
			// Se una parte non è leggibile, mostriamo comunque il peso realmente
			// riscontrato. Solo una cartella del tutto illeggibile resta “—”.
			if complete || sz > 0 {
				out[i].Size = sz
			} else {
				out[i].Size = -1
			}
		}
	}

	sort.SliceStable(out, func(i, j int) bool {
		return strings.ToLower(out[i].RelativePath) < strings.ToLower(out[j].RelativePath)
	})
	return out, nil
}

func visibleColumnsInOrder() []int {
	order := currentColumnOrder()
	result := make([]int, 0, len(order))
	for _, i := range order {
		if i >= 0 && i < len(columns) {
			visible := columns[i].DefaultVisible
			if i < len(columnVisible) {
				visible = columnVisible[i]
			}
			if visible {
				result = append(result, i)
			}
		}
	}
	return result
}

func renderTextColumns(root string, entries []Entry, selected []int) string {
	if len(selected) == 0 {
		selected = []int{0}
	}

	var b strings.Builder
	b.WriteString("GoList!\r\n")
	b.WriteString("Cartella: " + root + "\r\n")
	b.WriteString(fmt.Sprintf("Elementi: %d", len(entries)))
	if len(entries) > 0 {
		b.WriteString(" · Peso totale: " + humanSize(totalListedSize(entries, currentRecursive)))
		if total := formatTotalDuration(totalAudioDuration(entries)); total != "" {
			b.WriteString(" · Durata audio: " + total)
		}
	}
	b.WriteString("\r\n\r\n")

	for pos, i := range selected {
		if pos > 0 {
			b.WriteString(" | ")
		}
		b.WriteString(columns[i].Title)
	}
	b.WriteString("\r\n")
	for pos, i := range selected {
		if pos > 0 {
			b.WriteString("-+-")
		}
		b.WriteString(strings.Repeat("-", utf8.RuneCountInString(columns[i].Title)))
	}
	b.WriteString("\r\n")

	for _, e := range entries {
		for pos, i := range selected {
			if pos > 0 {
				b.WriteString(" | ")
			}
			v := tableCell(e, i)
			v = strings.NewReplacer("\r", " ", "\n", " ", "\t", " ").Replace(v)
			b.WriteString(v)
		}
		b.WriteString("\r\n")
	}
	return b.String()
}

func renderText(root string, entries []Entry) string {
	return renderTextColumns(root, entries, visibleColumnsInOrder())
}

func tableCell(e Entry, col int) string {
	switch col {
	case 0:
		return e.Name
	case 1:
		return e.Title
	case 2:
		return e.Artist
	case 3:
		return e.Album
	case 4:
		return formatDuration(e.DurationSeconds)
	case 5:
		return e.SampledBy
	case 6:
		return e.Year
	case 7:
		return e.Track
	case 8:
		return e.Genre
	case 9:
		return e.Comment
	case 10:
		return e.AlbumArtist
	case 11:
		return e.Composer
	case 12:
		return e.DiscNumber
	case 13:
		return e.Publisher
	case 14:
		return e.Copyright
	case 15:
		return e.ISRC
	case 16:
		return e.TagVersion
	case 17:
		return e.ExtraID3
	case 18:
		return e.Kind
	case 19:
		if e.Size < 0 {
			return "—"
		}
		return humanSize(e.Size)
	case 20:
		return e.Created
	case 21:
		return e.Modified
	case 22:
		return e.RelativePath
	}
	return ""
}

func numericPrefix(s string) int64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return -1
	}
	end := 0
	for end < len(s) && s[end] >= '0' && s[end] <= '9' {
		end++
	}
	if end == 0 {
		return -1
	}
	n, err := strconv.ParseInt(s[:end], 10, 64)
	if err != nil {
		return -1
	}
	return n
}

func totalAudioDuration(entries []Entry) int64 {
	var total int64
	for _, e := range entries {
		total += e.DurationSeconds
	}
	return total
}

func totalListedSize(entries []Entry, recursive bool) int64 {
	var total int64
	for _, e := range entries {
		if e.Size < 0 {
			continue
		}
		// Con Sottocartelle attivo le directory contengono già la somma dei
		// figli: sommarle di nuovo conterebbe gli stessi byte più volte.
		if recursive && e.Kind == "DIR" {
			continue
		}
		total += e.Size
	}
	return total
}

func formatTotalDuration(sec int64) string {
	if sec <= 0 {
		return ""
	}
	h := sec / 3600
	m := (sec % 3600) / 60
	s := sec % 60
	return fmt.Sprintf("%d:%02d:%02d", h, m, s)
}

func headerWndProc(hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	if msg == WM_RBUTTONUP {
		showColumnMenu()
		return 0
	}
	r, _, _ := pCallWindowProcW.Call(oldHeaderWndProc, hwnd, uintptr(msg), wParam, lParam)
	return r
}

func showColumnMenu() {
	if hwndTable == 0 {
		return
	}

	// Keep the column chooser effectively open while the user toggles
	// several columns. A native popup menu closes after each command, so
	// we recreate it immediately at the same position until the user clicks
	// outside, presses Esc, or chooses "Chiudi".
	var pt POINT
	pGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	closeID := IDColumnBase + len(columns) + 100

	for {
		menu, _, _ := pCreatePopupMenu.Call()
		if menu == 0 {
			return
		}

		for i, c := range columns {
			flags := uintptr(MF_STRING)
			if i < len(columnVisible) && columnVisible[i] {
				flags |= MF_CHECKED
			}
			pAppendMenuW.Call(menu, flags, uintptr(IDColumnBase+i), uintptr(unsafe.Pointer(wptr(c.Title))))
		}
		pAppendMenuW.Call(menu, MF_SEPARATOR, 0, 0)
		pAppendMenuW.Call(menu, MF_STRING, uintptr(closeID), uintptr(unsafe.Pointer(wptr("Chiudi"))))

		cmd, _, _ := pTrackPopupMenu.Call(menu, TPM_RETURNCMD|TPM_RIGHTBUTTON, uintptr(int64(pt.X)), uintptr(int64(pt.Y)), 0, hwndMain, 0)
		pDestroyMenu.Call(menu)

		if int(cmd) == closeID || cmd == 0 {
			return
		}
		i := int(cmd) - IDColumnBase
		if i < 0 || i >= len(columns) {
			return
		}
		toggleColumn(i)
	}
}

func toggleColumn(i int) {
	if i < 0 || i >= len(columns) {
		return
	}
	if len(columnVisible) != len(columns) {
		columnVisible = make([]bool, len(columns))
		for j := range columnVisible {
			columnVisible[j] = true
		}
	}
	if len(columnSavedWidths) != len(columns) {
		columnSavedWidths = make([]int32, len(columns))
		for j, c := range columns {
			columnSavedWidths[j] = c.Width
		}
	}
	if columnVisible[i] {
		visibleCount := 0
		for _, v := range columnVisible {
			if v {
				visibleCount++
			}
		}
		if visibleCount <= 1 {
			return
		}
		r, _, _ := pSendMessageW.Call(hwndTable, LVM_GETCOLUMNWIDTH, uintptr(i), 0)
		if int32(r) >= 30 {
			columnSavedWidths[i] = int32(r)
		}
		columnVisible[i] = false
		pSendMessageW.Call(hwndTable, LVM_SETCOLUMNWIDTH, uintptr(i), 0)
	} else {
		columnVisible[i] = true
		w := columnSavedWidths[i]
		if w < 30 || w > 2000 {
			w = columns[i].Width
		}
		pSendMessageW.Call(hwndTable, LVM_SETCOLUMNWIDTH, uintptr(i), uintptr(w))
	}
	saveTableLayout()
	if currentFolder != "" && len(currentEntries) > 0 {
		currentText = renderText(currentFolder, currentEntries)
	}
}

func setupTable() {
	columnVisible = make([]bool, len(columns))
	columnSavedWidths = make([]int32, len(columns))
	for i, c := range columns {
		columnVisible[i] = c.DefaultVisible
		columnSavedWidths[i] = c.Width
		txt := wptr(c.Title)
		col := LVCOLUMN{Mask: LVCF_FMT | LVCF_WIDTH | LVCF_TEXT | LVCF_SUBITEM, Fmt: c.Align, Cx: c.Width, PszText: txt, ISubItem: int32(i)}
		pSendMessageW.Call(hwndTable, LVM_INSERTCOLUMNW, uintptr(i), uintptr(unsafe.Pointer(&col)))
	}
	ex := uintptr(LVS_EX_FULLROWSELECT | LVS_EX_HEADERDRAGDROP | LVS_EX_GRIDLINES | LVS_EX_DOUBLEBUFFER)
	pSendMessageW.Call(hwndTable, LVM_SETEXTENDEDLISTVIEWSTYLE, 0, ex)
	loadTableLayout()
	for i := range columns {
		w := columnSavedWidths[i]
		if !columnVisible[i] {
			w = 0
		}
		pSendMessageW.Call(hwndTable, LVM_SETCOLUMNWIDTH, uintptr(i), uintptr(w))
	}

	h, _, _ := pSendMessageW.Call(hwndTable, LVM_GETHEADER, 0, 0)
	hwndHeader = h
	if hwndHeader != 0 {
		headerWndProcCallback = syscall.NewCallback(headerWndProc)
		idx := int32(GWLP_WNDPROC)
		oldHeaderWndProc, _, _ = pSetWindowLongPtrW.Call(hwndHeader, uintptr(idx), headerWndProcCallback)
	}
}

func refreshTable() {
	if hwndTable == 0 {
		return
	}
	pSendMessageW.Call(hwndTable, LVM_DELETEALLITEMS, 0, 0)
	for i, e := range currentEntries {
		first := wptr(tableCell(e, 0))
		item := LVITEM{Mask: LVIF_TEXT, IItem: int32(i), ISubItem: 0, PszText: first}
		pSendMessageW.Call(hwndTable, LVM_INSERTITEMW, 0, uintptr(unsafe.Pointer(&item)))
		for col := 1; col < len(columns); col++ {
			txt := wptr(tableCell(e, col))
			sub := LVITEM{Mask: LVIF_TEXT, IItem: int32(i), ISubItem: int32(col), PszText: txt}
			pSendMessageW.Call(hwndTable, LVM_SETITEMW, 0, uintptr(unsafe.Pointer(&sub)))
		}
	}
	currentText = renderText(currentFolder, currentEntries)
	if currentFolder == "" {
		setText(hwndStatus, "Pronto.")
		pShowWindow.Call(hwndHint, SW_SHOW)
	} else {
		total := formatTotalDuration(totalAudioDuration(currentEntries))
		weight := humanSize(totalListedSize(currentEntries, currentRecursive))
		if total != "" {
			setText(hwndStatus, fmt.Sprintf("%d elementi · %s · %s audio — %s", len(currentEntries), weight, total, currentFolder))
		} else {
			setText(hwndStatus, fmt.Sprintf("%d elementi · %s — %s", len(currentEntries), weight, currentFolder))
		}
		pShowWindow.Call(hwndHint, SW_HIDE)
	}
}

func sortCurrent(col int) {
	if col < 0 || col >= len(columns) {
		return
	}
	if currentSortColumn == col {
		currentSortAscending = !currentSortAscending
	} else {
		currentSortColumn = col
		currentSortAscending = true
	}
	asc := currentSortAscending
	less := func(a, b Entry) bool {
		cmp := 0
		switch col {
		case 4:
			if a.DurationSeconds < b.DurationSeconds {
				cmp = -1
			} else if a.DurationSeconds > b.DurationSeconds {
				cmp = 1
			}
		case 6, 7, 12:
			an, bn := numericPrefix(tableCell(a, col)), numericPrefix(tableCell(b, col))
			if an >= 0 && bn >= 0 {
				if an < bn {
					cmp = -1
				} else if an > bn {
					cmp = 1
				}
			} else {
				cmp = strings.Compare(strings.ToLower(tableCell(a, col)), strings.ToLower(tableCell(b, col)))
			}
		case 19:
			if a.Size < b.Size {
				cmp = -1
			} else if a.Size > b.Size {
				cmp = 1
			}
		case 20:
			cmp = strings.Compare(a.Created, b.Created)
		case 21:
			cmp = strings.Compare(a.Modified, b.Modified)
		default:
			cmp = strings.Compare(strings.ToLower(tableCell(a, col)), strings.ToLower(tableCell(b, col)))
		}
		if cmp == 0 {
			cmp = strings.Compare(strings.ToLower(a.RelativePath), strings.ToLower(b.RelativePath))
		}
		if asc {
			return cmp < 0
		}
		return cmp > 0
	}
	sort.SliceStable(currentEntries, func(i, j int) bool { return less(currentEntries[i], currentEntries[j]) })
	refreshTable()
}

func layoutPath() string {
	exe, err := os.Executable()
	if err != nil {
		return "GoList.layout.json"
	}
	exe, _ = filepath.Abs(exe)
	return filepath.Join(filepath.Dir(exe), "GoList.layout.json")
}

func saveTableLayout() {
	if hwndTable == 0 {
		return
	}
	if len(columnSavedWidths) != len(columns) {
		columnSavedWidths = make([]int32, len(columns))
		for i, c := range columns {
			columnSavedWidths[i] = c.Width
		}
	}
	for i := range columns {
		if i < len(columnVisible) && columnVisible[i] {
			r, _, _ := pSendMessageW.Call(hwndTable, LVM_GETCOLUMNWIDTH, uintptr(i), 0)
			w := int32(r)
			if w >= 30 && w <= 2000 {
				columnSavedWidths[i] = w
			}
		}
	}
	order := make([]int32, len(columns))
	if len(order) > 0 {
		pSendMessageW.Call(hwndTable, LVM_GETCOLUMNORDERARRAY, uintptr(len(order)), uintptr(unsafe.Pointer(&order[0])))
	}
	visible := append([]bool(nil), columnVisible...)
	widths := append([]int32(nil), columnSavedWidths...)
	data, err := json.MarshalIndent(TableLayout{Widths: widths, Order: order, Visible: visible}, "", "  ")
	if err == nil {
		_ = os.WriteFile(layoutPath(), append(data, '\n'), 0644)
	}
}

func loadTableLayout() {
	data, err := os.ReadFile(layoutPath())
	if err != nil {
		return
	}
	var l TableLayout
	if json.Unmarshal(data, &l) != nil {
		return
	}
	if len(l.Widths) == len(columns) {
		for i, w := range l.Widths {
			if w >= 30 && w <= 2000 {
				columnSavedWidths[i] = w
			}
		}
	}
	if len(l.Visible) == len(columns) {
		nVisible := 0
		for _, v := range l.Visible {
			if v {
				nVisible++
			}
		}
		if nVisible > 0 {
			copy(columnVisible, l.Visible)
		}
	}
	for i := range columns {
		w := columnSavedWidths[i]
		if !columnVisible[i] {
			w = 0
		}
		pSendMessageW.Call(hwndTable, LVM_SETCOLUMNWIDTH, uintptr(i), uintptr(w))
	}
	if len(l.Order) == len(columns) {
		seen := make(map[int32]bool)
		valid := true
		for _, x := range l.Order {
			if x < 0 || int(x) >= len(columns) || seen[x] {
				valid = false
				break
			}
			seen[x] = true
		}
		if valid {
			pSendMessageW.Call(hwndTable, LVM_SETCOLUMNORDERARRAY, uintptr(len(l.Order)), uintptr(unsafe.Pointer(&l.Order[0])))
		}
	}
}

func doList() {
	root := strings.TrimSpace(getText(hwndPath))
	if root == "" {
		msgBox("Scegli o trascina una cartella.", appTitle, MB_OK|MB_ICONINFORMATION)
		return
	}
	r, _, _ := pSendMessageW.Call(hwndRecursive, BM_GETCHECK, 0, 0)
	recursive := r == BST_CHECKED
	entries, err := scanFolder(root, recursive)
	if err != nil {
		msgBox("Impossibile leggere la cartella:\r\n"+err.Error(), appTitle, MB_OK|MB_ICONERROR)
		return
	}
	currentFolder, currentEntries = root, entries
	currentRecursive = recursive
	currentSortColumn = -1
	currentSortAscending = true
	refreshTable()
}

func copyClipboard(text string) error {
	if text == "" {
		return fmt.Errorf("nulla da copiare")
	}
	r, _, _ := pOpenClipboard.Call(hwndMain)
	if r == 0 {
		return fmt.Errorf("clipboard non disponibile")
	}
	defer pCloseClipboard.Call()
	pEmptyClipboard.Call()
	u := utf16z(text)
	bytes := uintptr(len(u) * 2)
	h, _, _ := pGlobalAlloc.Call(GMEM_MOVEABLE, bytes)
	if h == 0 {
		return fmt.Errorf("memoria clipboard")
	}
	ptr, _, _ := pGlobalLock.Call(h)
	if ptr == 0 {
		return fmt.Errorf("memoria clipboard")
	}
	dst := unsafe.Slice((*uint16)(unsafe.Pointer(ptr)), len(u))
	copy(dst, u)
	pGlobalUnlock.Call(h)
	r, _, _ = pSetClipboardData.Call(CF_UNICODETEXT, h)
	if r == 0 {
		return fmt.Errorf("scrittura clipboard")
	}
	return nil
}

func chooseExportFormat() string {
	menu, _, _ := pCreatePopupMenu.Call()
	if menu == 0 {
		return ""
	}
	defer pDestroyMenu.Call(menu)
	pAppendMenuW.Call(menu, MF_STRING, IDExportHTML, uintptr(unsafe.Pointer(wptr("HTML — pagina web (.html)"))))
	pAppendMenuW.Call(menu, MF_SEPARATOR, 0, 0)
	pAppendMenuW.Call(menu, MF_STRING, IDExportTXT, uintptr(unsafe.Pointer(wptr("TXT — testo (.txt)"))))
	pAppendMenuW.Call(menu, MF_STRING, IDExportTSV, uintptr(unsafe.Pointer(wptr("TSV — Excel/tabulato (.tsv)"))))
	pAppendMenuW.Call(menu, MF_STRING, IDExportCSV, uintptr(unsafe.Pointer(wptr("CSV — valori separati (.csv)"))))
	pAppendMenuW.Call(menu, MF_STRING, IDExportJSON, uintptr(unsafe.Pointer(wptr("JSON — dati strutturati (.json)"))))
	var pt POINT
	pGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	cmd, _, _ := pTrackPopupMenu.Call(menu, TPM_RETURNCMD|TPM_RIGHTBUTTON, uintptr(int64(pt.X)), uintptr(int64(pt.Y)), 0, hwndMain, 0)
	switch int(cmd) {
	case IDExportHTML:
		return ".html"
	case IDExportTXT:
		return ".txt"
	case IDExportTSV:
		return ".tsv"
	case IDExportCSV:
		return ".csv"
	case IDExportJSON:
		return ".json"
	}
	return ""
}

func saveDialogFor(ext string) string {
	labels := map[string]string{
		".html": "HTML (*.html)",
		".txt":  "Testo (*.txt)",
		".tsv":  "TSV (*.tsv)",
		".csv":  "CSV (*.csv)",
		".json": "JSON (*.json)",
	}
	patterns := map[string]string{
		".html": "*.html", ".txt": "*.txt", ".tsv": "*.tsv", ".csv": "*.csv", ".json": "*.json",
	}
	label := labels[ext]
	pattern := patterns[ext]
	if label == "" {
		return ""
	}
	buf := make([]uint16, 32768)
	copy(buf, syscall.StringToUTF16("GoList"+ext))
	filter := label + "\x00" + pattern + "\x00\x00"
	fu := append(utf16.Encode([]rune(filter)), 0)
	defExt := strings.TrimPrefix(ext, ".")
	of := OPENFILENAME{LStructSize: uint32(unsafe.Sizeof(OPENFILENAME{})), HwndOwner: hwndMain, LpstrFilter: &fu[0], NFilterIndex: 1, LpstrFile: &buf[0], NMaxFile: uint32(len(buf)), LpstrTitle: wptr("Esporta " + strings.ToUpper(defExt)), LpstrDefExt: wptr(defExt), Flags: OFN_OVERWRITEPROMPT | OFN_PATHMUSTEXIST}
	r, _, _ := pGetSaveFileNameW.Call(uintptr(unsafe.Pointer(&of)))
	if r == 0 {
		return ""
	}
	path := syscall.UTF16ToString(buf)
	if filepath.Ext(path) == "" {
		path += ext
	}
	return path
}

func currentColumnOrder() []int {
	order := make([]int32, len(columns))
	if hwndTable != 0 && len(order) > 0 {
		pSendMessageW.Call(hwndTable, LVM_GETCOLUMNORDERARRAY, uintptr(len(order)), uintptr(unsafe.Pointer(&order[0])))
	}
	result := make([]int, 0, len(columns))
	seen := make([]bool, len(columns))
	for _, v := range order {
		i := int(v)
		if i >= 0 && i < len(columns) && !seen[i] {
			result = append(result, i)
			seen[i] = true
		}
	}
	for i := range columns {
		if !seen[i] {
			result = append(result, i)
		}
	}
	return result
}

func htmlDialogWndProc(hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case WM_CREATE:
		htmlDlgHwnd = hwnd
		createControl("STATIC", "Seleziona gli attributi da esportare in "+exportDlgFormatLabel+":", WS_CHILD|WS_VISIBLE, 20, 16, 650, 22, hwnd, 0)
		order := currentColumnOrder()
		for pos, i := range order {
			col := pos / 12
			row := pos % 12
			x := int32(20 + col*335)
			y := int32(48 + row*28)
			h := createControl("BUTTON", columns[i].Title, WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_AUTOCHECKBOX, x, y, 310, 23, hwnd, IDHTMLColumnBase+i)
			checked := i < len(htmlDlgSelection) && htmlDlgSelection[i]
			if checked {
				pSendMessageW.Call(h, BM_SETCHECK, BST_CHECKED, 0)
			}
		}
		createControl("STATIC", "Ordine: disposizione corrente delle colonne", WS_CHILD|WS_VISIBLE, 20, 390, 320, 22, hwnd, 0)
		createControl("BUTTON", "Tutte", WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_PUSHBUTTON, 20, 425, 90, 28, hwnd, IDHTMLAll)
		createControl("BUTTON", "Nessuna", WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_PUSHBUTTON, 118, 425, 90, 28, hwnd, IDHTMLNone)
		createControl("BUTTON", "Esporta "+exportDlgFormatLabel, WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_PUSHBUTTON, 440, 425, 120, 28, hwnd, IDHTMLOK)
		createControl("BUTTON", "Annulla", WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_PUSHBUTTON, 568, 425, 90, 28, hwnd, IDHTMLCancel)
		return 0
	case WM_COMMAND:
		switch int(loword(wParam)) {
		case IDHTMLAll, IDHTMLNone:
			state := uintptr(0)
			if int(loword(wParam)) == IDHTMLAll {
				state = BST_CHECKED
			}
			getDlg := user32.NewProc("GetDlgItem")
			for i := range columns {
				h, _, _ := getDlg.Call(hwnd, uintptr(IDHTMLColumnBase+i))
				if h != 0 {
					pSendMessageW.Call(h, BM_SETCHECK, state, 0)
				}
			}
			return 0
		case IDHTMLOK:
			getDlg := user32.NewProc("GetDlgItem")
			selected := make([]bool, len(columns))
			n := 0
			for i := range columns {
				h, _, _ := getDlg.Call(hwnd, uintptr(IDHTMLColumnBase+i))
				if h != 0 {
					r, _, _ := pSendMessageW.Call(h, BM_GETCHECK, 0, 0)
					if r == BST_CHECKED {
						selected[i] = true
						n++
					}
				}
			}
			if n == 0 {
				pMessageBoxW.Call(hwnd, uintptr(unsafe.Pointer(wptr("Seleziona almeno un attributo."))), uintptr(unsafe.Pointer(wptr("Esporta "+exportDlgFormatLabel))), MB_OK|MB_ICONINFORMATION)
				return 0
			}
			htmlDlgSelection = selected
			htmlDlgAccepted = true
			pDestroyWindow.Call(hwnd)
			return 0
		case IDHTMLCancel:
			htmlDlgAccepted = false
			pDestroyWindow.Call(hwnd)
			return 0
		}
	case WM_CLOSE:
		htmlDlgAccepted = false
		pDestroyWindow.Call(hwnd)
		return 0
	case WM_DESTROY:
		htmlDlgDone = true
		htmlDlgHwnd = 0
		return 0
	}
	r, _, _ := pDefWindowProcW.Call(hwnd, uintptr(msg), wParam, lParam)
	return r
}

func chooseExportColumns(formatLabel string) ([]int, bool) {
	exportDlgFormatLabel = strings.ToUpper(strings.TrimSpace(formatLabel))
	if exportDlgFormatLabel == "" {
		exportDlgFormatLabel = "HTML"
	}
	if len(columns) == 0 {
		return nil, false
	}
	if !htmlDlgClassRegistered {
		hInst, _, _ := pGetModuleHandleW.Call(0)
		className := wptr("GoListHTMLExportClass")
		cursor, _, _ := pLoadCursorW.Call(0, IDC_ARROW)
		wc := WNDCLASSEX{CbSize: uint32(unsafe.Sizeof(WNDCLASSEX{})), LpfnWndProc: syscall.NewCallback(htmlDialogWndProc), HInstance: hInst, HIcon: hIconBig, HCursor: cursor, HbrBackground: COLOR_WINDOW + 1, LpszClassName: className, HIconSm: hIconSmall}
		if r, _, _ := pRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc))); r == 0 {
			return nil, false
		}
		htmlDlgClassRegistered = true
	}

	htmlDlgSelection = make([]bool, len(columns))
	for i := range columns {
		if i < len(columnVisible) {
			htmlDlgSelection[i] = columnVisible[i]
		} else {
			htmlDlgSelection[i] = columns[i].DefaultVisible
		}
	}
	htmlDlgDone = false
	htmlDlgAccepted = false

	hInst, _, _ := pGetModuleHandleW.Call(0)
	className := wptr("GoListHTMLExportClass")
	hwnd, _, _ := pCreateWindowExW.Call(1, uintptr(unsafe.Pointer(className)), uintptr(unsafe.Pointer(wptr("Esporta "+exportDlgFormatLabel+" — Attributi"))), WS_OVERLAPPED|WS_CAPTION|WS_SYSMENU|WS_VISIBLE, CW_USEDEFAULT, CW_USEDEFAULT, 700, 510, hwndMain, 0, hInst, 0)
	if hwnd == 0 {
		return nil, false
	}
	pEnableWindow.Call(hwndMain, 0)
	pShowWindow.Call(hwnd, SW_SHOW)
	pUpdateWindow.Call(hwnd)

	var m MSG
	for !htmlDlgDone {
		r, _, _ := pGetMessageW.Call(uintptr(unsafe.Pointer(&m)), 0, 0, 0)
		if int32(r) <= 0 {
			break
		}
		pTranslateMessage.Call(uintptr(unsafe.Pointer(&m)))
		pDispatchMessageW.Call(uintptr(unsafe.Pointer(&m)))
	}
	pEnableWindow.Call(hwndMain, 1)
	pSetForegroundWindow.Call(hwndMain)
	if !htmlDlgAccepted {
		return nil, false
	}
	selected := make([]int, 0, len(columns))
	for _, i := range currentColumnOrder() {
		if i >= 0 && i < len(htmlDlgSelection) && htmlDlgSelection[i] {
			selected = append(selected, i)
		}
	}
	return selected, len(selected) > 0
}

func chooseHTMLColumns() ([]int, bool) {
	return chooseExportColumns("HTML")
}

func htmlExport(path string, selected []int) error {
	if len(selected) == 0 {
		return fmt.Errorf("nessun attributo selezionato")
	}
	var b strings.Builder
	b.WriteString("<!doctype html><html lang=\"it\"><head><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width,initial-scale=1\"><title>GoList!</title>")
	b.WriteString("<style>html,body{margin:0;padding:0;background:#000040;color:#fff;font-family:Arial,Helvetica,sans-serif}body{padding:18px 20px 26px}.golist{font-family:'Arial Black',Arial,sans-serif;font-size:42px;line-height:1;color:#004080;font-weight:900;margin:4px 0 10px}.rule{height:1px;background:#ffbf00;width:90%;margin:0 0 12px}.summary{text-align:right;color:#409fff;font-size:12px;line-height:1.5;margin:0 1% 20px}.summary .value{color:#ffbf00}.folder{overflow-wrap:anywhere}.table-wrap{overflow:auto;width:100%}table{border-collapse:collapse;width:98%;font-size:12px}th{color:#ffbf00;text-align:left;padding:6px 8px;border-bottom:1px solid #ffbf00;white-space:nowrap}td{color:#fff;padding:4px 8px;border-bottom:1px dotted rgba(64,159,255,.20);vertical-align:top;white-space:normal;overflow-wrap:anywhere}tr:nth-child(even) td{background:rgba(255,255,255,.015)}.footer{display:flex;justify-content:flex-end;align-items:center;gap:7px;margin-top:34px;opacity:.58;color:#bfc8df;font-size:11px}.footer img{width:30px;height:30px;object-fit:contain}</style></head><body>")
	b.WriteString("<div class=\"golist\">GoList!</div><div class=\"rule\"></div>")
	b.WriteString("<div class=\"summary\"><span class=\"value\">" + strconv.Itoa(len(currentEntries)) + "</span> elementi")
	weight := humanSize(totalListedSize(currentEntries, currentRecursive))
	b.WriteString(" · peso totale <span class=\"value\">" + html.EscapeString(weight) + "</span>")
	if total := formatTotalDuration(totalAudioDuration(currentEntries)); total != "" {
		b.WriteString(" · durata audio <span class=\"value\">" + html.EscapeString(total) + "</span>")
	}
	b.WriteString("<br><span class=\"folder\">" + html.EscapeString(currentFolder) + "</span></div>")
	b.WriteString("<div class=\"table-wrap\"><table><thead><tr>")
	for _, i := range selected {
		b.WriteString("<th>" + html.EscapeString(columns[i].Title) + "</th>")
	}
	b.WriteString("</tr></thead><tbody>")
	for _, e := range currentEntries {
		b.WriteString("<tr>")
		for _, i := range selected {
			b.WriteString("<td>" + html.EscapeString(tableCell(e, i)) + "</td>")
		}
		b.WriteString("</tr>")
	}
	b.WriteString("</tbody></table></div>")
	logo := base64.StdEncoding.EncodeToString(botoloPNG)
	b.WriteString("<div class=\"footer\"><span>ShiduLab 2002 - 2026</span><img alt=\"Botolo\" src=\"data:image/png;base64," + logo + "\"></div>")
	b.WriteString("</body></html>")
	return os.WriteFile(path, []byte(b.String()), 0644)
}

func escTSV(s string) string { return strings.NewReplacer("\t", " ", "\r", " ", "\n", " ").Replace(s) }

func exportHeaders() []string {
	h := make([]string, len(columns))
	for i, c := range columns {
		h[i] = c.Title
	}
	return h
}

func exportRow(e Entry) []string {
	r := make([]string, len(columns))
	for i := range columns {
		r[i] = tableCell(e, i)
	}
	return r
}

func filterColumnOrder(wanted map[int]bool) []int {
	out := make([]int, 0, len(wanted))
	for _, i := range currentColumnOrder() {
		if wanted[i] {
			out = append(out, i)
		}
	}
	return out
}

func presetColumns(preset string) []int {
	switch strings.ToLower(strings.TrimSpace(preset)) {
	case "base":
		return filterColumnOrder(map[int]bool{0: true, 18: true, 19: true, 20: true, 21: true, 22: true})
	case "music", "musica":
		return filterColumnOrder(map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 8: true, 11: true, 13: true})
	case "all", "tutti":
		return currentColumnOrder()
	case "last", "ultimi", "":
		fallthrough
	default:
		return visibleColumnsInOrder()
	}
}

func selectedHeaders(selected []int) []string {
	out := make([]string, 0, len(selected))
	for _, i := range selected {
		if i >= 0 && i < len(columns) {
			out = append(out, columns[i].Title)
		}
	}
	return out
}

func selectedRow(e Entry, selected []int) []string {
	out := make([]string, 0, len(selected))
	for _, i := range selected {
		if i >= 0 && i < len(columns) {
			out = append(out, tableCell(e, i))
		}
	}
	return out
}

func exportSelectedTo(path string, selected []int) error {
	if len(selected) == 0 {
		selected = []int{0}
	}
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".html", ".htm":
		return htmlExport(path, selected)
	case ".txt":
		return os.WriteFile(path, []byte(renderTextColumns(currentFolder, currentEntries, selected)), 0644)
	case ".csv":
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		w := csv.NewWriter(f)
		defer w.Flush()
		_ = w.Write(selectedHeaders(selected))
		for _, e := range currentEntries {
			_ = w.Write(selectedRow(e, selected))
		}
		return w.Error()
	case ".tsv":
		var b strings.Builder
		head := selectedHeaders(selected)
		for i, h := range head {
			if i > 0 {
				b.WriteByte('\t')
			}
			b.WriteString(escTSV(h))
		}
		b.WriteString("\r\n")
		for _, e := range currentEntries {
			row := selectedRow(e, selected)
			for i, v := range row {
				if i > 0 {
					b.WriteByte('\t')
				}
				b.WriteString(escTSV(v))
			}
			b.WriteString("\r\n")
		}
		return os.WriteFile(path, []byte(b.String()), 0644)
	case ".json":
		attrs := selectedHeaders(selected)
		rows := make([]map[string]string, 0, len(currentEntries))
		for _, e := range currentEntries {
			m := make(map[string]string, len(selected))
			for pos, i := range selected {
				if pos < len(attrs) && i >= 0 && i < len(columns) {
					m[attrs[pos]] = tableCell(e, i)
				}
			}
			rows = append(rows, m)
		}
		data, err := json.MarshalIndent(struct {
			Folder     string              `json:"folder"`
			Attributes []string            `json:"attributes"`
			Entries    []map[string]string `json:"entries"`
		}{currentFolder, attrs, rows}, "", "  ")
		if err != nil {
			return err
		}
		return os.WriteFile(path, append(data, '\n'), 0644)
	default:
		return fmt.Errorf("formato non supportato: %s", ext)
	}
}

func exportTo(path string) error {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".json":
		data, err := json.MarshalIndent(struct {
			Folder  string  `json:"folder"`
			Entries []Entry `json:"entries"`
		}{currentFolder, currentEntries}, "", "  ")
		if err != nil {
			return err
		}
		return os.WriteFile(path, append(data, '\n'), 0644)
	case ".csv":
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
		w := csv.NewWriter(f)
		defer w.Flush()
		_ = w.Write(exportHeaders())
		for _, e := range currentEntries {
			_ = w.Write(exportRow(e))
		}
		return w.Error()
	case ".tsv":
		var b strings.Builder
		head := exportHeaders()
		for i, h := range head {
			if i > 0 {
				b.WriteByte('\t')
			}
			b.WriteString(escTSV(h))
		}
		b.WriteString("\r\n")
		for _, e := range currentEntries {
			row := exportRow(e)
			for i, v := range row {
				if i > 0 {
					b.WriteByte('\t')
				}
				b.WriteString(escTSV(v))
			}
			b.WriteString("\r\n")
		}
		return os.WriteFile(path, []byte(b.String()), 0644)
	case ".html", ".htm":
		selected, ok := chooseHTMLColumns()
		if !ok {
			return nil
		}
		return htmlExport(path, selected)
	default:
		return os.WriteFile(path, []byte(renderText(currentFolder, currentEntries)), 0644)
	}
}

func regPathHKCU(key string) (string, error) {
	const prefix = `HKCU\`
	if !strings.HasPrefix(strings.ToUpper(key), strings.ToUpper(prefix)) {
		return "", fmt.Errorf("chiave Registro non supportata: %s", key)
	}
	return key[len(prefix):], nil
}

func regCreateHKCU(key string) (uintptr, error) {
	const (
		HKEY_CURRENT_USER       = uintptr(0x80000001)
		KEY_WRITE               = uintptr(0x20006)
		REG_OPTION_NON_VOLATILE = uintptr(0)
	)
	rel, err := regPathHKCU(key)
	if err != nil {
		return 0, err
	}
	var h uintptr
	r, _, _ := pRegCreateKeyExW.Call(
		HKEY_CURRENT_USER,
		uintptr(unsafe.Pointer(wptr(rel))),
		0, 0,
		REG_OPTION_NON_VOLATILE,
		KEY_WRITE,
		0,
		uintptr(unsafe.Pointer(&h)),
		0,
	)
	if r != 0 {
		return 0, fmt.Errorf("RegCreateKeyExW %s: errore %d", key, r)
	}
	return h, nil
}

func regSetStringHKCU(key, name, value string) error {
	const REG_SZ = uintptr(1)
	h, err := regCreateHKCU(key)
	if err != nil {
		return err
	}
	defer pRegCloseKey.Call(h)

	val := syscall.StringToUTF16(value)
	var namePtr uintptr
	if name != "" {
		namePtr = uintptr(unsafe.Pointer(wptr(name)))
	}
	var dataPtr uintptr
	if len(val) > 0 {
		dataPtr = uintptr(unsafe.Pointer(&val[0]))
	}
	r, _, _ := pRegSetValueExW.Call(
		h,
		namePtr,
		0,
		REG_SZ,
		dataPtr,
		uintptr(len(val)*2),
	)
	if r != 0 {
		return fmt.Errorf("RegSetValueExW %s\\%s: errore %d", key, name, r)
	}
	return nil
}

func regDeleteTreeHKCU(key string) error {
	const HKEY_CURRENT_USER = uintptr(0x80000001)
	rel, err := regPathHKCU(key)
	if err != nil {
		return err
	}
	r, _, _ := pRegDeleteTreeW.Call(HKEY_CURRENT_USER, uintptr(unsafe.Pointer(wptr(rel))))
	// ERROR_FILE_NOT_FOUND (2) e ERROR_PATH_NOT_FOUND (3): delete-if-exists riuscito.
	if r != 0 && r != 2 && r != 3 {
		return fmt.Errorf("RegDeleteTreeW %s: errore %d", key, r)
	}
	return nil
}

// Compatibilità interna con le chiamate esistenti, ma senza lanciare reg.exe.
// Evitiamo così decine di processi console durante la creazione del menu a cascata.
func runReg(args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("comando Registro incompleto")
	}
	switch strings.ToLower(args[0]) {
	case "delete":
		return regDeleteTreeHKCU(args[1])
	case "add":
		key := args[1]
		name := ""
		value := ""
		hasValue := false
		for i := 2; i < len(args); i++ {
			switch strings.ToLower(args[i]) {
			case "/ve":
				name = ""
				hasValue = true
			case "/v":
				if i+1 >= len(args) {
					return fmt.Errorf("/v senza nome valore")
				}
				name = args[i+1]
				hasValue = true
				i++
			case "/d":
				if i+1 >= len(args) {
					return fmt.Errorf("/d senza dati")
				}
				value = args[i+1]
				i++
			}
		}
		if hasValue {
			return regSetStringHKCU(key, name, value)
		}
		h, err := regCreateHKCU(key)
		if err == nil {
			pRegCloseKey.Call(h)
		}
		return err
	default:
		return fmt.Errorf("operazione Registro non supportata: %s", args[0])
	}
}

func regMenuLabel(key, label string) error {
	return runReg("add", key, "/v", "MUIVerb", "/d", label, "/f")
}

func regCascade(key, label string) (string, error) {
	if err := regMenuLabel(key, label); err != nil {
		return "", err
	}
	// Menu statico per-user: SubCommands vuoto + sottochiave Shell.
	// È più semplice per Explorer e, soprattutto, evita che il nodo padre
	// venga interpretato come un verbo eseguibile privo di associazione.
	if err := runReg("add", key, "/v", "SubCommands", "/d", "", "/f"); err != nil {
		return "", err
	}
	shellKey := key + `\Shell`
	if err := runReg("add", shellKey, "/f"); err != nil {
		return "", err
	}
	return shellKey, nil
}

func regCommandItem(key, label, command string) error {
	if err := regMenuLabel(key, label); err != nil {
		return err
	}
	return runReg("add", key+`\command`, "/ve", "/d", command, "/f")
}

func addContextTree(base, folderToken, exePath, icon string) error {
	// Ripartiamo puliti: converte anche la vecchia voce singola nel nuovo sottomenu.
	_ = runReg("delete", base, "/f")
	rootShell, err := regCascade(base, "GoList!")
	if err != nil {
		return err
	}
	if err := runReg("add", base, "/v", "Icon", "/d", icon, "/f"); err != nil {
		return err
	}

	openKey := rootShell + `\00Open`
	openCmd := `"` + exePath + `" --folder "` + folderToken + `"`
	if err := regCommandItem(openKey, "Apri GoList!", openCmd); err != nil {
		return err
	}
	_ = runReg("add", openKey, "/v", "Icon", "/d", icon, "/f")

	exportKey := rootShell + `\10Export`
	exportShell, err := regCascade(exportKey, "Esporta")
	if err != nil {
		return err
	}

	formats := []struct {
		Key, Label, Arg string
	}{
		{"10HTML", "HTML", "html"},
		{"20TXT", "TXT", "txt"},
		{"30TSV", "TSV", "tsv"},
		{"40CSV", "CSV", "csv"},
		{"50JSON", "JSON", "json"},
	}
	presets := []struct {
		Key, Label, Arg string
	}{
		{"10Last", "Ultimi attributi usati", "last"},
		{"20Base", "Base", "base"},
		{"30Music", "Musica", "music"},
		{"40All", "Tutti", "all"},
		{"50Choose", "Scegli attributi...", "choose"},
	}

	for _, f := range formats {
		formatKey := exportShell + `\` + f.Key
		formatShell, err := regCascade(formatKey, f.Label)
		if err != nil {
			return err
		}
		for _, p := range presets {
			itemKey := formatShell + `\` + p.Key
			cmd := `"` + exePath + `" --folder "` + folderToken + `" --export ` + f.Arg + ` --preset ` + p.Arg
			if err := regCommandItem(itemKey, p.Label, cmd); err != nil {
				return err
			}
		}
	}
	return nil
}

func addContextMenu() error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	exePath, _ = filepath.Abs(exePath)
	folderBase := `HKCU\Software\Classes\Directory\shell\GoList`
	bgBase := `HKCU\Software\Classes\Directory\Background\shell\GoList`
	iconPath := ensureContextIcon()
	icon := `"` + exePath + `"`
	if iconPath != "" {
		icon = `"` + iconPath + `"`
	}
	if err := addContextTree(folderBase, "%1", exePath, icon); err != nil {
		return err
	}
	if err := addContextTree(bgBase, "%V", exePath, icon); err != nil {
		return err
	}
	return nil
}

func removeContextMenu() error {
	_ = runReg("delete", `HKCU\Software\Classes\Directory\shell\GoList`, "/f")
	_ = runReg("delete", `HKCU\Software\Classes\Directory\Background\shell\GoList`, "/f")
	return nil
}

func runStartupExport() {
	format := strings.ToLower(strings.TrimSpace(startupExportFormat))
	if format == "" || currentFolder == "" || len(currentEntries) == 0 {
		return
	}
	ext := "." + strings.TrimPrefix(format, ".")
	switch ext {
	case ".html", ".txt", ".tsv", ".csv", ".json":
	default:
		msgBox("Formato di esportazione non riconosciuto: "+format, appTitle, MB_OK|MB_ICONERROR)
		return
	}

	var selected []int
	if strings.EqualFold(startupExportPreset, "choose") {
		var ok bool
		selected, ok = chooseExportColumns(strings.ToUpper(strings.TrimPrefix(ext, ".")))
		if !ok {
			return
		}
	} else {
		selected = presetColumns(startupExportPreset)
	}
	if len(selected) == 0 {
		selected = []int{0}
	}

	path := saveDialogFor(ext)
	if path == "" {
		return
	}
	if err := exportSelectedTo(path, selected); err != nil {
		msgBox("Esportazione fallita:\r\n"+err.Error(), appTitle, MB_OK|MB_ICONERROR)
		return
	}
	msgBox("Lista esportata in "+strings.ToUpper(strings.TrimPrefix(ext, "."))+".", appTitle, MB_OK|MB_ICONINFORMATION)
}

func onSize() {
	var rc RECT
	pGetClientRect.Call(hwndMain, uintptr(unsafe.Pointer(&rc)))
	w, h := rc.Right-rc.Left, rc.Bottom-rc.Top
	if w < 760 {
		w = 760
	}
	if h < 400 {
		h = 400
	}
	pMoveWindow.Call(hwndPath, 12, 12, uintptr(w-390), 26, 1)
	getDlg := user32.NewProc("GetDlgItem")
	move := func(id int, x, y, cw, ch int32) {
		hw, _, _ := getDlg.Call(hwndMain, uintptr(id))
		pMoveWindow.Call(hw, uintptr(x), uintptr(y), uintptr(cw), uintptr(ch), 1)
	}
	move(IDBrowse, w-368, 12, 92, 26)
	move(IDList, w-268, 12, 92, 26)
	move(IDRecursive, w-168, 15, 155, 24)
	move(IDCopy, 12, 48, 100, 28)
	move(IDDeleteList, 120, 48, 110, 28)
	move(IDSave, 238, 48, 100, 28)
	move(IDAddContext, 346, 48, 180, 28)
	move(IDRemoveContext, 534, 48, 195, 28)
	pMoveWindow.Call(hwndTable, 12, 86, uintptr(w-24), uintptr(h-120), 1)
	hintW := int32(360)
	hintH := int32(28)
	hintX := (w - hintW) / 2
	hintY := 86 + (h-120-hintH)/2
	pMoveWindow.Call(hwndHint, uintptr(hintX), uintptr(hintY), uintptr(hintW), uintptr(hintH), 1)
	brandW := int32(190)
	pMoveWindow.Call(hwndStatus, 14, uintptr(h-28), uintptr(w-brandW-42), 18, 1)
	pMoveWindow.Call(hwndBrand, uintptr(w-brandW-14), uintptr(h-28), uintptr(brandW), 18, 1)
}

func wndProc(hwnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case WM_CREATE:
		hwndMain = hwnd
		hwndPath = createControl("EDIT", "", WS_CHILD|WS_VISIBLE|WS_BORDER|WS_TABSTOP|ES_AUTOHSCROLL, 12, 12, 420, 26, hwnd, IDPath)
		createControl("BUTTON", "Sfoglia…", WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_PUSHBUTTON, 440, 12, 92, 26, hwnd, IDBrowse)
		createControl("BUTTON", "GoList!", WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_PUSHBUTTON, 540, 12, 92, 26, hwnd, IDList)
		hwndRecursive = createControl("BUTTON", "Sottocartelle", WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_AUTOCHECKBOX, 640, 15, 155, 24, hwnd, IDRecursive)
		createControl("BUTTON", "Copia", WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_PUSHBUTTON, 12, 48, 100, 28, hwnd, IDCopy)
		createControl("BUTTON", "Delete list", WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_PUSHBUTTON, 120, 48, 110, 28, hwnd, IDDeleteList)
		createControl("BUTTON", "Esporta ▼", WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_PUSHBUTTON, 238, 48, 100, 28, hwnd, IDSave)
		createControl("BUTTON", "Aggiungi menu Windows", WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_PUSHBUTTON, 346, 48, 180, 28, hwnd, IDAddContext)
		createControl("BUTTON", "Rimuovi menu Windows", WS_CHILD|WS_VISIBLE|WS_TABSTOP|BS_PUSHBUTTON, 534, 48, 195, 28, hwnd, IDRemoveContext)
		hwndTable = createControl("SysListView32", "", WS_CHILD|WS_VISIBLE|WS_BORDER|WS_TABSTOP|WS_VSCROLL|WS_HSCROLL|LVS_REPORT|LVS_SHOWSELALWAYS, 12, 86, 820, 440, hwnd, IDTable)
		hwndHint = createControl("STATIC", "Trascina qui una cartella", WS_CHILD|WS_VISIBLE|SS_CENTER, 240, 286, 360, 28, hwnd, IDHint)
		hwndStatus = createControl("STATIC", "Pronto.", WS_CHILD|WS_VISIBLE, 14, 530, 600, 18, hwnd, IDStatus)
		hwndBrand = createControl("STATIC", "ShiduLab 2002 - 2026", WS_CHILD|WS_VISIBLE|SS_RIGHT, 650, 530, 190, 18, hwnd, IDBrand)
		setupTable()
		pDragAcceptFiles.Call(hwnd, 1)
		if currentFolder != "" {
			setText(hwndPath, currentFolder)
			doList()
		}
		return 0

	case WM_COMMAND:
		switch int(loword(wParam)) {
		case IDBrowse:
			if p := browseFolder(); p != "" {
				setText(hwndPath, p)
				doList()
			}
		case IDList:
			doList()
		case IDCopy:
			if currentText == "" {
				doList()
			}
			text := renderText(currentFolder, currentEntries)
			if err := copyClipboard(text); err != nil {
				msgBox(err.Error(), appTitle, MB_OK|MB_ICONERROR)
			} else {
				msgBox("Lista copiata negli appunti.", appTitle, MB_OK|MB_ICONINFORMATION)
			}
		case IDDeleteList:
			currentEntries = nil
			currentText = ""
			currentSortColumn = -1
			pSendMessageW.Call(hwndTable, LVM_DELETEALLITEMS, 0, 0)
			pShowWindow.Call(hwndHint, SW_SHOW)
			setText(hwndStatus, "Lista cancellata. Cartella selezionata mantenuta.")
		case IDSave:
			if currentText == "" {
				doList()
			}
			if currentText != "" {
				ext := chooseExportFormat()
				if ext == "" {
					break
				}
				var selected []int
				if ext == ".html" {
					var ok bool
					selected, ok = chooseHTMLColumns()
					if !ok {
						break
					}
				}
				p := saveDialogFor(ext)
				if p == "" {
					break
				}
				var err error
				if ext == ".html" {
					err = htmlExport(p, selected)
				} else {
					err = exportTo(p)
				}
				if err != nil {
					msgBox("Esportazione fallita:\r\n"+err.Error(), appTitle, MB_OK|MB_ICONERROR)
				} else {
					msgBox("Lista esportata in "+strings.ToUpper(strings.TrimPrefix(ext, "."))+".", appTitle, MB_OK|MB_ICONINFORMATION)
				}
			}
		case IDAddContext:
			if contextMenuBusy {
				break
			}
			contextMenuBusy = true
			getDlg := user32.NewProc("GetDlgItem")
			btnAdd, _, _ := getDlg.Call(hwnd, IDAddContext)
			btnRemove, _, _ := getDlg.Call(hwnd, IDRemoveContext)
			pEnableWindow.Call(btnAdd, 0)
			pEnableWindow.Call(btnRemove, 0)
			setText(hwndStatus, "Registro il menu Windows…")

			err := addContextMenu()

			pEnableWindow.Call(btnAdd, 1)
			pEnableWindow.Call(btnRemove, 1)
			contextMenuBusy = false
			if err != nil {
				setText(hwndStatus, "Registrazione menu Windows fallita.")
				msgBox("Non riesco ad aggiungere GoList! al menu contestuale:\r\n"+err.Error(), appTitle, MB_OK|MB_ICONERROR)
			} else {
				setText(hwndStatus, "Menu Windows registrato.")
				msgBox("GoList! aggiunto al menu contestuale con scorciatoie di esportazione.\r\n\r\nWindows 11 può mostrarlo sotto “Mostra altre opzioni”.", appTitle, MB_OK|MB_ICONINFORMATION)
			}
		case IDRemoveContext:
			if contextMenuBusy {
				break
			}
			contextMenuBusy = true
			_ = removeContextMenu()
			contextMenuBusy = false
			msgBox("Voce GoList! rimossa dal menu contestuale.", appTitle, MB_OK|MB_ICONINFORMATION)
		}
		return 0

	case WM_NOTIFY:
		if lParam != 0 {
			hdr := (*NMHDR)(unsafe.Pointer(lParam))
			if hdr.HwndFrom == hwndTable && int32(hdr.Code) == LVN_COLUMNCLICK {
				nmlv := (*NMLISTVIEW)(unsafe.Pointer(lParam))
				sortCurrent(int(nmlv.ISubItem))
				return 0
			}
		}

	case WM_DROPFILES:
		hdrop := wParam
		n, _, _ := pDragQueryFileW.Call(hdrop, 0, 0, 0)
		if n > 0 {
			buf := make([]uint16, n+1)
			pDragQueryFileW.Call(hdrop, 0, uintptr(unsafe.Pointer(&buf[0])), n+1)
			p := syscall.UTF16ToString(buf)
			if st, err := os.Stat(p); err == nil {
				if !st.IsDir() {
					p = filepath.Dir(p)
				}
				setText(hwndPath, p)
				doList()
			}
		}
		pDragFinish.Call(hdrop)
		return 0

	case WM_SIZE:
		onSize()
		return 0

	case WM_DESTROY:
		saveTableLayout()
		pPostQuitMessage.Call(0)
		return 0
	}
	r, _, _ := pDefWindowProcW.Call(hwnd, uintptr(msg), wParam, lParam)
	return r
}

func main() {
	runtime.LockOSThread()
	pCoInitializeEx.Call(0, 2)
	defer ole32.NewProc("CoUninitialize").Call()

	icc := INITCOMMONCONTROLSEX{DwSize: uint32(unsafe.Sizeof(INITCOMMONCONTROLSEX{})), DwICC: ICC_LISTVIEW_CLASSES}
	pInitCommonControlsEx.Call(uintptr(unsafe.Pointer(&icc)))

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--folder":
			if i+1 < len(args) {
				currentFolder = args[i+1]
				i++
			}
		case "--export":
			if i+1 < len(args) {
				startupExportFormat = args[i+1]
				i++
			}
		case "--preset":
			if i+1 < len(args) {
				startupExportPreset = args[i+1]
				i++
			}
		}
	}

	// Give Windows a stable application identity for taskbar grouping/icon handling.
	pSetCurrentProcessExplicitAppUserModelID.Call(uintptr(unsafe.Pointer(wptr("ShiduLab.GoList"))))

	hInst, _, _ := pGetModuleHandleW.Call(0)
	className := wptr("GoList2WindowClass")
	cursor, _, _ := pLoadCursorW.Call(0, IDC_ARROW)
	hIconBig = createIconFromICO(botoloICO, 32)
	hIconSmall = createIconFromICO(botoloICO, 16)
	wc := WNDCLASSEX{CbSize: uint32(unsafe.Sizeof(WNDCLASSEX{})), LpfnWndProc: syscall.NewCallback(wndProc), HInstance: hInst, HIcon: hIconBig, HCursor: cursor, HbrBackground: COLOR_WINDOW + 1, LpszClassName: className, HIconSm: hIconSmall}
	if r, _, _ := pRegisterClassExW.Call(uintptr(unsafe.Pointer(&wc))); r == 0 {
		panic("RegisterClassExW")
	}
	hwnd, _, _ := pCreateWindowExW.Call(0, uintptr(unsafe.Pointer(className)), uintptr(unsafe.Pointer(wptr(appTitle))), WS_OVERLAPPED|WS_CAPTION|WS_SYSMENU|WS_THICKFRAME|WS_MINIMIZEBOX|WS_MAXIMIZEBOX|WS_VISIBLE, CW_USEDEFAULT, CW_USEDEFAULT, 980, 640, 0, 0, hInst, 0)
	if hwnd == 0 {
		panic("CreateWindowExW")
	}
	if hIconBig != 0 {
		pSendMessageW.Call(hwnd, WM_SETICON, ICON_BIG, hIconBig)
	}
	if hIconSmall != 0 {
		pSendMessageW.Call(hwnd, WM_SETICON, ICON_SMALL, hIconSmall)
	}
	pShowWindow.Call(hwnd, SW_SHOW)
	pUpdateWindow.Call(hwnd)
	if startupExportFormat != "" {
		runStartupExport()
	}
	var m MSG
	for {
		r, _, _ := pGetMessageW.Call(uintptr(unsafe.Pointer(&m)), 0, 0, 0)
		if int32(r) <= 0 {
			break
		}
		pTranslateMessage.Call(uintptr(unsafe.Pointer(&m)))
		pDispatchMessageW.Call(uintptr(unsafe.Pointer(&m)))
	}
	if hIconSmall != 0 {
		pDestroyIcon.Call(hIconSmall)
	}
	if hIconBig != 0 && hIconBig != hIconSmall {
		pDestroyIcon.Call(hIconBig)
	}
}
