package org.holicc.collection;

import java.util.*;
import java.util.stream.IntStream;

public class Array {

    public static void main(String[] args) {
        Array array = new Array();
        System.out.println(array.constructArr_02(new int[]{1, 2, 3, 4, 5}));
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

    /**
     * 最小的k个数
     * 时间复杂度 O(n log(n))
     * 空间复杂度 O(k)
     * <p>
     * https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
     */
    public int[] getLeastNumbers_01(int[] arr, int k) {
        Arrays.sort(arr);
        int[] r = new int[k];
        System.arraycopy(arr, 0, r, 0, k);
        return r;
    }

    /**
     * 最小的k个数
     * 时间复杂度 O(n)
     * 空间复杂度 O(1)
     * <p>
     * https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
     */
    public int[] getLeastNumbers_02(int[] arr, int k) {
        if (k == 0 || arr.length == 0) {
            return new int[0];
        }
        // 最后一个参数表示我们要找的是下标为k-1的数
        return quickSearch(arr, 0, arr.length - 1, k - 1);
    }

    private int[] quickSearch(int[] nums, int lo, int hi, int k) {
        // 每快排切分1次，找到排序后下标为j的元素，如果j恰好等于k就返回j以及j左边所有的数；
        int j = partition(nums, lo, hi);
        if (j == k) {
            return Arrays.copyOf(nums, j + 1);
        }
        // 否则根据下标j与k的大小关系来决定继续切分左段还是右段。
        return j > k ? quickSearch(nums, lo, j - 1, k) : quickSearch(nums, j + 1, hi, k);
    }

    // 快排切分，返回下标j，使得比nums[j]小的数都在j的左边，比nums[j]大的数都在j的右边。
    private int partition(int[] nums, int lo, int hi) {
        int v = nums[lo];
        int i = lo, j = hi + 1;
        while (true) {
            while (++i <= hi && nums[i] < v) ;
            while (--j >= lo && nums[j] > v) ;
            if (i >= j) {
                break;
            }
            int t = nums[j];
            nums[j] = nums[i];
            nums[i] = t;
        }
        nums[lo] = nums[j];
        nums[j] = v;
        return j;
    }

    /**
     * 构建乘积数组
     * 时间复杂度 O(n^2) 超时
     * 空间复杂度 O(n)
     * <p>
     * https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/
     */
    public int[] constructArr_01(int[] a) {
        int[] r = new int[a.length];
        //
        for (int i = 0; i < r.length; i++) {
            int t = a[i];
            int mul = 1;
            for (int j = 0; j < a.length; j++) {
                if (j != i) mul *= a[j];
            }
            r[i] = mul;
        }
        return r;
    }

    /**
     * 构建乘积数组
     * <p>
     * https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/
     */
    public int[] constructArr_02(int[] a) {
        if (a.length == 0) return new int[0];
        int[] r = new int[a.length];
        //
        r[0] = 1;
        //
        for (int i = 1; i < r.length; i++) {
            r[i] = r[i - 1] * a[i - 1];
        }
        //
        int t = 1;
        for (int i = a.length - 2; i >= 0; i--) {
            t *= a[i + 1];
            r[i] *= t;
        }
        return r;
    }

    /**
     * 在排序数组中查找数字 I
     * <p>
     * https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
     */
    public int search_01(int[] nums, int target) {
        if (nums.length == 0) return 0;
        Map<Integer, Integer> table = new HashMap<>();
        //
        for (int num : nums) {
            table.put(num, table.getOrDefault(num, 0) + 1);
        }
        //
        return table.getOrDefault(target, 0);
    }

    public int search(int[] nums, int target) {
        return helper(nums, target) - helper(nums, target - 1);
    }

    int helper(int[] nums, int tar) {
        int i = 0, j = nums.length - 1;
        while (i <= j) {
            int m = (i + j) / 2;
            if (nums[m] <= tar) i = m + 1;
            else j = m - 1;
        }
        return i;
    }

    /**
     * 旋转数组的最小数字
     * <p>
     * https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
     */
    public int minArray(int[] numbers) {
        return IntStream.of(numbers).min().orElse(-1);
    }

    public int minArray_01(int[] numbers) {
        for (int i = numbers.length - 1; i >= 0; i--) {
            if (i == 0) return numbers[i];
            if (numbers[i - 1] > numbers[i]) return numbers[i - 1];
        }
        return -1;
    }

    /**
     * 顺时针打印矩阵
     * <p>
     * https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
     */
    public int[] spiralOrder(int[][] matrix) {
        if (matrix.length == 0) return new int[0];
        int l = 0, r = matrix[0].length - 1, t = 0, b = matrix.length - 1, x = 0;
        int[] res = new int[(r + 1) * (b + 1)];
        while (true) {
            for (int i = l; i <= r; i++) res[x++] = matrix[t][i]; // left to right.
            if (++t > b) break;
            for (int i = t; i <= b; i++) res[x++] = matrix[i][r]; // top to bottom.
            if (l > --r) break;
            for (int i = r; i >= l; i--) res[x++] = matrix[b][i]; // right to left.
            if (t > --b) break;
            for (int i = b; i >= t; i--) res[x++] = matrix[i][l]; // bottom to top.
            if (++l > r) break;
        }
        return res;
    }

    /**
     * 0～n-1中缺失的数字
     * 时间复杂度 O(n)
     * 空间复杂度 O(1)
     * <p>
     * https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
     */
    public int missingNumber(int[] nums) {
        if (nums.length == 1 && nums[0] == 0) return 1;
        int sum = (nums.length * (nums.length - 1)) / 2;
        int real = 0;
        for (int num : nums) {
            real += num;
        }
        return sum - real;
    }

    /**
     * 扑克牌中的顺子
     * 时间复杂度 O (log(n)
     * 空间发展度 O (1)
     * <p>
     * https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
     */
    public boolean isStraight(int[] nums) {
        int zero = 0;
        int tmp = -1;
        Arrays.sort(nums);
        for (int num : nums) {
            if (zero < 0) return false;
            if (num == 0) {
                zero++;
            } else {
                if (tmp == num) return false;
                if (tmp != -1) zero -= (Math.abs(num - tmp) - 1);
                tmp = num;
            }
        }
        return zero >= 0;
    }

    /**
     * 滑动窗口的最大值
     * 时间复杂 O( n^k)
     * 空间复杂度 O(n)
     * <p>
     * https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
     */
    public int[] maxSlidingWindow(int[] nums, int k) {
        if (nums.length == 0) return new int[0];
        int[] r = new int[nums.length - k + 1];
        for (int i = 0; i < r.length; i++) {
            int max = Integer.MIN_VALUE;
            for (int j = i; j < i + k; j++) {
                max = Math.max(max, nums[j]);
            }
            r[i] = max;
        }
        return r;
    }

    public int[] maxSlidingWindow_01(int[] nums, int k) {
        if (nums.length == 0) return new int[0];
        //
        int[] r = new int[nums.length - k + 1];
        for (int i = 0, j = k - 1, maxIndex = -1, max = Integer.MIN_VALUE; j < nums.length; i++, j++) {
            if (maxIndex == -1) {
                maxIndex = findMaxIndex(nums, i, j);
                max = nums[maxIndex];
            } else {
                //
                if (maxIndex >= i && maxIndex <= j) {
                    max = Math.max(max, nums[j]);
                } else {
                    maxIndex = findMaxIndex(nums, i, j);
                    max = nums[maxIndex];
                }
            }
            //
            r[i] = max;
        }
        return r;
    }

    private int findMaxIndex(int[] nums, int i, int j) {
        int max = Integer.MIN_VALUE;
        int maxIndex = -1;
        for (int l = i; l <= j; l++) {
            if (max < nums[l]) {
                maxIndex = l;
                max = nums[l];
            }
        }
        return maxIndex;
    }



    /**
     * 二维数组中的查找
     * 时间复杂度 O(n^2)
     * 空间复杂度 O(1)
     * <p>
     * https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
     */
    public boolean findNumberIn2DArray(int[][] matrix, int target) {
        if (matrix.length == 0) return false;
        for (int i = 0; i < matrix.length; i++) {
            for (int j = 0; j < matrix[i].length; j++) {
                if (matrix[i][j] == target) return true;
            }
        }
        return false;
    }

    /**
     * 二维数组中的查找
     * 时间复杂度 O(n)
     * 空间复杂度 O(1)
     * <p>
     * https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
     */
    public boolean findNumberIn2DArray_01(int[][] matrix, int target) {
        if (matrix.length == 0) return false;
        int rows = matrix.length, cols = matrix[0].length;
        int row = 0, col = cols - 1;
        while (row < rows && col >= 0) {
            int t = matrix[row][col];
            if (t == target) {
                return true;
            } else if (t > target) {
                col--;
            } else {
                row++;
            }
        }
        return false;
    }
}

