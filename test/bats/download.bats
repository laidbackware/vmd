#!/usr/bin/env bats

load test_helpers

setup() {
  setup_command
  export_errors
  export_yamls
  export TEMP_DIR="$(mktemp -dt bats.XXXXX)"
}

teardown() {
  rm -rf "${TEMP_DIR}"
  rm $HOME/vmd-downloads/VMware-Tools-darwin-11.3.0-*.zip
  echo ""
}

@test "download single file successfully to temp" {
  $VMD_CMD logout
  rm -f $TEMP_DIR/*
  local cmd="$VMD_CMD download -p vmware_horizon_clients -s cart+andrd_x8632 -v 2106 -f VMware-Horizon-Client-AndroidOS-x86-*-store.apk --accepteula -o $TEMP_DIR"
  echo $cmd
  run $cmd
  echo $output
  [[ "$output" != *"No output directory set."* ]]
  [[ "$output" == *"Collecting download payload"* ]]
  [[ "$output" == *"Download started to"* ]]
  [[ "$output" == *"Download finished"* ]]
  [ "$status" -eq 0 ]
  [ -f $TEMP_DIR/VMware-Horizon-Client-*.apk ]
}

@test "download single file successfully to user vmd-downloads" {
  rm -f $TEMP_DIR/*
  local cmd="$VMD_CMD download -p vmware_tools -s vmtools -v 11.3.0 -f VMware-Tools-darwin-*.zip --accepteula"
  echo $cmd
  run $cmd
  echo $output
  [[ "$output" == *"No output directory set."* ]]
  [[ "$output" == *"Collecting download payload"* ]]
  [[ "$output" == *"Download started to"* ]]
  [[ "$output" == *"Download finished"* ]]
  [ "$status" -eq 0 ]
  [ -f $HOME/vmd-downloads/VMware-Tools-darwin-11.3.0-*.zip ]
}

@test "download multiple files successfully to temp" {
  rm -f $TEMP_DIR/*
  run $VMD_CMD download -p vmware_tools -s vmtools -v 11.3.0 -f VMware-Tools-darwin-* --accepteula -o $TEMP_DIR
  echo $output
  [[ "$output" != *"No output directory set."* ]]
  [[ "$output" == *"Collecting download payload"* ]]
  [[ "$output" == *"Download started to"* ]]
  [[ "$output" == *"Download finished"* ]]
  [ "$status" -eq 0 ]
  [ -f $TEMP_DIR/VMware-Tools-darwin-*.zip ]
  [ -f $TEMP_DIR/VMware-Tools-darwin-*.tar.gz ]
}

@test "download from manifest" {
  rm -f $TEMP_DIR/*
  run $VMD_CMD download -m <(echo "$VALID_YAML") --accepteula -o $TEMP_DIR
  echo $output
  [[ "$output" == *"Opening manifest file:"* ]]
  [[ "$output" == *"Collecting download payload"* ]]
  [[ "$output" == *"Download started to"* ]]
  [[ "$output" == *"Download finished"* ]]
  [ "$status" -eq 0 ]
  ls -l $TEMP_DIR
  [ -f $TEMP_DIR/VMware-Tools-darwin-*.zip ]
  [ -f $TEMP_DIR/VMware-Tools-darwin-*.tar.gz ]
  [ -f $TEMP_DIR/VMware-Tools-other-*.tar.gz ]
}

@test "download from manifest missing field" {
  run $VMD_CMD download -m <(echo "$INVALID_YAML_MISSING_FIELD") --accepteula -o $TEMP_DIR
  echo $output
  [[ "$output" == *"Opening manifest file:"* ]]
  [[ "$output" == *"Manifest entry 0 does not have the 4 required keys!"* ]]
  [[ "$output" != *"Collecting download payload"* ]]
  [[ "$output" != *"Download started to"* ]]
  [[ "$output" != *"Download finished"* ]]
  [ "$status" -eq 1 ]
}

@test "download from manifest invalid type" {
  run $VMD_CMD download -m <(echo "$INVALID_YAML_INVALID_TYPE") --accepteula -o $TEMP_DIR
  echo $output
  [[ "$output" == *"Opening manifest file:"* ]]
  [[ "$output" == *"Parsing file failed with error:"* ]]
  [[ "$output" != *"Collecting download payload"* ]]
  [[ "$output" != *"Download started to"* ]]
  [[ "$output" != *"Download finished"* ]]
  [ "$status" -eq 1 ]
}

@test "download with invalid product" {
  run $VMD_CMD download -p INVALID -s vmtools -v 11.3.0 -f VMware-Tools-darwin-*.zip --accepteula
  echo $output
  [[ "$output" == *"$ERRORINVALIDSLUG"* ]]
  [ "$status" -eq 1 ]
}

@test "download with invalid subproduct" {
  run $VMD_CMD download -p vmware_tools -s INVALID -v 11.3.0 -f VMware-Tools-darwin-*.zip --accepteula
  echo $output
  [[ "$output" == *"$ERRORINVALIDSUBPRODUCT"* ]]
  [ "$status" -eq 1 ]
}

@test "download with invalid version" {
  run $VMD_CMD download -p vmware_tools -s vmtools -v INVALID -f VMware-Tools-darwin-*.zip --accepteula
  echo $output
  [[ "$output" == *"$ERRORINVALIDVERSION"* ]]
  [ "$status" -eq 1 ]
}

@test "download with invalid credentials" {
  $VMD_CMD logout
  run $VMD_CMD download -p vmware_tools -s vmtools -v 11.3.0 -f VMware-Tools-darwin-*.zip --accepteula --user invalid --pass invalid
  echo $output
  [[ "$output" == *"$ERRORAUTHENTICATIONFAILURE"* ]]
  [ "$status" -eq 1 ]
}

@test "download when not entitled" {
  $VMD_CMD logout
  run $VMD_CMD download -p vmware_vsan -s esxi -v 7.* -f VMware-VMvisor-Installer-*.iso --accepteula
  echo $output
  [[ "$output" == *"$ERRORNOTENTITLED"* ]]
  [ "$status" -eq 1 ]
}

@test "download with invalid output directory" {
  run $VMD_CMD download -p vmware_tools -s vmtools -v 11.3.0 -f VMware-Tools-darwin-*.zip --accepteula -o /tmp/stilton/on/toast
  echo $output
  [[ "$output" == *"ERROR: Output directory"* ]]
  [ "$status" -eq 1 ]
}