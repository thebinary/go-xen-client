# go-xen-client
XenAPI client generated using xenapi.json

__Work In Progress__. Not fully tested.

Please do not modify any root source files manually. Every *.go files in the root
is generated using the source at builder/

to __re-generate__ the client source after any modification in builder/ source, run the:
```
make generate
```

to __clean/remove__ all client source, run:
```
make clean
```

TODO:
- To<object> fully correct implementation
