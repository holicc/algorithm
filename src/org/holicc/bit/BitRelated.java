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

}
