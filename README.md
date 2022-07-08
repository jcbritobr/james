![James](images/james.png)
# James

James is a cli for create gnome desktop launchers. It can add all entries in .desktop files, as its multiple values.

* Build
```
$ go build -v -ldflags "-s -w"
```

* Test
```
$ go test dengine/* -v
=== RUN   TestItWorks
--- PASS: TestItWorks (0.00s)
=== RUN   TestBuildFileData
--- PASS: TestBuildFileData (0.00s)
=== RUN   TestNewDesktopData
--- PASS: TestNewDesktopData (0.00s)
=== RUN   TestFillDesktopDataBuffer
--- PASS: TestFillDesktopDataBuffer (0.00s)
PASS
ok      command-line-arguments  0.006s
```
