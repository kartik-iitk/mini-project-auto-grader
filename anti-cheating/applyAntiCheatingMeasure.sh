#!/bin/sh
cd ../submission-data
function for_loop {
    for f in *;  do
        if [ -d $f  -a ! -h $f ];
        then
            cd -- "$f";
	        # create/overwrite existing task_test.go with our file as an anti-cheating measure.
	        cp ../../anti-ceating/task_test.go ./
            cd ..;
        fi;
    done;
};
for_loop
