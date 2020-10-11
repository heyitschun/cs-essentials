import random

class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

    def __str__(self):
        return f"<Node: data={self.data} next={self.next}>"

class LinkedList:
    head = None

    def append(self, data):
        if self.head == None:
            self.head = Node(data)
            return

        current = self.head
        while current.next != None:
            current = current.next

        current.next = Node(data)

    def prepend(self, data):
        new_head = Node(data)
        new_head.next = self.head
        self.head = new_head

    def delete(self, data):
        if self.head == None:
            return

        if self.head.data == data:
            self.head = self.head.next

        current = self.head
        while current.next != None:
            if current.next.data == data:
                current.next = current.next.next
                return
            current = current.next

    def view_tree(self):
        current = self.head
        while current.next != None:
            print(f"<Node: data={current.data}")
            current = current.next


ll = LinkedList()

for i in range(10):
    ll.append(Node(random.randint(0,100)))

ll.view_tree()

print("END")
