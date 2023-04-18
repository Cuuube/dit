package cli

const (
	ResetStyle = "\033[0m"

	StyleNormal      = "\033[0m"
	StyleBold        = "\033[1m"
	StyleDark        = "\033[2m"
	StyleItalic      = "\033[3m"
	StyleUnderline   = "\033[4m"
	StyleShine       = "\033[5m"
	StyleRevertColor = "\033[7m"
	StyleInvisible   = "\033[8m"

	ColorCodeBlack   = "0"
	ColorCodeRed     = "1"
	ColorCodeGreen   = "2"
	ColorCodeYellow  = "3"
	ColorCodeBlue    = "4"
	ColorCodeMagenta = "5"
	ColorCodeCyan    = "6"
	ColorCodeWhite   = "7"

	ColorBlack   = "\033[30m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorWhite   = "\033[37m"

	BgColorBlack   = "\033[40m"
	BgColorRed     = "\033[41m"
	BgColorGreen   = "\033[42m"
	BgColorYellow  = "\033[43m"
	BgColorBlue    = "\033[44m"
	BgColorMagenta = "\033[45m"
	BgColorCyan    = "\033[46m"
	BgColorWhite   = "\033[47m"

	ColorFormat     = "\033[3%dm"
	BgColorFormat   = "\033[4%dm"
	FullColorFormat = "\033[3%d;4%dm"

	CursorUpLine1     = "\033[1A"
	CursorUpLine2     = "\033[2A"
	CursorUpFormat    = "\033[%dA"
	CursorDownFormat  = "\033[%dB"
	CursorRightFormat = "\033[%dC"
	CursorLeftFormat  = "\033[%dD"

	CursorSavePosition = "\033[s"
	CursorLoadPosition = "\033[u"

	CursorSetPositionFormat = "\033[%d;%dH" // y, x

	ClearLine   = "\033[K"
	ClearScreen = "\033[2J"
)
