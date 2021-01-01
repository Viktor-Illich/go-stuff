#!/bin/bash

go run find_ip_v4.go ./auth.log |sort -rn | uniq -c | sort -rn
