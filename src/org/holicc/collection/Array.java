package org.holicc.collection;

public class Array {

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
     * 时间复杂度O(n)
     * 空间复杂度O(n)
     * <p>
     * https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
     */
    public int lastRemaining_01(int n, int m) {
        if (n == 1) return 0;
        int x = lastRemaining_01(n - 1, m);
        return (m + x) % n;
    }

    /**
     * 圆圈中最后剩下的数字
     * 时间复杂度O(n)
     * 空间复杂度O(1)
     * <p>
     * https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
     */
    public int lastRemaining_02(int n, int m) {
        int f = 0;
        for (int i = 2; i != n + 1; ++i)
            f = (m + f) % i;
        return f;
    }

    /**
     * 调整数组顺序使奇数位于偶数前面
     * <p>
     * 时间复杂度 O(n)
     * 空间复杂度 O(n)
     * <p>
     * https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
     */
    public int[] exchange_01(int[] nums) {
        int[] r = new int[nums.length];
        int k = 0;
        for (int num : nums) {
            if ((num & 1) == 1) {
                r[k++] = num;
            }
        }
        //
        for (int num : nums) {
            if ((num & 1) == 0) {
                r[k++] = num;
            }
        }
        return r;
    }

    /**
     * 调整数组顺序使奇数位于偶数前面
     * <p>
     * 时间复杂度 O(n)
     * 空间复杂度 O(1)
     * <p>
     * https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
     */
    public int[] exchange_02(int[] nums) {
        for (int i = 0, j = nums.length - 1; i < j; ) {
            if ((nums[i] & 1) == 1) i++;
            else if ((nums[j] & 1) == 0) j--;
            else {
                int t = nums[i];
                nums[i] = nums[j];
                nums[j] = t;
            }
        }
        return nums;
    }

    /**
     * 和为s的两个数字
     * <p>
     * 时间复杂度 O(n)
     * 空间复杂度 O(1)
     * <p>
     * https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
     */
    public int[] twoSum_01(int[] nums, int target) {
        int[] r = new int[2];
        for (int i = 0; i < nums.length; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                if (nums[i] + nums[j] == target) return new int[]{nums[i], nums[j]};
            }
        }
        return r;
    }

    /**
     * 和为s的两个数字
     * <p>
     * 时间复杂度 O(n)
     * 空间复杂度 O(1)
     * <p>
     * https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof/
     */
    public int[] twoSum_02(int[] nums, int target) {
        int i = 0;
        int j = nums.length - 1;
        while (i < j) {
            int s = nums[i] + nums[j];
            if (s < target) i++;
            else if (s > target) j--;
            else return new int[]{nums[i], nums[j]};
        }
        return new int[0];
    }
}
