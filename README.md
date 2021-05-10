## How to embed python in golang (MacOS as an example)
What this repo try to explain
- How to call python code in golang
- How to call python distribution in golang
  

  Reference doc about python in go
  https://www.datadoghq.com/blog/engineering/cgo-and-python/

#### 1. install pkg-config

```
brew install pkg-config
```

#### 2. setup env PKG_CONFIG_PATH
You also can run this command in the terminal directly.
```
export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/System/Library/Frameworks/Python.framework/Versions/2.7/lib/pkgconfig
```

#### 3. build

```
go build main.go 
```

#### 4. Run

Once built, we need to set the PYTHONPATH environment variable to the current working dir so that the import statement
will be able to find the foo.py module. From a shell, the command would look like this:

```
PYTHONPATH=. ./main
```




#### 5. A way to run python3 distribution
##### 1. build python distribution
```
python3 setup.py sdist
```
##### 2. specify the python distribution path while runing go exec file
```
PYTHONPATH=./dist/foo-1.0.0 ./main 
```



