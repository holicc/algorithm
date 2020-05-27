package org.holicc.collection;

public class Array {

    public static void main(String[] args) {
        Array array = new Array();
        int i = array.maxSubArray_02(new int[]{-2, 1});
        System.out.println(i);
    }

    /**
     * 连续子数组的最大和
     * 时间复杂度 O(n^2)
     * <p>
     * https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
     */
    public int maxSubArray_01(int[] nums) {
        int max = Integer.MIN_VALUE;
        for (int num : nums) {
            int sum = num;
            for (int i : nums) {
                sum += i;
                max = Math.max(max, sum);
            }
        }
        return max;
    }

    /**
     * 连续子数组的最大和
     * 时间复杂度 O(n)
     * <p>
     * https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
     */
    public int maxSubArray_02(int[] nums) {
        int max = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (nums[i - 1] >= 0) {
                nums[i] += nums[i - 1];
            }
            max = Math.max(nums[i], max);
        }
        return max;
    }

    /**
     * 圆圈中最后剩下的数字
     * <p>
     * https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
     */
    public int lastRemaining(int n, int m) {
        int[] ary = new int[n];
        for (int i = 0; i < n; i++) {
            ary[i] = i;
        }
        //
        int i = 0;
        while (n-- > 0) {
            while (ary[i++ + m] != -1) { }
            if (i >= n) i = 0;
        }
        return ary[i];
    }

}
