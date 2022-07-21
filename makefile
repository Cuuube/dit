CURRENT_DIR=$(shell pwd)
TARGET_DIR=$(CURRENT_DIR)/bin

# makefile用法：https://www.ruanyifeng.com/blog/2015/02/make.html

# 不去检查同名文件
.PHONY: cmd
.DEFAULT: all # 输入make执行make all


all: cmd

cmd:
	echo "$(TARGET_DIR)"
	go build -o $(TARGET_DIR)/dit $(CURRENT_DIR)/cmd/dit/main.go
	echo "[SUCCESS] make cmd OK!"