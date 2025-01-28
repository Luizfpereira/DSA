class Node:
    def __init__(self, value=None):
        self.value = value
        self.next = None

    def __str__(self):
        return str(self.value)

class LinkedList:
    def __init__(self):
        self.head = None
        self.tail = None
    
    def __iter__(self):
        curNode = self.head
        while curNode:
            yield curNode
            curNode = curNode.next

class Queue:
    def __init__(self):
        self.linkedList = LinkedList()

    def __str__(self):
        values = [str(x) for x in self.linkedList]
        return ' '.join(values)

    def enqueue(self, value):
        newNode = Node(value)
        if not self.linkedList.head:
            self.linkedList.head = newNode
            self.linkedList.tail = newNode
        else:
            self.linkedList.tail.next = newNode
            self.linkedList.tail = newNode

    def peek(self):
        if not self.linkedList.head:
            return "Queue is empty"
        else:
            return self.linkedList.head
    
    def isEmpty(self):
        if not self.linkedList.head:
            return True
        return False
    
    def dequeue(self):
        if self.isEmpty():
            return

        tempNode = self.linkedList.head
        if self.linkedList.head == self.linkedList.tail:
            self.linkedList.head = None
            self.linkedList.tail = None
        else:
            self.linkedList.head = self.linkedList.head.next
            tempNode.next = None


queue = Queue()
queue.enqueue(1)
queue.enqueue(2)
queue.enqueue(3)
print(queue) #1 2 3
queue.dequeue()
print(queue.peek()) # 2
queue.enqueue(4)
print(queue) # 2 3 4
queue.dequeue()
queue.dequeue()
queue.dequeue()
queue.dequeue()
print(queue) # 
