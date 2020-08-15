package main

import "fmt"

type Node struct {
	val      int
	nextNode *Node
}

//func reverseList(head *Node) *Node {
//	if(head.nextNode == nil){
//		return head
//	}
//	p:= reverseList(head.nextNode)
//	head.nextNode.nextNode = head
//	head.nextNode = nil  //避免尾部环 1->2->3->4->5           5->4->3->2->1->2->1
//	return p    //返回尾节点
//}

func reserList(head *Node) *Node {
	var pre *Node = nil

	nextNode := new(Node)
	nextNode = nil

	if head != nil {
		nextNode = head.nextNode
		head.nextNode = pre
		pre = head
		head = nextNode
	}

	//if cur != nil {
	//	pre, cur, cur.nextNode = cur, cur.nextNode, pre
	//}

	return pre
}

func main() {
	head := new(Node)
	head.val = 1
	l2 := new(Node)
	l2.val = 2
	l3 := new(Node)
	l3.val = 3
	l4 := new(Node)
	l4.val = 4
	l5 := new(Node)
	l5.val = 5

	head.nextNode = l2
	l2.nextNode = l3
	l3.nextNode = l4
	l4.nextNode = l5

	for {
		fmt.Println(head.val)
		if head.nextNode == nil {
			break
		}
		head = head.nextNode

	}

	pre := reverseList(head)

	for {
		fmt.Println(head.val)
		if head.nextNode == nil {
			break
		}
		head = head.nextNode

	}

	for {
		fmt.Println(pre.val)
		if pre.nextNode == nil {
			break
		}
		pre = pre.nextNode

	}

}

func reverseList(head *Node) *Node {
	if head == nil || head.nextNode == nil {
		return head
	}
	p := reverseList(head.nextNode)
	head.nextNode.nextNode = head
	head.nextNode = nil
	return p
}

func reverseLinkList(Head *Node) {
	cur := Head.nextNode
	Head.nextNode = nil //可省略，循环体中会指向为nil的prev
	var prev *Node
	var tmp *Node
	for cur != nil {
		tmp = cur.nextNode
		cur.nextNode = prev
		prev = cur
		cur = tmp
	}
	Head.nextNode = prev
}
