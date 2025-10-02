#!/usr/bin/env bash

# Get path to script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

. "${SCRIPT_DIR}/.env"

#Config
export WIKI_URL="https://bluearchive.wiki/w/"
export COOKIE_JAR="cookie.txt"
export API_URL="api.php"
export BOT_VERSION="1.0"

#User Creds
export WIKI_USER="$BOT_USER"
unset BOT_USER
export WIKI_PASSWORD="$BOT_PASSWORD"
unset BOT_PASSWORD
# <client name>/<version> (<contact information>) <library/framework name>/<version> [<library name>/<version> ...]
# omit parts not applicable 
export WIKI_AGENT="$WIKI_USER/$BOT_VERSION ($BOT_CONTACT) $(curl --version | head -n 1 | awk '{ print $1"/"$2 }')"

