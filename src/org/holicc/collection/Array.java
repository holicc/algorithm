package org.holicc.collection;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Array {

    public static void main(String[] args) {
        System.out.println(Math.pow(10, 3) - 1);
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


    /**
     * 数组中重复的数字
     * 时间复杂度 O(n log(n))
     * <p>
     * https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
     */
    public int findRepeatNumber(int[] nums) {
        Arrays.sort(nums);
        for (int i = 1; i < nums.length; i++) {
            if (nums[i - 1] == nums[i]) {
                return nums[i];
            }
        }
        return -1;
    }

    /**
     * 数组中出现次数超过一半的数字
     * 时间复杂度 O(n log(n))
     * <p>
     * https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/
     */
    public int majorityElement(int[] nums) {
        Arrays.sort(nums);
        return nums[nums.length / 2];
    }

    /**
     * 和为s的连续正数序列
     * 时间复杂度 O(n^2)
     * <p>
     * https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
     */
    public int[][] findContinuousSequence_01(int target) {
        List<int[]> r = new ArrayList<>();
        for (int i = 1; i + 2 < target; i++) {
            List<Integer> t = new ArrayList<>();
            int sum = 0;
            for (int j = i; sum <= target; j++) {
                sum += j;
                t.add(j);
                if (sum == target) break;
            }
            if (sum == target) {
                //
                int[] ary = new int[t.size()];
                for (int k = 0; k < ary.length; k++) {
                    ary[k] = t.get(k);
                }
                //
                r.add(ary);
            }
        }
        return r.toArray(new int[0][0]);
    }


    /**
     * 和为s的连续正数序列
     * 时间复杂度 O(n)
     * 使用数列求和公式 S(10)=(1+10)*(10)/2
     * <p>
     * https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
     */
    public int[][] findContinuousSequence_02(int target) {
        List<int[]> r = new ArrayList<>();
        for (int i = 1, j = 2; i < j; ) {
            int sum = (i + j) * (j - i + 1) / 2;
            if (sum < target) j++;
            else if (sum > target) i++;
            else {
                int[] t = new int[j - i + 1];
                for (int n = i, k = 0; k < t.length; k++, n++) {
                    t[k] = n;
                }
                r.add(t);
                i++;
            }
        }
        return r.toArray(new int[0][0]);
    }

    /**
     * 打印从1到最大的n位数
     * 时间复杂度 O(n)
     * 空间复杂度 O(10^n)
     * <p>
     * https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/
     */
    public int[] printNumbers(int n) {
        int size = (int) (Math.pow(10, n) - 1);
        int[] r = new int[size];
        for (int i = 0; i < r.length; i++) {
            r[i] = i + 1;
        }
        return r;
    }

}
