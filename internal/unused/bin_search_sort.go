package unused

import (
	"strconv"
	"strings"
)

func init() {

}

// BinSearch 在排序数组中执行二分查找操作。
// 参数:
// nums - 一个已排序的整数切片，用于执行查找操作。
// target - 需要查找的目标整数值。
// 返回值:
// 如果找到目标值，则返回其在数组中的索引；如果未找到，则返回 -1。
func BinSearch(nums []int, target int) int {
	// 定义查找范围的左右边界
	left, right := 0, len(nums)-1

	// 当左边界不超过右边界时，继续查找
	for left <= right {
		// 计算中间位置
		mid := (right + left) / 2

		if nums[mid] == target {
			// 判断中间值是否等于目标值
			return mid
		} else if nums[mid] < target {
			// 如果中间值小于目标值，调整左边界以缩小查找范围
			left = mid + 1
		} else {
			// 如果中间值大于目标值，调整右边界以缩小查找范围
			right = mid - 1
		}
	}

	// 如果未找到目标值，返回 -1
	return -1
}

// FindPeakElement 寻找峰值元素的索引
// nums 是一个整数数组，其中峰值元素被定义为大于其邻居的元素。
// 该函数返回任一峰值元素的索引。
// 参数:
//
//	nums: 一个整数数组，不为空，且至少包含一个元素。
//
// 返回值:
//
//	int: 峰值元素的索引。
func FindPeakElement(nums []int) int {
	// 初始化左右指针
	left, right := 0, len(nums)-1

	// 使用二分查找法寻找峰值
	for left < right {
		// 计算中间索引
		mid := (left + right) / 2

		// 如果中间元素小于其右侧元素，说明峰值在右侧
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			// 否则，峰值在左侧或就是中间元素
			right = mid
		}
	}

	// 当左右指针相遇时，即找到峰值，返回左指针即可
	return left
}

const mod = 1000000007

// InversePairs 计算数组中逆序对的数量。
// 逆序对定义为数组中的一对数值 (i, j) ，满足 i < j 且 nums[i] > nums[j]。
// 该函数使用归并排序的思想来计算逆序对，通过分治法将数组不断分割成子数组，直到每个子数组为单个元素。
// 然后两两合并子数组，在合并的过程中计算逆序对的数量。
// 参数:
//
//	nums - 待计算逆序对的整数数组。
//
// 返回值:
//
//	逆序对的数量。
func InversePairs(nums []int) int {
	// 调用 mergeSort 函数，传入初始的左边界 0 和右边界 len(nums)-1。
	// mergeSort 函数将执行实际的归并排序和逆序对计数。
	return mergeSort(nums, 0, len(nums)-1)
}

// mergeSort 对数组 nums 的子数组进行归并排序，并返回排序过程中的逆序对数量。
// 该函数使用分治算法，将数组分成两半分别排序，然后合并。
// 参数:
//
//	nums: 待排序的数组
//	start: 子数组的起始索引
//	end: 子数组的结束索引
//
// 返回值:
//
//	返回排序过程中的逆序对数量。
func mergeSort(nums []int, start, end int) int {
	// 当 start >= end 时，说明数组只有一个元素或者为空，不需要排序，返回 0。
	if start >= end {
		return 0
	}
	// 计算数组的中间索引 mid。
	mid := start + (end-start)/2
	// 递归处理左半部分
	leftCount := mergeSort(nums, start, mid)
	// 递归处理右半部分
	rightCount := mergeSort(nums, mid+1, end)
	// 合并并统计跨越左右两部分的逆序对
	crossCount := merge(nums, start, mid, end)
	// 返回总的逆序对数量，并对结果取模以防止溢出。
	return (leftCount + rightCount + crossCount) % mod
}

