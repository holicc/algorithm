package org.holicc.string;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;

public class Str {

    public static void main(String[] args) {
        Str str = new Str();
        str.minNumber(new int[]{10, 2});
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

    /**
     * 把数组排成最小的数
     * <p>
     * https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
     */
    public String minNumber(int[] nums) {
        String[] strs = new String[nums.length];
        for (int i = 0; i < nums.length; i++)
            strs[i] = String.valueOf(nums[i]);
        Arrays.sort(strs, (x, y) -> (x + y).compareTo(y + x));
        StringBuilder res = new StringBuilder();
        for (String s : strs)
            res.append(s);
        return res.toString();
    }

    public String reverseWords(String s) {
        List<String> list = Arrays.stream(s.split(" "))
                .map(w -> {
                    char[] chars = w.toCharArray();
                    for (int i = 0, j = chars.length - 1; i < j; i++, j--) {
                        char t = chars[i];
                        chars[i] = chars[j];
                        chars[j] = t;
                    }
                    return new String(chars);
                })
                .collect(Collectors.toList());
        StringBuilder builder = new StringBuilder();
        for (String c : list) {
            builder.append(c).append(" ");
        }
        builder.deleteCharAt(builder.length() - 1);
        return builder.toString();
    }

    /**
     * 反转字符串 II
     * <p>
     * https://leetcode-cn.com/problems/reverse-string-ii/
     */
    public String reverseStr(String s, int k) {
        char[] a = s.toCharArray();
        for (int start = 0; start < a.length; start += 2 * k) {
            int i = start, j = Math.min(start + k - 1, a.length - 1);
            while (i < j) {
                char tmp = a[i];
                a[i++] = a[j];
                a[j--] = tmp;
            }
        }
        return new String(a);
    }


    /**
     * 仅仅反转字母
     * <p>
     * https://leetcode-cn.com/problems/reverse-only-letters/
     */
    public String reverseOnlyLetters(String S) {
        Stack<Character> letters = new Stack();
        for (char c : S.toCharArray())
            if (Character.isLetter(c))
                letters.push(c);

        StringBuilder ans = new StringBuilder();
        for (char c : S.toCharArray()) {
            if (Character.isLetter(c))
                ans.append(letters.pop());
            else
                ans.append(c);
        }
        return ans.toString();
    }
}
