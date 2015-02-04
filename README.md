stat [![Build Status](https://img.shields.io/travis/rjeczalik/stat/master.svg)](https://travis-ci.org/rjeczalik/stat "linux_amd64") [![Build status](https://img.shields.io/appveyor/ci/rjeczalik/stat.svg)](https://ci.appveyor.com/project/rjeczalik/stat "windows_amd64")
======

Statistics and time series toys for the command line.

*Installation*

```
~ $ go get github.com/rjeczalik/stat/cmd/...
```

### cmd/dln

Prints derivative computed out of line-separated numbers.

*Example usage*

```
~ $ curl -sS http://cdimage.ubuntu.com/daily-live/current/vivid-desktop-amd64.iso -o vivid-amd64.iso &
[1] 21496
```
```
~ $ while sleep 1; do du -BK vivid-amd64.iso ; done | dln
1008
1172
1548
2332
3200
4756
6572
7056
7052
7048
7060
7036
7048
7056
^C
```

### cmd/hist

Prints histogram for line-separated data points. It sorts the result set by the number of occurances in descending order.

*Example usage*

```
~ $ log=https://gist.githubusercontent.com/rjeczalik/f18349ad629f07d19839/raw/b8089282fdd5a8ea8589fe33bc88cc6d29db7026/lazyvm.log
```
```
~ $ curl -sS $log | dln | hist
  0	962	░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
  1	5	
 18	3	
  3	2	
  9	1	
  6	1	
 49	1	
 23	1	
 22	1	
 21	1	
 59	1	
 11	1	
 78	1	
  4	1	
 42	1	
  2	1
```
```
~ $ curl -sS $log | dln | hist -slice 1:
  1	5	░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
 18	3	░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░
  3	2	░░░░░░░░░░░░░░░░░░░░░░░░░
  9	1	░░░░░░░░░░░░
  6	1	░░░░░░░░░░░░
 49	1	░░░░░░░░░░░░
 23	1	░░░░░░░░░░░░
 22	1	░░░░░░░░░░░░
 21	1	░░░░░░░░░░░░
 59	1	░░░░░░░░░░░░
 11	1	░░░░░░░░░░░░
 78	1	░░░░░░░░░░░░
  4	1	░░░░░░░░░░░░
 42	1	░░░░░░░░░░░░
  2	1	░░░░░░░░░░░░
```
