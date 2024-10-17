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

        def test_prepend_multiple_node(self):
            new_list = LinkedList()
            new_list.prepend(10)
            new_list.prepend(20)
            self.assertEqual(new_list.head.value, 20)
            self.assertEqual(new_list.tail.value, 10)
            self.assertEqual(new_list.length, 2)

        def test_insert_empty_linked_list(self):
            new_list = LinkedList()
            new_list.insert(0, 50)
            self.assertEqual(new_list.head.value, 50)
            self.assertEqual(new_list.tail.value, 50)
            self.assertEqual(new_list.length, 1)

        def test_insert_beginning_linked_list(self):
            new_list = LinkedList()
            new_list.append(10)
            new_list.append(20)
            new_list.append(30)
            new_list.insert(0, 50)
            self.assertEqual(new_list.head.value, 50)
            self.assertEqual(new_list.tail.value, 30)
            self.assertEqual(new_list.length, 4)
        
        def test_insert_end_linked_list(self):
            new_list = LinkedList()
            new_list.append(10)
            new_list.append(20)
            new_list.append(30)
            new_list.insert(3, 50)
            self.assertEqual(new_list.head.value, 10)
            self.assertEqual(new_list.tail.value, 50)
            self.assertEqual(new_list.length, 4)

        def test_insert_middle_linked_list(self):
            new_list = LinkedList()
            new_list.append(10)
            new_list.append(20)
            new_list.append(30)
            new_list.insert(2, 50)
            expected = '10 -> 20 -> 50 -> 30'
            self.assertEqual(expected, str(new_list))
            self.assertEqual(new_list.head.value, 10)
            self.assertEqual(new_list.tail.value, 30)
            self.assertEqual(new_list.length, 4)

        def test_insert_negative_index(self):
            new_list = LinkedList()
            new_list.append(10)
            new_list.append(20)
            new_list.append(30)
            res = new_list.insert(-1, 50)
            self.assertEqual(new_list.head.value, 10)
            self.assertEqual(new_list.tail.value, 30)
            self.assertEqual(new_list.length, 3)
            self.assertFalse(res)
        
        def test_insert_index_out_of_bound(self):
            new_list = LinkedList()
            new_list.append(10)
            new_list.append(20)
            new_list.append(30)
            res = new_list.insert(6, 50)
            self.assertEqual(new_list.head.value, 10)
            self.assertEqual(new_list.tail.value, 30)
            self.assertEqual(new_list.length, 3)
            self.assertFalse(res)

        def search_empty_list(self):
            new_list = LinkedList()
            res = new_list.search(50)
            self.assertEqual(res, -1)

        def search_filled_list(self):
            new_list = LinkedList()
            new_list.append(10)
            new_list.append(20)
            new_list.append(30)
            res = new_list.search(30)
            self.assertEqual(res, 2)

        def search_non_existant(self):
            new_list = LinkedList()
            new_list.append(10)
            new_list.append(20)
            new_list.append(30)
            res = new_list.search(50)
            self.assertEqual(res, -1)


if __name__ == '__main__':
    unittest.main()