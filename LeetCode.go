package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(l *ListNode) {
	for l != nil {
		if l.Next == nil {
			fmt.Printf("%d", l.Val)
		} else {
			fmt.Printf("%d->", l.Val)
		}
		l = l.Next
	}
	fmt.Println("")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret, node *ListNode
	var last int

	for l1 != nil || l2 != nil {
		var val1, val2, val int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		a := val1 + val2 + last
		if a >= 10 {
			last = 1
			val = a - 10
		} else {
			last = 0
			val = a
		}
		n := &ListNode{Val: val, Next: nil}
		if node == nil {
			ret = n
			node = n
		} else {
			node.Next = n
			node = n
		}
	}
	if last != 0 {
		node.Next = &ListNode{last, nil}
	}
	return ret
}

func testAddTwoNum() {
	l1 := &ListNode{2, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{3, nil}

	l2 := &ListNode{5, nil}
	l2.Next = &ListNode{6, nil}
	l2.Next.Next = &ListNode{4, nil}
	printList(addTwoNumbers(l1, l2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0.0
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func hammingDistance(x int, y int) int {
	f := x ^ y
	ret := 0
	for f > 0 {
		f &= f - 1
		ret++
	}
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	return &TreeNode{t1.Val + t2.Val, mergeTrees(t1.Left, t2.Left), mergeTrees(t1.Right, t2.Right)}
}

func arrayPairSum(nums []int) int {
	m := make([]int, 20001)
	for _, val := range nums {
		m[val+10000]++
	}
	a := 0
	ret := 0
	for k, val := range m {
		for i := 0; i < val; i++ {
			//fmt.Printf("val=%d, k=%d\n", val, k)
			if a%2 == 0 {
				ret += k - 10000
			}
			a++
		}
	}
	return ret
}

func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a ^= v
	}
	return a
}

func findMaxConsecutiveOnes(nums []int) int {
	ret := 0
	l := 0
	for _, v := range nums {
		if v == 0 {
			l = 0
		} else {
			l++
			if l > ret {
				ret = l
			}
		}
	}
	return ret
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	if r > l {
		l = r
	}
	return 1 + l
}

func detectCapitalUse(word string) bool {
	c := word[0]
	up, lower := false, false
	for i := 1; i < len(word); i++ {
		if word[i] > 'Z' {
			lower = true
		}
		if word[i] < 'a' {
			up = true
		}
	}
	return (c > 'Z' && !up) || (c < 'a' && !lower) || (c < 'a' && !up)
}

func findDisappearedNumbers(nums []int) []int {
	for k, v := range nums {
		tmp := nums[v-1]
		nums[v-1] = v
		nums[k] = tmp
	}
	for k, v := range nums {
		tmp := nums[v-1]
		nums[v-1] = v
		nums[k] = tmp
	}
	ret := make([]int, 0)
	fmt.Println(nums)
	for k, v := range nums {
		if v != k+1 {
			ret = append(ret, k+1)
		}
	}
	return ret
}

func sumOfLeftLeaves2(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	ret := 0
	if root.Left == nil && root.Right == nil && isLeft {
		ret += root.Val
	}
	l := sumOfLeftLeaves2(root.Left, true)
	r := sumOfLeftLeaves2(root.Right, false)
	return ret + l + r
}

func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeaves2(root, false)
}

func maximumProduct(nums []int) int {
	a, b, c, x, y := math.MinInt32, math.MinInt32, math.MinInt32, math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if v > a {
			a, b, c = v, a, b
		} else if v <= a && v > b {
			b, c = v, b
		} else if v <= b && v > c {
			c = v
		}
		if v < x {
			x, y = v, x
		} else if v >= x && v < y {
			y = v
		}
	}
	fmt.Println(a, b, c, x, y)
	if a*b*c > a*x*y {
		return a * b * c
	} else {
		return a * x * y
	}
}

func numberOfBoomerangs(points [][]int) int {
	ret := 0
	for _, p1 := range points {
		m := make(map[int]int, len(points))
		for _, p2 := range points {
			dis := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
			if count, ok := m[dis]; ok {
				ret += 2 * count
				m[dis]++
			} else {
				m[dis] = 1
			}
		}
		fmt.Println(m)
	}
	return ret
}

