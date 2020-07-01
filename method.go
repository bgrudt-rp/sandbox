package main

import "fmt"

type methodExample struct {
	id   int
	text string
}

func testMethod() {
	test1 := methodExample{1, "Test one"}
	test2 := methodExample{2, "Test two"}
	test3 := methodExample{3, "Test three"}

	fmt.Println(test1.text)
	test1.appendMethod("appended")
	fmt.Println(test1.text)

	fmt.Println(test2.text)
	test2 = test2.appendMethod2("appended")
	fmt.Println(test2.text)

	fmt.Println(test3.text)
	test3.badAppendMethod("appended")
	fmt.Println(test3.text)

}

//appendMethod uses a pointer to the var of type methodExample to manipulate that data in memory
func (m *methodExample) appendMethod(s string) {
	m.text = m.text + " " + s
}

//appendMethod2 requires the explicit reset of the input var by returning its new value
//and expecting the caller to set as needed
func (m methodExample) appendMethod2(s string) methodExample {
	m.text = m.text + " " + s
	return m
}

//badAppendMethod will not work, as it does not use a pointer
//and does not return the new value.  The changes made by this method
//are confined to the method itself
func (m methodExample) badAppendMethod(s string) {
	m.text = m.text + " this won't work " + s
}
