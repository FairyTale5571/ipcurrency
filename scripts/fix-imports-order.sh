#!/usr/bin/env bash
## Check if gogroup is installed
if ! tool_loc="$(type -p gogroup)" || [[ -z ${tool_loc} ]]; then
      echo "gogroup is not installed. installing...."
      go install github.com/Bubblyworld/gogroup@latest
fi

gogroup -order std,other,prefix=ipcurrency --rewrite $(find . -type f -name "*.go" | grep -v /vendor/ |grep -v /.git/)
