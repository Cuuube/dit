package main

import (
	"errors"
	"os"
	"strings"

	"github.com/Cuuube/dit/pkg/cli"
	"github.com/Cuuube/dit/pkg/mark"
)

var (
	topic *mark.Topic
	err   error
)

func main() {
	args := os.Args
	if len(args) < 2 {
		cli.Println("参数不足")
		return
	}

	srv := mark.NewMarkService()

	topicKey := cleanString(args[1])
	switch topicKey {

	// 特殊模式：展示所有topic
	case "list", "li":
		topics, err := srv.GetTopics()
		if err != nil {
			cli.Println("获取数据错误", err)
			return
		}
		table := make([][]string, len(topics))
		for _, t := range topics {
			table = append(table, []string{t.Name, t.Key})
		}
		cli.PrintTableWithHeader([]string{"Topic Name", "Topic Key"}, table)
		return

	// 创建Topic
	case "create", "ins", "insert":
		topicKey = ""
		if len(args) >= 3 {
			topicKey = cleanString(args[2])
		}
		err := goCreateTopic(srv, topicKey)
		if err != nil {
			cli.Println("删除topic失败：", err)
		}
		cli.Println("创建成功！")
		return

	// 删除key模式
	case "delete", "del", "remove":
		if len(args) < 3 {
			cli.Println("参数不足")
			return
		}
		topicKey = cleanString(args[2])
		err := srv.DeleteTopic(topicKey)
		if err != nil {
			cli.Println("删除topic失败：", err)
		}
		return

	default: // 校验topic是否存在
		// 查找topic合法性
		topic, err = srv.GetTopicByTopicKey(topicKey)
		if err != nil || topic.Key == "" {
			err = goCreateTopic(srv, topicKey)
			if err != nil {
				cli.Println("操作终止")
				return
			}
		}
	}

	if len(args) < 3 {
		cli.Println("参数不足")
		return
	}
	data := cleanString(args[2])
	switch data {

	// 特殊模式：输出list
	case "list", "li":
		li, err := srv.GetContents(topic.Key)
		if err != nil {
			cli.Println("获取数据错误", err)
			return
		}
		table := make([][]string, len(li))
		for _, l := range li {
			table = append(table, []string{l.Key, l.Data})
		}
		cli.PrintTableWithHeader([]string{"Key", topic.Key + " Data"}, table)
		return

	// 删除key模式
	case "delete", "del", "remove":
		if len(args) < 4 {
			cli.Println("参数不足")
			return
		}
		dateKey := args[3]
		err := srv.DeleteContent(topic.Key, dateKey)
		if err != nil {
			cli.Println("删除key失败：", err)
		}
		return

	// 输出统计
	case "analysis":
		cli.Println("功能开发中，敬请期待")
		return

	default:
		var key, value string
		switch topic.Type {
		case mark.TopicTypeDaliy:
			key = mark.TodayDatekey()
			value = data
		case mark.TopicTypeMinutely:
			key = mark.CurrentMinuteDatekey()
			value = data
		case mark.TopicTypeSecondly:
			key = mark.CurrentSecondDatekey()
			value = data
		case mark.TopicTypeKV:
			if len(args) < 4 {
				cli.Println("参数不足")
				return
			}
			key = args[2]
			value = args[3]
		}
		// 执行upsert
		err := srv.UpsertContent(topic.Key, key, value)
		if err != nil {
			cli.Println("记录出错：", err)
		} else {
			cli.Println("记录成功")
		}
	}
}

var (
	TopicTypeSelectMap = map[string]string{
		"1": mark.TopicTypeKV,
		"2": mark.TopicTypeDaliy,
		"3": mark.TopicTypeMinutely,
		"4": mark.TopicTypeSecondly,
	}
)

func goCreateTopic(srv mark.MarkService, topicKey string) error {
	for topicKey == "" {
		cli.Println("请键入【topicid】：")
		topicKey = cli.ReadInput()
		if topicKey == "n" {
			return errors.New("操作终止")
		}
	}
	cli.Println("请键入【名称】：")
	name := strings.ToLower(cli.ReadInput())
	if name == "" || name == "n" {
		return errors.New("操作终止")
	}
	cli.Println("请输入Topic类型：\n1.自由输入（需要手动指定key和value）\n2.每日限一条（不需要指定key）\n3.每分钟限一条（不需要指定key）\n4.每秒限一条（不需要指定key）")
	typInp := strings.ToLower(cli.ReadInput())
	typ, valid := TopicTypeSelectMap[typInp]
	if !valid {
		return errors.New("输入有误，操作终止")
	}

	return srv.CreateTopic(topicKey, name, typ)
}

func cleanString(str string) string {
	return strings.Trim(str, " 	\n")
}