func testNumbersof() {
	points := [][]int{[]int{0, 0}, []int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
	fmt.Println(numberOfBoomerangs(points) == 20)
}

func countBinaryOne(num int) int {
	ret := 0
	for num > 0 {
		ret += num & 1
		num = num >> 1
	}
	return ret
}

func readBinaryWatch(num int) []string {
	if num == 0 {
		return []string{"0:00"}
	}
	m := make(map[int][]int)

	for i := 0; i < 60; i++ {
		bit := countBinaryOne(i)

		if arr, ok := m[bit]; ok {
			m[bit] = append(arr, i)
		} else {
			m[bit] = []int{i}
		}
	}
	ret := make([]string, 0, 12*60)

	for i := 0; i <= num && i < 5; i++ {
		hours := m[i]
		for _, h := range hours {
			if h > 11 {
				break
			}
			mins := m[num-i]
			for _, m := range mins {
				var mdis string
				if m < 10 {
					mdis += "0"
				}
				ret = append(ret, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ret
}

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	ret := make([]int, 0, len(m))
	for _, v := range nums2 {
		if freq, ok := m[v]; ok {
			if freq > 0 {
				ret = append(ret, v)
				m[v]--
			}
		}
	}
	return ret
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	base := 7
	flag := false
	if num < 0 {
		num = -num
		flag = true
	}
	length := int(math.Log10(float64(num))/math.Log10(float64(base))) + 1
	if flag {
		length++
	}
	ret := make([]byte, length)
	if flag {
		ret[0] = '-'
	}
	i := 1
	for num > 0 {
		ret[length-i] = strconv.Itoa(num % base)[0]
		num = num / base
		i++
	}

	return string(ret)
}

func missingNumber(nums []int) int {
	sum1 := 0
	sum2 := 0
	for k, v := range nums {
		sum1 += k
		sum2 += v
	}
	sum1 += len(nums)
	return sum1 - sum2
}

func reverseStr(s string, k int) string {
	ret := []byte(s)
	i := 0
	for i < len(s) {
		end := i + k
		if end > len(s) {
			end = len(s)
		}
		for j := i; j < end; j++ {
			ret[j] = s[end-j+i-1]
		}

		i += 2 * k
	}
	return string(ret)
}

func checkRecord(s string) bool {
	a, l := 0, 0
	for _, ch := range s {
		if ch == 'A' {
			a++
			if a > 1 {
				return false
			}
		}
		if ch == 'L' {
			l++
			if l > 2 {
				return false
			}
		} else {
			l = 0
		}
	}
	return true
}

func diameterOfBinaryTree2(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	ld, ls := diameterOfBinaryTree2(root.Left)
	rd, rs := diameterOfBinaryTree2(root.Right)

	s := 0
	if ls > rs {
		s = ls + 1
	} else {
		s = rs + 1
	}

	d := 0
	if ld > d {
		d = ld
	}
	if rd > d {
		d = rd
	}
	dlr := 0
	if root.Left != nil {
		dlr += ls
	}
	if root.Right != nil {
		dlr += rs
	}
	if dlr > d {
		d = dlr
	}
	return d, s
}

func diameterOfBinaryTree(root *TreeNode) int {
	d, _ := diameterOfBinaryTree2(root)
	return d
}

func testDiameterOfBTree() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 7}
	//root.Left.Left.Left = &TreeNode{Val: 9}
	//root.Left.Left.Left.Left = &TreeNode{Val: 9}
	//9root.Left.Left.Left.Left.Left = &TreeNode{Val: 9}
	fmt.Println(diameterOfBinaryTree2(root))
}

func maxProfit(prices []int) int {
	min := math.MaxInt32
	max := 0
	for _, v := range prices {
		if v-min > max {
			max = v - min
		}
		if v < min {
			min = v
		}
	}
	return max
}

func findErrorNums(nums []int) []int {
	s := 0
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		s += v
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}
	s2 := (len(nums) + 1) * len(nums) / 2
	fmt.Println(s, s2)
	for k, v := range nums {
		if v > 0 {
			return []int{s + k + 1 - s2, k + 1}
		}
	}
	return []int{0, 0}
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil && t != nil {
		return false
	}
	if s != nil && t == nil {
		return true
	}

	if s.Val == t.Val {
		return (isSubtree(s.Left, t.Left) && isSubtree(s.Right, t.Right)) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
	} else {
		return isSubtree(s.Left, t) || isSubtree(s.Right, t)
	}
}

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	n1, n2 := 2, 1
	for j := 3; j <= n; j++ {
		n1, n2 = n1+n2, n1
	}
	return n1
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	return pathSumWithRoot(root, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func pathSumWithRoot(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	ret := 0
	if root.Val == sum {
		ret++
	}
	return ret + pathSumWithRoot(root.Left, sum-root.Val) + pathSumWithRoot(root.Right, sum-root.Val)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var last *ListNode
	p := head
	for p != nil {
		fmt.Println("last:", last, "p:", p)
		if last != nil && last.Val == p.Val {
			last.Next = p.Next
			p = p.Next
		} else {
			last = p
			p = p.Next
		}
	}
	return head
}

func testDeleteDuplicates() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	printList(head)
	head = deleteDuplicates(head)
	printList(head)
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret, sum := nums[0], 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > ret {
			ret = sum
		}
		if sum < 0 {
			sum = 0
		}
		fmt.Printf("sum:%d, ret=%d, num[i]=%d\n", sum, ret, nums[i])

	}
	return ret
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	if buckets == 1 {
		return 0
	}
	n := minutesToTest/minutesToDie + 1
	tt := 1
	ret := 0
	for tt < buckets {
		tt *= n
		ret++
	}
	return ret
}

