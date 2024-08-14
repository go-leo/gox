package election

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
)

var count int

var nodeID *int

// 将etcd的地址解析成slice of string
//endpoints := strings.Split(*addr, ",")
// 生成一个etcd的clien
//cli, err := clientv3.New(clientv3.Config{Endpoints: endpoints})
//if err != nil {
//log.Fatal(err)
//}
//defer cli.Close()
// 创建session,如果程序宕机导致session断掉，etcd能检测到
//session, err := concurrency.NewSession(cli)
//defer session.Close()
// 生成一个选举对象。下面主要使用它进行选举和查询等操作
// 另一个方法ResumeElection可以使用既有的leader初始化Election
//e1 := concurrency.NewElection(session, *electName)

// 选leader
func elect(e1 *concurrency.Election, electName string) {
	log.Println("acampaigning for ID:", *nodeID)
	// 调用Campaign方法选主,主的值为value-<主节点ID>-<count>
	if err := e1.Campaign(context.Background(), fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Println(err)
	}
	log.Println("campaigned for ID:", *nodeID)
	count++
}

// 为leader设置新值
func proclaim(e1 *concurrency.Election, electName string) {
	log.Println("proclaiming for ID:", *nodeID)
	// 调用Proclaim方法设置新值,新值为value-<主节点ID>-<count>
	if err := e1.Proclaim(context.Background(), fmt.Sprintf("value-%d-%d", *nodeID, count)); err != nil {
		log.Println(err)
	}
	log.Println("proclaimed for ID:", *nodeID)
	count++
}

// 重新选leader，有可能另外一个节点被选为了主
func resign(e1 *concurrency.Election, electName string) {
	log.Println("resigning for ID:", *nodeID)
	// 调用Resign重新选主
	if err := e1.Resign(context.TODO()); err != nil {
		log.Println(err)
	}
	log.Println("resigned for ID:", *nodeID)
}

// 查询leader的信息
func query(e1 *concurrency.Election, electName string) {
	// 调用Leader返回主的信息，包括key和value等信息
	resp, err := e1.Leader(context.Background())
	if err != nil {
		log.Printf("failed to get the current leader: %v", err)
	}
	log.Println("current leader:", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
}

// 可以直接查询leader的rev信息
func rev(e1 *concurrency.Election, electName string) {
	rev := e1.Rev()
	log.Println("current rev:", rev)
}

// 监控leader的变动。
func watch(e1 *concurrency.Election, electName string) {
	ch := e1.Observe(context.TODO())
	log.Println("start to watch for ID:", *nodeID)
	for i := 0; i < 10; i++ {
		resp := <-ch
		log.Println("leader changed to", string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))
	}
}
