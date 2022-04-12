package goalgo

/* Introduction to Algorithms

Jason Ku
MIT 6.006 Spring 2020

https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-006-introduction-to-algorithms-spring-2020/
https://www.youtube.com/playlist?list=PLUl4u3cNGP63EdVPNLG3ToM6LaEUuStEY

- Solve computational problems
- Communicate of ideas
- Prove correctness of results
- Argue efficiency of algorithms

What is a problem? What is an algorithm?

A problem is a binary relation between inputs and outputs

        Inputs           Outputs
			 x    ---------    x
			 x     \     /-    x
			 x       \ /       x
			 x    -- / \--     x
			 x                 x

Probabilities (like in 042) ask "Is there likely to be a result like this ..."

In this course, we want a discrete, general "answer"
- General problems
- arbitrarily sized inputs
- i want a machine or procedure that will generate 'an' output
- ... a correct output

efficiency - how well an algorithm performs ... as compared to other
methods ... **independent of the equipment and implementation**

"asymptotic analysis"



/// This is why the 'O' stuff is important ...

- Big O notation 	: corresponds to upper bounds
- Omega (Ω) 		: corresponds to lower bounds
- Theta (Θ)			: corresponds to both upper and lower bounds

O(1) 		constant
O(log n) 	logarithmic
O(n) 		linear
/// In some cases (other classes, big data sets) above this line is considered "efficient"
O(n log n) 	"log linear" or "nlogn"
O(n^2) 		quadratic
O(n^c) 		polynomial time (general case expanded from quadratic)
/// In 6.006, above this line is what we mean by "efficient"
2 ^O(n)  	exponential time

//* like in machine learning & AI

an Algorithm is ...
- gives a correct output for any given input in our problem domain

//* example was "interview each student to determine birthday"
- maintain an ordered list based on student interviews

//? why not have each person find their closest match?
//? or have them stand in certain areas based on their birth month?

*/
