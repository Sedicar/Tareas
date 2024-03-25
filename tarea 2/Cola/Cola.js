class Node {
    constructor(data, next = null) {
        this.data = data;
        this.next = next;
    }
}

class Queue {
    constructor() {
        this.head = null;
        this.size = 0;
        this.Length = 1;
    }
    Empty() {
        return this.size === 0;
    }
    Enqueue(data) {
        let node = new Node(data);
        if (!(this.Empty())) {
            node.next = this.head;
        }
        this.head = node;
        this.size++;
    }
    Dequeue() {
        this.Length++;
    }
    Peek() {
        if (this.Empty() || this.Length>this.size) {
            console.log("No hay nadie en la cola");
        } else {
            let current = this.head;
            for (let i = this.size; i > this.Length; i--) {
                current = current.next;
            }
            console.log(current.data);
        }
    }
}

let Lista = new Queue();
Lista.Enqueue({
    Nombre: "Hugo Chavez",
    Trabajo: "Mejor presidente del mundo"
});
Lista.Enqueue({
    Nombre: "Mario Bros",
    Trabajo: "Fontanero"
});
Lista.Enqueue({
    Nombre: "Cristiano Ronaldo",
    Trabajo: "Futbolista"
});
Lista.Enqueue({
    Nombre: "Jhon Wick",
    Trabajo: "La muerte"
});
Lista.Enqueue({
    Nombre: "Jonesy",
    Trabajo: "Bailarin"
});
Lista.Enqueue({
    Nombre: "Samus Aran",
    Trabajo: "Caza Recompensas"
});

Lista.Dequeue();


Lista.Peek();