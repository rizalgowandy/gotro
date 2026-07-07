#!/usr/bin/env bash
set -euo pipefail

if rg -n 'github\.com/tarantool/go-openssl|github\.com/tarantool/go-tarantool([[:space:]]|$)' go.mod go.sum; then
  echo 'legacy Tarantool/OpenSSL module dependency is still present' >&2
  exit 1
fi

go list -m github.com/tarantool/go-tarantool/v2
go test ./D/Tt ./X
