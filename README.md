# vmd
`vmd` is command line tool to download software from customerconnect.vmware.com. </br>
Influenced by [om](https://github.com/pivotal-cf/om) and [pivnet](https://github.com/pivotal-cf/pivnet-cli). </br>
Uses [vmware-download-sdk](https://github.com/laidbackware/vmware-download-sdk) which was heavily inspired by [vmw-sdk](https://github.com/apnex/vmw-sdk) by [apnex](https://github.com/apnex). </br>
![Test Status](https://github.com/laidbackware/vmd/actions/workflows/tests.yml/badge.svg?branch=main)

# Installation
`vmd` is a go binary and can be downloaded from the [releases](https://github.com/laidbackware/vmd/releases) page.</br>
On Linux/Mac the file just needs to be made executable.</br>
To make it available to all users `sudo mv vmd-<os>-<version> /usr/local/bin/vmd`

# Usage
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
Example manifest to use with `vmd download -m <filename>.yml`
```
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
# Known Issues
- When working in a shell if you add a * to the filename arguement of the download command and you are in a directory where a file matches the pattern, your shell will replace the * to pass in the full file name. This can be worked around by wrapping the file name in single quotes, or by defining the download in a manifest yaml.
- Some products such as horizon will not return the latest version when only a glob is provided. This is because the product switched naming standards meaning it breaks the sort of the version.
- Some product descriptions don't display fully. This is especially true for the horizon products as they are inconsistently named, meaning it's difficult to extract the version number without taking out part of the product name.
- NSX-T currently does not work with a glob as the last charactor of the version. This is due to the sorting prioritising the LE versions of the product.

# Testing
To run commands against source use `alias vmd="go run main.go"`</br>
Run go tests `go test ./...`</br>
Run BATS tests `bats test/bats`

# Development
Update the SDK `go get -u github.com/laidbackware/vmware-download-sdk`</br>
Ensure that your IDE exports `VMD_USER` and `VMD_PASS` to be able to run tests and debug.
