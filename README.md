# A ToDo CLI App powered by CobraCLI
This is a 3-day project for Go Todo App running on the [spf13/cobracli](https://github.com/spf13/cobra). It allows you to keep all your tasks in check, and keeps you productive.

![Todo CLI Image](https://i.ibb.co/3m16hTP/gitpic.png)
*A screenshot of the TodoCLI in operation*

To get started, clone this repo and run the command in the root directory:
```
go mod tidy
go build
./cli-todo-go
```

That should get you started and display the main application menu.

---

### Commands
Typical usage is `./cli-todo-go [command]`
Commands include:
- list: to list all your tasks
- create: to create a task
- delete: delete a task
- complete: mark a task as complete

Those will allow you to manage your tasks.

Todo for the todo:
- [x] Build the core app
- [ ] migrate store from file to db
- [ ] add input validation
- [ ] add date formatting
- [ ] write go tests
- [ ] make the cli interactive

Keep your tasks nice and done, as they should be.

*CLI applications are great for building internal tools to streamline workflows, and for application SDKs as well.*
