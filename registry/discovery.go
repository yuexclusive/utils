package registry

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Discovery struct {
	cli   *clientv3.Client             //etcd client
	data  map[string]map[string]string //services
	lock  sync.Mutex
	names []string
}

//NewDiscovery
func NewDiscovery(endpoints []string, name string) *Discovery {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	d := &Discovery{
		cli:  cli,
		data: make(map[string]map[string]string),
	}

	d.watch(name)

	return d
}

//watch
func (s *Discovery) watch(prefix string) error {
	//get data from etcd by prefix
	resp, err := s.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	for _, ev := range resp.Kvs {
		s.put(prefix, string(ev.Key), string(ev.Value))
	}

	//watch
	go func(prefix string) {
		rch := s.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
		log.Printf("watching prefix:%s now...", prefix)
		for wresp := range rch {
			for _, ev := range wresp.Events {
				switch ev.Type {
				case mvccpb.PUT: //修改或者新增
					s.put(prefix, string(ev.Kv.Key), string(ev.Kv.Value))
				case mvccpb.DELETE: //删除
					s.del(prefix, string(ev.Kv.Key))
				}
			}
		}

	}(prefix)
	return nil
}

//put
func (s *Discovery) put(prefix, key, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.data[prefix]; !ok {
		s.data[prefix] = make(map[string]string)
	}
	s.data[prefix][key] = val
	log.Println("put key :", key, "val:", val)
}

//del
func (s *Discovery) del(prefix, key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.data[prefix], key)
	log.Println("del key:", key)
}

//get service list by name
func (s *Discovery) List(name string) []string {
	addrs := make([]string, 0)

	for _, v := range s.data[name] {
		addrs = append(addrs, v)
	}
	return addrs
}

//get service by name
func (s *Discovery) Get(name string) (string, error) {
	m := s.data[name]
	var slice []string
	for _, v := range m {
		slice = append(slice, v)
	}
	length := len(slice)
	if length == 0 {
		return "", fmt.Errorf("service discovery error: can't find any address for name %s", name)
	}
	// rand
	index := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(length)

	index %= length
	return slice[index], nil
}

//Close
func (s *Discovery) Close() error {
	return s.cli.Close()
}
