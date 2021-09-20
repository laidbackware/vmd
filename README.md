# vmd
CLI tool to download software from customerconnect.vmware.com

alias vmd="go run main.go"

list products
list subproducts -p
list versions -p -s
list files -p -s -v
download -p -s -v -g -a
download -f -a

go get -u github.com/laidbackware/vmware-download-sdk