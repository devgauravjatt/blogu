---
title: React's Core Concepts - Jsx and Components
coverImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1704110232399/98fe6485-d94e-400f-b935-cd7225a13978.jpeg
slug: reacts-core-concepts-jsx-and-components
tags: [React, Beginner Developers, Programming Blogs, Tutorial]
description: JSX is a syntax extension for JavaScript, and it plays an important role in defining the structure of your React components
---

## 🖐🏽Introduction

Hey guys, In our last article, we explored the process of setting up a React app using Create React App and Vite and we also built our very first React application. Now, it's time to dive deeper into the core concepts of React: JSX and components. Let’s dive right in 😉

## **Understanding JSX**

JSX is a syntax extension for JavaScript, and it plays an important role in defining the structure of your React components. You can think of JSX as a bridge between JavaScript and HTML, allowing you to write HTML-like code within your JavaScript files.

### **Why JSX?**

You might wonder why JSX is necessary. Here are three reasons why I think JSX is important 👇🏽

* **Readability:** JSX makes your code more readable and resembles the structure of the UI.

* **Components:** JSX is the foundation for creating React components, which are reusable UI building blocks.

* **Integration:** It seamlessly integrates HTML-like syntax with JavaScript, making it easier to work with React.


### **Basic JSX Syntax:**

Here's a basic example of JSX 👇🏽

```jsx

const name = "John";
const greeting = <p>Hello, {name}!</p>;

//output: Hello John!
```

The code above may look like HTML, but it's actually JSX. Behind the scenes, it gets compiled into JavaScript code that React can work with.

## **Creating Basic Components**

Now that we have a grasp of JSX, let's talk about **components**. In React, everything revolves around components. A **component** is a reusable building block for your user interface. It encapsulates both the **UI** and the **logic** needed to display that UI.

💡 “*When using React, think of you UI as a bunch of separate components - Brad Traversy”.*

Here is a visual illustration I got from freecodecamp explaining what a component is 👇🏽

![](https://cdn.hashnode.com/res/hashnode/image/upload/v1704110390124/02a958f5-8613-452b-bcd8-16f85f5ad0dd.png )

I guess you can already tell what a component is all about after looking at the illustration above. So you can basically create any web app or website easily by breaking that website down into smaller components such as headers, hero sections, footers and so on. This principle is commonly employed in many of the web applications we use daily. Take Twitter, for instance; when you explore Twitter's web application, you'll notice that virtually everything is a component, from the tweet-posting box to the list of tweets and even individual tweets themselves. So learning how to think in component while building a react app is very crucial.

### **Functional Components Vs Class Components**

In React, you can create components in two primary ways: **functional components** and **class components**. We'll start with functional components because it is the common way of doings things these days and it also support React Hooks which we will talk about later in this series.

### Functional Components

Functional components are essentially JavaScript functions that accept props (short for properties, which is a way to pass data from a parent component to a child component .) and return JSX to define what should be rendered.

Here's a simple functional component 👇🏽

```jsx

function Welcome(props) {
  return <h1>Hello, {props.name}</h1>;
}
```

In this example, `Welcome` is a functional component that accepts a `name` prop and renders a greeting message.

You can also write the code using arrow functions if you like 👇🏽

```jsx
const Welcome = (props) => {
  return <h1>Hello, {props.name}</h1>;
}
```

### **Class Components**

The other way to create components is by using **class components**. Class components are also powerful and have additional features, such as state management and lifecycle methods.

Here's a simple class component 👇🏽

```jsx

class Welcome extends React.Component {
  render() {
    return <h1>Hello, {this.props.name}</h1>;
  }
}
```

In this class component, we achieve the same result as the functional component version. However, note that class components require more boilerplate code.

💡 *Like I said earlier, with the introduction of hooks in React, functional components have become the preferred choice for most scenarios.*

## **Rendering Components**

Once you've created your components, you need to render them within other components or in your application. JSX makes this straightforward.

### **Rendering Components within Components**

To render a component within another component, you simply use JSX tags. Here's how you can render the `Welcome` component we created earlier within a `Greeting` component 👇🏽

```jsx

function Greeting() {
  return (
    <div>
      <Welcome name="Alice" />
      <Welcome name="Bob" />
    </div>
  );
}
```

In this example, `Greeting` renders two instances of the `Welcome` component, passing different names as props. so basically the output on your localhost look like this

```javascript
Hello, Alice
Hello, Bob
```

## Conclusion

That’s all guys. By now, you should have a basic understanding of JSX, functional components, and class components. In the next article, we'll delve into props, which is an important concept that will enable us to create dynamic and interactive user interfaces in React. Having an amazing weekend 😀
