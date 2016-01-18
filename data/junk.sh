#!/bin/bash
        for i in `seq 1 300`;
        do
                cat /dev/urandom | tr -cd 'a-f0-9' | head -c 60  >> 3.txt
        	printf "\n" >> something.txt

        done
