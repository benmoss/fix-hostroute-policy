# fix-houstroute-policy

Fixes an HNS network by adding a `HostRoute` policy to one that was created without it. A workaround for if you're running Flannel 0.11.0 and therefore 
missing the patch from https://github.com/coreos/flannel/pull/1096

## Usage
```
curl.exe -LO https://github.com/benmoss/fix-hostroute-policy/releases/download/v0.1.0/fix-hostroutepolicy.exe
./fix-hostroutepolicy.exe vxlan0
```

The argument is just your HNS Network name, which if you're using the https://github.com/kubernetes-sigs/sig-windows-tools/tree/master/kubeadm/ script will be `vxlan0`.
