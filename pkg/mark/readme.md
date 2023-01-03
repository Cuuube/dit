# mark

命令行速记工具

## 概念

### 1、topic

话题。

想要创建速记，需要指定建在哪个话题下。
话题支持以下几种类型：

1. kv类型。速记时需要同时指定key和value。例如：`mark words apple 苹果`
2. 每日一条型。速记时不需要指定key（按照今日的datekey作为key）。例如：`mark weight 82.2kg`
3. 每分一条型。速记时不需要指定key（按照当前分钟的key作为key）。例如：`mark temperature 37度`
4. 每秒一条型。速记时不需要指定key（按照当前秒的key作为key）。最为方便、精细，容易使用。例如：`mark todo 一会倒垃圾`

**注意：除了kv类型，其他类型都可能让你覆写旧版本内容（根据选择的时间粒度），造成数据丢失！**

## 命令用法

### 1、话题（topic）篇：

### 创建topic：

`mark`

创建时会选择topic类型。

#### 列出topic

`mark list`

#### 删除topic：

`mark delete <topicKey>`
<!-- 
#### rename topic：

`mark rename <topicKey> <newTopicKey>` -->

### 2、速记内容（content）篇：

#### 速记

`mark <topicKey> <data>`

如果是kv类型：
`mark <topicKey> <key> <data>`

#### 查看话题下所有内容：

`mark <topicKey> list`

#### 删除速记

`mark <topicKey> delete <contentKey>`

<!-- #### 统计：

`mark <topicKey> analysis` -->
