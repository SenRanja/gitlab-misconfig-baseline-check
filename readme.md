

### go-bindata
使用方法：
```shell
go-bindata -o=bindata/bindata.go -ignore="\\.DS_Store|desktop.ini|README.md" -pkg=bindata -prefix=rules rules/... 
```

### 子模块同步
```shell
git clone --recursive https://github.com/xanzy/go-gitlab gitlab
```

### 参考
- [https://jaycechant.info/2020/go-bindata-golang-static-resources-embedding/](https://jaycechant.info/2020/go-bindata-golang-static-resources-embedding/)