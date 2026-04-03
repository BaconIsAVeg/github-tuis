package diffstyles

import "github.com/charmbracelet/lipgloss"

type DiffPalette struct {
	Add        lipgloss.Style
	Delete     lipgloss.Style
	Header     lipgloss.Style
	FileHeader lipgloss.Style
}

type diffColors struct {
	diffAdd        lipgloss.Color
	diffDelete     lipgloss.Color
	diffHeader     lipgloss.Color
	diffFileHeader lipgloss.Color
}

func NewDiffPalette(darkBackground bool) *DiffPalette {
	if darkBackground {
		return newDiffPalette(darkDiffColors())
	}
	return newDiffPalette(lightDiffColors())
}

func darkDiffColors() diffColors {
	return diffColors{
		diffAdd:        lipgloss.Color("34"),
		diffDelete:     lipgloss.Color("160"),
		diffHeader:     lipgloss.Color("36"),
		diffFileHeader: lipgloss.Color("15"),
	}
}

func lightDiffColors() diffColors {
	return diffColors{
		diffAdd:        lipgloss.Color("28"),
		diffDelete:     lipgloss.Color("124"),
		diffHeader:     lipgloss.Color("24"),
		diffFileHeader: lipgloss.Color("0"),
	}
}

func newDiffPalette(c diffColors) *DiffPalette {
	return &DiffPalette{
		Add: lipgloss.NewStyle().
			Foreground(c.diffAdd),
		Delete: lipgloss.NewStyle().
			Foreground(c.diffDelete),
		Header: lipgloss.NewStyle().
			Foreground(c.diffHeader).
			Bold(true),
		FileHeader: lipgloss.NewStyle().
			Foreground(c.diffFileHeader).
			Bold(true),
	}
}
