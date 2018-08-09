GO = go
BIN = ./shorturl

define build
	$(GO) build -o $(BIN) main.go
endef

help:
	@echo "make [help|build|testing|release]"

# 编译 web server
build:
	$(call build)
	chmod +x $(BIN)

# 发布到 10.1.22.162
h162:
	-@rm -rf ./release/*; mkdir -p ./release/conf ./release/logs
	$(call build)
	cp $(BIN) ./release/
	cp ./shorturl.sh ./release/
	cp -r ./conf/app.conf ./release/conf/app.conf
	rsync -avp ./release/* root@10.1.22.162:/data/vhosts/shorturl

# 发布到测试环境
testing:
	-@rm -rf ./release/*; mkdir -p ./release/conf ./release/logs
	$(call build)
	cp $(BIN) ./release/
	cp ./shorturl.sh ./release/
	cp -r ./conf/app.testing.conf ./release/conf/app.conf
	rsync -avp ./release/* testing.myhost.com::projects/shorturl
