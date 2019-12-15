#!/bin/bash

# Setup


# Serve
tcpserver -v -t 30 -RHl0 0.0.0.0 1337 sh -c 'exec ./$PROB 2>&1'