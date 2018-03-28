"""
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.
"""
class Solution(object):
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) == 1:
            return nums[0]
        
        candidate = nums[0]
        counter = 1
        
        for num in nums:
            if num == candidate:
                counter += 1
            else:
                counter -= 1
                if counter == 0:
                    candidate = num
                    counter = 1
                    
        return candidate
  
