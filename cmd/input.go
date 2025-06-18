package cmd

import (
  "fmt"
  "os"
  "strings"

  "github.com/charmbracelet/bubbles/textinput"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/charmbracelet/lipgloss"
  "github.com/spf13/cobra"
)

var (
  placeholder       string
  placeholderStyle  string
  inputStyle        string
)

func init() {
  inputCmd.Flags().StringVar(&placeholder, "placeholder", "", "Placeholder text")
  inputCmd.Flags().StringVar(&placeholderStyle, "placeholder-style", "italic", "Comma-separated placeholder style: bold,italic,faint,underline")
  inputCmd.Flags().StringVar(&inputStyle, "input-style", "bold", "Comma-separated input style")

  rootCmd.AddCommand(inputCmd)
}

var inputCmd = &cobra.Command{
  Use:   "input",
  Short: "Prompt for a single line of text",
  Run: func(cmd *cobra.Command, args []string) {
    ti := textinput.New()
    ti.Placeholder = placeholder
    ti.Focus()
    ti.CharLimit = 512
    ti.Width = 40

    // Apply styles
    ti.PromptStyle = lipgloss.NewStyle().Bold(true)
    ti.TextStyle = parseStyle(inputStyle)
    ti.PlaceholderStyle = parseStyle(placeholderStyle)

    p := tea.NewProgram(inputModel{
      input:       ti,
      placeholder: placeholder,
    })

    if m, err := p.Run(); err == nil {
      if model, ok := m.(inputModel); ok && model.done {
        fmt.Println(model.input.Value())
        os.Exit(0)
      }
    }
    os.Exit(1)
  },
}

type inputModel struct {
  input       textinput.Model
  placeholder string
  done        bool
}

func (m inputModel) Init() tea.Cmd {
  return textinput.Blink
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "esc":
      return m, tea.Quit
    case "enter":
      m.done = true
      return m, tea.Quit
    case "tab":
      if strings.TrimSpace(m.input.Value()) == "" && m.placeholder != "" {
        m.input.SetValue(m.placeholder)
      }
    }
  case tea.WindowSizeMsg:
    m.input.Width = msg.Width - 4
  }
  var cmd tea.Cmd
  m.input, cmd = m.input.Update(msg)
  return m, cmd
}

func (m inputModel) View() string {
  if m.done {
    return ""
  }
  return fmt.Sprintf(
    "\n  %s\n\n  (Enter to submit, Tab to use placeholder, Esc to cancel)\n",
    m.input.View(),
  )
}

// parseStyle converts comma-separated style string into lipgloss.Style
func parseStyle(s string) lipgloss.Style {
  style := lipgloss.NewStyle()
  tokens := strings.Split(s, ",")
  for _, tok := range tokens {
    switch strings.ToLower(strings.TrimSpace(tok)) {
    case "bold":
      style = style.Bold(true)
    case "italic":
      style = style.Italic(true)
    case "faint":
      style = style.Faint(true)
    case "underline":
      style = style.Underline(true)
    case "strikethrough":
      style = style.Strikethrough(true)
    }
  }
  return style
}

// vim:tabstop=2:softtabstop=2:shiftwidth=2:expandtab:filetype=go:
