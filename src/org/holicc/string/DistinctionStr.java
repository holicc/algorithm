package org.holicc.string;

import java.util.Objects;

public class DistinctionStr {

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

}