func testPathSum() {
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: -3}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Left = &TreeNode{Val: -2}
	root.Left.Right.Right = &TreeNode{Val: 1}
	fmt.Println(pathSum(root, 8))

}

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ret := float64(sum) / float64(k)
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		avg := float64(sum) / float64(k)
		if avg > ret {
			ret = avg
		}
	}
	return ret
}

func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	if num == 1 {
		return true
	}
	if num%2 == 0 {
		return isUgly(num / 2)
	}
	if num%3 == 0 {
		return isUgly(num / 3)
	}
	if num%5 == 0 {
		return isUgly(num / 5)
	}
	return false
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := &ListNode{Val: 0}
	last := prev
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			last.Next = l1
			last = l1
			l1 = l1.Next
		} else {
			last.Next = l2
			last = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		last.Next = l1
	} else {
		last.Next = l2
	}
	return prev.Next
}

func arrayToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	root := &ListNode{Val: nums[0]}
	last := root
	for i := 1; i < len(nums); i++ {
		node := &ListNode{Val: nums[i]}
		last.Next = node
		last = node
	}
	return root
}

func testMergeList() {
	l1 := arrayToList([]int{1, 3, 5})
	printList(l1)
	l2 := arrayToList([]int{3, 4, 6})
	printList(l2)
	printList(mergeTwoLists(l1, l2))
}

func removeElement(nums []int, val int) int {
	ret := 0
	for _, v := range nums {
		if v != val {
			nums[ret] = v
			ret++
		}
	}
	nums = nums[:ret]
	fmt.Println(nums)
	return ret
}

func generate(numRows int) [][]int {

	ret := make([][]int, numRows)
	if numRows == 0 {
		return nil
	}
	ret[0] = []int{1}
	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j-1 >= 0 {
				row[j] += ret[i-1][j-1]
			}
			if j < i {
				row[j] += ret[i-1][j]
			}
		}
		ret[i] = row
	}
	return ret
}

func reverseVowels(s string) string {
	m := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	ret := []byte(s)
	begin, end := 0, len(s)-1
	for begin < end {
		bOK, eOK := false, false
		if _, ok := m[s[begin]]; ok {
			bOK = true
		} else {
			bOK = false
			begin++
		}
		if _, ok := m[s[end]]; ok {
			eOK = true
		} else {
			eOK = false
			end--
		}
		if bOK && eOK {
			ret[begin], ret[end] = s[end], s[begin]
			begin++
			end--
		}
	}
	return string(ret)
}

