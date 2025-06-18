package cmd

import (
  "fmt"
  "os"

  // "github.com/charmbracelet/bubbles/selector"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/spf13/cobra"
)

var confirmMessage string

func init() {
  confirmCmd.Flags().StringVar(&confirmMessage, "message", "Are you sure?", "Confirmation prompt message")
  rootCmd.AddCommand(confirmCmd)
}

var confirmCmd = &cobra.Command{
  Use:   "confirm",
  Short: "Prompt for yes/no confirmation",
  Run: func(cmd *cobra.Command, args []string) {
    model := confirmModel{
      options: []string{"Yes", "No"},
      message: confirmMessage,
      cursor:  0,
    }
    p := tea.NewProgram(model)
    if m, err := p.Run(); err == nil {
      if model, ok := m.(confirmModel); ok {
        if model.cancelled {
          os.Exit(130) // 用户取消
        }
        if model.options[model.cursor] == "Yes" {
          fmt.Println("yes")
          os.Exit(0)
        }
      }
    }
    fmt.Println("no")
    os.Exit(1)
  },
}

type confirmModel struct {
  options []string
  message string
  cursor  int
  cancelled bool
}

func (m confirmModel) Init() tea.Cmd {
  return nil
}

func (m confirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c", "esc":
      m.cancelled = true
      return m, tea.Quit
    case "enter":
      return m, tea.Quit
    case "up", "k":
      if m.cursor > 0 {
        m.cursor--
      }
    case "down", "j", "tab":
      if m.cursor < len(m.options)-1 {
        m.cursor++
      } else {
        m.cursor = 0
      }
    }
  }
  return m, nil
}

func (m confirmModel) View() string {
  s := "\n" + m.message + "\n\n"
  for i, option := range m.options {
    cursor := " "
    if i == m.cursor {
      cursor = "▶"
    }
    s += fmt.Sprintf("  %s %s\n", cursor, option)
  }
  return s
}

// vim:tabstop=2:softtabstop=2:shiftwidth=2:expandtab:filetype=go:
