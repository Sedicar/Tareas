class Node {
    constructor(data, next = null) {
        this.data = data;
        this.next = next;
    }
}

class LinkedList {
    constructor() {
        this.size = 0;
        this.head = null;
    }

    getSize() {
        return this.size;
    }

    isEmpty() {
        return this.size === 0;
    }

    insertFirst(data) {
        const node = new Node(data);
        if (this.isEmpty()) {
            this.head = node;
        } else {
            node.next = this.head;
            this.head = node;
        }
        this.size++;
    }

    insertLast(data) {
        const node = new Node(data);
        if (this.isEmpty()) {
            this.head = node;
        } else {
            let current = this.head;
            while (current.next) {
                current = current.next;
            }
            current.next = node;
        }
        this.size++;
    }

    insert(data, index) {
        if (index < 0 || index > this.size) {
            console.log("El indice no es valido");
            return;
        }
        if (index === 0) {
            this.insertFirst(data);
        } else {
            const node = new Node(data);
            let prev = this.head;
            for (let i = 0; i < index - 1; i++) {
                prev = prev.next;
            }
            node.next = prev.next;
            prev.next = node;
            this.size++;
        }
    }

    removeIndex(index) {
        index--;
        if (index < 0 || index >= this.size) {
            console.log("El indice no es valido");
            return null;
        }
        if (index === 0) {
            this.head = this.head.next;
        } else {
            let prev = this.head;
            for (let i = 0; i < index - 1; i++) {
                prev = prev.next;
            }
            prev.next = prev.next.next;
        }
        this.size--;
    }

    searchIndex(index) {
        index--;
        if (index < 0 || index > this.size) {
            console.log("El indice no es valido");
            return null;
        }
        let current = this.head;
        for (let i = 0; i < index; i++) {
            current = current.next;
        }
        return current.data;
    }


    // Actividad:
    searchValue(value) {
        let compare = 0;
        let tries = 0;
        for (let i = 1; i < this.size + 1; i++) {
            compare = this.searchIndex(i);
            if (value === compare) {
                return i;
            }
            tries++;
        }
        if (tries === this.size) {
            console.log("No se encontro el valor");
            return null;
        }
    }
    removeValue(value) {
        let compare = "";
        let i = 0;
        let j = 0;
        do {
            compare = this.searchIndex(i + 1);
            if (i === this.size) {
                break;
            }
            if (value === compare) {
                this.size--;
                if ((this.searchValue(compare) - 1) === 0) {
                    this.head = this.head.next;
                    break;
                } else {
                    let container = this.head;
                    do {
                        container = container.next;
                        j++;
                    } while (j < i - 1);
                    container.next = container.next.next;
                    break;
                }
            }
            i++;
        } while (i < this.size);
        if (i === this.size && value !== compare) {
            console.log("No se encontro el valor");
        }
    }
    // *****

    printAll() {
        if (this.isEmpty()) {
            console.log("No hay nada en la lista");

        } else {
            let current = this.head;
            let result = "";
            while (current) {
                result += current.data + ' -> ';
                current = current.next;
            }
            console.log(result + ' null')
        }

    }

}

let lista = new LinkedList();
lista.printAll();
lista.insertFirst("Messi");
lista.printAll();
lista.insertLast("El Pepe");
lista.printAll();
lista.insertFirst("Ash");
lista.printAll()
lista.insertLast("Dross");
lista.printAll();
lista.insert("Goku", 3);
lista.printAll();
lista.removeIndex(3);
lista.printAll();
lista.insert("Osama bin Laden", 2)
lista.insertLast("Jose Gregorio Hernandez"
)
lista.insertLast("Raheem Sterling"
)
lista.insertLast("Pikachu"
)
lista.removeValue("Goku");
lista.removeValue("Pikachu");
console.log("índice:" + lista.searchValue("Raheem Sterling"));
lista.printAll();