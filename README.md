# trim

Trim output to a terminal height and width. 

The size of the terminal is fixed to default (20x80) when the file goes from stdin, for some reason term.GetSize returns an error, but when a file is specified as an argument it works as expected.

I need something that can proces large files more eficiently than awk version see`old/trim.sh` and in one step will return line count too, so call `exit()` after this pseudo head output is unsuitable. Golang is the tool for that. 


Below are tests: 


```bash
$ wc zmaz.txt 
     88884  877907268 4389625224 zmaz.txt
$ stat zmaz.txt 
  File: zmaz.txt
  Size: 4389625224	Blocks: 8573496    IO Block: 4096   regular file
$ ls -lh zmaz.txt 
-rw-r--r-- 1 adrian adrian 4.1G Jun 28 11:57 zmaz.txt

# sh/awk version
$ time cat zmaz.txt | trim.sh
real	0m19.603s
user	0m19.743s
sys	0m9.862s
$ time cat zmaz.txt | ./trimgo 
real	0m1.628s
user	0m0.375s
sys	0m2.092s
$ time ./trimgo zmaz.txt
real	0m0.712s
user	0m0.270s
sys	0m0.451s
```
