#!/bin/sh
cd ../submission-data
function for_loop {
    for f in *;  do
        if [ -d $f  -a ! -h $f ];
        then
            cd -- "$f";

            # delete all test files from user submission.
            find . -type f -iname \*_test.go -delete

            # copy out task_test.go with an anti-cheating measure.
            cp ../../anti-cheating/task_test.go ./

            cd ..;
        fi;
    done;
};
for_loop
