#!/usr/bin/env bash

set -e

workspaceId=$(evans cli call --file samples/evans-create-workspace.json ecm.ContentEngine.CreateWorkspace | jq .objectId)
echo "Workspace $workspaceId Created."
sed 