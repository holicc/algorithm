package org.holicc.collection;

import java.util.Stack;

/**
 * 包含min函数的栈
 * <p>
 * https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/submissions/
 */
public class MinStack {

    private Stack<Integer> stackA;
    private Stack<Integer> stackB;

    /**
     * initialize your data structure here.
     */
    public MinStack() {
        stackA = new Stack<>();
        stackB = new Stack<>();
    }

    public void push(int x) {
        if (stackA.isEmpty() || x <= stackB.peek()) stackB.push(x);
        stackA.push(x);
    }

    public void pop() {
        if (stackA.pop().equals(stackB.peek())) stackB.pop();
    }

    public int top() {
        return stackA.peek();
    }

    public int min() {
        return stackB.peek();
    }

}
