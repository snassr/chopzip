## chopzip, a zipping tool


## Instructions
```bash
# create unzipped directory (if any)
mkdir -p /go/src/github.com/snassr/chopzip/unzipped
# delete old files (if any)
rm /go/src/github.com/snassr/chopzip/unzipped/*
# cd to chopzip directory
cd /go/src/github.com/snassr/chopzip
# zip our example hello.txt file
bin/chopzip -in ./hello.txt -out ./unzipped/zipped.zip
# cd to unzipped folder
cd /go/src/github.com/snassr/chopzip/unzipped
# use unzip
unzip zipped.zip
# check if our file was written
cat /go/src/github.com/snassr/chopzip/unzipped/zippedfile && echo ""
```

### run as go file
```bash
cd /go/src/github.com/snassr/chopzip && go run main.go -in ./hello.txt -out ./unzipped/zipped.zip
```

### run as bin
```bash
cd /go/src/github.com/snassr/chopzip && bin/chopzip -in ./hello.txt -out ./unzipped/zipped.zip
```

### build
```bash
cd /go/src/github.com/snassr/chopzip && go build -o bin/chopzip main.go
```

## Package Basket (can only use these go packages!)
```txt
---------------Basket 2---------------
text
os
sync
archive
unicode
>>> Staple Pantry <<<
strings
math
errors
time
testing
flag
io
context
log
```
