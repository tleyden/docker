# Running Couchbase Server 2.2 under CoreOS Fleet

## Launch CoreOS instances via AWS Cloud Formation

Launch 3 CoreOS instances on AWS Cloud Formation via: [this link](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#cstack=sn%7ECouchbase-CoreOS%7Cturl%7Ehttp://tleyden-misc.s3.amazonaws.com/couchbase-coreos/coreos-stable-pv.template)

Ssh into one of the coreos instances, it doesn't matter which one.

## Update the "memlock" setting

It will work without updating the memlock limit setting, but at best it will be less performant, and at worst it will lead to crashes.  Unfortunately this step is still manual.

* Ssh into each CoreOS node in the cluster and:
 * `cp /usr/lib/systemd/system/docker.service /etc/systemd/system/`
 * edit /etc/systemd/system/docker.service and add line: `LimitMEMLOCK=infinity`
 * `sudo systemctl daemon-reload`
 * `sudo systemctl restart docker`

## 

## Generate unit files

We already have a special unit file for the "bootstrap node", but we'll need unit files for the N other Couchbase Server nodes.

In the default CloudFormation, it will launch three nodes total.  Which means we want two other Couchbase Server nodes.  Generate the unit files for those two nodes with:

```
$ create_node_services.sh 2
```

## Add Couchbase credentials to etcd

```
$ etcdctl set /services/couchbase/userpass "user:passw0rd"
```

Replace `user:passw0rd` with a sensible username and password.  It **must** be colon separated, with no spaces.

## Launch CoreOS Fleet

* 
* wget https://raw.githubusercontent.com/tleyden/docker/master/couchbase-server-2.2/fleet-unit-files/couchbase_bootstrap_node.service
* wget https://raw.githubusercontent.com/tleyden/docker/master/couchbase-server-2.2/fleet-unit-files/couchbase_bootstrap_node_announce.service
* wget https://raw.githubusercontent.com/tleyden/docker/master/couchbase-server-2.2/fleet-unit-files/couchbase_node.service


# TODO -- use volume container

Currently data is stored inside the container filesystem.  It should be changed to store the data on a volume mounted from the CoreOS instance.

# References

* https://registry.hub.docker.com/u/ncolomer/couchbase/