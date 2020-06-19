package org.holicc.bit;

/**
 * 位运算和二进制相关
 */
public class BitRelated {

    public static void main(String[] args) {
        System.out.println(Integer.toBinaryString(9));
        System.out.println(Integer.toBinaryString(8));
        System.out.println(Integer.toBinaryString(7));
        System.out.println(Integer.toBinaryString(9 & (8)));
    }

    /**
     * 二进制中1的个数
     * 时间复杂度 O(n)
     * <p>
     * https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
     */
    public int hammingWeight(int n) {
        int r = 0;
        while (n != 0) {
            r++;
            n = n & (n - 1);
        }
        return r;
    }

    /**
     * 不用加减乘除做加法
     * 时间复杂度 O(1)
     * 空间复杂度 O(1)
     * <p>
     * https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
     */
    public int add(int a, int b) {
        while (b != 0) { // 当进位为 0 时跳出
            int c = (a & b) << 1;  // c = 进位
            a ^= b; // a = 非进位和
            b = c; // b = 进位
        }
        return a;
    }

}