// merge 函数用于合并两个已排序的子数组，并计算逆序对的数量。
// 参数 nums 是待合并的数组，start、mid、end 是数组的索引，用于定义两个子数组。
// 返回值是合并后的数组中逆序对的数量。
func merge(nums []int, start, mid, end int) int {
	// 创建一个临时数组，用于存放合并后的结果。
	temp := make([]int, end-start+1)
	// 初始化两个指针 i 和 j，分别指向两个子数组的开始位置。
	// 初始化一个指针 k，用于指向临时数组的当前位置。
	i, j, k := start, mid+1, 0
	// 初始化逆序对计数器。
	count := 0

	// 遍历两个子数组，将较小的元素放入临时数组，并移动相应的指针。
	for i <= mid && j <= end {
		// 当左半部分的当前元素大于右半部分的当前元素时，
		// 左半部分当前元素到中间元素的所有元素都与右半部分当前元素构成逆序对。
		if nums[i] > nums[j] {
			count += mid - i + 1
			temp[k] = nums[j]
			j++
		} else {
			temp[k] = nums[i]
			i++
		}
		k++
	}

	// 如果左半部分还有剩余元素，将其复制到临时数组中。
	for i <= mid {
		temp[k] = nums[i]
		i++
		k++
	}

	// 如果右半部分还有剩余元素，将其复制到临时数组中。
	for j <= end {
		temp[k] = nums[j]
		j++
		k++
	}

	// 将临时数组中的元素复制回原数组。
	for k, v := range temp {
		nums[start+k] = v
	}

	// 返回逆序对的数量。
	return count
}

// MinNumberInRotateArray 用于在旋转排序数组中找到最小的数字。
// 参数：
//
//	nums - 一个旋转排序数组，可能包含重复元素。
//
// 返回值：
//
//	返回数组中的最小数字。
func MinNumberInRotateArray(nums []int) int {
	// 初始化左右指针，分别指向数组的起始和末尾位置
	left, right := 0, len(nums)-1

	// 使用二分查找法来寻找最小值
	for left < right {
		// 计算中间位置
		mid := (left + right) / 2

		if nums[mid] < nums[right] {
			// 如果中间值小于右边界值，说明最小值在左半部分（包括中间值）
			right = mid
		} else if nums[mid] > nums[right] {
			// 如果中间值大于右边界值，说明最小值在右半部分（不包括中间值）
			left = mid + 1
		} else {
			// 如果中间值等于右边界值，无法确定最小值在哪一侧，因此缩小右边界
			right -= 1
		}
	}

	// 当循环结束时，left 和 right 指向同一个位置，该位置即为最小值的位置
	return nums[left]
}

// CompareVersion 比较两个版本号的大小。
// 版本号以字符串形式提供，格式为"X.Y.Z..."，其中X、Y、Z等是数字。
// 返回值为整数，当version1大于version2时返回1，当version1小于version2时返回-1，当两者相等时返回0。
func CompareVersion(version1 string, version2 string) int {
	// 将版本号version1按"."分割为多个部分
	part1 := strings.Split(version1, ".")
	// 将版本号version2按"."分割为多个部分
	part2 := strings.Split(version2, ".")

	// 初始化长度为version1的段数
	length := len(part1)
	// 如果version2的段数多于version1，则更新长度为version2的段数
	if len(part2) > len(part1) {
		length = len(part2)
	}

	// 遍历每个版本号的部分进行比较
	for i := 0; i < length; i++ {
		// 初始化n1为当前version1的部分，如果当前部分不存在，则为0
		var n1 int64
		if i < len(part1) {
			n1, _ = strconv.ParseInt(part1[i], 10, 64)
		}

		// 初始化n2为当前version2的部分，如果当前部分不存在，则为0
		var n2 int64
		if i < len(part2) {
			n2, _ = strconv.ParseInt(part2[i], 10, 64)
		}

		// 如果version1的部分大于version2的部分，则返回1
		if n1 > n2 {
			return 1
		} else if n1 < n2 { // 如果version1的部分小于version2的部分，则返回-1
			return -1
		}
	}

	// 如果所有部分都相等，则返回0
	return 0
}
