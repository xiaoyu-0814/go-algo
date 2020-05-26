package leetcode

//287.寻找重复数
/*
	给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

示例 1:

输入: [1,3,4,2,2]
输出: 2
示例 2:

输入: [3,1,3,4,2]
输出: 3
说明：

不能更改原数组（假设数组是只读的）。
只能使用额外的 O(1) 的空间。
时间复杂度小于 O(n2) 。
数组中只有一个重复的数字，但它可能不止重复出现一次。
*/

//1.利用hash解法
func findDuplicate_01(nums []int) int {
	hash := map[int]struct{}{}
	for _,v := range nums{
		if _,ok := hash[v];ok {
			return v
		}else {
			hash[v]= struct{}{}
		}
	}
	return -1
}

//2.暴力破解法
func findDuplicate_02(nums []int) int {
	for k,v := range nums{
		for j :=0;j< len(nums)&&j != k;j++{
			if v == nums[j] {
				return v
			}
		}
	}
	return -1
}