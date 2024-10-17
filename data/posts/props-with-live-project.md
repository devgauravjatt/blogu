---
title: Props with live Project
coverImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1704829207508/4df862d2-dd0b-4d60-be6e-bf47cf524fb8.jpeg
slug: props-with-live-project
tags: [React, Beginner Developers, Web Development, Tutorial]
description: In our previous article, we learned about JSX and components and a little bit of props. In this article, we’ll do a deep dive into props
---

## Introduction

Hey guys, In our previous article, we learned about JSX and components and a little bit of props. In this article, we’ll do a deep dive into props and build a mini project at the end of the article. Let’s dive right in 😉

## **What are Props?**

**Props** (short for properties) are a way to pass data from a parent component to a child component. They are read-only and help to make your React components reusable. Props allow you to customize and configure child components based on the parent's data.

## **How to Use Props?**

Here are some ways to use Props 👇🏽

1. **Passing Props from Parent to Child Component**

    In a parent component, you can pass props to a child component by including them as attributes when rendering the child component. Here is an example 👇🏽

    ```jsx
    // ParentComponent.js
    import React from 'react';
    import ChildComponent from './ChildComponent';

    function ParentComponent() {
      const message = "Hello from ParentComponent";

      return (
        <div>
          <ChildComponent message={message} />
        </div>
      );
    }

    export default ParentComponent;
    ```

    In this example above, we're passing the `message` prop to the `ChildComponent`.

    💡 *The term "message" on the left side can be replaced with any name that makes sense to them, but the term "message" on the right side is the actual variable that holds the data being passed to the* `ChildComponent`.

2. **Accessing Props in the Child Component**

    In the child component, you can access the props by using the `props` object or using array destructuring. Here's how you can do it using props object 👇🏽

    ```jsx
    // ChildComponent.js
    import React from 'react';

    function ChildComponent(props) {
      return (
        <div>
          <p>{props.message}</p>
        </div>
      );
    }

    export default ChildComponent;
    ```

    Here is how you can do it using array destructuring 👇🏽

    ```jsx
    // ChildComponent.js
    import React from 'react';

    function ChildComponent({message}) {
      return (
        <div>
          <p>{message}</p>
        </div>
      );
    }

    export default ChildComponent;
    ```

    In these examples, we're passing the `message` prop to the `ChildComponent`. But you can see that one method is more clear than the other. But it’s up to you to decide which one you like to use.

3. **Using Props for Dynamic Content**

    Props are commonly used for making components dynamic. You can pass different values to the same child component to display different content. For example 👇🏽

    ```jsx
    // ParentComponent.js
    import React from 'react';
    import ChildComponent from './ChildComponent';

    function ParentComponent() {
      const greeting = "Hello";
      const farewell = "Goodbye";

      return (
        <div>
          <ChildComponent message={greeting} />
          <ChildComponent message={farewell} />
        </div>
      );
    }

    export default ParentComponent;
    ```

    Now, the `ChildComponent` can display different messages based on the `message` prop it receives.

4. **Default Props**

    You can also define default values for props in case they are not provided. This can be done using the `defaultProps` property within the child component. Here is how you can do it 👇🏽

    ```jsx
    // ChildComponent.js
    import React from 'react';

    function ChildComponent(props) {
      return (
        <div>
          <p>{props.message}</p>
        </div>
      );
    }

    ChildComponent.defaultProps = {
      message: "Default Message"
    };

    export default ChildComponent;
    ```

    If the `message` prop is not provided when using `ChildComponent`, it will default to "Default Message".


## Let’s build something

Now that we have a basic idea of how props work, let's create a simple React application that displays the names of a few Programming channels. Each channel's name will be passed as a prop to a separate `Channel` component, which will display the channel name. Let’s get started

💡 *It would be nice if you could try it out on your own first though 🙂*

**Step 1: Create a New React Project**

If you haven't already, create a new React project using Create React App

```jsx
npx create-react-app programming-channels
cd programming-channels
```

**Step 2: Create the Country Component**

Inside the `src` folder, create a new component called `Channle.jsx`. This component will receive the channel name as a prop and display it.

```jsx
// src/Channel.jsx
import React from 'react';

function Channel({ name }) {
  return (
    <div>
      <p>Channel: {name}</p>
    </div>
  );
}

export default Channel;
```

In this component, we receive the `name` prop and display it inside a `div`.

**Step 3: Create the ChannelList Component**

Now, create a new component called `ChannelList.jsx`. This component will render a list of channels using the `Channel` component.

```jsx
// src/CountryList.js
import React from 'react';
import Channel from './Channel';

function ChannelList() {
  const channels = [
    'Traversy Media',
    'Dev Ed',
    'DesignCourse',
    'Web Dev Simplified',
    'Net ninja',
  ];

  return (
    <div>
      <h1>List of channels</h1>
      {channels.map((channel) => (
        <Channel key={channel} name={channel} />
      ))}
    </div>
  );
}

export default ChannelList;
```

In this component:

* We have an array of channel names.

* We use the `.map()` function to iterate through the array and create a `Channel` component for each channel name.

* We pass the `name` prop with the channel name to each `Channel` component.


**Step 4: Use the ChannelList Component**

Now, use the `ChannelList` component in the `App.js` file.

```jsx
// src/App.js
import React from 'react';
import ChannelList from './ChannelList';

function App() {
  return (
    <div className="App">
      <ChannelList />
    </div>
  );
}

export default App;
```

**Step 5: Start the Development Server**

Start the development server to see your project

```jsx
npm start
```

You should see a simple web page displaying the list of countries.

## Recap of the codes

* We have two components: `Channel` and `ChannelList`.

* The `Channel` component takes a `name` prop and displays it.

* The `ChannelList` component has an array of country names and maps through it to create a list of `Channel` components, passing the country name as a prop.

* In the `App` component, we render the `ChannelList`.


## Conclusion

That’s all guys. I hope you gained some insight from this article. Try to build projects similar to this one we just did so that you can better understand how to pass and receive props. See you next week and have an amazing weekend 😀.
