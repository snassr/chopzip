## chopzip, an unzipping tool


## Instructions
```bash
# delete old files
rm /go/src/github.com/snassr/choppd/unzipped/*
# cd to choppd directory
cd /go/src/github.com/snassr/choppd
# zip our example hello.txt file
bin/chopzip -in ./hello.txt -out ./unzipped/zipped.zip
# cd to unzipped folder
cd /go/src/github.com/snassr/choppd/unzipped
# use unzip
unzip zipped.zip
# check if our file was written
cat /go/src/github.com/snassr/choppd/unzipped/zippedfile && echo ""
```

### run as go file
```bash
cd /go/src/github.com/snassr/choppd && go run main.go -in ./hello.txt -out ./unzipped/zipped.zip
```

### run as bin
```bash
cd /go/src/github.com/snassr/choppd && bin/chopzip -in ./hello.txt -out ./unzipped/zipped.zip
```

### build
```bash
cd /go/src/github.com/snassr/choppd && go build -o bin/chopzip main.go
```

### basket
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