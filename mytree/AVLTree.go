package mytree

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type avlnode struct {
	val    int
	lchild *avlnode
	rchild *avlnode
	height int // 树高
}

// AVLTree AVL树
type AVLTree struct {
	root *avlnode
}

func (avlt *AVLTree) getHeight(node *avlnode) int {
	if node != nil {
		return node.height
	}
	return 0
}

func (avlt *AVLTree) llRotate(k2 *avlnode) *avlnode {
	k1 := k2.lchild
	k2.lchild = k1.rchild
	k1.rchild = k2
	k2.height = max(avlt.getHeight(k2.lchild), avlt.getHeight(k2.rchild)) + 1
	k1.height = max(avlt.getHeight(k1.lchild), k2.height) + 1
	return k1
}
func (avlt *AVLTree) rrRotate(k1 *avlnode) *avlnode {
	k2 := k1.rchild
	k1.rchild = k2.lchild
	k2.lchild = k1
	k1.height = max(avlt.getHeight(k1.lchild), avlt.getHeight(k2.rchild)) + 1
	k2.height = max(avlt.getHeight(k1), avlt.getHeight(k2.rchild)) + 1
	return k2
}
func (avlt *AVLTree) rlRotate(k1 *avlnode) *avlnode {
	k1.rchild = avlt.llRotate(k1.rchild)
	return avlt.rrRotate(k1)
}
func (avlt *AVLTree) lrRotate(k3 *avlnode) *avlnode {
	k3.lchild = avlt.rrRotate(k3.lchild)
	return avlt.llRotate(k3)
}

// Insert AVL树插入
func (avlt *AVLTree) Insert(data int) {
	avlt.root = avlt.insert(avlt.root, data)
}

func (avlt *AVLTree) insert(root *avlnode, data int) *avlnode {
	if root == nil {
		root = &avlnode{val: data, lchild: nil, rchild: nil, height: 0} // 新建节点的高度定为0
	} else { // avl树和二叉搜索具有一些共同的性质
		if data < root.val {
			root.lchild = avlt.insert(root.lchild, data)
			if avlt.getHeight(root.lchild)-avlt.getHeight(root.rchild) >= 2 { // 将节点插入左子树，自然要判断左子树的高度是否大于右子树，下同
				if data < root.lchild.val {
					root = avlt.llRotate(root)
				} else {
					root = avlt.lrRotate(root)
				}
			}
		} else if data > root.val {
			root.rchild = avlt.insert(root.rchild, data)
			if avlt.getHeight(root.rchild)-avlt.getHeight(root.lchild) >= 2 {
				if data > root.rchild.val {
					root = avlt.rrRotate(root)
				} else {
					root = avlt.rlRotate(root)
				}
			}
		} else {
			fmt.Println("节点存在，不允许添加节点")
		}
	}
	root.height = max(avlt.getHeight(root.lchild), avlt.getHeight(root.rchild)) + 1
	return root
}

// Remove 移除avl树中的节点
func (avlt *AVLTree) Remove(data int) (int, error) {
	removenode := avlt.search(avlt.root, data)
	if removenode == nil {
		return -1, fmt.Errorf("找不到删除节点")
	}
	removenode = avlt.remove(avlt.root, removenode)
	return 0, nil
}

