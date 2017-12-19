#!/bin/bash

set -eu

PANDOC_VERSION="2.0.5"
PANDOC_CHECKSUM="751c505deee6c5d1eecf3230cbea3bbdeafc69ee924daf89edd7e9438d9c22c4"

function version {
  local tag=$(git describe --exact-match --tags HEAD 2>/dev/null || true)
  if [[ -n "$tag" ]];
  then
    echo "$tag"
  else
    git reflog show -1 |cut -d' ' -f1
  fi
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

function task_prepare_ci {
  sudo apt-get -qq update
  sudo apt-get install \
    -y \
    texlive \
    texlive-luatex \
    texlive-xetex \
    texlive-latex-extra \
    texlive-latex-recommended \
    texlive-fonts-recommended \
    fonts-lmodern
  sudo luaotfload-tool --update
  wget -O pandoc.deb https://github.com/jgm/pandoc/releases/download/${PANDOC_VERSION}/pandoc-${PANDOC_VERSION}-1-amd64.deb
  echo "${PANDOC_CHECKSUM}  pandoc.deb" | sha256sum -c
  sudo dpkg -i pandoc.deb
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
  prepare-ci) task_prepare_ci ;;
  *) task_usage ;;
esac
