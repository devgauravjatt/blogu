---
title: React Hooks - UseState
coverImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1707311477343/86f84f2f-b6bf-4185-a706-89fe3d6c6723.jpeg
slug: react-hooks-usestate
tags: [React, Beginner Developers, Tutorial, articles, Web Development]
description: State is vital for building interactive applications. It enables your components to respond to user interactions, data fetching
---

## Introduction

Hey guys, In our last article, we learned about props and did a mini-project. In this article, we’ll be talking about the concept of state as well as building a little project. Let’s dive right in 😉

## What is state?

In React, **state** basically refers to the data that a component manages which can change over time.

Imagine you're building a weather app. The temperature at a location is a perfect example of what state represents. It can change throughout the day, and your app needs to keep track of that change and update the display accordingly.

State is essential for creating dynamic user interfaces because it allows you to store and update information that influences how your components render and behaves.

## Importance of state

State is vital for building interactive applications. It enables your components to respond to user interactions, data fetching, and other events. When the state of a component changes, React automatically re-renders the component, updating the user interface to reflect the new state.

This dynamic behavior is what makes React so powerful for building modern web applications.

## **Using** `useState` Hook

Like I have said in this series, functional components are the modern way of building UI in React. To manage state in functional components, you can use the `useState` hook. The `useState` hook allows you to add state to your functional components easily.

💡 *Hooks are special functions that are only available while React is rendering. They let you “hook into” different React features. - React Docs*

*There are other Hooks apart from* `useState` but we’ll use `useState` in today’s project.

## **How to Use** `useState` in React ?

I am a project based learner and I also hope you are too. So in order to learn how `useState` works, we’ll be building this simple counter project together. Let’s get started 👇🏽

**Step 1: Import React and** `useState` hook

```jsx
import React, { useState } from 'react';
```

In this line, we import React and the `useState` hook from the React library. These are necessary for creating functional components and managing state.

*You can also decide to remove the React and just import only the* `useState` which will still work just fine.

**Step 2: Initialize State**

```jsx
const [count, setCount] = useState(0);
```

Here, we use the `useState` hook to initialize a state variable called `count` with an initial value of `0`. This is where we'll store and manage our count. and `setCount` is used to update `count` by providing a new value.

**Step 3: Create Increment and Decrement Functions**

```jsx
function increment() {
  setCount((prevCount) => prevCount + 1);
}

function decrement() {
  setCount((prevCount) => prevCount - 1);
}
```

We define two functions, `increment` and `decrement`, which are called when the corresponding buttons are clicked(which we’ll be creating shortly). Inside these functions, we use the functional form of `setCount`. This form takes a function as its argument, which receives the previous state value (`prevCount`) and returns the new state value. By using this approach, we ensure that we're updating the state based on the previous state correctly.

**Step 4: Render the Component**

```jsx
return (
  <div>
    <h1>Counter</h1>
    <p>Count: {count}</p>
    <button onClick={increment}>Increment</button>
    <button onClick={decrement}>Decrement</button>
  </div>
);
```

In the `return` statement, we define the component's JSX. It displays the current count (`{count}`), along with two buttons for incrementing and decrementing the count. When these buttons are clicked, the `increment` and `decrement` functions are called, updating the state and causing the component to re-render with the new count value.

Here is the code in full 👇🏽

```jsx
import React, { useState } from 'react';

function Counter() {
  // Initialize state with a count of 0
  const [count, setCount] = useState(0);

  // Function to increment the count
  function increment() {
    setCount((prevCount) => prevCount + 1);
  }

  // Function to decrement the count
  function decrement() {
    setCount((prevCount) => prevCount - 1);
  }

  return (
    <div>
      <h1>Counter</h1>
      <p>Count: {count}</p>
      <button onClick={increment}>Increment</button>
      <button onClick={decrement}>Decrement</button>
    </div>
  );
}

export default Counter;
```

## Conclusion

That’s all guys, Hope you learned something from this article, Remember that State allows you to manage dynamic data within your components, making them responsive to user interactions. As you persist in learning and applying these concepts, they will gradually become more intuitive.

Have an amazing weekend and see you next week 😀
