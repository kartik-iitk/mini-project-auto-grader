### Introduction
More than a recruitment task, this assignment is made to help you better assess if this project is a good fit for you. The project would proceed similar to this task, i.e, you'll need to read through official documentation to complete small parts of the project. The role of mentors would be to 
* Discuss the architecture of the service.
* Point you towards specific documentation that you need to read.
* Doubt clearing

At first glance, it might look like Go was a prerequisite for the project, and we forgot to mention it in the presentation. **This is not true**. This task is designed for people with zero experience in Go.

> Remember to keep it simple. Each function should only take 4 to 5 lines. If it's more than that, you might be overcomplicating things.


### Task
* Install **Go**
* Install **Git**
* Install **Goland** IDE.
* Create account on **Github**
* Apply for **Student Developer Pack** (optional)
* Do the [official tutorial](https://go.dev/tour/) only till the **Methods and Interfaces** part. 

After this, you should be able to write Go code that, at the very least, compiles.

Next, download the repo, and open the project in Goland.

Before starting with this exercise, you should brush up on the **Linked List** concepts from **ESC101**. A linked list has the concept of "nodes" which are linked to other nodes. We can do operations on these nodes recursively (likely inserting elements at the back, traversal, etc). Similarly, An HTML document contains "tags" (such as `div`, `a`, `meta`, etc). Google about the document structure to learn more. These tags can then be represented as nodes, and instead of a linear linked list, we have a **tree**, where each node can contain some amount of children. The traversal of children from the same parent are done in the exact same way as Linked lists.

In this exercise, you need to determine certain properties of the HTML document, but from its tree structure.

* Open the official documentation of [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html). All the code that you need for this exercise is present in this documentation, either directly or via hyperlinks.
* Scroll down to the [index](https://pkg.go.dev/golang.org/x/net/html#pkg-index) section to get a high level overview. There are a couple of functions and `type` declaration. To read a short description about any of these functions/types/receiver-functions, click on the name. To go back to index, click on `index` from the left sidebar.
* As you can see from the worksheet, we are extensively using `html.Node` struct. Navigate to this struct, and inspect its definition, summary, and example-usage. If you encounter new keywords during inspection of `Node` (for example, `io.Reader`, `NodeType`, etc), just click on the hyperlinks, which would take you to their official documentation. Make sure to read all linked examples carefully.
* Now, you have a good understanding of what an HTML node is, what data does it contain, and what receiver functions can be invoked on it. This should be sufficient to complete the exercise.

More keywords that you might need to google for.

* Struct
* Interface
* make
* range
* append
* map
* slice
* Variadic functions
* Conditionals (if, for)
* Buffers
* Receiver functions
* Depth first search (DFS)

You need to edit `task.go` only. After you make any small changes, open `task_test.go`, and click on the `Run` icon to the left. You should see the results of the test at the bottom of the screen.

Once all tests are green, you're good to go.

### Submission
Install `git` and commit your changes via

```bash
git add .
git commit -m "Your message goes here"
git push
```

Then, open the github repository, and go to `Actions` tab. If you see a green tick, it means that you've successfully completed the assignment. If you see a red cross, your tests failed. Expand it to view more details, and make the necessary changes.

### Help
First, **Google, Google, Google**. All parts of the puzzle are present on the internet (in fact, on the documentation page itself). You just need to join those pieces.

With that said, don't be blocked on logistical issues, syntax or compilation errors for too long. If you are, you can ask your doubts on the `cf-stress` channel on `pclub` discord server.

Depending upon how many submissions are made, we might even provide more hints on the channel.

P.S: If you aren't selected officially, or are unable to complete both the tasks for some reason, and would like to do the project anyway, drop an email to `variety.jones.silk@gmail.com`, mentioning the reason, as well as your motivation for doing the project. We are always looking for enthusiastic folks to work with. (However, you should still satisfy the prerequisites). Note that this is strictly on an individual capacity, and pClub is not associated with it. So even if you complete the project unofficially, there would be no ratification from pclub.