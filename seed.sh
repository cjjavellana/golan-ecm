#!/usr/bin/env bash

set -e

workspaceId=$(evans cli call --file samples/evans-create-workspace.json ecm.ContentEngine.CreateWorkspace | jq .objectId | sed -e 's/"//g')
echo "Workspace $workspaceId Created."

dcId=$(sed -e "s/{{workspaceId}}/$workspaceId/g" samples/evans-create-docclass.json | evans cli --call ecm.ContentEngine.CreateDocumentClass | jq .objectId)
echo "Document class $dcId Created."