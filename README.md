# tcp-analyzer

- a golang tcp analyzer 
-  tcp,udp, 4tcp, 6tcp network analyzer using golang and awk.
-  it prepares the files for network struct that i will code tomorrow prior to the binary release. 
- The worst part of the network file is that the file is not being well written using the linux file.io utils and instead of writing and modifying the kernel approach, i implemented the awk for faster code write.

```
gauravsablok@gaurav-sablok ~/Desktop/codecreatede/golang/go-network-struct » 
go run main.go -i tcp6file.txt2024-10-01\ 22:01:16.000454544\ +0200\ CEST\ m=+0.156402181
COMMAND
kdeconnec
gitkraken
gitkraken
trunk
trunk
trunk
thunderbi
cursor
trunk
firefox
2024-10-02 15:39:09.797671689 +0200 CEST m=+0.000426027
PID
2349
3188
3188
5890
5890
5890
90451
95185
96642
102998
2024-10-02 15:39:09.797671689 +0200 CEST m=+0.000426027
```
- binary verion
  ```
  gauravsablok@gaurav-sablok ~/Desktop/codecreatede/golang/tcp-analyzer ±main⚡ » ./tcp-analyzer -h
Creating the network structs

Usage:
  flags [flags]

Flags:
  -h, --help           help for flags
  -i, --input string   network (default "path to the networkfile")
gauravsablok@gaurav-sablok ~/Desktop/codecreatede/golang/tcp-analyzer ±main⚡ » ./tcp-analyzer -i tcp6file.txt2024-10-01\ 22:01:16.000454544\ +0200\ CEST\ m=+0.156402181
```

Gaurav Sablok
