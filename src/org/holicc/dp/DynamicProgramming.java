package org.holicc.dp;

public class DynamicProgramming {

    /**
     * n个骰子的点数
     * <p>
     * https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/
     */
    public double[] twoSum(int n) {
        if (n == 0) return new double[0];
        int code = 6;//每个骰子有6个数字
        int[][] count = new int[2][n * code + 1];
        int flag = 0;
        //i为骰子数,j为点数,k为移动的位数(1=<k<=6)
        for (int j = 1; j <= code; j++) count[flag][j] = 1;//第1枚骰子,i=1
        for (int i = 2; i <= n; i++) {//第2枚到第n枚骰子
            for (int j = 1; j < i; j++) count[1 - flag][j] = 0;//第1~i-1置0,第0位没有进行修改过，始终为初始值0;
            for (int j = i; j <= i * code; j++) {
                count[1 - flag][j] = 0;//这里必须有清空操作,否则后面的+=操作会基于上上次循环的初始值累加
                for (int k = 1; k <= 6 && j - k >= 1; k++) {
                    count[1 - flag][j] += count[flag][j - k];
                }
            }
            flag = 1 - flag;
        }
        double base = Math.pow(code, n);
        double[] p = new double[n * code - n + 1];
        for (int j = 0; j < p.length; j++) p[j] = count[flag][n + j] / base;//base已经是double类型了，这里前面可不用加(double)转换了
        return p;
    }

    /**
     * 青蛙跳台阶问题
     * <p>
     * https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
     */
    public int numWays(int n) {
        int a = 1, b = 1, sum;
        for (int i = 0; i < n; i++) {
            sum = (a + b) % 1000000007;
            a = b;
            b = sum;
        }
        return a;
    }

    /**
     * 斐波那契数列
     * <p>
     * 时间复杂度 O(n^2)
     * 空间复杂度 O(n)
     * <p>
     * https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
     */
    public int fib(int n) {
        int a = 0, b = 1, sum = 0;
        for (int i = 0; i < n; i++) {
            sum = (a + b) % 1000000007;
            a = b;
            b = sum;
        }
        return sum;
    }
}
