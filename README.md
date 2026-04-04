# github-tuis

Shared Go library for building GitHub terminal UI applications

---

## Overview

A reusable library providing common UI components and GitHub client utilities for terminal-based applications. Built on the Charmbracelet ecosystem (Bubble Tea, Lipgloss, Bubbles) for consistent, beautiful TUIs.

---

## Install

```bash
go get github.com/BaconIsAVeg/github-tuis
```

---

## Features

- **UI Components**: Pre-built header, status bar, and notification components
- **Styles Palette**: Centralized color scheme with light/dark terminal support
- **GitHub Client**: Unified REST and GraphQL API client with auto-authentication
- **Debug Utilities**: Conditional logging for development and troubleshooting
- **Overlay Helpers**: Utilities for positioning UI elements

---

## Usage

### Import the library

```go
import (
    "github.com/BaconIsAVeg/github-tuis/ui/styles"
    "github.com/BaconIsAVeg/github-tuis/ui/header"
    "github.com/BaconIsAVeg/github-tuis/ui/statusbar"
    "github.com/BaconIsAVeg/github-tuis/ui/notification"
    tuisclient "github.com/BaconIsAVeg/github-tuis/github/client"
)
```

---

### Initialize styles

Create a palette that adapts to the terminal background:

```go
palette := styles.NewPalette(lipgloss.HasDarkBackground())
```

---

### Create a GitHub client

The client automatically detects authentication from environment variables or gh CLI:

```go
client, err := tuisclient.NewClient(tuisclient.ClientOptions{})
if err != nil {
    log.Fatal(err)
}
defer client.Close()

// Use REST API
repos, _, err := client.REST().Repositories.List(ctx, "org", nil)

// Use GraphQL API
var query struct { ... }
err := client.GraphQL().Query("QueryName", &query, nil)
```

---

### Build UI components

Pass the palette to each component:

```go
hdr := header.New(palette)
hdr.SetLeft("My App")
hdr.SetRight("v1.0.0")
hdr.SetWidth(80)

bar := statusbar.New(palette)
bar.SetMode("normal")
bar.SetKeybindings([]statusbar.KeyBinding{
    {Key: "q", Desc: "quit"},
    {Key: "j/k", Desc: "navigate"},
})

notif := notification.New(palette)
```

---

### Display notifications

Show toast messages with different types:

```go
// Success notification (3s default)
cmd := notif.Show("Operation completed")

// Info notification
cmd := notif.ShowInfo("Fetching data...")

// Warning notification
cmd := notif.ShowWarning("Rate limit approaching")

// Custom timeout
cmd := notif.ShowWithTimeout("Saved", notification.TypeSuccess, 5*time.Second)
```

---

### Handle authentication

The client supports multiple authentication methods:

```go
// Option 1: Environment variables
// Set GH_TOKEN or GITHUB_TOKEN

// Option 2: gh CLI configuration
// Automatically detected if gh is configured

// Option 3: Explicit token
client, err := tuisclient.NewClient(tuisclient.ClientOptions{
    Token: "ghp_xxx",
    Host:  "github.com", // or GitHub Enterprise host
})
```

---

### Enable debug logging

Set the `GH_DEBUG` environment variable to enable logging:

```bash
export GH_DEBUG=1
```

Debug logs are written to `ghr-debug.log` in the current directory.

---

## Components

### Styles palette

Centralized color definitions supporting both light and dark terminal backgrounds. All UI components accept a palette pointer for consistent styling.

### Header

Three-section header with optional text input for search or filtering. Supports left, middle, and right content areas.

### Status bar

Mode indicator with keybinding hints. Displays current mode and available keyboard shortcuts.

### Notification

Toast-style notifications with success, info, and warning variants. Auto-hides after configurable timeout.

### Overlay helpers

Utilities for placing UI elements on top of background content with optional shadow effects.

---

## Dependencies

- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/lipgloss` - Styling and layout
- `github.com/charmbracelet/bubbles` - Pre-built components
- `github.com/google/go-github/v82` - GitHub REST API
- `github.com/cli/go-gh/v2` - GitHub CLI SDK (GraphQL)

---

## Projects using this library

- **gh-purview** - Terminal UI for browsing and approving GitHub pull requests

---

## License

MIT License - see [LICENSE](LICENSE) for details.
