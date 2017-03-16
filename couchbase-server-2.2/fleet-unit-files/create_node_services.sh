#!/bin/bash

# Generate couchbase_node.n.service unit files for however many 
# other couchbase nodes you want (in addition to the bootstrap node)
#
# Usage
#
# ./create_node_services.sh 2

NUM_NODES=$1

re='^[0-9]+$'
if ! [[ $NUM_NODES =~ $re ]] ; then
   echo "error: Not a number" >&2; exit 1
fi

for i in `seq 1 $NUM_NODES`;
do
    cp couchbase_node.service.template couchbase_node.$i.service
    echo "Created couchbase_node.$i.service"
done    
