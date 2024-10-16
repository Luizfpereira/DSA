import unittest
from linked_list import LinkedList

class TestLinkedList(unittest.TestCase):

        def test_initial_state(self):
            new_list = LinkedList()
            self.assertIsNone(new_list.head)
            self.assertIsNone(new_list.tail)
            self.assertEqual(new_list.length, 0)

        def test_append_single_node(self):
            new_list = LinkedList()
            new_list.append(10)
            self.assertEqual(new_list.head.value, 10)
            self.assertEqual(new_list.tail.value, 10)
            self.assertEqual(new_list.length, 1)

        def test_append_multiple_nodes(self):
            new_list = LinkedList()
            new_list.append(10)
            new_list.append(20)
            self.assertEqual(new_list.head.value, 10)
            self.assertEqual(new_list.tail.value, 20)
            self.assertEqual(new_list.length, 2)

        def test_str_multiple_nodes(self):
            new_list = LinkedList()
            new_list.append(10)
            new_list.append(20)
            expected = '10 -> 20'
            self.assertEqual(expected, str(new_list))


if __name__ == '__main__':
    unittest.main()