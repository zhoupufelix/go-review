##归并排序

归并操作的工作原理如下：
第一步：申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列
第二步：设定两个指针，最初位置分别为两个已经排序序列的起始位置
第三步：比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置
重复步骤3直到某一指针超出序列尾
将另一序列剩下的所有元素直接复制到合并序列尾

时间复杂度O(n log n)  空间复杂度T（n)


```go

func MergeSort(nums []int)[]int{
  result := mergeSort(nums)
  return result
}

func mergeSort(nums []int)[]int{
   if len(nums) <= 1{
   	  return nums
   }
   mid := len(nums)/2
   left := mergeSort(nums[:mid])
   right := mergeSort(nums[mid:])
   result := merge(left,right)
    return result
}

func merge(left,right []int)(result []int){
    l:=0
    r:=0
    for l<len(left) && r<len(right){
    	if  left[l] < right[r]{
            result = append(result,left[l])
            l++
        }else{
            result = append(result,right[r])
            r++
        }
    }
    result = append(result,left[l:])
    result = append(result,left[r:])
    return result
}

```