.PHONY: verify-tarantool-stack verify-core status-changed-files commit-push pullpush

verify-tarantool-stack:
	bash ./scripts/verify-tarantool-stack.sh

verify-core:
	bash ./scripts/verify-core.sh

status-changed-files:
	bash ./scripts/status-changed-files.sh

commit-push:
	@test -n "$(msg)" || (echo 'Usage: make commit-push msg="the commit message"' && exit 1)
	bash ./scripts/commit-push.sh "$(msg)"

pullpush:
	@test -n "$(msg)" || (echo 'Usage: make pullpush msg="the commit message"' && exit 1)
	./pullpush.sh "$(msg)"
