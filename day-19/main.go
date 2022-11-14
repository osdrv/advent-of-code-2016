package main

type Elf struct {
	id    int
	gifts int
	next  *Elf
	prev  *Elf
}

func NewElf(id int) *Elf {
	return &Elf{id: id, gifts: 1}
}

func makeElves(nElves int) *Elf {
	var head, tail *Elf
	for i := 0; i < nElves; i++ {
		e := NewElf(i + 1)
		if head == nil {
			head = e
			tail = e
		}
		tail.next = e
		e.prev = tail
		tail = e
	}
	tail.next = head
	head.prev = tail

	return head
}

func solve1(nElves int) int {
	head := makeElves(nElves)

	ptr := head
	for ptr.next != ptr {
		debugf("playing elf: %d (gifts: %d)", ptr.id, ptr.gifts)
		ptr.gifts += ptr.next.gifts
		ptr.next.gifts = 0
		ptr.next = ptr.next.next
		ptr = ptr.next
	}
	return ptr.id
}

func solve2(nElves int) int {
	head := makeElves(nElves)

	player, counter := head, head
	totCnt := nElves

	dist := totCnt / 2
	for i := 0; i < dist; i++ {
		counter = counter.next
	}

	for player.next != player {
		debugf("playing elf %d stealing from %d", player.id, counter.id)
		player.gifts += counter.gifts
		counter.prev.next = counter.next
		counter.next.prev = counter.prev
		counter = counter.next
		totCnt--
		dist--
		for dist < totCnt/2 {
			dist++
			counter = counter.next
		}
		player = player.next
	}
	return player.id
}

func main() {
	//nElves := 5
	nElves := 3005290

	printf("===== part 1 =====")
	res1 := solve1(nElves)
	printf("last elf id: %d", res1)

	printf("===== part 2 =====")
	res2 := solve2(nElves)
	printf("last elf id: %d", res2)
}
