# Find top X words
To search top X words in a given file

## Languages and Tools
  1.  The reference platform for the task is a Linux 64bit system.  
  2. The solution is to be written entirely in Go.  
  3. Your program should also handle binary files (e.g. /boot/vmlinuz) without crashing.  
 
## The Problem
Given the attached text file as an argument, your program will read the file, and output the 20 most frequently used words in the file in order, along with their frequency. The output should be similar to that of the following bash program:  

```bash
#!/usr/bin/env bash
cat $1 | tr -cs 'a-zA-Z' '[\n*]' | grep -v "^$" | tr '[:upper:]' '[:lower:]'| sort | uniq -c | sort -nr | head -20
 ```

Sample output (this output is not from the reference text):
9 the
5 you
4 lessons 
4 can
4 a
3 vim
3 of
3 is
3 file
2 users
2 tutorial 
2 tutor
2 the
2 that
2 new
2 make
2 it
2 in
2 have
2 for
