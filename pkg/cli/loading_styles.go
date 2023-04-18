package cli

import (
	"strings"
)

type LoadingAnimation interface {
	Print()
	Clear()
	Reset()
}

func NewTextWithDotLoading(text string) *TextWithDotLoading {
	return &TextWithDotLoading{text: text}
}

var _ (LoadingAnimation) = (*TextWithDotLoading)(nil)

// TextWithDotLoading 文字+省略号的简单加载动画。例如：Loading...
type TextWithDotLoading struct {
	text   string
	cursor uint8
}

func (obj *TextWithDotLoading) Print() {
	var suffix strings.Builder
	var i uint8 = 0
	for ; i <= obj.cursor; i++ {
		suffix.WriteRune('.')
	}
	if i >= 6 {
		i = 0
	}
	obj.cursor = i

	Println(obj.text + suffix.String())
}

func (obj *TextWithDotLoading) Clear() {
	Print(CursorUpLine1 + ClearLine)
}

func (obj *TextWithDotLoading) Reset() {
	obj.cursor = 0
}
