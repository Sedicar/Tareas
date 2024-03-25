class Node {
    constructor(data, next = null) {
        this.data = data;
        this.next = next;
    }
}

class Stack {
    constructor() {
        this.head = null;
        this.size = 0;
        this.Length = 0;
    }
    Empty() {
        return this.size === 0;
    }
    Push(data) {
        let node = new Node(data);
        if (!(this.Empty())) {
            node.next = this.head;
        }
        this.head = node;
        this.size++;
        this.Length++;
    }
    Pop() {
        this.size--;
    }
    Peek() {
        if (this.Empty()) {
            console.log("Esta vacia");
        } else {
            let Limit = this.Length - this.size;
            let current = this.head;
            for (let i = 0; i < Limit; i++) {
                current = current.next;
            }
            console.log(current.data);
        }
    }
}