setup_command() {
  SCRIPT_DIR="$( cd "$( dirname "$BATS_TEST_FILENAME" )" >/dev/null 2>&1 && pwd )/../.."
  export VMD_CMD="go run $SCRIPT_DIR/main.go"
}

export_errors() {
  export ERRORINVALIDSLUG="Invalid slug provided"
  export ERRORINVALIDSUBPRODUCT="Invalid sub-product provided"
	export ERRORINVALIDVERSION="Invalid version provided"
	export ERRORNOMATCHINGVERSIONS="Version glob did not match any files"
	export ERRORNOMATCHINGFILES="No matching files for provided glob"
	export ERRORMULTIPLEMATCHINGFILES="Glob matches multiple files, must be restricted to match a single file"
	export ERROREULAUNACCEPTED="Eula has not been accepted for this sub-product"
	export ERRORNOTENTITLED="You are not entitled to download this sub-product"
	export ERRORNOVERSINGLOB="No version glob provided"
	export ERRORMULTIPLEVERSIONGLOB="Multiple version globs not supported"
}