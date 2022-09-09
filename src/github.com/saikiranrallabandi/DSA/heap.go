//Suppose you have a hospital in which patients are attended based on their ages. The oldest are always attended first, no matter when he/she got in the queue.
//You can't just keep track of the oldest one because
//if you pull he/she out, you don't know the next oldest one. In order to solve this hospital problem, you implement a max heap. This heap is, by definition, partially ordered. This means you cannot sort the patients by their age, but you know that the oldest ones are always in the top, so you can pull a patient out in constant time O(1) and re-balance the heap in log time O(log N).

//Suppose you have a sequence of integers and you want to keep track of the median. The median is the number that is in the middle of an ordered array.

//Example:

//[1, 2, 5, 7, 23, 27, 31]
//In the above case, 7 is the median because the array containing the smaller numbers [1, 2, 5] is of the same size of the one containing the bigger numbers [23, 27, 31]. Normally, if the array has an even number of elements, the median is the arithmetic average of the 2 elements in the middle, e.g (5 + 7)/2.

//Now, how do you keep track of the median? By having 2 heaps, one min heap containing the numbers smaller than the current median and a max heap containing the numbers bigger than the current median. Now, if these heaps are always balanced, the 2 heaps will contain the same number of elements or one will have 1 element more than the other, the most.

//When you add a new element to the sequence, if the number is smaller than the current median, you add it to the min heap, otherwise, you add it to the max heap. Now, if the heaps are unbalanced (one heap has more than 1 element more than the other), you pull an element from the biggest heap and add to the smallest. Now they're balanced.