func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}
	if num < 2 {
		return true
	}
	begin, end := 0, num/2
	for begin <= end {
		mid := begin + (end-begin)/2
		fmt.Printf("begin=%d, end=%d, mid=%d\n", begin, end, mid)
		s := mid * mid
		if s == num {
			return true
		}
		if s < num {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

func isPerfectSquare2(num int) bool {
	s := num
	for s*s > num {
		s = (s + num/s) / 2
	}
	fmt.Println(s)
	return s*s == num
}

func findMode(root *TreeNode) []int {
	max := 0
	var findModeHelper func(root *TreeNode, ret *[]int, last **TreeNode, lastCount *int)
	findModeHelper = func(root *TreeNode, ret *[]int, last **TreeNode, lastCount *int) {
		if root == nil {
			return
		}
		findModeHelper(root.Left, ret, last, lastCount)
		if (*last) != nil && root.Val == (*last).Val {
			(*lastCount)++
		} else {
			*lastCount = 1
		}
		if *lastCount > max {
			max = *lastCount
			*ret = []int{root.Val}
		} else if *lastCount == max {
			*ret = append(*ret, root.Val)
		}
		*last = root
		findModeHelper(root.Right, ret, last, lastCount)
	}
	ret := []int{}
	var node *TreeNode
	lastCount := 0
	findModeHelper(root, &ret, &node, &lastCount)
	return ret
}

func testFindMode() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 1}
	//root.Right.Left = &TreeNode{Val: 2}
	fmt.Println(findMode(root))

}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var treeDepth func(root *TreeNode) int

	treeDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l, r := treeDepth(root.Left), treeDepth(root.Right)
		if l == -1 || r == -1 {
			return -1
		}
		if l > r+1 || r > l+1 {
			return -1
		}
		if l > r {
			return l + 1
		} else {
			return r + 1
		}
	}
	return treeDepth(root) != -1
}

func testIsBalance() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 1}
	//root.Right.Left = &TreeNode{Val: 2}
	fmt.Println(isBalanced(root))

}

func countSegments(s string) int {
	ret := 0
	last := '\r'
	for _, v := range s {
		if v == ' ' && last != ' ' && last != '\r' {
			ret++
		}
		last = v
	}
	if last != ' ' && last != '\r' {
		ret++
	}
	return ret
}

func testCountSegments() {
	fmt.Println(countSegments(" Hello, my name is John") == 5)
	fmt.Println(countSegments("") == 0)

}

func getRow(rowIndex int) []int {
	ret := make([]int, rowIndex+1)
	ret[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j >= 1; j-- {
			ret[j] += ret[j-1]
		}
	}
	return ret
}

func arrangeCoins(n int) int {

	low, high := 0, n

	for low <= high {
		mid := low + (high-low)/2
		sum := (1 + mid) * mid / 2
		fmt.Printf("low:%d, high=%d, mid=%d, sum=%d\n", low, high, mid, sum)
		if sum == n {
			return mid
		}
		if sum > n {
			high = mid - 1
		} else {
			if n-sum < mid+1 {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return 0
}

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}
	return trailingZeroes(n/5) + n/5
}

func nxx(n int) int {
	if n == 0 {
		return 1
	}
	ret := 1
	for n > 1 {
		ret *= n
		n--
	}
	return ret
}

func removeDuplicates(nums []int) int {
	dup := 0
	for k, v := range nums {
		if k != 0 && v == nums[k-1] {
			dup++
		} else {
			nums[k-dup] = v
		}
	}
	return len(nums) - dup
}

func isPalindrome(x int) bool {
	if 0 <= x && x < 10 {
		return true
	}
	r, s := 0, x
	x = x / 10
	b := 10
	for x >= 10 {
		r = 10*r + x%10
		x = x / 10
		b *= 10
	}
	fmt.Println(b, r, x, s%10, (s-b*x)/10)
	return r == (s-b*x)/10 && x == s%10
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	var say func(last []byte) []byte

	say = func(last []byte) []byte {
		ret := []byte{}
		count, prev := 1, last[0]
		for k := 1; k < len(last); k++ {
			if prev == last[k] {
				count++
			} else {
				ret = append(ret, []byte(strconv.Itoa(count))...)
				ret = append(ret, prev)
				count = 1
			}
			prev = last[k]
		}
		ret = append(ret, []byte(strconv.Itoa(count))...)
		ret = append(ret, prev)
		return ret
	}

	ret := []byte{'1'}
	for i := 1; i < n; i++ {
		//fmt.Println(string(ret))
		ret = say(ret)
	}
	return string(ret)
}

func hasPathSum(root *TreeNode, sum int) bool {
	var findPath func(root *TreeNode, prev, sum int) bool
	findPath = func(root *TreeNode, prev, sum int) bool {
		if root == nil {
			return sum == 0
		}
		prev += root.Val
		if root.Left == nil && root.Right == nil {
			return prev == sum
		}
		return findPath(root.Left, prev, sum) || findPath(root.Right, prev, sum)
	}
	return findPath(root, 0, sum)
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1, m2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		fmt.Println(m1, m2)
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}
	return true
}

