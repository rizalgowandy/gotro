#!/usr/bin/env bash
set -euo pipefail

if [[ $# -lt 1 ]]; then
  echo 'Usage: scripts/commit-push.sh "commit message" [files...]' >&2
  exit 1
fi

MSG="$1"
shift

if [[ $# -eq 0 ]]; then
  git add Makefile scripts
else
  git add "$@"
fi

git status --short
git diff --cached --stat
git commit -m "$MSG"
git push origin master
