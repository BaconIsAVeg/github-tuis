package statusbar

import (
	"github.com/BaconIsAVeg/github-tuis/ui/styles"
	"github.com/charmbracelet/lipgloss"
)

type KeyBinding struct {
	Key  string
	Desc string
}

type Model struct {
	mode          string
	middleContent string
	keybindings   []KeyBinding
	width         int
	styles        *styles.Palette
}

func New(s *styles.Palette) Model {
	return Model{
		styles: s,
	}
}

func (m *Model) SetMode(mode string) {
	m.mode = mode
}

func (m *Model) SetMiddleContent(s string) {
	m.middleContent = s
}

func (m *Model) SetKeybindings(kb []KeyBinding) {
	m.keybindings = kb
}

func (m *Model) SetWidth(width int) {
	m.width = width
}

func (m Model) Mode() string {
	return m.mode
}

func (m Model) View() string {
	barBg := m.styles.SecondaryBg

	modeContent := m.styles.StatusMode.Render(m.mode)

	keysText := ""
	for i, k := range m.keybindings {
		if i > 0 {
			keysText += m.styles.StatusSep.Render("  ")
		}
		keysText += m.styles.StatusKey.Render(k.Key) + m.styles.StatusDesc.Render(" "+k.Desc)
	}

	keysContent := m.styles.StatusBar.Render(keysText)

	leftWidth := lipgloss.Width(modeContent)
	rightWidth := lipgloss.Width(keysContent)
	middleWidth := max(m.width-leftWidth-rightWidth, 0)

	middleContent := lipgloss.NewStyle().
		Background(barBg).
		Width(middleWidth).
		Render(m.middleContent)

	statusBar := lipgloss.JoinHorizontal(lipgloss.Top, modeContent, middleContent, keysContent)

	if lipgloss.Width(statusBar) < m.width {
		statusBar = lipgloss.NewStyle().Width(m.width).Background(barBg).Render(statusBar)
	}

	return statusBar
}
