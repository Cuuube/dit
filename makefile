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

mark:
	echo "$(TARGET_DIR)"
	go build -o $(TARGET_DIR)/mark $(CURRENT_DIR)/cmd/mark/main.go
	# sqlite3 $(TARGET_DIR)/mark_sqlite.db
	echo "[SUCCESS] make cmd OK!"

initmark:
	echo "$(TARGET_DIR)"
	go build -o $(TARGET_DIR)/mark $(CURRENT_DIR)/cmd/mark/main.go
	sqlite3 $(TARGET_DIR)/mark_sqlite.db 'create table topic (key text primary_key,name text,type text,create_time text,update_time text);create table content (id integer primary_key AUTO_INCREMENT,topic text,data text,key text, create_time text,update_time text);'
	echo "[SUCCESS] make cmd OK!"
