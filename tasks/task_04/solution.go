package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	if len(nums) < 2 {
		return Stats{}
	}
	res := Stats{
		Count: len(nums) - 1,
		Sum:   nums[len(nums)-1] - nums[0],
		Min:   nums[1] - nums[0],
		Max:   nums[1] - nums[0],
	}

	for i := range len(nums) - 1 {
		res.Min = min(nums[i+1]-nums[i], res.Min)
		res.Max = max(res.Max, nums[i+1]-nums[i])
	}
	return res
}
