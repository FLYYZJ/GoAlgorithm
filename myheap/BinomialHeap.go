package myheap

import "fmt"

type binomialNode struct {
	val    int
	degree int
	lchild *binomialNode // 左子节点
	parent *binomialNode // 父节点
	next   *binomialNode // 链表上的连接节点
}

// BinomialHeap 二项堆
type BinomialHeap struct {
	root *binomialNode
}

/*
 * 将h1, h2中的根表合并成一个按度数递增的链表，返回合并后的根节点
 * 两个二项堆的根链表合并成一个链表，合并后的新链表按照'节点的度数'单调递增排序
 */
func (bih *BinomialHeap) merge(h1 *binomialNode, h2 *binomialNode) *binomialNode {
	var head *binomialNode
	pos := &head // 使用这个二级指针可以减少一个判断，即找出head节点的判断，后续的写法可以统一
	for h1 != nil && h2 != nil {
		// 找最小度的节点作为首节点
		if h1.degree < h2.degree {
			*pos, h1 = h1, h1.next
		} else {
			*pos, h2 = h2, h2.next
		}
		pos = &((*pos).next)
	}
	if h1 != nil {
		*pos = h1
	} else {
		*pos = h2
	}
	return head
}

// 连接child到指定二项树中
func (bih *BinomialHeap) link(child *binomialNode, root *binomialNode) {
	child.parent = root
	// 将原root的左孩子记为自己的兄弟节点，这里的next可视为兄弟
	// 二项树的内部也维持着一个链表，链接root下的各个兄弟节点
	child.next = root.lchild
	root.lchild = child
	root.degree++
}

// Union 合并两个二项堆
func (bih *BinomialHeap) Union(other *BinomialHeap) {
	bih.root = bih.union(bih.root, other.root)
}

// 合并操作主程序，删除/插入操作相关
func (bih *BinomialHeap) union(h1 *binomialNode, h2 *binomialNode) *binomialNode {
	var heap, prev_x, x, next_x *binomialNode // x当前节点
	heap = bih.merge(h1, h2)                  // 按度排列，最小度在前
	if heap == nil {
		return nil
	}
	prev_x, x = nil, heap
	next_x = x.next

	for next_x != nil {
		if (x.degree != next_x.degree) || (next_x.next != nil && next_x.degree == next_x.next.degree) {
			// 当前度不等于下一个节点的度 或者  下一个节点的度和下下个节点的度一样（放到下一个节点再合并，因为度较小的节点要在前面）
			prev_x, x = x, next_x
		} else if x.val <= next_x.val { // 当前优先级高于下一优先级
			x.next = next_x.next
			bih.link(next_x, x) // 将next_x作为x的子节点
		} else { // 当前优先级低于下一优先级，此时next_x会作为合并二项树的根节点
			if prev_x == nil { // 处于链表头，则重置链表首节点
				heap = next_x
			} else {
				prev_x = next_x
			}
			bih.link(x, next_x)
			x = next_x
		}
		next_x = x.next
	}
	return heap
}

// Insert 二项堆插入
func (bih *BinomialHeap) Insert(data int) {
	if bih.search(bih.root, data) != nil {
		fmt.Println("当前节点已存在，不能重复插入")
	}
	bih.root = bih.union(bih.root, &binomialNode{lchild: nil, val: data, next: nil, parent: nil})
}

//
func (bih *BinomialHeap) search(root *binomialNode, data int) *binomialNode {
	for root != nil {
		if root.val == data {
			return root
		}
		if child := bih.search(root.lchild, data); child != nil {
			return child
		}
		root = root.next
	}
	return nil
}

// 删除操作相关
func (bih *BinomialHeap) reverse(root *binomialNode) *binomialNode {
	// 此处的入参 root为某个二项树的左孩子，后续会变成新二项堆的尾结点
	// 因为一开始从链表尾部开始构建，因此该操作也被称为二项树反转，会得到新的二项堆
	if root == nil {
		return nil
	}
	root.parent = nil // 断开子树和父节点（根节点）
	var next, tail *binomialNode
	for root.next != nil { // 拆开二项树的每个节点
		next = root.next // 获得兄弟节点
		root.next = tail
		tail = root
		root = next
		root.parent = nil
	}
	root.next = tail
	return root
}

