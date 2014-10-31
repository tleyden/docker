# Running under CoreOS Fleet


* Launch 3 CoreOS instances on AWS Cloud Formation via: [this link](https://console.aws.amazon.com/cloudformation/home?region=us-east-1#cstack=sn%7ECouchbase-CoreOS%7Cturl%7Ehttp://tleyden-misc.s3.amazonaws.com/couchbase-coreos/coreos-stable-pv.template)
* ssh into one of the coreos instances
* On each node in the cluster, update the docker memlock setting 
 * `cp /usr/lib/systemd/system/docker.service /etc/systemd/system/`
 * edit /etc/systemd/system/docker.service and add line: `LimitMEMLOCK=infinity`
 * `sudo systemctl daemon-reload`
 * `sudo systemctl restart docker`
* wget https://raw.githubusercontent.com/tleyden/docker/master/couchbase-server-2.2/fleet-unit-files/couchbase_bootstrap_node.service
* wget https://raw.githubusercontent.com/tleyden/docker/master/couchbase-server-2.2/fleet-unit-files/couchbase_bootstrap_node_announce.service
* wget https://raw.githubusercontent.com/tleyden/docker/master/couchbase-server-2.2/fleet-unit-files/couchbase_node.service


# TODO

* It cannot discover the ip, gets the docker ip.  Needs to pass in real ip.
* Have it store data on a volume

# References

* https://registry.hub.docker.com/u/ncolomer/couchbase/