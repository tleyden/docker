Caffe is a deep learning framework.   See the [Caffe homepage](http://caffe.berkeleyvision.org/) for more info.  

The following docker image variations are available:

* CPU
    * Develop branch: [tleyden5iwx/caffe-cpu-develop](https://registry.hub.docker.com/u/tleyden5iwx/caffe-cpu-develop/)
    * Master branch: [tleyden5iwx/caffe-cpu-master](https://registry.hub.docker.com/u/tleyden5iwx/caffe-cpu-master/)
* GPU 
    * Develop branch: [tleyden5iwx/caffe-gpu-develop](https://registry.hub.docker.com/u/tleyden5iwx/caffe-gpu-develop/)
    * Master branch: [tleyden5iwx/caffe-gpu-master](https://registry.hub.docker.com/u/tleyden5iwx/caffe-gpu-master/)

**Pull and run:**

    $ sudo docker pull tleyden5iw/caffe-cpu-master
    $ sudo docker run -i -t tleyden5iwx/caffe-cpu-master /bin/bash

**Verify:**

    $ cd /opt/caffe/data/mnist
    $ ./get_mnist.sh
    $ cd ../../examples/mnist
    $ sed -i 's/solver_mode: GPU/solver_mode: CPU/' lenet_solver.prototxt
    $ ./create_mnist.sh
    $ ./train_lenet.sh

**Expected output:**

    libdc1394 error: Failed to initialize libdc1394 
    I1018 17:02:23.552733    66 caffe.cpp:90] Starting Optimization 
    I1018 17:02:23.553583    66 solver.cpp:32] Initializing solver from parameters:
    ... lots of output ...
    I1207 03:17:50.054651    57 solver.cpp:247] Iteration 10000, Testing net (#0)
    I1207 03:17:55.369581    57 solver.cpp:298]     Test net output #0: accuracy = 0.9904
    I1207 03:17:55.370614    57 solver.cpp:298]     Test net output #1: loss = 0.029635 (* 1 = 0.029635 loss)
    I1018 17:17:58.684598    66 caffe.cpp:102] Optimization Done.

**Troubleshooting:**

If you get the error "error while loading shared libraries: libglog.so.0: cannot open shared object file: No such file or directory", try running:

    $ ldconfig

