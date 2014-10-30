# Running under CoreOS Fleet

* Launch 3 CoreOS instances on AWS Cloud Formation: See https://coreos.com/docs/running-coreos/cloud-providers/ec2/  
* ssh into one of the coreos instances
* Update the docker memlock setting 
 * `cp /usr/lib/systemd/system/docker.service /etc/systemd/system/`
 * edit /etc/systemd/system/docker.service and add line: `LimitMEMLOCK=infinity`
 * `sudo systemctl daemon-reload`
 * `sudo systemctl restart docker`


# TODO

* Have it store data on a volume

# References

* https://registry.hub.docker.com/u/ncolomer/couchbase/