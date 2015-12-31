#!/usr/bin/env bash
set -e

SRC_DIR=src
PKG_DIR=pkg
BIN_DIR=bin

DEV_PKGS=(
  server
  fileapi
)

mode=$1

clean_dir() {
  # clear dir
  rm -rf "$SRC_DIR" "$PKG_DIR" "$BIN_DIR"
  # make dir
  mkdir "$SRC_DIR" "$PKG_DIR" "$BIN_DIR"
  # make links
  cd src
  for d in ${DEV_PKGS[@]} ; do
    ln -sf "../$d" $d
  done
  go env
}

run_go_command() {
  for d in ${DEV_PKGS[@]} ; do
     go "$mode" "$d"
  done
}

if [ -z "$mode" ]; then
  cat <<EOF
usage:
   ./make.sh 
         clean    ... clean setup build environment
         build    ... execute go build command
         install  ... execute go install command
EOF
elif [ "$mode" = "clean" ]; then
  clean_dir
else
  run_go_command
fi

# EOF
