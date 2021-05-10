## How to embed python in golang (MacOS as an example)
#### 1. install pkg-config
```
brew install pkg-config
```


#### 2. setup env PKG_CONFIG_PATH
```
export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/System/Library/Frameworks/Python.framework/Versions/2.7/lib/pkgconfig
```


#### 3. Go build and run
Once built, we need to set the PYTHONPATH environment variable to the current working dir so that the import statement will be able to find the foo.py module. From a shell, the command would look like this:
```
go build main.go && PYTHONPATH=. ./main
```
