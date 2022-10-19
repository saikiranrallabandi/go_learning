/*
Problem: How will you find the middle element
of a singly linked list without iterating the list more than once

Solution: we can use the two-pointer method.
one is fast and one slow, The faster pointer travels two nodes per step,
while the slow pointer only moves one. The slow pointer will point to the
middle node of the linked list

when the fast pointer points to the last node, i.e when the next node is null.
*/

func (node Node*) getMiddle(Node *head) {
	struct Node *slow = head;
	struct Node *fast = head;

	if (head) {
		while (fast != NULL && fast->next !=NULL) {
			fast = fast->next->next;
			slow = slow->next;
		}
	}
	return slow
}