func (avlt *AVLTree) remove(root *avlnode, z *avlnode) *avlnode { // z 为根节点
	if root == nil || z == nil {
		return nil
	}
	if z.val < root.val {
		root.lchild = avlt.remove(root.lchild, z)                         // 删除左子树
		if avlt.getHeight(root.rchild)-avlt.getHeight(root.lchild) >= 2 { // 判定平衡
			r := root.rchild
			if avlt.getHeight(r.lchild) > avlt.getHeight(r.rchild) { // 判定旋转方式
				root = avlt.rlRotate(root)
			} else {
				root = avlt.rrRotate(root)
			}
		}
	} else if z.val > root.val {
		root.rchild = avlt.remove(root.rchild, z)
		if avlt.getHeight(root.lchild)-avlt.getHeight(root.rchild) >= 2 {
			l := root.lchild
			if avlt.getHeight(l.rchild) > avlt.getHeight(l.lchild) {
				root = avlt.lrRotate(root)
			} else {
				root = avlt.llRotate(root)
			}
		}
	} else { // 此时root就是要删除的节点
		if root.lchild != nil && root.rchild != nil {
			if avlt.getHeight(root.lchild) > avlt.getHeight(root.rchild) {
				lmax, _ := avlt.getMaxNode(root.lchild)
				root.val = lmax.val
				root.lchild = avlt.remove(root.lchild, lmax)
			} else {
				rmin, _ := avlt.getMinNode(root.rchild)
				root.val = rmin.val
				root.rchild = avlt.remove(root.rchild, rmin)
			}
		} else {
			if root.lchild != nil {
				root = root.lchild
			} else {
				root = root.rchild
			}
		}
	}
	return root
}

// GetMin 获取avl树中的最小的节点值
func (avlt *AVLTree) GetMin() (int, error) {
	if avlt.root == nil {
		return -1, fmt.Errorf("树为空")
	}
	p := avlt.root
	for p.lchild != nil {
		p = p.lchild
	}
	return p.val, nil
}

func (avlt *AVLTree) getMinNode(root *avlnode) (*avlnode, error) {
	if root == nil {
		return nil, fmt.Errorf("树为空")
	}
	for root.lchild != nil {
		root = root.lchild
	}
	return root, nil
}

// GetMax 获取avl树中的最大的节点值
func (avlt *AVLTree) GetMax() (int, error) {
	if avlt.root == nil {
		return -1, fmt.Errorf("树为空")
	}
	p := avlt.root
	for p.rchild != nil {
		p = p.rchild
	}
	return p.val, nil
}

func (avlt *AVLTree) getMaxNode(root *avlnode) (*avlnode, error) {
	if root == nil {
		return nil, fmt.Errorf("树为空")
	}
	for root.rchild != nil {
		root = root.rchild
	}
	return root, nil
}

func (avlt *AVLTree) search(root *avlnode, data int) *avlnode {
	if root == nil || root.val == data {
		return root
	}
	if data < root.val {
		return avlt.search(root.lchild, data)
	}
	return avlt.search(root.rchild, data)
}

// MidVisit 中序遍历AVL树
func (avlt *AVLTree) MidVisit() {
	if avlt.root == nil {
		fmt.Println("AVL树为空")
		return
	}
	avlt.midVisit(avlt.root)
}

func (avlt *AVLTree) midVisit(root *avlnode) {
	if root == nil {
		return
	}
	avlt.midVisit(root.lchild)
	fmt.Printf("%d  ", root.val)
	avlt.midVisit(root.rchild)
}

// PreVisit 先序遍历AVL树
func (avlt *AVLTree) PreVisit() {
	if avlt.root == nil {
		fmt.Println("AVL树为空")
		return
	}
	avlt.preVisit(avlt.root)
}

func (avlt *AVLTree) preVisit(root *avlnode) {
	if root == nil {
		return
	}
	fmt.Printf("%d  ", root.val)
	avlt.preVisit(root.lchild)
	avlt.preVisit(root.rchild)
}

// PrintTree 打印AVL树
func (avlt *AVLTree) PrintTree() {
	avlt.printTree(avlt.root, avlt.root.val, 0)
}

func (avlt *AVLTree) printTree(root *avlnode, data int, direction int) {
	if root != nil {
		if direction == 0 {
			fmt.Printf("%2d 是 root \n", root.val)

		} else {
			kid := ""
			if direction == 1 {
				kid = "右"
			} else {
				kid = "左"
			}
			fmt.Printf("%2d 是 %2d 的 %s孩子  \n", root.val, data, kid)
		}
		avlt.printTree(root.lchild, root.val, -1)
		avlt.printTree(root.rchild, root.val, 1)
	}

}
