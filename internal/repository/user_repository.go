package repository //handles data storage

import "github.com/mithun/gin/internal/model" //repository depend on these model because it needs to store and retrive them 

type UserRepository struct {
	users []model.User
}

func NewUserRepository() *UserRepository {   /*func NewUserRepository(): By convention, Go constructors are prefixed with New.
*UserRepository: The return type is a pointer (*) to the struct. This is crucial for avoiding unnecessary data copying when passing the repository around.
&UserRepository{...}: The & operator takes the address of the newly created struct, creating a pointer.
users: []model.User{},: Inside the brace, we initialize fields. Here, users is initialized to an empty slice []model.User{} to ensure it is not nil, preventing potential nil pointer dereference errors later. */
	return &UserRepository{
		users: []model.User{},
	}
}

func (r *UserRepository) Create(name string) model.User {
	user := model.User{
		ID : len(r.users)+1,
		Name : name,
	}
	r.users = append(r.users,user)
	return user
}
/*Instantiates a model.User: A new user struct is created.
ID is assigned based on the current length of the repository's users slice plus one (len(r.users) + 1), essentially acting as a simple, non-concurrent-safe auto-incrementing ID.
Name is assigned the value passed into the function.
Appends the new user: The newly created user is added to the end of the r.users slice using the append built-in function. This modifies the repository's data store in memory.
Returns the user: The function finishes by returning the new model.User struct instance it just created. */
func (r *UserRepository) GetAll() []model.User {
	return r.users
}

