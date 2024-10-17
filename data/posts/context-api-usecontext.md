---
title: Context Api - UseContext()
coverImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1708597804213/98153651-5f1f-45bd-a323-b94bcfd45dab.jpeg
slug: context-api-usecontext
tags: [Web Development, React, Beginner Developers, Programming Blogs, webdev]
description: Context in React is a mechanism to share values, users data, shopping cart or any global state, between components without the need to pass props through every level of the component tree. It helps in avoiding prop drilling
---

## Introduction

Hey guys, In our previous article in the React series, we learned about how useEffect works and we also talked about some of it’s use cases. In this article, we’ll be talking about the React Context Api. Let’s dive right in 😉

## What is Context ?

Context in React is a mechanism to share values, users data, shopping cart or any global state, between components without the need to pass props through every level of the component tree. It helps in avoiding prop drilling, where you pass data through several layers of components like passing props from a parent component such as `App` to a child component such `NavBar`.

## **The Basics of useContext**

The `useContext` hook is part of the React Hooks API, and its purpose is to access the value of a React context. Let’s see an example for better understanding 👇.

```javascript
// DemoContext.jsx
import { createContext } from 'react';

const DemoContext = createContext();

export const DemoProvider = ({ children }) => {
  const message = 'Testing this out to see if it works!';

  return (
    <DemoContext.Provider value={{ message }}>{children}</DemoContext.Provider>
  );
};

export default DemoContext;
```

In this example, `DemoProvider` sets up the context and provides the value “**Testing this out to see if it works!”** to its children using `DemoContext.Provider`.

In order for you to have access to the `message` in your `App` component, you have to wrap your `App` with the `DemoProvider` like this 👇🏽

```javascript
// App.jsx
import { DemoProvider } from './DemoContext';
import Home from './components/Home';

const App = () => {
  return (
    <>
      <DemoProvider>
        <Home />
      </DemoProvider>
    </>
  );
};

export default App;
```

You now have access to the `message` globally in your `App` what you need to do now is to use it where ever you want in your application. We’ll be using it on our home component 👇🏽

```javascript
// Home.jsx
import { useContext } from 'react';
import DemoContext from '../DemoContext';

const Home = () => {
  const { message } = useContext(DemoContext);
  return <div className="home">{message}</div>;
};

export default Home;
```

From the code above we imported `useContext` and `DemoContext` . We also use the `DemoContext` by passing it to the `useContext` and extracting the `message` from it and lastly, passing it to our UI.

This is just the basics of how context works with `useContext`.

Let’s see some of the use cases of when to use the `useContext` before we take a real world example 👇🏽

## **When to Use** `useContext`

Understanding when to use `useContext` is crucial for effective React development. Here are some scenarios where `useContext` shines 👇🏽

### **1\. Avoiding Prop Drilling**

`useContext` is especially useful when dealing with deeply nested components where passing props becomes impractical. Instead of passing data through each intermediate component, you can directly consume the context where it's needed.

### **2\. Managing Global State**

When you need to share state or functionality across multiple components, `useContext` can be an efficient way to create a centralized place for that data or behavior. This is particularly relevant for managing global state in your application.

### **3\. Cleaner Code**

`useContext` simplifies the code by making the codebase more readable and reducing the verbosity associated with context consumption.

### **4\. Dynamic Values**

When the context value is expected to change over time, such as when managing dynamic global states or themes, `useContext` provides an elegant solution for components to stay in sync with these changes.

Let’s now see an example of where it can be used in a real world application 👇🏽

### **The Shopping Cart Analogy for** `useContext`

Imagine you are building an e-commerce website with various components responsible for different parts of the user interface. One critical aspect of the website is the shopping cart, which keeps track of the items users want to purchase.

### Context as the Shopping Cart

In the context of our analogy, the shopping cart represents the shared data or state that needs to be accessed by multiple components. The Context API in React serves as this virtual shopping cart. It provides a way for components to share information without manually passing it through each one or passing it from the parent component down to where it is needed.

### The `useContext` Hook as the Shopper

Now, let's introduce the `useContext` hook as the shopper who wants to interact with the shopping cart. Instead of carrying the cart around and passing it to every store, the shopper uses the `useContext` hook to access the cart whenever needed.

### Setting up the Context (Creating the Cart)

Before shoppers can use the `useContext` hook, you need to set up the context, which is similar to creating the shopping cart. In React, you create a context using `createContext()`. 👇🏽

```javascript
import { createContext } from 'react';

const ShoppingCartContext = createContext();
```

This creates a new shopping cart, ready to be filled with items.

### Providing the Cart (Creating the Provider Component)

To fill the shopping cart with items, you need a provider component. This is where the `DemoProvider` component from the previous examples comes in. It wraps the part of your application that needs access to the shopping cart just like we talked about in previous examples. 👇🏽

```javascript

const ShoppingCartProvider = ({ children }) => {
  const [cart, setCart] = useState([]);

  const addToCart = (item) => {
    setCart([...cart, item]);
  };

  return (
    <ShoppingCartContext.Provider value={{ cart, addToCart }}>
      {children}
    </ShoppingCartContext.Provider>
  );
};

```

This provider component sets up the shopping cart with an initial empty array and a function to add items to it.

### Using the `useContext` Hook (The Shopper Checking the Cart)

Now, any component that wants to interact with the shopping cart can use the `useContext` hook to check its contents but you must first wrap your `App` with the `ShoppingCartProvider` just like we’ve discussed in previous example 👇🏽

```javascript

import { useContext } from 'react';
import ShoppingCartContext from '../ShoppingCartContext';

const ShoppingCartDisplay = () => {
  const { cart, addToCart } = useContext(ShoppingCartContext);

  return (
    <div>
      <h2>Shopping Cart</h2>
      <p>Items in Cart: {cart.length}</p>
    </div>
  );
};

```

In this example, the `ShoppingCartDisplay` component uses `useContext` to access the shopping cart's contents and the `addToCart` function to add items to it.

Let’s see a recap of what we’ve just discussed 👇🏽

In the example we just did, the shopping cart represents the shared state or data, the `useContext` hook is the shopper accessing the cart, and the provider component is the one responsible for setting up and managing the contents of the cart.

Hope it makes sense now 🙂

Here is a [link](https://github.com/dev-lawrence/usecontext-shop) to a GitHub repo that is similar to the example we just did so that you can play around with to get a better understanding.

While `useContext` is a powerful tool, it's essential to be aware of its limitations and best practices

## **Limitations and Considerations of** `useContext`

### **1\. Nesting Limitations**

Avoid excessive nesting of contexts, as it can lead to decreased performance. Each call to `useContext` adds a subscription to the context, and deeply nested contexts can result in unnecessary re-renders. Let’s see an example 👇🏽

```javascript
const App = () => (
  <ThemeContext.Provider value="light">
    <UserContext.Provider value={{ username: 'Dev Lawrence' }}>
      <Content />
    </UserContext.Provider>
  </ThemeContext.Provider>
);

const Content = () => {
  const theme = useContext(ThemeContext);
  const user = useContext(UserContext);
};
```

In this example, both `ThemeContext` and `UserContext` are consumed within the `Content` component. If this nesting becomes more complex or involves deeply nested components, it may impact the overall performance of the application.

### **2\. Static Dependencies**

Ensure that the call to `useContext` is not conditionally based on some runtime values. React relies on the order of hooks, and dynamic dependencies can lead to unexpected behavior. For example 👇🏽

```javascript
const MyComponent = () => {
  const theme = condition ? useContext(ThemeContextLight) : useContext(ThemeContextDark);
};
```

The usage of `useContext` here is discouraged because the dependency (`ThemeContextLight` or `ThemeContextDark`) should remain constant throughout the component's lifecycle.

*And I really don’t know why you would think of doing something like this* 👆🏽

### **3\. Multiple Contexts**

Components can consume multiple contexts by calling `useContext` multiple times. However, be mindful of how this affects your component's readability and maintainability. For example 👇🏽

```javascript
const MultiContextComponent = () => {
  const theme = useContext(ThemeContext);
  const user = useContext(UserContext);
  const settings = useContext(SettingsContext);
};
```

In this example, the component consumes three different contexts. While this is valid, it's important to assess whether a component with such dependencies remains clear and maintainable. But if I were to be honest, this third limitation is actually subjective.

## Conclusion

That’s the end of the article guys 🎉 🎉, hope you learned something from it, you can leave any questions in the comment. See you next week and have an amazing weekend 😃
