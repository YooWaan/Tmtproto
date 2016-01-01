#!/usr/bin/env bash
set -e

SRC_DIR=src
PKG_DIR=pkg
BIN_DIR=bin

DEV_PKGS=(
  types
  auth
  server
)

mode=$1

clean_dir() {
  # clear dir
  rm -rf "$PKG_DIR" "$BIN_DIR"
  for d in ${DEV_PKGS[@]} ; do
    rm -f "./$SRC_DIR/$d"
  done
  # make dir
  mkdir -p "$SRC_DIR" "$PKG_DIR" "$BIN_DIR"
  # make links
  cd src
  for d in ${DEV_PKGS[@]} ; do
    echo $d
    ln -sf "../$d" $d
  done
  go env
}

run_go_command() {
  for d in ${DEV_PKGS} ; do
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
