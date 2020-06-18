# Examples

1. overview\
Displays total number of hosts and VMs. Also displays resident VMs per host.

```
go run overview.go --hostname <xenserver-host> --username <xenserver-username> --password <xenserver-password>
```

Example output:

```
Total Hosts:  2
Total VMs:  4

Resident VMs per Host
-----------------------

======== node1.cluster1.example.com =========
1. vm2.example.com

======== node1.cluster1.example.com =========
1. vm1.example.com
2. vm9.example.com

```