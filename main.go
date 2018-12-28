package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"
)

const (
	EDEGE_NAME      = "parent"
	NODE_NAME       = "md5"
	TYPE_NODES_NAME = "nodes"
	TYPE_EDGES_NAME = "edges"
)

var maxUid int64
var num int64
var c int
var tp string

type ProcChain struct {
	// how many groutine num for run
	c int
	// which type should build  , the value is 'nodes' or 'edges'
	tp string

	//the uid random boundary value
	maxUid int64

	//  how many data to build , if this.total > this.num then stop process
	num int64
	// has build total length
	total int64
}

func main() {

	flag.Int64Var(&maxUid, "maxuid", 650000000, "the uid random boundary value") // 0.65 Billion
	flag.Int64Var(&num, "num", 10000, "how many data you want to build")
	flag.StringVar(&tp, "type", TYPE_NODES_NAME, "which type should build  , the value you can select 'nodes' or 'edges' ")
	flag.IntVar(&c, "c", 100, "how many goroutine num for run")
	flag.Parse()
	if c <= 0 {
		c = 100
	}
	if num <= 0 {
		num = 10000
	}
	if maxUid <= 0 {
		panic("should set maxUid gt 0")
	}

	p := &ProcChain{
		c:      c,
		tp:     tp,
		maxUid: maxUid,
		num:    num,
	}

	p.start() //start run
}

func (p *ProcChain) start() {

	switch p.tp {
	case TYPE_NODES_NAME:
		for {
			if p.total >= p.num {
				break
			}
			p.startNodes()
		}
	case TYPE_EDGES_NAME:
		for {
			if p.total >= p.num {
				break
			}
			p.startEdges()
		}
	}

}

func (p *ProcChain) startNodes() {
	ch := make(chan string, p.c)
	defer close(ch)
	for i := 0; i < p.c; i++ {
		go func() {
			ch <- p.outputNodes()

		}()
	}
	for i := 0; i < p.c; i++ {
		if v, ok := <-ch; ok {
			atomic.AddInt64(&p.total, 1)
			fmt.Println(v)
		}
	}

}

func (p *ProcChain) startEdges() {
	ch := make(chan []string, p.c)
	defer close(ch)
	for i := 0; i < p.c; i++ {
		go func() {
			ch <- p.outputEdegs()

		}()
	}
	for i := 0; i < p.c; i++ {
		if v, ok := <-ch; ok {
			atomic.AddInt64(&p.total, int64(len(v)))
			fmt.Println(strings.Join(v, "\n"))
		}
	}

}

// rand num create edegs
func (p *ProcChain) outputEdegs() (data []string) {
	var n int64 = 5
	randEdgeNum := int(Rand(1, n))
	data = make([]string, 0, n)
	for i := 0; i < randEdgeNum; i++ {
		start := assignUid(p.maxUid)
		end := assignUid(p.maxUid)
		data = append(data, edgesRdf(start, end))
	}
	return
}

func (p *ProcChain) outputNodes() string {
	uid := assignUid(p.maxUid)
	return nodeRdf(uid)
}

// return edges rdf format
func edgesRdf(start, end string) string {
	return fmt.Sprintf("<%s>\t<%s>\t<%s>\t.", start, EDEGE_NAME, end)
}

// return node rdf format
// when use md5 the same uid will create the same value
func nodeRdf(uid string) string {
	return fmt.Sprintf("<%s>\t<%s>\t\"%s\"\t.", uid, NODE_NAME, md5Bytes([]byte(uid)))
}

// rand to build uid
func assignUid(max int64) string {
	uid := Rand(1, max)
	return fmt.Sprintf("0x%x", uid)
}

// Rand between [min ~ max]
func Rand(min, max int64) int64 {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	v := max - min
	if v <= 0 {
		v = 1
	}
	return r.Int63n(v) + min
}

// md5
func md5Bytes(v []byte) string {
	return fmt.Sprintf("%x", md5.Sum(v))
}
