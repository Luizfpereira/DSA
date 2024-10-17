class Node:
    def __init__(self, value) -> None:
        self.value = value
        self.next = None

class LinkedList:
    def __init__(self) -> None:
        self.head = None
        self.tail = None
        self.length = 0

    def __str__(self) -> str:
        curr = self.head
        result = ''

        while curr is not None:
            result += str(curr.value)
            if curr.next is not None:
                result += ' -> '
            curr = curr.next

        return result

    def append(self, value):
        new_node = Node(value)
        if self.head is None:
            self.head = new_node
            self.tail = new_node
        else:
            self.tail.next = new_node
            self.tail = new_node
        self.length += 1

    def prepend(self, value):
        new_node = Node(value)
        if self.head is None:
            self.head = new_node
            self.tail = new_node
        else:
            new_node.next = self.head
            self.head = new_node
        self.length += 1

    def insert(self, index, value):
        if index < 0 or index > self.length:
            return False

        new_node = Node(value)
        if self.head is None:
            self.head = new_node
            self.tail = new_node
        elif index == 0:
            new_node.next = self.head
            self.head = new_node
        else:
            temp_node = self.head

            for _ in range(index-1):
                temp_node = temp_node.next
            new_node.next = temp_node.next
            temp_node.next = new_node

            if new_node.next is None:
                self.tail = new_node
        self.length += 1

        return True
    
    def traverse(self):
        current = self.head
        while current:
            print(current.value)
            current = current.next

    def search(self, target):
        current = self.head
        index = 0

        while current:
            if current.value == target:
                return index
            current = current.next
            index += 1
        return -1
            



# new_list = LinkedList()
# print(new_list.head, new_list.tail)
# print(new_list)

# new_list.append(10)
# new_list.append(20)
# new_list.insert(2, 50)

# print(new_list.head.value, new_list.tail.value)
# print(new_list.search(50))
