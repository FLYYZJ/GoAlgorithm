package mytree

import "fmt"

type color bool

const (
	RED   color = true
	BLACK color = false
)

type rbtnode struct {
	val       int
	lchild    *rbtnode
	rchild    *rbtnode
	nodeColor color    // 颜色，true为红，false为黑
	parent    *rbtnode // 父节点

}

// RedBlackTree 红黑树
type RedBlackTree struct {
	root *rbtnode
}

/*
 * 对红黑树的节点(x)进行左旋转
 *
 * 左旋示意图(对节点x进行左旋)：
 *      px                              px
 *     /                               /
 *    x                               y
 *   /  \      --(左旋)-->           / \                #
 *  lx   y                          x  ry
 *     /   \                       /  \
 *    ly   ry                     lx  ly
 *
 *
 */
func (rbt *RedBlackTree) lRotate(x *rbtnode) {
	// 处理好子节点关系
	y := x.rchild
	x.rchild = y.lchild
	if y.lchild != nil {
		y.lchild.parent = x
	}
	// 处理好父节点关系
	y.parent = x.parent
	if x.parent == nil { // 此时x为根节点
		rbt.root = y // y变成x的父节点，此时y就是根节点
	} else {
		if x.parent.lchild == x { // 变换父子关系
			x.parent.lchild = y
		} else {
			x.parent.rchild = y
		}
	}
	// 更新x 和 y 之间的父子关系
	y.lchild = x // x设为y的左孩子
	x.parent = y // x的父节点设为y
}

/*
 * 对红黑树的节点(y)进行右旋转
 *
 * 右旋示意图(对节点y进行左旋)：
 *            py                               py
 *           /                                /
 *          y                                x
 *         /  \      --(右旋)-->            /  \                     #
 *        x   ry                           lx   y
 *       / \                                   / \                   #
 *      lx  rx                                rx  ry
 *
 */
func (rbt *RedBlackTree) rRotate(y *rbtnode) {
	// 处理好子节点关系
	x := y.lchild
	y.lchild = x.rchild
	if x.rchild != nil {
		x.rchild.parent = y
	}
	// 处理好父节点关系
	x.parent = y.parent
	if y.parent == nil { // y 为根节点的情况
		rbt.root = x
	} else {
		if y == y.parent.rchild {
			y.parent.rchild = x
		} else {
			y.parent.lchild = x
		}
	}
	x.rchild = y
	y.parent = x
}

func (rbt *RedBlackTree) Insert(data int) {
	node := &rbtnode{
		lchild: nil, rchild: nil, parent: nil, val: data,
	}
	var y *rbtnode
	x := rbt.root
	// 二叉搜索树插入步骤
	for x != nil {
		y = x
		if node.val < x.val {
			x = x.lchild
		} else {
			x = x.rchild
		}
	}
	node.parent = y
	// 确定node为y的左右子节点
	if y != nil {
		if node.val < y.val {
			y.lchild = node
		} else {
			y.rchild = node
		}
	} else { // 此时node就是根节点
		rbt.root = node
	}
	node.nodeColor = RED // 默认置为红色
	rbt.insertFixup(node)

}

// 红黑树插入修正
func (rbt *RedBlackTree) insertFixup(node *rbtnode) {
	var parent, gparent *rbtnode
	for parent = node.parent; parent != nil && parent.nodeColor == RED; parent = node.parent {
		gparent = parent.parent
		if parent == gparent.lchild { // 父节点为祖父的左子
			uncle := gparent.rchild
			// case 1 叔叔节点为红色
			if uncle != nil && uncle.nodeColor == RED {
				uncle.nodeColor = BLACK
				parent.nodeColor = BLACK
				gparent.nodeColor = RED
				node = gparent
				continue
			}
			// case 2 叔叔节点为黑色或NIL，当前节点为右子， 将情况转为case3
			if parent.rchild == node {
				rbt.lRotate(parent)
				parent, node = node, parent
			}
			// case 3 叔叔节点为黑色或NIL，当前节点为左子
			parent.nodeColor = BLACK
			gparent.nodeColor = RED
			rbt.rRotate(gparent)
		} else { // 父节点为祖父的右子
			uncle := gparent.lchild
			if uncle != nil && uncle.nodeColor == RED {
				uncle.nodeColor = BLACK
				parent.nodeColor = BLACK
				gparent.nodeColor = RED
				node = gparent
				continue
			}
			if parent.lchild == node {
				rbt.rRotate(parent)
				parent, node = node, parent
			}
			parent.nodeColor = BLACK
			gparent.nodeColor = RED
			rbt.lRotate(gparent)
		}
	}
	rbt.root.nodeColor = BLACK
}

