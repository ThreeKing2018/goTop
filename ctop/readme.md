# Ctop
> [https://github.com/bcicen/ctop](https://github.com/bcicen/ctop)

# 使用方法

#### linux使用方法
```
#sudo wget https://github.com/bcicen/ctop/releases/download/v0.7.1/ctop-0.7.1-linux-amd64 -O /usr/local/bin/ctop
#sudo chmod +x /usr/local/bin/ctop

#ctop
```

#### mac使用方法
```
#brew install ctop

or

#sudo curl -Lo /usr/local/bin/ctop https://github.com/bcicen/ctop/releases/download/v0.7.1/ctop-0.7.1-darwin-amd64
#sudo chmod +x /usr/local/bin/ctop

#ctop
```

#### docker使用方法
```
docker run --rm -ti \
  --name=ctop \
  -v /var/run/docker.sock:/var/run/docker.sock \
  quay.io/vektorlab/ctop:latest
```

#效果

![](https://github.com/bcicen/ctop/raw/master/_docs/img/grid.gif)