# coursera-stanford-algorithms-divide-conquer
This repository contains my solutions to: [Divide and Conquer, Sorting and Searching, and Randomized Algorithms](https://www.coursera.org/specializations/algorithms), the first course in Coursera's: [Algorithms Specialization](https://www.coursera.org/specializations/algorithms) which seeks to help students "Learn To Think Like A Computer Scientist. Master the fundamentals of the design and analysis of algorithms."

# Table Of Contents
- [Language Choice: Go](#language-choice-go)
- [Weekly Work](#weekly-work)
  * [Week 1](#week-1)
    + [First Attempt: Integers](#first-attempt-integers)
      - [Aside: Working with Types in Go (as a JS developer)](#aside-working-with-types-in-go-as-a-js-developer)
    + [Second Attempt: Numbers as Strings](#second-attempt-numbers-as-strings)
    + [Testing](#testing)
      - [External Tests](#external-tests)
    + [Comparing My Results](#comparing-my-results)
  * [Week 2](#week-2)
    + [Analysis and Prior Work](#analysis-and-prior-work)
    + [Go Implementation Observations](#go-implementation-observations)
      - [A minor bug](#a-minor-bug)
      - [Ain't no `while` in golang](#aint-no-while-in-golang)
  * [Week 3](#week-3)
    + [Analysis and Solution](#analysis-and-solution)
      - [Setup](#setup)
      - [Partition](#partition)
      - [Slice Helpers](#slice-helpers)
      - [QuickSort](#quicksort)
      - [Choose Median Of Three](#choose-median-of-three)
        * [A Correct Implementation](#a-correct-implementation)
        * [An Efficient Implementation](#an-efficient-implementation)
        * [A Reliable Implementation?](#a-reliable-implementation)

# Language Choice: Go
I've chosen Go as the language to provide implementations for the programming assignments.  I'm relatively new to Golang but I'm finding it to be more important as I work with Kubernetes in my day job at GoDaddy.  I'm viewing this course as an opportunity to familiarize myself with the standard set of algorithms that every comp-sci major should know, and to also become familiar with Golang itself.  As such I will attempt to leverage Go's native package system, test system, and other basic features of the language and it's environment that are fundamental to use.

Go presents an interesting set of challenges that are not found when using the more common languages Java and Python for implementing the code challenges in the course.  I will discuss my experience of working through the course's coding elements in this README and I will focus when possible on the learnings I discover in Go as well as the differences I see between the other languages I've worked with in the past.

# Weekly Work

## Week 1
Week one's challenge asks students to [implement the Karatsubsa Multiplication algorithm for very large numbers](/src/week1/docs/Assignment1.png). 

### First Attempt: Integers
I began by attacking the problem using integers.  The actual Karatsuba algorithm is relatively straightforward and simple to implement.  The biggest challenge came when splitting an input number which had an odd number of digits.  It was easy enough to chop bigger numbers but when I got down to `len(x) == 1 && len(y) == 2` things went sour.  I was unsure how to approach the recursion at this point.  In hindsight it occurs to me that I could have just "zero padded" `x` by handling this edge case with specific code.  But as I explored the implementation in the cases that it could solve (even digits) I found that it was unable to deal with large integers.  It was clear that my implementation depended on the `float64` and `int` data types in Go.  Even when I explicitly called for an `int64` I didn't have enough significance to hold the large numbers that were under operation.  Clearly I needed another approach and there were hints in the description of the Karatsuba Multiplication algorithm itself.  

#### Aside: Working with Types in Go (as a JS developer) 
A quick aside on type casting in golang.  I've spent the last couple of years primarily using Javascript as my language of choice.  One of the things I love about Javascript development is that it facilitates writing code quickly.  I think of it as a shoot-from-the-hip language which allows you to play loose and fast with types.  Javascript doesn't care about the type of your variable at declaration time and will do its best (and sometimes its worst) at converting between types implicitly.  This is really nice if you trust your input and want to move fast.  Python offers similar flexibility, but it's a bit tighter with its rules.  Go, however, is more like the traditional languages.  Types matter.  This lead me down a path of cast-cast-cast calls, especially when working with Go's standard `math` library.  It wasn't uncommon for me to write something like the following:
```
func getPlaces(n int) int {
  if n == 0 {
    return 1
  }
  return int(math.Floor(math.Log10(float64(n))) + 1)
}
```
That return statement would be much simpler in a language that did more implicit type casting.  Of course, everything in computer science is a tradeoff and code complexity is traded for performance and safety.

### Second Attempt: Numbers as Strings
The description of the Karatsuba Algorithm includes the key to unlocking this solution.  It tells us that the number of digits in each of the numbers we work with should be a power of 2 and that we'll need grade school multiplication, addition, and subtraction.

In the case of multiplication we know that we'll only operate on terms of length == 1.  So we don't have to worry about integer overflow in that case.  But for our addition and subtraction the same cannot be said.  You could imagine adding 1 to a the max value of an int64 (9223372036854775807).  In that case our result would overflow and our implementation will fail.  So how can we solve both the "length is a power of 2" issue and the "adding large numbers" issue?  We simply treat all numbers as strings and perform "grade school addition/subtraction" on those large numbers when we encounter them.

So when we have 1234 + 5678 we just work our way through the terms from right to left "carrying the one" whenever we need to.  It's a relatively straightforward concept that we can all do in our heads.  Of course, writing out a specific implementation, especially of subtraction, is often more complicated than just crunching numbers in our heads.  

The final bit of complexity came at the end steps of the Karatsuba Multiplication when numbers need to be multiplied by 10 to a potentially very large power.  Initially I was using multiplication here as well but this fails with large numbers.  Fortunately it's easy to see that right padding `x` with `n` 0's is the same as `x * 10^n`.  Once that was in place the implementation was complete and the test case that was presented in the assignment was passed.

### Testing
I used Go's test facility to write some initial tests of my methods.  This was incredibly helpful for debugging some of the edge cases in my implementation.  As is the case in unit testing it made it easy for me to see where my Karatsuba implementation failed, for example in subtraction when subtracting a positive minus a negative.  Go's test system is interesting in that it doesn't provide a native means for passing test cases to a test method.  I solved this by creating an array of cases inside the test and then iterating over it and collecting failures.  This presents challenges as there is no `tuple` datatype in Go and we are forced to use a `struct` to represent a composite collection of different data types.  Fortunately Go provides the ability to create anonymous structs so they can be declared inline and one off.  Here's an example of one of the tests: 

```
func TestStrSub(t *testing.T) {
  // x, y, expected value
  cases := []struct {
    x string
    y string
    z string
  }{
    {"-10", "-20", "10"},
    {"-20", "-10", "-10"},
    {"10", "20", "-10"},
    {"20", "10", "10"},
    {"-20", "10", "-30"},
    {"-10", "20", "-30"},
    {"20", "-10", "30"},
    {"10", "-20", "30"},
  }

  for i := 0; i < len(cases); i++ {
    testCase := cases[i]
    actual := strmath.StrSub(testCase.x, testCase.y)
    if testCase.z != actual {
      t.Logf("StrSub failed to Subtract %s - %s = %s, expected %s", testCase.x, testCase.y, actual, testCase.z)
      t.Fail()
    }
  }
}
```

#### External Tests
In the forums for the course I found a link to this [repository of test cases for the algorithms](https://github.com/beaunus/stanford-algs).  This seemed like the best way to prove to myself that my implementation was complete and correct.  In order to work with the bash based test runner for these tests I needed to modify my `main` function so that it took filenames as input and then read the contents of those files as input to my algorithm.  Finally my driver would need to return a single line output of result.  I provided this implementation and then ran the script and all of the test passed.

### Comparing My Results
Now that I knew that my implementation was complete and correct I was curious what I could have done better.  This lead me to search the course's forums and github to see how other people solved this problem.  That lead to an interesting discovery.

Students who used Java and Python were at a significant advantage over those using a language like Go.  Java has a `BigInt` class which can hold ints > 64 bit.  Python allows integers of virtually unlimited length.  When I found code in Java and Python that implemented Karatsubsa Multiplication it was devoid of what I found to be the most challenging and interesting part of this exercise.  Those implementations were largely just a direct translation of the provided psuedo code into the appropriate language.  They were quick and clean and clear, but lacked the depth of challenge that was presented by the limitations of Go.  In Go I had no choice but to implement Addition and Subtraction on my own, and I had work with int-strings.  Without that part of the exercise there isn't much of a challenge.  One might as well just use in built math functions.

Interestingly, when we offer algorithmic questions during interviews at GoDaddy we will often create multi-tiered challenges.  In the case of Karatsuba Multiplication I could see asking a candidate who had just successfully solved this problem in Java or Python to go one level deeper.  I might ask them to take a second pass at their solution but to avoid the use of a `BigInt` in Java or to consider an explicit use of int-strings in python.

## Week 2
Week two's challenge asks students to [count the number of inversions in a very large list of integers](/src/week2/docs/Assignment2.png).

### Analysis and Prior Work
At the core this question is asking us to implement merge sort.  This will provide us with a fast running `O(n log n)` algorithm as long as the work that we do to count the inversions can be done in constant time.  I found this challenge to be far simpler than the prior week's work for two main reasons.  

First, the lectures spend a great deal of time working with merge sort.  The discussion of the algorithm, analysis of it, and finally the direct discussion of how to "piggy-back" on merge sort to count inversions.  The trick, being able to count the inversions in constant time, is relatively straightforward if you understand the implication of the merge on the pre-sorted left/right halves of the array.  I've added comments to the source files that I wrote to discuss this, but in essence the fact that both lists are sorted implies that when an inversion is found, all remaining elements in the left array are also inverted against the current element of the right array.  This allows for simple arithmetic to determine the count of inversions.

Second, I personally had an advantage because I've used this exact question to interview candidates at GoDaddy.  In fact, I have implemented a solution to this problem in Python for those interviews.  The solutions in these two languages are actually syntactically similar insofar as Python and Go have common expressions.  You're using similar `append` functions and array slicing syntax to do the same work.

Since the material was familiar to me, I made it a point to work through the solution without looking at the lecture notes or my past work.  Instead I relied on my understanding of the algorithm to work through exactly what it should do.  In the end I solved the problem quickly and only faced one minor bug.

### Go Implementation Observations
#### A minor bug
It's worth mentioning that when I wrote the initial implementation I failed to use an `if/else if` construct and instead used a pair of `if` statements in the loop in the main body of the `merge` routine.  This resulted in an `index out of bounds` error which was a bit confusing when I encountered it.  The index out of bounds occurred when the first condition was matched and the index `i` was incremented.  This would allow the second condition to evaluate with `i` being out of range.  It was silly and simple to fix once I saw it.

#### Ain't no `while` in golang
One interesting syntactic discovery while working on the `merge` routine was that Go lacks a `while` loop.  The documentation humorously states that "in Go, `while` is spelled `for`."  This is a factual statement.  Go uses a `for` construct which only includes the exit clause in place of a `while` statement.  When the initializer condition and increment statement are removed from the `for` declaration, all that remains is an endless loop which exits if it's boolean bounding statement resolves to false.

I'm no language designer but I have to think that having a `while` is a bit clearer syntactically, even if it is just a synonym for a bounded, non-initialized, non-incrementing `for` loop.  Regardless, the compiler spit out a reasonable message and a quick google search lead to Go's excellent documentation which cleared up my mistake.


## Week 3
Week three's challenge asks students to [implement quicksort using 3 different pivot routines and measure comparisons](/src/week3/docs/Assignment3.png).

### Analysis and Solution
#### Setup
I began by writng stubs and tests to cover the expected functionality of parts of the system.  The problem asks for two of the `ChoosePivot` routines to be incredibly simple; one returns the first element to pivot around and one returns the last element.  So I implemented and tested them immediately.  

I wrote a prototype for the `QuickSort` method and decided that I would pass the `ChoosePivot` function as a parameter.  Since I'm new to Go I had to do a little research to discover how to actually declare and implement this.   The syntax, though new to me, is straightforward enough.  One simply declares a new `type` and specifies a function prototype for that type: 
```
type fnChoosePivot func(k []int) int
```
Once that's done the new type can be specified as the type of a parameter on the target function: 
```
func QuickSort(choosePivot fnChoosePivot, k []int) ([]int, int) {
  ...
}
```
I thought a bit about what to actually test on the `QuickSort` routine.  The problem asks about counting the number of comparisons performed as a function of the `ChoosePivot` method that we use.  This implies that we'll need different test cases for each of the `ChoosePivot` routines, and that we'll need to calculate those beforehand.  There was enough complexity here for me to defer the implementation of the tests for `QuickSort` until later.

#### Partition
I took a stab at the `Partition` function as it is both central to the operation of `QuickSort` and simple to implement.  Interestingly, the nature of Go `slices` makes the routine even easier than it might be in another language.

I think of Go `slices` as little windows that look in on an underlying `array`.  They can be modified without consequence to the underlying array, and operating on their elements modifies the elements of that array.  This works really well in a problem like `QuickSort` where we know we want to look at segments of the array and make in-place changes.

The general implementation of `Partition` requires that we pass a `left-index` and `right-index` so that we can create a window over the elements that we actually want to work with.  By leveraging `slice`s we can drop those indicies and assume that we will partition the entire `slice` that is received as input.  

I initially figured that this would create additional memory consumption but this is not exactly the case.  Go `slice`s are passed by value and so each call to `Partition`, regardless of the specific `slice` we pass, will make a copy of that `slice`.  Effecitvely, we get something for nothing.  We get increased code simplicity and the same memory performance.

An alternative implementation might pass a pointer to the `slice` which would pervent the copying, or possibly one could pass the array itself.  The former would likely work, and be relatively straightforward albeit it would add a bunch of poitner dereferencing.  The latter though would present challenges as the size of an `array` in `Go` is a fundamental part of the `type` of that `array`.  Passing it as a parameter would require a fixed length for that `array`.  It is generally better to use `slice`s for just such a reason, and so that's what I chose to do.

#### Slice Helpers
Go `slice`s are pretty awesome but it was clear from the start that I was going to need some helper utilities in order to do some relatively simple tasks.  I broke off a `lib` package and collected a set of routines that would help me work with slices.  

Copying a `slice` in Go creates a new `slice` which points to the same `array`.  I needed to be able to clone a `slice` and produce a new copy of the `array` so I wrote a `CloneSlice` routine to do this.  

Equality is undefined for `slice`s so I had to implement an element by element comparison.

To prove that `Partition` worked correctly in the tests I needed to verify that the resulting arrays were all either greater or less than the appropriate paritioned results.  I added slice helpers for this.

Finally, swapping elements of a slice is a fundamental need for `Partition` so I wrote an implementation of that to keep the code as DRY as possible.

#### QuickSort
The implmentation of the actual `QuickSort` is incredibly simple.  This is one of the most impressive parts of the algorithm.  It's all of 8 lines of code.  With that said, once implemented, I faced a tough challenge on how to test this code.

I don't have a lot of experience testing recursive functions.  I suppose I could produce mock implementations of both `Partition` and `choosePivot` and design test cases where the expected output is generated.  But it seemed to me that this would be wasteful.  The core functionality of the `QuickSort` really exists in those two methods so appropriately testing them should work as expected.

We likely could test to make sure that `choosePivot` is called with the appropriate paramters, likewise for `Partition`, but things got a bit fuzzy when I started thinking about evaluating what gets passed to the recursive calls to `QuickSort`.  I'm unsure how to properly test a recursive function without significantly alerting the prototype or the environment.  I would love to discuss this further.  

For now, I removed the unit test on `QuickSort` and simply relied on the full tests to prove that `QuickSort` is behaving correctly.

#### Choose Median Of Three
The final piece of the puzzle was to produce an adequate implementation of the third `ChoosePivot` routine: `ChooseMedianOfThree`.  This case is interesting to me because it highlights a fundamental issue that I have with these kinds of problems (when examined in the real world and used as interview questions).

##### A Correct Implementation
It it relatively easy to implement correct solution to the median of three numbers.  If we define `a`, `b`, and `c` as the three numbers to compare we can simply take `a` and compare it to see if it is between `b` and `c`, if it is, return `a`. Next, do the same for `b` between `a` and `c`, and return `b`.  Finally, otherwise, return `c`.  [Code for this can be found here.](https://github.com/damiansilbergleithcunniff/coursera-stanford-algorithms-divide-conquer/blob/1a33fa186397a567e4bf1c6029332f6e1f382ab4/src/week3/algorithms/quicksort.go#L65-L84)

This solution is correct in that it will always return the median of the three numbers.  However, in its worst case it requires 8 comparisons to determine the result.  In its best case it still requires 4.  So as Dr. Roughgarden has taught us we must ask the question, "Can we do better?"

##### An Efficient Implementation
As it turns out it is possible to write a more efficient version of the Median of Three problem.  It's possible to determine the median of 3 with only 3 comparisons.  [This can be seen here.](https://github.com/damiansilbergleithcunniff/coursera-stanford-algorithms-divide-conquer/blob/a215c230e350f5607de5f6f61b14157317f97758/src/week3/algorithms/quicksort.go#L65-L97)

I found that, in order to develop this solution I needed to write out the permutations for the possible outcomes and ensure that the code covered those cases.  Likewise I needed to produce tests that would exercise them.

##### A Reliable Implementation?
In the real world we need to strive for our code to be both efficient and maintainable.  We want to make sure that developers, especially more junior developers, are able to work with the code that we produce.  Often these two potentials are at odds.  I made it a point to comment both of these code blocks, but if I hadn't it would be far easier to follow the first implementation than the second.  So what then is the right choice?

I would argue that the difference in performance here falls into the scope of the constant terms that we ignore when we're doing Big Oh notation.  As such, we should care more about the maintainablility of what we write.  Obviously on tiny embedded systems, or super performant code, we might care.  But in the general cases that we really deal with in most of our work we would prefer to be clear over efficient (at this scale).  We see this in the [Zen of Python](https://www.python.org/dev/peps/pep-0020/) and this is no doubt a contributor to that language's widespread success.

It is, perhaps, overly analytical to examine this case and debate this point, but I could easily see a candidate getting a "ding" for using the first solution and not the second.  It's important to understand the code that you're writing, what you're writing it for, and to target solutions to both the problem and the audience.- [coursera-stanford-algorithms-divide-conquer](#coursera-stanford-algorithms-divide-conquer)

