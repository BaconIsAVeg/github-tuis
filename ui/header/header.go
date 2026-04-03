package header

import (
	"github.com/BaconIsAVeg/github-tuis/ui/styles"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	leftContent   string
	middleContent string
	rightContent  string
	textInput     textinput.Model
	editing       bool
	width         int
	styles        *styles.Palette
}

func New(s *styles.Palette) Model {
	ti := textinput.New()
	ti.CharLimit = 500
	ti.Prompt = ""

	return Model{
		styles:    s,
		textInput: ti,
	}
}

func (m *Model) SetLeft(s string) {
	m.leftContent = s
}

func (m *Model) SetMiddle(s string) {
	m.middleContent = s
}

func (m *Model) SetRight(s string) {
	m.rightContent = s
}

func (m *Model) SetWidth(width int) {
	m.width = width
}

func (m *Model) StartEditing(currentText string) {
	m.textInput.SetValue(currentText)
	m.textInput.Focus()
	m.editing = true
}

func (m *Model) StopEditing() string {
	m.editing = false
	m.textInput.Blur()
	return m.textInput.Value()
}

func (m *Model) CancelEditing() {
	m.editing = false
	m.textInput.Blur()
}

func (m Model) IsEditing() bool {
	return m.editing
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if m.editing {
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) View() string {
	leftContent := m.styles.HeaderTitle.Render(m.leftContent)
	rightContent := m.styles.HeaderText.Padding(0, 1).Render(m.rightContent)

	leftWidth := lipgloss.Width(leftContent)
	rightWidth := lipgloss.Width(rightContent)
	middleWidth := m.width - leftWidth - rightWidth

	if middleWidth < 1 {
		middleWidth = 1
	}

	var middleContent string
	if m.editing {
		m.textInput.Width = middleWidth - 1
		middleContent = lipgloss.NewStyle().Width(middleWidth).Render(" " + m.textInput.View())
	} else {
		middleContent = m.styles.StatusBar.Width(middleWidth).Render(m.middleContent)
	}

	header := lipgloss.JoinHorizontal(lipgloss.Top, leftContent, middleContent, rightContent)

	if lipgloss.Width(header) < m.width {
		header = m.styles.Header.Width(m.width).Render(header)
	}

	return header
}
