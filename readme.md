# dit工具箱(开发中)

**强制要求golang版本 >=1.18**

编译：
```bash
make
```

运行：
```bash
./bin/dit sys overview
// ./bin/dit <module> <cmd>
```

范例：
```bash
# === dit sys ===
# overview 概览：
dit sys overview # 或者
dit sys

# === dit disk ===
# disk 查看硬盘使用
dit disk overview # 或者
dit disk

# === dit mem ===
# mem 查看内存使用
dit mem overview # 或者
dit mem

# === dit file ===
# mv 重命名：
dit file mv aaa.txt bbb.txt
dit file mv '(\w+).txt' '$1'




```

