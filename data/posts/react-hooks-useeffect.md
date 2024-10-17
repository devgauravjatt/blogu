---
title: React Hooks - UseEffect
coverImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1707310785327/bb232701-49c5-4723-aa2e-4eef72d9d2bd.jpeg
slug: react-hooks-useeffect
tags: [React, Beginner Developers, JavaScript, Web Development, ReactHooks]
description: React hook that enables you to perform side effects in functional components. Side effects are operations that affect
---

## Introduction

Hey guys, In our last article, we learnt about state(specifically useState) and also build a mini project. In this article, we’ll be talking about React hooks and talking more about the useEffect hook. Let’s dive right in 😉

### What is useEffect?

`useEffect` is a React hook that enables you to perform side effects in functional components. Side effects are operations that affect the external world, such as data fetching, subscriptions, or manually changing the DOM. Unlike lifecycle methods in class components, `useEffect` allows you to manage side effects in a declarative and concise manner.

### How does it work?

Basically, `useEffect` runs after the render is committed to the screen. It runs after every render by default.

**Basic Syntax**

Here is how the syntax of `useEffect` look like 👇🏽

```jsx
import { useEffect } from 'react';

const useEffectSyntax = () => {
  useEffect(() => {
    // Your side effect code here
    return () => {
      // Cleanup code (optional)
    };
  }, [/* dependencies */]);

  // Component render logic
}
export default useEffectSyntax;
```

* The first argument to `useEffect` is the function containing the side effect code. it is where the main logic of the code lives.

* The second argument is an array of dependencies. When any of these dependencies change, the effect is re-run.

* The return function (optional) is used for cleanup. It runs before the component unmounts or when the dependencies change.


### Examples

Let’s see some basic examples to understand `useEffect` better;

```jsx
import { useEffect } from 'react';

const SimpleEffectExample = () => {
  // useEffect to perform an action on component mount
  useEffect(() => {
    console.log('Component mounted!');

    // Cleanup function (optional)
    return () => {
      console.log('Component will unmount!');
    };
  }, []); // Empty dependency array means it runs only once on mount

  return (
    <div>
      <p>This is a simple component with useEffect.</p>
    </div>
  );
}

export default SimpleEffectExample;
```

From the example above, the `useEffect` is used to log messages to the console. The function inside it runs after the component is initially rendered. The dependency array `[]` is empty, meaning the effect runs only once when the component mounts. It simulates the behavior of `componentDidMount` in class components. Inside the `useEffect` function:

* It logs 'Component mounted!' to the console.

* It returns a cleanup function that logs 'Component will unmount!' to the console when the component is about to unmount.


Let’s see another example 👇🏽

```jsx
import { useState, useEffect } from 'react';

const SimpleEffectExample = () => {
  // State to hold a count value
  const [count, setCount] = useState(0);

  // useEffect to update the document title
  useEffect(() => {
    document.title = `Count: ${count}`;

    // Cleanup function (optional)
    return () => {
      document.title = 'React App'; // Reset the title on component unmount
    };
  }, [count]); // Re-run effect when count changes

  return (
    <div>
      <p>Current Count: {count}</p>
      <button onClick={() => setCount((prevCount) => prevCount + 1)}>
        Increase Count
      </button>
    </div>
  );
}

export default SimpleEffectExample;
```

From the example above, we have a state variable `count` initialized to `0` using `useState`, and a button to increase the count. The `useEffect` is used to update the document title with the current count. It takes a function as its first argument, and this function will run after the component renders. The dependency array `[count]` specifies that the effect should re-run whenever the `count` state changes. Inside the `useEffect` function:

* We set the document title to include the current count.

* We return a cleanup function that resets the document title to 'React App' when the component unmounts or when `count` changes.


The component renders a paragraph displaying the current count and a button. When the button is clicked, the `setCount` function is called, updating the `count` state and triggering the `useEffect` to update the document title.

Let’s see one more example 👇🏽

```jsx
import { useState, useEffect } from 'react';

const FetchData = () => {
  const [data, setData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch('<https://api.example.com/data>'); // just some dummy api endpoint
      const result = await response.json();
      setData(result);
    };

    fetchData();
  }, []); // Empty dependency array means it runs only once on mount

  return (
    <div>
      {data ? (
        <ul>
          {data.map(item => (
            <li key={item.id}>{item.name}</li>
          ))}
        </ul>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
}

```

This example is a bit more technical than the previous ones. but we’ll go through it together.

From the example above, we use the `useState` hook to create a state variable `data` and its corresponding updater function `setData`. The initial value of `data` is set to `null`.

Inside `useEffect`, an asynchronous function `fetchData` is declared. This function fetches data from a hypothetical API ([`https://api.example.com/data`](https://api.example.com/data)) using the `fetch` API.

The `fetchData` function is then invoked within the `useEffect`. Because the dependency array is empty (`[]`), this effect will run only once when the component mounts.

Once the data is fetched, it is stored in the `data` state using the `setData` function.

## Conclusion

That’s all guys, Hope you learned something from this article, Remember that useEffect basically allows you to perform side effects. It’s perfectly fine if you don’t understand it at first. Just try to go over each of the examples.

Have an amazing weekend and see you next week 😀
