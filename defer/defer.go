package _defer

import "fmt"

// 原则
// 1.defer 是 Go 语言提供的一种用于注册延迟调用的机制
//每一次 defer 都会把函数压入栈中，当前函数返回前再把延迟函数取出并执行。
// defer 定义的函数会先进入一个栈，函数 return 前，会按先进后出（FILO）的顺序执行。
//也就是说最先被定义的 defer 语句最后执行。
// 2.defer 语句定义时，对外部变量的引用是有两种方式的，分别是作为函数参数和作为闭包引用。
//作为函数参数，则在 defer定义时就把值传递给 defer，并被缓存起来；
//作为闭包引用的话，则会在 defer 函数真正调用时根据整个上下文确定当前的值。

// 1. 先执行赋值 r=0
// 2. 再执行 defer r++ 此时r=1
// 3. 再执行return r 此时r=1,所以return 1
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

// 1. 先执行赋值 r=t=5
// 2. 再执行 defer t++ 此时t=10 r不变还是5
// 3. 再执行return r 此时r=5,所以return 5
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// 0. 由于是函数调用，先将r赋值0，并缓存起来，此时r=0
// 1. 执行赋值 r=1
// 2. 再执行 defer r+5 此时r=0+5 但是函数调用是局部变量，外层不变
// 3. 再执行return r 此时r=1,所以return 1
func f3() (r int) {
	defer func(r int) { // 作为函数参数传入 defer 函数
		r = r + 5
	}(r)
	return 1
}
func f4() int {
	var i int
	defer func() {
		i++
	}()
	return i
}
func Handle1() {
	//res := f4()
	//fmt.Println(res)
	f5()
}

type Person struct {
	age int
}

// 1. person.age 这一行代码跟之前含义是一样的，此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，
// 等到最后执行该 defer 语句的时候取出，即输出 28；
// 2. defer 缓存的是结构体 Person{28} 的地址。所以是29。
// 如果这个地址指向的结构体没有被改变，最后 defer 语句后面的函数执行的时候取出仍是 28；
// 3. 闭包引用，person 的值已经被改变，指向结构体 Person{29}，所以输出 29.
// 输出结果是29 29 28
func f5() {
	person := &Person{28}
	fmt.Printf("The address of person is: %p\n", person)

	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
	fmt.Printf("The address of person is: %p\n", person)

}

// 参考答案及解析：29 28 28。
// 这道题在f5题目的基础上做了一点点小改动，前一题最后一行代码 person.age = 29 是修改引用对象的成员 age，
// 这题最后一行代码 person = &Person{29} 是修改引用对象本身，来看看有什么区别。
// 1处.person.age 这一行代码跟之前含义是一样的，此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，
// 等到最后执行该 defer 语句的时候取出，即输出 28；
// 2处.defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变，
// 最后 defer 语句后面的函数执行的时候取出仍是 28；
// 3处.闭包引用，person 的值已经被改变，指向结构体 Person{29}，所以输出 29.
// 由于 defer 的执行顺序为先进后出，即 3 2 1，所以输出 29 28 28。
// 核心的变化是最后把地址person指向的地址改变了，但是第二个defer中缓存的地址是旧地址
func f6() {
	person := &Person{28}
	fmt.Printf("The address of person is: %p\n", person)
	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
	fmt.Printf("The address of person is: %p\n", person)
}