// Remove 删除节点
func (rbt *RedBlackTree) Remove(data int) {

	node := rbt.iterativeSearch(rbt.root, data) // 找到待删除节点
	if node == nil {
		fmt.Println("在该红黑树中找不到删除节点")
	}
	if (node.lchild != nil) && (node.rchild != nil) { // 被删除节点的左右节点均不为空
		replace := node.rchild // 找到右子树的最小节点来替换
		for replace.lchild != nil {
			replace = replace.lchild
		}
		if node.parent != nil {
			if node.parent.lchild == node {
				node.parent.lchild = replace
			} else {
				node.parent.rchild = replace
			}
		} else {
			rbt.root = replace // 待删除节点为根节点
		}
		child := replace.rchild
		parent := replace.parent
		// 这里涉及到一个假设，即假设删除节点位置还有一层黑色，因此如果替代节点为黑色，则替代节点的右子树要比左子树少一个黑节点
		rep_color := replace.nodeColor
		if parent == node { // 如果删除的node恰好是replace的父节点
			parent = replace // 将作为其兄弟节点的父节点
		} else {
			if child != nil {
				child.parent = parent // replace给自己的子节点分配新父节点
			}
			parent.lchild = child
			replace.rchild = node.rchild // 分配删除节点的右子树关系
			node.rchild.parent = replace
		}
		replace.parent = node.parent       // 分配删除节点的左子树关系
		replace.nodeColor = node.nodeColor // 替换节点设置为删除节点的颜色
		replace.lchild = node.lchild
		node.lchild.parent = replace
		if rep_color == BLACK {
			rbt.removeFixup(child, parent)
		}
		return
	}
	// 当node仅有一个子节点或者只是一个叶子节点时
	var child, parent *rbtnode
	if node.lchild != nil {
		child = node.lchild
	} else {
		child = node.rchild
	}
	parent = node.parent
	node_color := node.nodeColor

	if child != nil {
		child.parent = parent
	}
	if parent != nil {
		if parent.lchild == node {
			parent.lchild = child
		} else {
			parent.rchild = child
		}
	}

	if node_color == BLACK {
		rbt.removeFixup(child, parent)
	}

}

func (rbt *RedBlackTree) removeFixup(node *rbtnode, node_parent *rbtnode) {
	var other *rbtnode
	for (node == nil || node.nodeColor == BLACK) && node != rbt.root { // 待处理节点未到根节点，且为黑
		if node_parent.lchild == node { // 删除节点为左子
			other = node_parent.rchild // 找到兄弟节点
			// case 1 兄弟节点为红色，可以转case 2，3，4
			if other.nodeColor == RED {
				other.nodeColor = BLACK
				node_parent.nodeColor = RED
				rbt.lRotate(node_parent)
				other = node_parent.rchild
			}
			if (other.lchild == nil || other.lchild.nodeColor == BLACK) && (other.rchild == nil || other.rchild.nodeColor == BLACK) {
				// case 2 兄弟节点为黑色，且兄弟节点的子节点均为黑色
				other.nodeColor = RED
				node = node_parent
				node_parent = node.parent
			} else {
				if other.rchild == nil || other.rchild.nodeColor == BLACK {
					// case 3 兄弟节点为黑色，且左孩子为红，右孩子为黑，可以转case4
					other.lchild.nodeColor = BLACK
					other.nodeColor = RED
					rbt.rRotate(other)
					other = node_parent.rchild
				}
				// case 4 兄弟节点为黑色，且w的右孩子为红色，左孩子为任意颜色
				other.nodeColor = node_parent.nodeColor
				node_parent.nodeColor = BLACK
				other.rchild.nodeColor = BLACK
				rbt.lRotate(node_parent)
				node = rbt.root
				break
			}
		} else {
			other = node_parent.lchild
			if other.nodeColor == RED {
				other.nodeColor = BLACK
				node_parent.nodeColor = RED
				rbt.rRotate(node_parent)
				other = node_parent.lchild
			}
			if (other.lchild == nil || other.lchild.nodeColor == BLACK) && (other.rchild == nil || other.rchild.nodeColor == BLACK) {
				other.nodeColor = RED
				node = node_parent
				node_parent = node.parent
			} else {
				if other.lchild == nil || other.lchild.nodeColor == BLACK {
					other.rchild.nodeColor = BLACK
					other.nodeColor = RED
					rbt.lRotate(other)
					other = node_parent.lchild
				}
				other.nodeColor = node_parent.nodeColor
				node_parent.nodeColor = BLACK
				other.lchild.nodeColor = BLACK
				rbt.rRotate(node_parent)
				node = rbt.root
				break
			}
		}
	}
	if node != nil {
		node.nodeColor = BLACK
	}
}

// PreVisit 先序遍历红黑树
func (rbt *RedBlackTree) PreVisit() {
	if rbt.root == nil {
		fmt.Println("红黑树为空")
		return
	}
	rbt.preVisit(rbt.root)
}

func (rbt *RedBlackTree) preVisit(root *rbtnode) {
	if root == nil {
		return
	}
	fmt.Printf("%d  ", root.val)
	rbt.preVisit(root.lchild)
	rbt.preVisit(root.rchild)
}

// MidVisit 中序遍历AVL树
func (rbt *RedBlackTree) MidVisit() {
	if rbt.root == nil {
		fmt.Println("红黑树为空")
		return
	}
	rbt.midVisit(rbt.root)
}

func (rbt *RedBlackTree) midVisit(root *rbtnode) {
	if root == nil {
		return
	}
	rbt.midVisit(root.lchild)
	fmt.Printf("%d  ", root.val)
	rbt.midVisit(root.rchild)
}

func (rbt *RedBlackTree) getMinNode(root *rbtnode) (*rbtnode, error) {
	if root == nil {
		return nil, fmt.Errorf("树为空")
	}
	for root.lchild != nil {
		root = root.lchild
	}
	return root, nil
}

func (rbt *RedBlackTree) getMaxNode(root *rbtnode) (*rbtnode, error) {
	if root == nil {
		return nil, fmt.Errorf("树为空")
	}
	for root.rchild != nil {
		root = root.rchild
	}
	return root, nil
}

func (rbt *RedBlackTree) search(root *rbtnode, val int) *rbtnode {
	if root == nil || root.val == val {
		return root
	}
	if val < root.val {
		return rbt.search(root.lchild, val)
	}
	return rbt.search(root.rchild, val)
}

func (rbt *RedBlackTree) iterativeSearch(root *rbtnode, val int) *rbtnode {
	for root != nil && root.val != val {
		if val < root.val {
			root = root.lchild
		} else {
			root = root.rchild
		}
	}
	return root
}
