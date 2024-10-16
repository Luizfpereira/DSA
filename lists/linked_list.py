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



new_list = LinkedList()
print(new_list.head, new_list.tail)
print(new_list)

new_list.append(10)
new_list.append(20)

print(new_list.head.value, new_list.tail.value)
print(new_list)

