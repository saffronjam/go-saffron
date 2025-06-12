package gui

import (
	"github.com/saffronjam/cimgui-go/imgui"
	"go.uber.org/zap/zapcore"
	"strings"
)

type Log struct {
	TextBuffer  strings.Builder
	LineOffsets []int
	LineLevels  []zapcore.Level
}

func NewLog() *Log {
	log := &Log{}
	// Always start with first line offset
	log.LineOffsets = append(log.LineOffsets, 0)
	return log
}

func (l *Log) AddEntry(entry zapcore.Entry) {
	startOffset := l.TextBuffer.Len()
	l.TextBuffer.WriteString(entry.Message)
	l.TextBuffer.WriteByte('\n')
	l.LineLevels = append(l.LineLevels, entry.Level)

	// Record line offsets
	for i := startOffset; i < l.TextBuffer.Len(); i++ {
		if l.TextBuffer.String()[i] == '\n' {
			l.LineOffsets = append(l.LineOffsets, i+1)
		}
	}
}

func (l *Log) Clear() {
	l.TextBuffer.Reset()
	l.LineOffsets = l.LineOffsets[:0]
	l.LineOffsets = append(l.LineOffsets, 0) // Reset to 0
}
func (l *Log) RenderUI() {
	imgui.Begin("Log")

	imgui.SetCursorPos(imgui.Vec2{X: 5, Y: 30})
	if imgui.Button("Clear") {
		l.Clear()
	}

	imgui.Separator()

	PushFont("roboto-mono", 18.0)
	autoscroll := false
	if imgui.BeginChildIDV(imgui.IDStr("scrolling"), imgui.Vec2{X: 0, Y: 0}, imgui.ChildFlagsNone, imgui.WindowFlagsHorizontalScrollbar) {
		imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 0, Y: 0})
		buf := l.TextBuffer.String()
		for i := 0; i < len(l.LineOffsets)-1; i++ {
			switch l.LineLevels[i] {
			case zapcore.DebugLevel:
				// Softer gray for debug
				imgui.TextColored(imgui.Vec4{X: 0.65, Y: 0.65, Z: 0.70, W: 1.0}, "[DEBUG]  ")
			case zapcore.InfoLevel:
				// Muted blue for info
				imgui.TextColored(imgui.Vec4{X: 0.30, Y: 0.60, Z: 0.85, W: 1.0}, "[INFO]   ")
			case zapcore.WarnLevel:
				// Warm amber for warning
				imgui.TextColored(imgui.Vec4{X: 0.95, Y: 0.75, Z: 0.20, W: 1.0}, "[WARN]   ")
			case zapcore.ErrorLevel:
				// Muted red for error
				imgui.TextColored(imgui.Vec4{X: 0.85, Y: 0.33, Z: 0.31, W: 1.0}, "[ERROR]  ")
			case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
				// Deep red for fatal
				imgui.TextColored(imgui.Vec4{X: 0.60, Y: 0.15, Z: 0.18, W: 1.0}, "[FATAL]  ")
			default:
				// Neutral gray for unknown
				imgui.TextColored(imgui.Vec4{X: 0.80, Y: 0.80, Z: 0.80, W: 1.0}, "[?]      ")
			}

			imgui.SameLine()

			line := buf[l.LineOffsets[i]:l.LineOffsets[i+1]]
			line = strings.TrimSuffix(line, "\n")
			imgui.TextUnformatted(line)
		}
		// Autoscroll logic
		if imgui.ScrollY() >= imgui.ScrollMaxY()-5 {
			autoscroll = true
		}
		imgui.PopStyleVar()
		if autoscroll {
			imgui.SetScrollHereYV(1.0)
		}
	}
	PopFont()
	imgui.EndChild()
	imgui.End()
}
