#!/bin/bash

top -l1 | grep -E "^CPU" | tail -1 | awk '{ print $3 + $5 }'