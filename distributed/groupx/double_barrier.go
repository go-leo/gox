package groupx

import (
	"bufio"
	"flag"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		addr        = flag.String("addr", "http://127.0.0.1:2379", "etcd addresses")
		barrierName = flag.String("name", "my-test-doublebarrier", "barrier name")
		count       = flag.Int("c", 2, "")
	)
	flag.Parse()
	// 解析etcd地址
	endpoints := strings.Split(*addr, ",")
	// 创建etcd的client
	cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	// 创建session
	s1, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s1.Close()
	// 创建/获取栅栏
	b := recipe.NewDoubleBarrier(s1, *barrierName, *count)
	// 从命令行读取命令
	consolescanner := bufio.NewScanner(os.Stdin)
	for consolescanner.Scan() {
		action := consolescanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case "enter": // 持有这个barrier
			// 当调用者调用Enter时，会被阻塞住，直到一共有count（初始化这个栅栏的时候设定的值）个节点调用了Enter，这count个被阻塞的节点才能继续执行。所以，你可以利用它编排一组节点，让这些节点在同一个时刻开始执行任务。
			b.Enter()
			fmt.Println("enter")
		case "leave": // 释放这个barrier
			// 同理，如果你想让一组节点在同一个时刻完成任务，就可以调用Leave方法。节点调用Leave方法的时候，会被阻塞，直到有count个节点，都调用了Leave方法，这些节点才能继续执行。
			b.Leave()
			fmt.Println("leave")
		case "quit", "exit": //退出
			return
		default:
			fmt.Println("unknown action")
		}
	}
}
