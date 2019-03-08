
public class DoublyLinkedList {
    private Node head;
    private Node tail;

    private class Node {
        int element;
        Node next;
        Node prev;
 
        public Node(int element, Node next, Node prev) {
            this.element = element;
            this.next = next;
            this.prev = prev;
        }
    }

    public void addFirst(int element) {
        Node tmp = new Node(element, head, null);
        if(head != null ) {head.prev = tmp;}
        head = tmp;
        if(tail == null) { tail = tmp;}
    }

    public void addLast(int element) {
         
        Node tmp = new Node(element, null, tail);
        if(tail != null) {tail.next = tmp;}
        tail = tmp;
        if(head == null) { head = tmp;}
    }

    void deleteNode(Node head_ref, Node del) 
    {  
        if (head == null || del == null) { 
            return; 
        } 
        if (head == del) { 
            head = del.next; 
        } 
        if (del.next != null) { 
            del.next.prev = del.prev; 
        } 

        if (del.prev != null) { 
            del.prev.next = del.next; 
        } 
        return; 
    }

    public void printNodes(){
        Node tmp = head;
        while(tmp != null){
            System.out.println(tmp.element);
            tmp = tmp.next;
        }
    }

    public static void main(String args[]){
         
        DoublyLinkedList dll = new DoublyLinkedList();
        dll.addFirst(10);
        dll.addFirst(34);
        dll.addLast(56);
        dll.addLast(364);
        dll.printNodes();
        dll.deleteNode(dll.head, dll.head);
        dll.printNodes();
    }
}