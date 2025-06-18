package cmd

import (
  "fmt"
  "os"
  "strings"

  "github.com/charmbracelet/bubbles/textarea"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/spf13/cobra"
)

var writePlaceholder string

func init() {
  writeCmd.Flags().StringVar(&writePlaceholder, "placeholder", "", "Placeholder text for the textarea")
  rootCmd.AddCommand(writeCmd)
}

var writeCmd = &cobra.Command{
  Use:   "write",
  Short: "Multi-line textarea input (Ctrl+J to submit)",
  Run: func(cmd *cobra.Command, args []string) {
    model := writeModel{
      area:      textarea.New(),
      cancelled: false,
      placeholder: writePlaceholder,
    }
    model.area.Placeholder = writePlaceholder
    model.area.Focus()
    model.area.CharLimit = 0
    model.area.ShowLineNumbers = true
    model.area.SetHeight(10)

    p := tea.NewProgram(model)
    if m, err := p.Run(); err == nil {
      if result, ok := m.(writeModel); ok {
        if result.cancelled {
          os.Exit(130)
        }
        fmt.Println(result.area.Value())
        os.Exit(0)
      }
    }
    os.Exit(1)
  },
}

type writeModel struct {
  area        textarea.Model
  cancelled   bool
  placeholder string
}

func (m writeModel) Init() tea.Cmd {
  return textarea.Blink
}

func (m writeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmd tea.Cmd

  switch msg := msg.(type) {
  case tea.KeyMsg:
    switch msg.Type {
    case tea.KeyCtrlC, tea.KeyEsc:
      m.cancelled = true
      return m, tea.Quit
    }

    if msg.String() == "ctrl+j" {
      return m, tea.Quit
    }
    if msg.String() == "tab" && strings.TrimSpace(m.area.Value()) == "" && m.placeholder != "" {
      m.area.SetValue(m.placeholder)
    }
  }

  m.area, cmd = m.area.Update(msg)
  return m, cmd
}

func (m writeModel) View() string {
  return "\n" + m.area.View() + "\n\n(Ctrl+J to submit, Esc to cancel, Tab to insert placeholder)\n"
}

// vim:tabstop=2:softtabstop=2:shiftwidth=2:expandtab:filetype=go:
