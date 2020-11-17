package main

import "fmt"

type subject interface {
	register(user User)
	notifyAll()
}

type Notifier struct {
	familyList []observer
	messages []string
}

func newNotifier() *Notifier {
	return &Notifier{
		messages: []string{},
		familyList: []observer{},
	}
}

func (n *Notifier) addNewMessage(message string){
	n.messages = append(n.messages, message)
	n.notifyAll(message)
}

func (n *Notifier) register(o observer) {
	n.familyList = append(n.familyList, o)
}

func (n *Notifier) notifyAll(message string) {
	for _, observer := range n.familyList {
		observer.update(message)
	}
}

func (n *Notifier) showAllMembers(){
	cnt := 1
	for _, observer := range n.familyList {
		fmt.Printf("%v)" + observer.getUserInfo(), cnt)
		fmt.Println()
		cnt++;
	}
	length := len(n.familyList)
	if length == 0 {
		fmt.Println("There are no family members!")
	}
}

func (n *Notifier) showMessageLog(){
	fmt.Println("All sended messages: ")
	for i, message := range n.messages {
		cnt := i+1
		fmt.Printf("%v) Message: "+message, cnt)
		fmt.Println()
	}
}

type observer interface {
	update(string)
	getUserInfo() string
}

func (u *User) getUserInfo() string{
	return ("Family member firstname: "+u.Firstname+", lastname: "+u.Lastname+", family status: "+u.Status)
}

func (u *User) update(message string) {
	fmt.Println("Hello Dear "+u.Firstname+" |(New Message!) "+message)
}
