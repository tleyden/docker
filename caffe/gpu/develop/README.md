Caffe is a deep learning framework.   See the [Caffe homepage](http://caffe.berkeleyvision.org/) for more info.  

This was created for the [elastic-thought](https://github.com/tleyden/elastic-thought) project, which is a REST api wrapper around Caffe, but should be useful for anyone running Caffe.  

There are CPU and GPU versions of this image:

This docker image is part of the following images generated from a single template:

* CPU: [tleyden5iwx/caffe-cpu-master](https://registry.hub.docker.com/u/tleyden5iwx/caffe-cpu-master/)
* GPU: [tleyden5iwx/caffe-gpu-master](https://registry.hub.docker.com/u/tleyden5iwx/caffe-gpu-master/)

**Verify CPU Version:**

    $ cd /opt/caffe/data/mnist
    $ ./get_mnist.sh
    $ cd ../../examples/mnist
    $ sed -i 's/solver_mode: GPU/solver_mode: CPU/' lenet_solver.prototxt
    $ cd ../../
    $ ./examples/mnist/create_mnist.sh
    $ ./examples/mnist/train_lenet.sh

**Expected output:**

    libdc1394 error: Failed to initialize libdc1394 
    I1018 17:02:23.552733    66 caffe.cpp:90] Starting Optimization 
    I1018 17:02:23.553583    66 solver.cpp:32] Initializing solver from parameters:
    ... lots of output ...
    I1207 03:17:50.054651    57 solver.cpp:247] Iteration 10000, Testing net (#0)
    I1207 03:17:55.369581    57 solver.cpp:298]     Test net output #0: accuracy = 0.9904
    I1207 03:17:55.370614    57 solver.cpp:298]     Test net output #1: loss = 0.029635 (* 1 = 0.029635 loss)
    I1018 17:17:58.684598    66 caffe.cpp:102] Optimization Done.

**How to launch GPU instances:**

* You will to run on hardware that has the nvidia kernel module installed
* You will need to pass in the nvidia devices in the `docker run` command

See [Running Caffe on AWS GPU Instance via Docker](http://tleyden.github.io/blog/2014/10/25/running-caffe-on-aws-gpu-instance-via-docker/) for instructions.  

**Verify GPU Version:**

    $ cd /opt/caffe/data/mnist
    $ ./get_mnist.sh
    $ ./examples/mnist/create_mnist.sh
    $ ./examples/mnist/train_lenet.sh

**Troubleshooting:**

If you get the error "error while loading shared libraries: libglog.so.0: cannot open shared object file: No such file or directory", try running:

    $ ldconfig

**References:**

* [Running Caffe on AWS GPU Instance via Docker](http://tleyden.github.io/blog/2014/10/25/running-caffe-on-aws-gpu-instance-via-docker/)
* [ElasticThought Caffe REST API](https://github.com/tleyden/elastic-thought)
