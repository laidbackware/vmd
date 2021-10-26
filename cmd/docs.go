package cmd

//  Usage Section
const (
	getProductsUsage = `  # List of available products
  vmd get products`

	getSubProductsUsage = `  # List of available sub-products of product vmware_tools
  vmd get subproducts -p vmware_tools`

	getVersions = `  # List of available versions of sub-products vmtools of vmware_tools
  vmd get versions -p vmware_tools -s vmtools`

	getFiles = `  # List of available files of version 11.3.0 of vmware_tools
  vmd get files -p vmware_tools -s vmtools -v 11.3.0`

	getManifestExample = `  # Display example manifest file
  vmd get manifestexample`

	downloadUsage = `  # Download the latest version of release 11 with a file matching the pattern
  vmd download -p vmware_tools -s vmtools -v 11.* -f VMware-Tools-darwin-*.zip --accepteula

  # Download files using a manifest file
  vmd download -m manifest.yml --accepteula`
)

const exampleManifest = `---
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
---`