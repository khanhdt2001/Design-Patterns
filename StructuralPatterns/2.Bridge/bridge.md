**Bridge** is a structural design pattern that lets you split a large class or a set of closely 
related classes into two separate hierarchies—abstraction and implementation—which can be developed 
independently of each other.

## Problem
Say you have a geometric `Shape` class with a pair of subclasses: `Circle` and `Square`. You want to extend this class
hierarchy to incorporate colors, so you plan to create Red and Blue shape subclasses. However, since you already have 
two subclasses, you’ll need to create four class combinations such as `BlueCircle` and `RedSquare`.
</br>
Adding new shape types and colors to the hierarchy will grow it exponentially. For example, to add a triangle shape 
you’d need to introduce two subclasses, one for each color. And after that, adding a new color would require creating 
three subclasses, one for each shape type. The further we go, the worse it becomes.

## Solution
The Bridge pattern attempts to solve this problem by switching from inheritance to the object composition.
**_What this means is that you extract one of the dimensions into a separate class hierarchy, so that the original
classes will reference an object of the new hierarchy, instead of having all of its state and behaviors within one class._**

- Abstraction (also called interface) is a high-level control layer for some entity. This layer isn’t supposed to do 
any real work on its own. It should delegate the work to the implementation layer (also called platform).</br>

Note that we’re not talking about interfaces or abstract classes from your programming language. These aren’t the same things.
When talking about real applications, the abstraction can be represented by a graphical user interface (GUI), and the
implementation could be the underlying operating system code (API) which the GUI layer calls in response to user interactions.
