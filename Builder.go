package main

type User struct {
	//Personal Info
	Firstname, Lastname string
	ID int
	//Family Status
	Status string
}
type UserBuilder struct {
	user *User
}
func NewUserBuilder() *UserBuilder {
	return &UserBuilder{&User{}}
}
func (b *UserBuilder) Build() *User {
	return b.user
}
func (b *UserBuilder) PersonalInfo() *PersonInfoBuilder  {
	return &PersonInfoBuilder{*b}
}
func (b *UserBuilder) FamilyStatus() *FamilyStatusBuilder  {
	return &FamilyStatusBuilder{*b}
}

type FamilyStatusBuilder struct {
	UserBuilder
}

func (fsb *FamilyStatusBuilder) InFamily(Status string) *FamilyStatusBuilder {
	fsb.user.Status = Status
	return fsb
}

type PersonInfoBuilder struct {
	UserBuilder
}
func (pib *PersonInfoBuilder) withFirstName(Firstname string) *PersonInfoBuilder {
	pib.user.Firstname = Firstname
	return pib
}
func (pib *PersonInfoBuilder) withLastname(Lastname string) *PersonInfoBuilder {
	pib.user.Lastname = Lastname
	return pib
}
func (pib *PersonInfoBuilder) WithID(ID int) *PersonInfoBuilder {
	pib.user.ID = ID
	return pib
}

//func main()  {
//	pb := NewUserBuilder()
//	p := pb.PersonalInfo().withFirstName("Aza").withLastname("Saiduly").WithID(1).FamilyStatus().InFamily("Brother").Build()
//	fmt.Println(p)
//}