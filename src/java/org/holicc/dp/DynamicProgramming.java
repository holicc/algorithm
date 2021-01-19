package org.holicc.dp;

public class DynamicProgramming {


    public static void main(String[] args) {
        DynamicProgramming programming = new DynamicProgramming();

        programming.maxValue(new int[][]{
                {1, 3, 1},
                {1, 5, 1},
                {4, 2, 1},
        });
    }

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

    /**
     * 礼物的最大价值
     * 时间复杂度 O(nm)
     * 空间复杂度 O(1)
     * <p>
     * https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof/
     */
    public int maxValue(int[][] grid) {
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (i - 1 >= 0 && j - 1 >= 0) {
                    grid[i][j] += Math.max(
                            grid[i - 1][j],
                            grid[i][j - 1]
                    );
                } else if (i - 1 >= 0) {
                    grid[i][j] += grid[i - 1][j];
                } else if (j - 1 >= 0) {
                    grid[i][j] += grid[i][j - 1];
                }
            }
        }
        return grid[grid.length - 1][grid[0].length - 1];
    }

    /**
     * 丑数
     * 时间复杂度 O(n)
     * 空间复杂度 O(n)
     * <p>
     * https://leetcode-cn.com/problems/chou-shu-lcof/
     */
    public int nthUglyNumber(int n) {
        int a = 0, b = 0, c = 0;
        int[] dp = new int[n];
        dp[0] = 1;
        for (int i = 1; i < n; i++) {
            int n2 = dp[a] * 2, n3 = dp[b] * 3, n5 = dp[c] * 5;
            dp[i] = Math.min(Math.min(n2, n3), n5);
            if (dp[i] == n2) a++;
            if (dp[i] == n3) b++;
            if (dp[i] == n5) c++;
        }
        return dp[n - 1];
    }

    /**
     * 股票的最大利润
     * 时间复杂度 O(n)
     * 空间复杂度 O(1)
     * <p>
     * https://leetcode-cn.com/problems/gu-piao-de-zui-da-li-run-lcof/
     */
    public int maxProfit_01(int[] prices) {
        int cost = Integer.MAX_VALUE, profit = 0;
        for (int price : prices) {
            cost = Math.min(cost, price);
            profit = Math.max(profit, price - cost);
        }
        return profit;
    }

    /**
    * 剪绳子
    * 时间复杂度 O(n)
    * 空间复杂度 O(n2)
    * <p>
    * https://leetcode-cn.com/problems/jian-sheng-zi-lcof/
    */
    public int cuttingRope(int n) {
        int[] dp = new int[n + 1];
        dp[1] = 1;
        for (int i = 2; i <= n; i++) {
            for (int j = 1; j < i; j++) {
                dp[i] = Math.max(dp[i], (Math.max(j, dp[j])) * (Math.max(i - j, dp[i - j])));
            }
        }
        return dp[n];
    }
}
