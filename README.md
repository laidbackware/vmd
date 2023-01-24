# <span style="color:red">**DEPRECATED - vmd is not vcc and can now be found [here](https://github.com/vmware-labs/vmware-customer-connect-cli)**</span>.
# vmd
`vmd` is command line tool to download software from customerconnect.vmware.com.

Influenced by [om](https://github.com/pivotal-cf/om) and [pivnet](https://github.com/pivotal-cf/pivnet-cli).

Uses [vmware-download-sdk](https://github.com/vmware-labs/vmware-customer-connect-sdk) which was heavily inspired by [vmw-sdk](https://github.com/apnex/vmw-sdk) by [apnex](https://github.com/apnex).

![Test Status](https://github.com/laidbackware/vmd/actions/workflows/tests.yml/badge.svg?branch=main)

## Installation

`vmd` is a go binary and can be downloaded from the [releases](https://github.com/laidbackware/vmd/releases) page.

On Linux/Mac the file just needs to be made executable.

To make it available to all users `sudo mv vmd-<os>-<version> /usr/local/bin/vmd`

## Authentication
The Customer Connect username and password can either be passed in as command argurements using `--user` and `--pass` or can be exported as environmental variables, making sure to encause passwords in a single quote in case of special charactors. 

Example below for Linux/Mac:

```
export VMD_USER='email@email.com'
export VMD_PASS='##'
```

## Usage
The examples below assume that the credentials have been exported as environmental variables.

```
  # Download the latest version of release 11 with a file matching the pattern
  vmd download -p vmware_tools -s vmtools -v 11.* -f VMware-Tools-darwin-*.zip --accepteula

  # Download files using a manifest file
  vmd download -m <filename>.yml --accepteula

  # List of available products
  vmd get products

  # List of available sub-products of product vmware_tools
  vmd get subproducts -p vmware_tools

  # List of available versions of sub-products vmtools of vmware_tools
  vmd get versions -p vmware_tools -s vmtools

  # List of available files of version 11.3.0 of vmware_tools
  vmd get files -p vmware_tools -s vmtools -v 11.3.0

  # Display example manifest file
  vmd get manifestexample
```

Multiple downloads can be specified in a manifest file as below and downloaded using `vmd download -m <filename>.yml`

``` yaml
---
# This section will download the latest version of vmware_tools
# Each glob pattern will download a single file each
product: vmware_tools
subproduct: vmtools
version: "*"
filename_globs:
        - "VMware-Tools-darwin-*.tar.gz"
        - "VMware-Tools-darwin-*.zip"
---
# This section will download the latest minor release from major version 10
# The single glob pattern will download 2 files
product: vmware_tools
subproduct: vmtools
version: "10.*"
filename_globs:
        - "VMware-Tools-other-*"
---
```

## Known Issues

- When working in a shell if you add a * to the filename arguement of the download command and you are in a directory where a file matches the pattern, your shell will replace the * to pass in the full file name. This can be worked around by wrapping the file name in single quotes, or by defining the download in a manifest yaml.
- Some products such as horizon will not return the latest version when only a glob is provided. This is because the product switched naming standards meaning it breaks the sort of the version.
- Some product descriptions don't display fully. This is especially true for the horizon products as they are inconsistently named, meaning it's difficult to extract the version number without taking out part of the product name.

## Testing
Tests assume that you have exported credentials as environmental variables and are run from the root of the repo.

- Run go tests `go test ./...`</br>
- Run [BATS](https://github.com/bats-core/bats-core) tests with `bats test/bats`
- To run commands against source use `alias vmd="go run main.go"`

## Development

Update the SDK `go get -u github.com/vmware-labs/vmware-customer-connect-sdk`

Ensure that your IDE exports `VMD_USER` and `VMD_PASS` to be able to run tests and debug.
