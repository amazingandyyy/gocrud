#!/usr/bin/env bash

function prepare {
	echo "initing"
}

function get_user_input {
	read -r -p "Enter repo name (e.g. gobase): " REPO_NAME
	read -r -p "Enter previous repo name (default: __REPO_NAME__): " PREVIOUS_REPO_NAME
	PREVIOUS_REPO_NAME=${PREVIOUS_REPO_NAME:-"__REPO_NAME__"}
}

function replace_repo_name {
	# shellcheck disable=SC2038
	find -E . -type f -regex '.*\.(go|yaml|md)' | xargs sed -i "" "s/$PREVIOUS_REPO_NAME/${REPO_NAME}/g"
	# shellcheck disable=SC2038
	find -E . -type f -regex '.*\install.sh' | xargs sed -i "" "s/$PREVIOUS_REPO_NAME/${REPO_NAME}/g"
}

function go_mod {
	rm -rf go.mod
	go mod init "${REPO_NAME}"
}

function setup_structure {
	read -r -p "Config as a what tool? " TOOL

	case "${TOOL}" in
	[cli])
		echo "Config as a CLI tool"
		;;
	*)
		echo "Not adding CLI tool"
		;;
	esac
}

function post {
	echo "DONE!"
}

prepare
get_user_input
replace_repo_name
go_mod
setup_structure
post