func checkPerfectNumber(num int) bool {
	sum := 0
	for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			d := num / i
			sum += i
			if d != i {
				sum += d
			}
		}
	}
	sum++
	fmt.Println(sum)
	return sum == num
}

func findAnagrams(s string, p string) []int {
	m := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		m[p[i]]++
	}
	ret := []int{}
	//abcdea
	left, right, count := 0, 0, len(p)
	for right < len(s) {
		fmt.Println(m)
		fmt.Printf("left:%d, right:%d, count:%d\n", left, right, count)
		if m[s[right]] > 0 {
			count--
		}
		m[s[right]]--
		right++
		if count == 0 {
			ret = append(ret, left)
		}
		if right-left == len(p) {
			if m[s[left]] >= 0 {
				count++
			}
			m[s[left]]++
			left++
		}
	}
	return ret
}

func isValid(s string) bool {
	st := list.New()
	m := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := 0; i < len(s); i++ {
		v := s[i]
		switch v {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			st.PushBack(v)
		case ')':
			fallthrough
		case ']':
			fallthrough
		case '}':
			if st.Len() == 0 {
				return false
			}
			node := st.Back()
			fmt.Println(node.Value)
			if node.Value.(byte) != m[v] {
				return false
			}
			st.Remove(node)
		}
	}
	return st.Len() == 0
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var helper func(root *TreeNode, depth int, min *int)
	helper = func(root *TreeNode, depth int, min *int) {
		if root == nil {
			return
		}
		depth++
		if root.Left == nil && root.Right == nil {
			if depth < *min {
				*min = depth
			}
		}
		helper(root.Left, depth, min)
		helper(root.Right, depth, min)
	}
	min := math.MaxInt32
	helper(root, 0, &min)
	return min
}

func testMinDepth() {
	root := &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 20}
	root.Left = &TreeNode{Val: 30}
	root.Left.Left = &TreeNode{Val: 33}
	fmt.Println(minDepth(root))
}

func wordPattern(pattern string, str string) bool {
	i, j, last := 0, 0, 0
	m := make(map[byte]string)
	m2 := make(map[string]byte)
	for i < len(pattern) && j < len(str) {
		fmt.Println(m)
		if str[j] == ' ' || j == len(str)-1 {
			ss := strings.TrimSpace(str[last : j+1])
			if s, ok := m[pattern[i]]; ok {
				if s != ss {
					fmt.Println(s, str[last:j])
					return false
				}
			} else {
				m[pattern[i]] = str[last:j]
			}
			if ch, ok := m2[ss]; ok {
				if ch != pattern[i] {
					return false
				}
			} else {
				m2[ss] = pattern[i]
			}
			i++
			last = j + 1
		}
		j++
	}
	fmt.Println(j, i)
	return i == len(pattern) && j == len(str)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if k <= 0 {
		return false
	}
	m := make(map[int]int)
	if k >= len(nums) {
		k = len(nums) - 1
	}
	for i, n := range nums {
		//if i > k {
		//	delete(m, nums[i-k-1])
		//}
		if j, ok := m[n]; ok && i-j <= k {
			return true
		}
		m[n] = i
	}
	return false
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	for head != nil && head.Val == val {
		head = head.Next
	}
	prev, p := head, head
	for p != nil {
		if p.Val == val {
			prev.Next = p.Next
		} else {
			prev = p
		}
		p = p.Next
	}
	return head
}

