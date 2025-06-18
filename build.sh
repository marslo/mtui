#!/usr/bin/env bash

set -euo pipefail

function clean() {
  rm -f go.mod go.sum
  rm -f mtui
}

function build() {
  go mod init github.com/marslo/mtui
  go get github.com/charmbracelet/bubbles
  go get github.com/charmbracelet/bubbletea
  go get github.com/charmbracelet/lipgloss
  go get github.com/spf13/cobra

  go mod tidy
  go build -o mtui
}

usage="""
USAGE
  \$ $0 [option]

OPTIONS
  -c, --clean       Run clean only
  -b, --build       Run clean and then build
  -h, --help        Show this help message
"""

function showHelp() {
  echo -e "${usage}"
}

function main() {
  [[ $# -eq 0 ]] && { clean && build; return $?; }

  while [[ $# -gt 0 ]]; do
    case $1 in
      -h | --help  ) showHelp                      ; return 0  ;;
      -c | --clean ) clean                         ; return $? ;;
      -b | --build ) clean && build                ; return $? ;;
      *            ) echo "unknown option: $1" >&2 ; return 1  ;;
    esac
  done
}

main "$@"

# vim:tabstop=2:softtabstop=2:shiftwidth=2:expandtab:filetype=sh:
