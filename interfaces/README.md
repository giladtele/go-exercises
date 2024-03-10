# Interfaces Drills

1. __Basic Interface Implementation:__
* Create a simple interface called Shape with a method Area() that returns a float64.
* Implement this interface in two structs: Circle and Rectangle.
* Write a function that takes in a Shape and returns its area.

2. __Interface Composition:__
* Define two new interfaces, Solid with a method Volume() and Material with a method Material() returning a string.
* Create a new struct Cube that implements both Shape and Solid interfaces and has an attribute defining its material.
* Implement the Material interface in the Cube struct.

3. __Empty Interface Usage:__
* Write a function that takes an empty interface as an argument and demonstrates type assertion to identify different types including int, string, and your Shape interface.

4. __Type Switches:__
* Write a function that uses a type switch to handle different types: int, string, and a custom type implementing the Shape interface.

5. __Interface Satisfaction:__
* Create a function that checks if an object satisfies the Shape interface without explicitly implementing it.

6. __Advanced Interface Concepts:__
* Explore the concept of embedding interfaces. Create a new interface that embeds the Shape interface along with a new method. Implement this new interface in a struct.

7. __Practical Application:__
* Create a program that uses interfaces to handle a variety of different geometric shapes. The program should be able to calculate the area, and if applicable, the volume of a set of shapes provided by the user.