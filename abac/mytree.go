package abac

import (
	"fmt"
)

// 节点（字典树）
type MyNode struct {
	// Data
	children map[interface{}]*MyNode //子节点  interface{}:mynode
	isEnd    bool
}

// 二叉树
type MyTree struct {
	Root *MyNode
}

// 构造访问策略树
func NewTree() *MyTree {
	return &MyTree{Root: NewTreeNode()}
}

// 构造树节点
func NewTreeNode() *MyNode {
	return &MyNode{
		// Data:  data,
		children: make(map[interface{}]*MyNode),
		isEnd:    false,
	}
}

// 向访问策略树插入一个访问策略
func (mt *MyTree) Append(policy []interface{}) {
	node := mt.Root
	for i := 0; i < len(policy); i++ {
		_, ok := node.children[policy[i]]
		// 如果该节点不存在，则构造一个节点
		if !ok {
			node.children[policy[i]] = NewTreeNode()
		}
		node = node.children[policy[i]]
	}
	node.isEnd = true
}

// 搜索树中是否存在指定单词
func (mt *MyTree) Search(tra []interface{}) bool {
	node := mt.Root
	// fmt.Println(tra)
	// 用来记录需要进行回溯的节点和属性
	var preNode []*MyNode
	noNeedAttrIndex := []int{}
	// resultAttr := []string{}
	for i := 0; i < len(tra); i++ {
		// for _, j := range tra {
		// 判断该属性是否满足策略
		_, ok := node.children[tra[i]]
		// fmt.Println(ok, tra[i], node.children)

		// 判断是否有另一条策略不需要该属性
		var key interface{} = nil
		// 比如（w,%!s(<nil>),c2,5,x2,y2）和（w,b2,c2,9,x2,y2），x2 y2 w b2 c2 5其实是可以进行访问的
		_, ok2 := node.children[key]
		// 如果不需要该属性，则进行记录，以便当不满足当前策略时可以回溯到这个策略
		// 因为有可能有多个属性不需要，所以要记录的是这个节点，不能是key
		if ok2 {
			preNode = append(preNode, node.children[key])
			noNeedAttrIndex = append(noNeedAttrIndex, i)
		}
		// fmt.Println(ok, tra[i], ok2, preNode, noNeedAttrIndex)

		// 当前属性并不能满足这个策略时
		if !ok {
			// 如果之前也没有不需要任意一个属性的策略
			if len(noNeedAttrIndex) == 0 {
				// fmt.Println("----", node.children[key])
				return false
			}
			i = noNeedAttrIndex[len(noNeedAttrIndex)-1]
			node = preNode[len(preNode)-1]
			noNeedAttrIndex = noNeedAttrIndex[:len(noNeedAttrIndex)-1]
			preNode = preNode[:len(preNode)-1]
		} else {
			node = node.children[tra[i]]
			// resultAttr = append(resultAttr, fmt.Sprintf("%s", tra[i]))
		}
	}
	return node.isEnd
}

// 判断树中是否有指定前缀的单词
func (mt *MyTree) StartsWith(prefix []interface{}) bool {
	node := mt.Root
	for i := 0; i < len(prefix); i++ {
		_, ok := node.children[prefix[i]]
		if !ok {
			return false
		}
		node = node.children[prefix[i]]
	}
	return true
}

// // 把policy变成tree存储
// // var attr_list = []string{"org", "role", "researce", "action"}
// func PolicyToTree(policies []model.Policy, attr_list []string) *MyTree {
// 	tree := NewTree()
// 	// 第几个op
// 	m := 0
// 	//1. subrules首先生成一个字典树，然后作为一个整体节点再进行和env这些的排序
// 	for _, v := range policies {
// 		// 其中，每条规则应为一个单独的子树
// 		// rule := "{" + policy.SubRules[i] + "\"owner\":\"" + policy.Owner + "\",\"env\":[" + policy.Env.CreatedTime + "," + policy.Env.EndTime + "]}"
// 		rule := "{" + policy.SubRules[i] + "}"
// 		// "role":"u2","action":"r&w","org":"北大六院","researce":"结缔组织疾病"
// 		// fmt.Println("role is: ", i, rule)
// 		values := make(map[string]interface{})
// 		err := json.Unmarshal([]byte(rule), &values)
// 		if err != nil {
// 			fmt.Println("参数化失败: ", err.Error())
// 		}
// 		// values := strings.Split(rule, ",")

// 		// value作为一个树的一个枝
// 		var value []interface{}

// 		// var attr_list = []string{"org", "role", "researce", "action"}
// 		// 根据属性进行顺序的更改
// 		for _, v := range attr_list {
// 			if v == "action" {
// 				ops := strings.Split(values[v].(string), "&")
// 				value = append(value, ops[m])
// 				if m < len(ops)-1 {
// 					m += 1
// 					i -= 1
// 				} else {
// 					m = 0
// 				}
// 			} else {
// 				value = append(value, values[v])

// 			}
// 			// fmt.Println("k is ", k, "个v", value)
// 		}
// 		// fmt.Println("-----第", i, "条value is ", value, "---")
// 		tree.Append(value)
// 	}

// 	return tree
// }

func PreorderPrint(mn *MyNode) {
	if mn == nil {
		return
	}
	// fmt.Print(len(mn.children), " ")
	// fmt.Print(mn.children, " ")
	for k, child := range mn.children {
		fmt.Print(k)
		if !child.isEnd {
			fmt.Print("->")
		} else {
			fmt.Print("\n")
		}
		// fmt.Println("child is ", child, " ")
		// fmt.Println(mn.children[child], " ")
		PreorderPrint(child)
	}
	fmt.Print("=")
}
