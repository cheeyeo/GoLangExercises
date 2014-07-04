## Problems

* How are integers stored on a computer?

Integers are stored as sequence of bytes on a computer.

* We know that (in base 10) the largest 1 digit number is 9 and the largest 2 digit number is 99. Given that in binary the largest 2 digit number is 11 (3), the largest 3 digit number is 111 (7) and the largest 4 digit number is 1111 (15) what's the largest 8 digit number? (hint: 101-1 = 9 and 102-1 = 99)

* Although overpowered for the task you can use Go as a calculator. Write a program that computes 32132 × 42452 and prints it to the terminal. (Use the * operator for multiplication)

* What is a string? How do you find its length?

A string is a sequence of characters with a fixed length. To find the length of a string we can use the following function:

len(String)

* What's the value of the expression (true && false) || (false && true) || !(false && false)?

(true && false) => false
(false && true) => false
(false && false) => false

(false) || (false) || !(false) => true
