#!/usr/bin/env bats

load test_helpers

setup() {
  setup_command
  export_errors
}

@test "get eula successfully" {
  run $VMD_CMD get eula -p vmware_tools -s vmtools -v 11.3.0
  echo $output
  [[ "$output" == *"Open the URL in your browser: http://www.vmware.com"* ]]
  [ "$status" -eq 0 ]
}

@test "get eula with invalid product" {
  run $VMD_CMD get eula -p INVALID -s vmtools -v 11.3.0
  echo $output
  [[ "$output" == *"$ERRORINVALIDSLUG"* ]]
  [ "$status" -eq 1 ]
}

@test "get eula with invalid subproduct" {
  run $VMD_CMD get eula -p vmware_tools -s INVALID -v 11.3.0
  echo $output
  [[ "$output" == *"$ERRORINVALIDSUBPRODUCT"* ]]
  [ "$status" -eq 1 ]
}

@test "get eula with invalid version" {
  run $VMD_CMD get eula -p vmware_tools -s vmtools -v INVALID
  echo $output
  [[ "$output" == *"$ERRORINVALIDVERSION"* ]]
  [ "$status" -eq 1 ]
}

@test "get eula with invalid credentials" {
  $VMD_CMD logout
  run $VMD_CMD get eula -p vmware_tools -s vmtools -v 11.3.0 --user invalid --pass invalid
  echo $output
  [[ "$output" == *"$ERRORAUTHENTICATIONFAILURE"* ]]
  [ "$status" -eq 1 ]
}