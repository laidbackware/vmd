#!/usr/bin/env bats

load test_helpers

setup() {
  setup_command
  export_errors
}

@test "get files with successfully" {
  run $VMD_CMD get files -p vmware_tools -s vmtools -v 11.3.0
  echo $output
  [[ "$output" == *"Eula Accepted:"* ]]
  [[ "$output" == *"Eligable to Download:  true"* ]]
  [[ "$output" == *"VMware-Tools-windows-11.3.0-18090558.zip"* ]]
  [ "$status" -eq 0 ]
}

@test "get files with invalid product" {
  run $VMD_CMD get files -p INVALID -s vmtools -v 11.3.0
  echo $output
  [[ "$output" == *"$ERRORINVALIDSLUG"* ]]
  [ "$status" -eq 1 ]
}

@test "get files with invalid subproduct" {
  run $VMD_CMD get files -p vmware_tools -s INVALID -v 11.3.0
  echo $output
  [[ "$output" == *"$ERRORINVALIDSUBPRODUCT"* ]]
  [ "$status" -eq 1 ]
}

@test "get files with invalid version" {
  run $VMD_CMD get files -p vmware_tools -s vmtools -v INVALID
  echo $output
  [[ "$output" == *"$ERRORINVALIDVERSION"* ]]
  [ "$status" -eq 1 ]
}