func testDeleteList() {
	head := arrayToList([]int{1, 2, 2, 1, 1})
	printList(removeElements(head, 1))
}

func addBinary(a string, b string) string {
	alen, blen, l := len(a), len(b), 0
	if alen > blen {
		l = alen
	} else {
		l = blen
	}
	ret := make([]byte, l+1)
	var carry byte
	for i, j, k := alen-1, blen-1, l; i >= 0 || j >= 0; {
		sum := carry
		if i >= 0 {
			sum += a[i] - '0'
			i--
		}
		if j >= 0 {
			sum += b[j] - '0'
			j--
		}
		carry = sum / 2
		ret[k] = sum%2 + '0'
		k--
	}
	if carry == 1 {
		ret[0] = '1'
		return string(ret)
	} else {
		return string(ret[1:])
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	i := -1
	for j := 0; j < len(s); j++ {
		if s[j] == ' ' {
			i = j
		}
	}
	if i == -1 {
		return len(s)
	}
	return len(s) - i - 1
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ret := []byte{}
	for i := 0; i < len(strs[0]); i++ {
		ch := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return string(ret)
			}
			if strs[j][i] != ch {
				return string(ret)
			}
		}
		ret = append(ret, ch)
	}
	return string(ret)
}

func judgeSquareSum(c int) bool {
	start, end := 0, int(math.Sqrt(float64(c)))
	for start <= end {
		sum := start*start + end*end
		if sum == c {
			return true
		}
		if sum > c {
			end--
		} else {
			start--
		}
	}
	return false
}

func findNthDigit(n int) int {
	count, bit := 1, 1
	for n > bit*count*9 {
		fmt.Println(n)
		n -= bit * count * 9
		count *= 10
		bit++
	}
	fmt.Printf("n=%d, count=%d, bit=%d, div=%d, mod=%d\n", n, count, bit, n/bit, n%bit)
	div, mod := (n-1)/bit, (n-1)%bit
	return (int)(strconv.Itoa(count + div)[mod]) - '0'
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n > len(flowerbed) {
		return false
	}
	if n == 0 {
		return true
	}
	for i := 0; i < len(flowerbed); i++ {
		if (i == 0 || flowerbed[i-1] == 0) && flowerbed[i] == 0 && ((i == len(flowerbed)-1) || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
			if n == 0 {
				return true
			}
		}

	}
	return false
}

func findUnsortedSubarray(nums []int) int {
	begin, end := 0, len(nums)-1
	for begin <= end {

	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findRadius(houses []int, heaters []int) int {
	//sort.Ints(houses)
	sort.Ints(heaters)
	fmt.Println(houses)
	fmt.Println(heaters)
	ret := 0
	for _, house := range houses {
		low, high := 0, len(heaters)-1
		fmt.Println("-----------")
		for low < high-1 {
			mid := low + (high-low)/2
			fmt.Printf("house:%d, low:%d, high:%d, mid=%d, heater[mid]:%d\n", house, low, high, mid, heaters[mid])
			if heaters[mid] == house {
				low, high = mid, mid
			}
			if heaters[mid] < house {
				low = mid
			} else {
				high = mid
			}
		}
		fmt.Printf("house:%d, low:%d, high:%d\n", house, low, high)
		c := min(abs(heaters[low]-house), abs(heaters[high]-house))
		if c > ret {
			ret = c
		}
	}
	return ret
}

type NumArray struct {
	Data []int
}

func Constructor1(nums []int) NumArray {
	data := make([]int, len(nums))
	copy(data, nums)
	for i := 1; i < len(nums); i++ {
		data[i] += data[i-1]
	}
	return NumArray{Data: data}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.Data[j]
	}
	return this.Data[j] - this.Data[i-1]
}

func thirdMax(nums []int) int {
	max3, max2, max1 := math.MinInt32-1, math.MinInt32-1, math.MinInt32-1
	for _, n := range nums {
		//fmt.Println(max3, max2, max1)
		if n > max1 {
			max3, max2, max1 = max2, max1, n
			continue
		}
		if n == max1 {
			continue
		}
		if n > max2 {
			max3, max2 = max2, n
			continue
		}
		if n == max2 {
			continue
		}
		if n > max3 {
			max3 = n
		}
	}
	if max3 != math.MinInt32-1 {
		return max3
	}
	return max1
}

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	nums := make([]int, 0, n)
	nums = append(nums, 2)
	for i := 3; i < n; i++ {
		flag := true
		for _, n := range nums {
			if n*n > i {
				break
			}
			if i%n == 0 {
				flag = false
				break
			}
		}
		if flag {
			nums = append(nums, i)
		}
	}
	//fmt.Println(nums)
	return len(nums)
}

