#!/bin/bash

set -eu

function version {
  git reflog show -1 |cut -d' ' -f1
}

function task_build {
  echo "Building resume.pdf"
  pandoc \
    --pdf-engine=lualatex \
    --template=template.latex \
    -f markdown \
    -o resume.pdf \
    --metadata=build_date:$(date +%Y-%m-%d) \
    --metadata=version:$(version) \
    resume.md
  echo "done"
}

function task_watch {
  task_build || true
  while inotifywait resume.md template.latex;
  do
    task_build || true
  done
}

function task_usage {
  echo "usage: $0 build|watch"
  exit 1
}

cmd=${1:-}
shift || true
case "$cmd" in
  build) task_build ;;
  watch) task_watch ;;
  *) task_usage ;;
esac
