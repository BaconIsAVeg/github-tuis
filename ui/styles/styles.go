package styles

import "github.com/charmbracelet/lipgloss"

type Palette struct {
	Header           lipgloss.Style
	HeaderTitle      lipgloss.Style
	HeaderText       lipgloss.Style
	StatusBar        lipgloss.Style
	StatusMode       lipgloss.Style
	StatusKey        lipgloss.Style
	StatusDesc       lipgloss.Style
	StatusSep        lipgloss.Style
	StatusPass       lipgloss.Style
	StatusFail       lipgloss.Style
	StatusPending    lipgloss.Style
	PreviewBorder    lipgloss.Style
	PreviewContent   lipgloss.Style
	Notification     lipgloss.Style
	NotificationInfo lipgloss.Style
	NotificationWarn lipgloss.Style
	SecondaryBg      lipgloss.Color
	ShadowFg         lipgloss.Color
}

type colors struct {
	primaryBg      lipgloss.Color
	secondaryBg    lipgloss.Color
	primaryFg      lipgloss.Color
	dimFg          lipgloss.Color
	popFg          lipgloss.Color
	notificationBg lipgloss.Color
	notificationFg lipgloss.Color
	infoBg         lipgloss.Color
	warnBg         lipgloss.Color
	shadowFg       lipgloss.Color
	statusPass     lipgloss.Color
	statusFail     lipgloss.Color
	statusPending  lipgloss.Color
}

func NewPalette(darkBackground bool) *Palette {
	if darkBackground {
		return newPalette(darkColors())
	}
	return newPalette(lightColors())
}

func darkColors() colors {
	return colors{
		primaryBg:      lipgloss.Color("97"),
		secondaryBg:    lipgloss.Color("234"),
		primaryFg:      lipgloss.Color("15"),
		dimFg:          lipgloss.Color("244"),
		popFg:          lipgloss.Color("178"),
		notificationBg: lipgloss.Color("28"),
		notificationFg: lipgloss.Color("15"),
		infoBg:         lipgloss.Color("25"),
		warnBg:         lipgloss.Color("124"),
		shadowFg:       lipgloss.Color("#333333"),
		statusPass:     lipgloss.Color("2"),
		statusFail:     lipgloss.Color("1"),
		statusPending:  lipgloss.Color("4"),
	}
}

func lightColors() colors {
	return colors{
		primaryBg:      lipgloss.Color("97"),
		secondaryBg:    lipgloss.Color("254"),
		primaryFg:      lipgloss.Color("0"),
		dimFg:          lipgloss.Color("242"),
		popFg:          lipgloss.Color("172"),
		notificationBg: lipgloss.Color("28"),
		notificationFg: lipgloss.Color("15"),
		infoBg:         lipgloss.Color("25"),
		warnBg:         lipgloss.Color("124"),
		shadowFg:       lipgloss.Color("#cccccc"),
		statusPass:     lipgloss.Color("2"),
		statusFail:     lipgloss.Color("1"),
		statusPending:  lipgloss.Color("4"),
	}
}

func newPalette(c colors) *Palette {
	return &Palette{
		Header: lipgloss.NewStyle().
			Background(c.primaryBg).
			Foreground(c.primaryFg).
			Padding(0, 1),
		HeaderTitle: lipgloss.NewStyle().
			Background(c.primaryBg).
			Foreground(c.primaryFg).
			Bold(true).
			Padding(0, 1),
		HeaderText: lipgloss.NewStyle().
			Background(c.primaryBg).
			Foreground(c.primaryFg),
		StatusBar: lipgloss.NewStyle().
			Background(c.secondaryBg).
			Foreground(c.primaryFg).
			Padding(0, 1),
		StatusMode: lipgloss.NewStyle().
			Background(c.primaryBg).
			Foreground(c.primaryFg).
			Bold(true).
			Padding(0, 1),
		StatusKey: lipgloss.NewStyle().
			Background(c.secondaryBg).
			Foreground(c.popFg).
			Bold(true),
		StatusDesc: lipgloss.NewStyle().
			Background(c.secondaryBg).
			Foreground(c.dimFg),
		StatusSep: lipgloss.NewStyle().
			Background(c.secondaryBg),
		StatusPass: lipgloss.NewStyle().
			Foreground(c.statusPass).
			Bold(true),
		StatusFail: lipgloss.NewStyle().
			Foreground(c.statusFail).
			Bold(true),
		StatusPending: lipgloss.NewStyle().
			Foreground(c.statusPending).
			Bold(true),
		PreviewBorder: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(c.dimFg).
			BorderBottom(false),
		PreviewContent: lipgloss.NewStyle().
			Foreground(c.primaryFg),
		Notification: lipgloss.NewStyle().
			Background(c.notificationBg).
			Foreground(c.notificationFg).
			Padding(0, 1).
			Bold(true),
		NotificationInfo: lipgloss.NewStyle().
			Background(c.infoBg).
			Foreground(c.notificationFg).
			Padding(0, 1).
			Bold(true),
		NotificationWarn: lipgloss.NewStyle().
			Background(c.warnBg).
			Foreground(c.notificationFg).
			Padding(0, 1).
			Bold(true),
		SecondaryBg: c.secondaryBg,
		ShadowFg:    c.shadowFg,
	}
}
