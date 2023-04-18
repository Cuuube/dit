package cli

import (
	"fmt"
	"strings"
)

type LoadingAnimation interface {
	Print()
	Clear()
	Reset()
}

type Progress struct {
	Total   int
	Current int
	// Percent float64
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

func NewTextWithStringsLoopLoading(text string, strs []string) *TextWithStringsLoopLoading {
	return &TextWithStringsLoopLoading{text: text, strs: strs}
}

var _ (LoadingAnimation) = (*TextWithStringsLoopLoading)(nil)

// TextWithStringsLoopLoading 文字+省略号的简单加载动画。例如：Loading...
type TextWithStringsLoopLoading struct {
	text   string
	cursor int
	strs   []string
}

func (obj *TextWithStringsLoopLoading) Print() {
	if obj.cursor >= len(obj.strs) {
		obj.cursor = 0
	}
	curStr := obj.strs[obj.cursor]
	obj.cursor += 1

	Println(obj.text + curStr)
}

func (obj *TextWithStringsLoopLoading) Clear() {
	Print(CursorUpLine1 + ClearLine)
}

func (obj *TextWithStringsLoopLoading) Reset() {
	obj.cursor = 0
}

func NewTextWithProgressLoading(text string) *TextWithProgressLoading {
	return &TextWithProgressLoading{
		text:     text,
		Progress: &Progress{},
	}
}

var _ (LoadingAnimation) = (*TextWithProgressLoading)(nil)

// TextWithProgressLoading 文字+进度条的简单加载动画。例如：Loading[=========----]
type TextWithProgressLoading struct {
	text     string
	Progress *Progress
}

func (obj *TextWithProgressLoading) Print() {
	fmt.Print(obj.text)
	cur := obj.Progress.Current
	persent := cur * 100 / obj.Progress.Total
	if persent >= 100 {
		persent = 100
	}
	var sb strings.Builder
	sb.WriteRune('[')
	for i := 1; i <= 10; i++ {
		if persent >= i*10 {
			sb.WriteRune('=')
		} else {
			sb.WriteRune('-')
		}
	}
	sb.WriteRune(']')
	fmt.Println(sb.String())
}

func (obj *TextWithProgressLoading) Clear() {
	Print(CursorUpLine1 + ClearLine)
}

func (obj *TextWithProgressLoading) Reset() {
	obj.Progress.Current = 0
}