// Remove 删除二项堆中的某个节点
func (bih *BinomialHeap) Remove(data int) {
	if bih.root == nil {
		fmt.Println("当前二项堆为空")
		return
	}
	node := bih.search(bih.root, data) // 找到待删除的节点
	if node == nil {
		fmt.Println("找不到待删除的节点")
		return
	}
	// 将待删除节点移到对应二项树的根
	parent := node.parent
	for parent != nil {
		node.val, parent.val = parent.val, node.val
		node = parent
		parent = node.parent
	}
	// 找到对应的二项树
	prev, pos := bih.root, bih.root
	for pos != node {
		prev = pos
		pos = pos.next
	}
	// 移除该二项树
	if prev != nil {
		prev.next = node.next
	} else {
		bih.root = node.next
	}
	// 对应二项树删除指定节点后，将该二项树重新整合入二项堆中
	bih.root = bih.union(bih.root, bih.reverse(node.lchild))
}

// 二项堆更新，修改某个节点的值

// 二项树中节点减小后调整
func (bih *BinomialHeap) decrease(node *binomialNode, data int) {
	if data >= node.val || bih.search(bih.root, data) == nil {
		fmt.Println("指定节点已存在， 或当前节点值比预设值要低")
		return
	}
	node.val = data
	c, p := node, node.parent
	for p != nil && c.val < p.val {
		p.val, c.val = c.val, p.val
		c, p = p, c.parent
	}
}

// 二项树中节点加大后调整
func (bih *BinomialHeap) increase(node *binomialNode, data int) {
	if data <= node.val || bih.search(bih.root, data) != nil {
		fmt.Println("指定节点已存在， 或当前节点值比预设值要高")
		return
	}
	node.val = data
	var least *binomialNode
	cur, child := node, node.lchild
	for child != nil {
		// 每个二项树都是最小堆
		// 每次均在子节点及子节点的兄弟中选取最小节点进行替换
		if cur.val > child.val {
			least = child
			for child.next != nil {
				if least.val > child.next.val {
					least = child.next // 找兄弟中的最小值
				}
				child = child.next
			}
			least.val, cur.val = cur.val, least.val
			cur, child = least, cur.lchild
		} else {
			child = child.next
		}
	}

}

// GetMin 获得最小值
func (bih *BinomialHeap) GetMin() (int, error) {
	if bih.root == nil {
		fmt.Println("当期二项堆为空")
		return -1, fmt.Errorf("heap is empty")
	}
	minNode := bih.getMinNode()
	return minNode.val, nil
}

// 返回最小值对应的节点，以及其前缀
func (bih *BinomialHeap) getMinNode() *binomialNode {
	cur, minNode := bih.root, bih.root
	for cur != nil {
		if cur.val < minNode.val {
			minNode = cur
		}
		cur = cur.next
	}
	return minNode
}

// PopMin 获得最小值，并删除这个节点
func (bih *BinomialHeap) PopMin() (int, error) {
	if bih.root == nil {
		fmt.Println("当期二项堆为空")
		return -1, fmt.Errorf("heap is empty")
	}
	minNode := bih.getMinNode()
	result := minNode.val
	bih.Remove(minNode.val)
	return result, nil
}

// PrintHeap 输出heap
func (bih *BinomialHeap) PrintHeap() {
	if bih.root == nil {
		fmt.Println("当期二项堆为空")
	}
	p, treeCount := bih.root, 1
	for p != nil {
		fmt.Printf("第 %d 二项树度为 %d, 根为 %d\n", treeCount, p.degree, p.val)
		bih.printHeap(p.lchild, p, 1)
		p = p.next
	}

}

func (bih *BinomialHeap) printHeap(node *binomialNode, prev *binomialNode, direction int) {
	for node != nil {
		if direction == 1 {
			fmt.Printf("\t节点%02d(度为 %02d) 为 %02d的子节点\n", node.val, node.degree, prev.val)
		} else {
			fmt.Printf("\t节点%02d(度为 %02d) 为 %02d的兄弟节点\n", node.val, node.degree, prev.val)
		}
		if node.lchild != nil {
			bih.printHeap(node.lchild, node, 1)
		}
		prev, node, direction = node, node.next, 2
	}

}