func countPrimes2(n int) int {
	if n <= 2 {
		return 0
	}
	count := 0
	dic := make([]bool, n)
	for i := 2; i < n; i++ {
		if dic[i] == false {
			count++
			for j := 2; i*j < n; j++ {
				dic[i*j] = true
			}
		}
	}
	return count
}

func isPalindromeStr(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if (s[i] > 'Z' && s[i] < 'a') || s[i] > 'z' || (s[i] < 'A' && s[i] > '9') || s[i] < '0' {
			i++
			continue
		}
		if (s[j] > 'Z' && s[j] < 'a') || s[j] > 'z' || (s[j] < 'A' && s[j] > '9') || s[j] < '0' {
			j--
			continue
		}
		si, sj := s[i], s[j]
		if si > 'Z' {
			si -= 32
		}
		if sj > 'Z' {
			sj -= 32
		}
		if si != sj {
			return false
		}
		i++
		j--
	}
	return true
}

func convertToTitle(n int) string {
	if n < 1 {
		return ""
	}
	dic := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret := make([]byte, 10)
	count := 0
	for n > 0 {
		n--
		ret[9-count] = dic[(n % 26)]
		n = n / 26
		count++
	}
	return string(ret[10-count:])
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse := func(arr []int) {
		low, high := 0, len(arr)-1
		for low < high {
			arr[low], arr[high] = arr[high], arr[low]
			low++
			high--
		}
	}
	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:])
}

func reverse(x int) int {
	nagative := false
	if x < 0 {
		x = -x
		nagative = true
	}
	ret := 0
	for x > 0 {
		ret *= 10
		ret += x % 10
		x = x / 10
	}
	if ret > (math.MaxInt32) {
		return 0
	}
	if nagative {
		ret = -ret
		if ret < math.MinInt32 {
			return 0
		}
		return ret
	}
	return ret
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	findMaxIndex := func(arr []int) (int, int) {
		idx, max := -1, math.MinInt32
		for i, n := range arr {
			if n > max {
				idx, max = i, n
			}
		}
		return idx, max
	}
	fmt.Println(findMaxIndex([]int{1, 7, 0, 5, 4}))
	return nil
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(
		people,
		func(i, j int) bool {
			return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
		},
	)
	insertSlice := func(list [][]int, i int, ele []int) {
		for j := len(list) - 2; j >= i; j-- {
			list[j+1] = list[j]
		}
		list[i] = ele
	}

	ret := make([][]int, len(people))
	for i := 0; i < len(people); i++ {
		pair := people[i]
		insertSlice(ret, pair[1], pair)
		//fmt.Println(ret)
	}
	return ret
}

func findDuplicates(nums []int) []int {
	ret := make([]int, len(nums))
	count := 0
	for _, i := range nums {
		if i < 0 {
			i = -i
		}
		if nums[i-1] < 0 {
			ret[count] = i
			count++
		}
		nums[i-1] = -nums[i-1]
	}
	return ret[:count]
}

func optimalDivision(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	if len(nums) == 1 {
		return fmt.Sprintf("%d", nums[0])
	}
	if len(nums) == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	str := fmt.Sprintf("%d/(", nums[0])

	for i := 1; i < len(nums); i++ {
		str = fmt.Sprintf("%s%d/", str, nums[i])
	}
	ret := []byte(str)
	ret[len(ret)-1] = ')'
	return string(ret)
}

