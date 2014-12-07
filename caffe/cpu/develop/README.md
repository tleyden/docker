Caffe is a deep learning framework.   See the [Caffe homepage](http://caffe.berkeleyvision.org/) for more info.  

The following variations are available:

* CPU
    * Develop branch: tleyden5iwx/caffe-cpu-develop
    * Master branch: tleyden5iwx/caffe-cpu-master
* GPU 
    * Develop branch: tleyden5iwx/caffe-gpu-develop
    * Master branch: tleyden5iwx/caffe-gpu-master

**Pull and run:**

    $ sudo docker pull tleyden5iwx/
    $ sudo docker run -i -t tleyden5iwx/caffe /bin/bash

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
    I1018 17:17:58.684598    66 caffe.cpp:102] Optimization Done.

**Troubleshooting:**

If you get the error "error while loading shared libraries: libglog.so.0: cannot open shared object file: No such file or directory", try running:

    $ ldconfig

