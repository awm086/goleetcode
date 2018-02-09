package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

// Next returns Node n's next node

// Append appends node n to list s
func (listNode *ListNode) Insert(i int) {
	if listNode == nil {
		listNode.Val = i
		listNode.Next = nil
	} else {
		for ;; listNode = listNode.Next {
			if listNode.Next == nil {
				listNode.Next = &ListNode{i, nil}
				break
			}
		}
	}
}

func Insert2(listNode *ListNode, i int) *ListNode {
	// If it's nil, create it
	if listNode == nil {
		listNode = &ListNode{i, nil}
	} else {
		l := listNode
		for ; l.Next != nil; l = l.Next {
		}
		l.Next = &ListNode{i, nil}
	}
	return listNode
}

func New() *ListNode {
	return nil
}

func  Insert(listNode *ListNode,i int) {
	if listNode == nil {
		listNode.Val = i
		listNode.Next = nil
	} else {
		for ;; listNode = listNode.Next {
			if listNode.Next == nil {
				listNode.Next = &ListNode{i, nil}
				break
			}
		}
	}
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var resultList ListNode
	carry := 0
	sum := 0
	l1, l2 = l1.Next, l2.Next

	for ;;{

		sum, carry = sumDigits(l1.Val, l2.Val, carry)
		Insert(&resultList, sum)

		if l1.Next == nil && l2.Next == nil {
			break
		}
		// add
		if l1.Next == nil {
			l1.Next = &ListNode{0,nil}
		}
		if  l2.Next == nil {
			l2.Next = &ListNode{0,nil}
		}
		//resultList = *resultList.Next
		l1 = l1.Next
		l2 = l2.Next
	}


	if carry != 0 {
		Insert(&resultList, carry)
	}


	return &resultList

}


func sumDigits(i int, k int, j int) (int, int){
	carryon :=  (i + k +j)/10
	sum := (i + k + j) % 10;
	return sum, carryon
}

func main () {
	// try out insert
	vals := []int{1,3,4,5}
	// init list
	var list ListNode
	for _, i := range vals{
		Insert(&list, i)
	}
	test := New()
	for _, i := range vals{
		test = Insert2(test, i)
	}
	PrintListNodeIter(*test)

	vals2 := []int{1,3,4,5}
	// init list
	var list2 ListNode
	for _, i := range vals2{
		Insert(&list2, i)
	}
	l3 := addTwoNumbers(&list,&list2)
	PrintListNodeIter(list)
	fmt.Println()
	PrintListNodeIter(list2)
	fmt.Println()
	PrintListNodeIter(*l3)

	vals = []int{0, 1}
	// init list
	var list3 ListNode
	for _, i := range vals{
		Insert(&list3, i)
	}

	vals2 = []int{0,1,2}
	// init list
	var list4 ListNode
	for _, i := range vals2{
		Insert(&list4, i)
	}
	fmt.Println()
	PrintListNodeIter(list3)
	fmt.Println()
	PrintListNodeIter(list4)
	result := addTwoNumbers(&list3,&list4)
	fmt.Println()
	PrintListNodeIter(*result)


	vals = []int{9, 9}
	// init list
	var list5 ListNode
	for _, i := range vals{
		Insert(&list5,i)
	}

	vals2 = []int{1}
	// init list
	var list6 ListNode
	for _, i := range vals2{
		Insert(&list6,i)
	}
	fmt.Println()
	PrintListNodeIter(list5)
	fmt.Println()
	PrintListNodeIter(list6)
	result = addTwoNumbers(&list5,&list6)
	fmt.Println()
	PrintListNodeIter(*result)

}

func PrintlistNodeRec (list ListNode) {
	if list.Next == nil {
		fmt.Print(list.Val)
		return
	}
	fmt.Print(list.Val)
	fmt.Print(" ")

	PrintlistNodeRec(*list.Next)
}

func PrintListNodeIter(list ListNode) {
	// skip the first element
	//list = *list.Next
	for _ = list.Val;;{
		if list.Next == nil {
			fmt.Print(list.Val)
			break
		}
		fmt.Print(list.Val)
		fmt.Print(" ")
		list = *list.Next

	}
}
