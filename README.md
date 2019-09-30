# go-xen-client
XenAPI client generated using xenapi.json

Please do not modify any root source files manually. Every *.go files in the root
is generated using the source at builder/

to __re-generate__ the client source after any modification in builder/ source, run the following:
```
cd builder
go run *.go
cd ..
go generate .
```