func singleNonDuplicate(nums []int) int {
	ret := 0
	for _, n := range nums {
		ret ^= n
	}
	return ret
}

func singleNonDuplicate1(nums []int) int {
	count, last := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		fmt.Println("nums[i]:", nums[i], "last:", last)
		if nums[i] == last {
			count++
		} else {
			if count == 1 {
				return last
			} else {
				count = 1
			}
		}
		last = nums[i]
	}
	if count == 1 {
		return nums[len(nums)-1]
	}
	return -1
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := list.New()
	q.PushBack(root)
	ret := make([]int, 0, 100)
	for q.Len() > 0 {
		size := q.Len()
		max := math.MinInt32
		for i := 0; i < size; i++ {
			ele := q.Front()
			node := ele.Value.(*TreeNode)
			q.Remove(ele)
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				q.PushBack(node.Left)
			}
			if node.Right != nil {
				q.PushBack(node.Right)
			}

		}
		ret = append(ret, max)
	}
	return ret
}

func testLargestValues() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 9}
	fmt.Println(largestValues(root))

	//[1,3,2,5,3,null,9]
}

func listToTree(s string) *TreeNode {
	list := strings.Split(s, ",")
	if len(list) == 0 {
		return nil
	}
	val, _ := strconv.Atoi(list[0])
	root := &TreeNode{Val: val}
	i := 1
	q := []*TreeNode{root}
	fmt.Println(list)
	for i < len(list) {
		size := len(q)
		for j := 0; j < size; j++ {
			if q[j] != nil {
				if list[i] != "null" {
					val, _ := strconv.Atoi(list[i])
					q[j].Left = &TreeNode{Val: val}
				}
				q = append(q, q[j].Left)
				if list[i+1] != "null" {
					val, _ := strconv.Atoi(list[i+1])
					q[j].Right = &TreeNode{Val: val}
				}
				q = append(q, q[j].Right)
			} else {
				q = append(q, nil, nil)
			}
			i += 2
		}
		q = q[size:]
	}
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("<nil> Tree")
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			if q[i] == nil {
				fmt.Printf(" ")
			} else {
				fmt.Printf("%d ", q[i].Val)
				q = append(q, q[i].Left)
				q = append(q, q[i].Right)
			}
		}
		fmt.Printf("\n")
		q = q[size:]
	}
}

func printTreePretty(root *TreeNode, indent int) {
	if root == nil {
		return
	}

	printTreePretty(root.Right, indent+4)
	if indent > 0 {
		format := fmt.Sprintf("%%%ds", indent)
		fmt.Printf(format, " ")
	}
	fmt.Printf("%d\n", root.Val)
	printTreePretty(root.Left, indent+4)
}

func findFrequentTreeSum(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	dic := make(map[int]int)

	var post func(root *TreeNode) int
	post = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		s := post(root.Left) + post(root.Right) + root.Val
		dic[s]++
		return s
	}

	post(root)
	max := math.MinInt32
	ret := []int{}
	for sum, count := range dic {
		if count > max {
			ret = []int{sum}
			max = count
		} else if count == max {
			ret = append(ret, sum)
		}
	}
	return ret
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return root
	}
	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}
	type Node struct {
		ele *TreeNode
		dep int
	}
	stack := []*Node{&Node{root, 1}}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.ele == nil {
			continue
		}
		if node.dep == d-1 {
			l, r := node.ele.Left, node.ele.Right
			node.ele.Left, node.ele.Right = &TreeNode{Val: v}, &TreeNode{Val: v}
			node.ele.Left.Left, node.ele.Right.Right = l, r
		} else {
			stack = append(stack, &Node{node.ele.Left, node.dep + 1})
			stack = append(stack, &Node{node.ele.Right, node.dep + 1})
		}
	}
	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	return nil
}

func main() {
	root := listToTree("4,2,6,3,1,5,null")
	printTreePretty(root, 0)
	root = addOneRow(root, 1, 0)
	fmt.Println("---------------------------------")
	printTreePretty(root, 0)
}
