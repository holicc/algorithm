package org.holicc.string;

import java.util.Objects;

public class Str {

    public static void main(String[] args) {
        Str str = new Str();
        System.out.println(str.reverseWords("a good   example"));
    }

    /**
     * 第一个只出现一次的字符
     * 时间复杂度 O(n)
     * 空间复杂度 O(n)
     * <p>
     * https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
     */
    public char firstUniqChar(String s) {
        if (Objects.isNull(s) || s.equals("")) return ' ';
        char[] table = new char[26];
        //
        for (char c : s.toCharArray()) {
            table[c - 'a'] += 1;
        }
        //
        for (char c : s.toCharArray()) {
            if (table[c - 'a'] == 1) return c;
        }
        return ' ';
    }

    /**
     * 替换空格
     * <p>
     * https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
     */
    public String replaceSpace(String s) {
        return s == null ? null : s.replaceAll(" ", "%20");
    }

    /**
     * 翻转单词顺序
     * <p>
     * https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
     */
    public String reverseWords(String s) {
        String[] strs = s.trim().split(" "); // 删除首尾空格，分割字符串
        StringBuilder res = new StringBuilder();
        for (int i = strs.length - 1; i >= 0; i--) { // 倒序遍历单词列表
            if (strs[i].equals("")) continue; // 遇到空单词则跳过
            res.append(strs[i]).append(" "); // 将单词拼接至 StringBuilder
        }
        return res.toString().trim(); // 转化为字符串，删除尾部空格，并返回
    }

}
