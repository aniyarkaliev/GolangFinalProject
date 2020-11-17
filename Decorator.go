package main

type BasicRoom interface {
	getDesignInfo() string
	makeAction(n int)
}

type withFirePlace struct {
	basicRoom BasicRoom
}

func (wF *withFirePlace) getDesignInfo() string{
	message := wF.basicRoom.getDesignInfo() + "with Fire place "
	return message
}

func (wF *withFirePlace) makeAction(n int) {
	wF.basicRoom.makeAction(n)
}

type withAquarium struct {
	basicRoom BasicRoom
}

func (wA *withAquarium) getDesignInfo() string{
	message := wA.basicRoom.getDesignInfo() + "with Aquarium "
	return message
}

func (wF *withAquarium) makeAction(n int) {
	wF.basicRoom.makeAction(n)
}

type withCoffeeTable struct {
	basicRoom BasicRoom
}

func (wC *withCoffeeTable) getDesignInfo() string{
	message := wC.basicRoom.getDesignInfo() + "with Coffee table "
	return message
}

func (wF *withCoffeeTable) makeAction(n int) {
	wF.basicRoom.makeAction(n)
}