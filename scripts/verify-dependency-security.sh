#!/usr/bin/env bash
set -euo pipefail

mods="$(go list -m -f '{{.Path}} {{.Version}}' all)"

version_lt() {
  local left="${1#v}"
  local right="${2#v}"
  [ "$(printf '%s\n%s\n' "$left" "$right" | sort -V | head -n1)" = "$left" ] && [ "$left" != "$right" ]
}

version_ge() {
  ! version_lt "$1" "$2"
}

selected_version() {
  awk -v mod="$1" '$1 == mod { print $2; exit }' <<<"$mods"
}

require_absent() {
  local mod="$1"
  local version
  version="$(selected_version "$mod")"
  if [ -n "$version" ]; then
    echo "$mod is selected at $version; remove it from the module graph" >&2
    exit 1
  fi
}

require_min() {
  local mod="$1"
  local min="$2"
  local version
  version="$(selected_version "$mod")"
  if [ -n "$version" ] && version_lt "$version" "$min"; then
    echo "$mod is $version, expected >= $min" >&2
    exit 1
  fi
}

require_not_otel_vulnerable() {
  local version
  version="$(selected_version go.opentelemetry.io/otel)"
  if [ -n "$version" ] && version_ge "$version" v1.36.0 && version_lt "$version" v1.41.0; then
    echo "go.opentelemetry.io/otel is $version, expected outside vulnerable range v1.36.0-v1.40.x" >&2
    exit 1
  fi
}

require_absent github.com/docker/docker
require_min github.com/docker/cli v29.3.1
require_min github.com/opencontainers/runc v1.3.6
require_min go.mongodb.org/mongo-driver v1.17.7
require_min github.com/gofiber/fiber/v2 v2.52.13
require_min github.com/shamaton/msgpack/v2 v2.4.1
require_min golang.org/x/crypto v0.53.0
require_min golang.org/x/image v0.18.0
require_not_otel_vulnerable

echo "dependency security graph ok"
