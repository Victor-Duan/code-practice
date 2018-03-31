/*
225. Implement Stack using Queues

Implement the following operations of a stack using queues.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
empty() -- Return whether the stack is empty.
Notes:
You must use only standard operations of a queue -- which means only push to back, peek/pop from front, size, and is empty operations are valid.
Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue), as long as you use only standard operations of a queue.
You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).
*/

type MyStack struct {
    values []int 
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{values: []int{}} 
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.values = append(this.values, x) 
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    if this.Empty() {
        return 0
    }
    
    topVal := this.values[len(this.values) - 1]
    this.values = this.values[:len(this.values) - 1] 
    return topVal
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.values[len(this.values) - 1]  
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.values) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
