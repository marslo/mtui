#!/usr/bin/env bash
#=============================================================================
#     FileName : build.sh
#       Author : marslo.jiao@gmail.com
#      Created : 2025-06-17 21:15:06
#   LastChange : 2025-06-18 01:09:58
#=============================================================================

set -euo pipefail

declare -r APP_NAME="mtui"
declare -r DIST_DIR="dist"

declare usage="""
USAGE
  \$ $0 [option]

OPTIONS
  -c, --clean        Clean workspace
  -b, --build        Clean and build for local (host) platform
      --macos        Build for macOS (darwin/arm64)
      --linux        Build for Linux  (linux/amd64)
      --win          Build for Windows (windows/amd64)
  -h, --help         Show this help
"""

function showHelp() { echo -e "${usage}"; }

function clean() {
  rm -rf go.mod go.sum "${APP_NAME}" "${DIST_DIR}"
}

function build() {
  [[ -d "${DIST_DIR}" ]] || mkdir -p "${DIST_DIR}"
  go mod init github.com/marslo/mtui 2>/dev/null || true
  go get github.com/charmbracelet/bubbles
  go get github.com/charmbracelet/bubbletea
  go get github.com/charmbracelet/lipgloss
  go get github.com/spf13/cobra
  go mod tidy
  go build -o "${APP_NAME}"
}

function crossBuild() {
  local goos="$1"
  local goarch="$2"
  local ext="${3:-}"
  local output="${DIST_DIR}/${APP_NAME}-${goos}-${goarch}${ext}"

  echo "Building ${output} ..."
  GOOS="${goos}" GOARCH="${goarch}" go build -o "${output}"
}

function buildMac() { crossBuild darwin arm64; }
function buildLinux() { crossBuild linux amd64; }
function buildWindows() { crossBuild windows amd64 ".exe"; }

function main() {
  [[ $# -eq 0 ]] && { clean && build; return $?; }

  while [[ $# -gt 0 ]]; do
    case "$1" in
      -h | --help  ) showHelp                  ; return 0  ;;
      -c | --clean ) clean                     ; return $? ;;
      -b | --build ) clean && build            ; return $? ;;
      --macos      ) buildMac                  ; return $? ;;
      --linux      ) buildLinux                ; return $? ;;
      --win        ) buildWindows              ; return $? ;;
      *            ) echo "unknown option: $1" ; return 1  ;;
    esac
  done
}

main "$@"

# vim:tabstop=2:softtabstop=2:shiftwidth=2:expandtab:filetype=sh:
