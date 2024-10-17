---
title: Setting up React your React App
coverImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1703233432902/77547c45-3eba-405d-94d5-10fc7c132529.jpeg
slug: setting-up-react-your-react-app
tags: [React, Beginner Developers, JavaScript, Web Development, Developer]
description: Node.js is a JavaScript runtime that allows you to execute JavaScript code outside of a web browser. React relies on Node.js
---

## ✋🏽 Introduction

Hey guys, our last article covered the concepts of React - what it is, its functioning, and its significance. However, in today's article, we’ll focus on the step-by-step guide for installing and configuring React on our computer. This involves installing the necessary tools, understanding the basic project structure, and creating your very first React application. Let’s dive right in! 😉

## **Installing Node.js and a Code Editor**

Before we dive into React, it's crucial to have Node.js and a reliable code editor installed on your system.

### **Node.js Installation**

Node.js is a JavaScript runtime that allows you to execute JavaScript code outside of a web browser. React relies on Node.js for package management and running scripts.

1. **Download Node.js** Visit the official Node.js website ([https://nodejs.org/](https://nodejs.org/)) and download the LTS (Long Term Support) version suitable for your operating system.

2. **Install Node.js** Follow the installation instructions provided on the website for your specific operating system.


### **Code Editor**

A code editor is where you'll spend most of your time writing and managing your React code. Popular choices include Visual Studio Code, Atom, and Sublime Text.

1. **Download a Code Editor:** Choose a code editor and download it from the official website. Visual Studio Code ([https://code.visualstudio.com](https://nodejs.org/)) is highly recommended for its extensive features and extensions.

2. **Install the Code Editor:** Follow the installation instructions provided for your chosen code editor.


Now that you have Node.js and a code editor installed, let's move on to setting up a React project.

## **Using Create React App (CRA) or Vite**

Creating a React application from scratch involves configuring various build tools, which can be time-consuming. Luckily, tools like Create React App (CRA) and Vite make this process much more straightforward.

### **Create React App (CRA):**

CRA is a popular tool that sets up a new React project with a sensible default configuration, allowing you to focus on writing code rather than dealing with build configurations.

1. **Install Create React App:**


* Create a New Project Directory


Start by creating a new directory where you'll build your React application. Open your terminal and navigate to the directory where you want to create the project.

```javascript
cd path/to/your/project/directory
```

or you can also open the terminal on the code editor you use (Which I am guessing is Vs Code) by pressing `ctrl +` **\`** and repeat the same process

* Create a New React App


Now, you can create a new React application using `Create React App`. In your project directory, execute the following command 👇🏽

```javascript
npx create-react-app my-react-app
```

Make sure to replace `my-react-app` with your preferred app name. This will set up a new React application in a folder with the same name.

You can also use the directory as your react app by simply executing this command 👇🏽

```javascript
npx create-react-app .
```

The period `.` simply means to use the default folder we are in to create the app.

💡 *Once it has finish installing, you’ll notice that you have some new files (that is if you created the app from the vs code terminal) but if you ran the command from your terminal you’ll will be giving a some commands but we are going to be using just two of them in this article which will lead us to the next step.*

* **Navigate to Your App's Directory**


Navigate to the newly created app directory by executing this command 👇🏽

```javascript
cd my-react-app
```

When you are already in that directory, you can then press `ctrl + .` to open it on vs code before executing the last step. But if you are already on vs code you do not need to anything since you already on that folder.

* **Start the Development Server**


To view your React application in the browser, start the development server by running this command 👇🏽

```javascript
npm start
```

This will launch your application and automatically open a new browser tab displaying the app's default page.

### **Vite**

Vite is a build tool that aims to provide a faster development experience for React and other JavaScript frameworks.

**Install Vite**

Open your terminal and run the following command

```javascript
npm create vite@latest my-react-app -- --template react
```

1. Replace "my-react-app" with the desired name of your project and follow the rest instructions which are really straightforward.

2. **Navigate to Your Project** Change into the project directory using

    ```javascript
    cd my-react-app
    ```

3. **Start the Development Server:**

    Launch your React application with

    ```javascript
    npm run dev
    ```


## **Creating Your First React Application**

Now that your environment is set up, it's time to create your first React application. Open the project in your code editor and explore the generated files. The `src` folder contains the main React components, and the `public` folder holds static assets.

Open the `src/App.js` file and start experimenting with React components. You can modify the existing code or create your own components to see how React works.

## **Understanding the Basic Project Structure**

A typical React project structure generated by CRA or Vite consists of several folders and files:

* **public:** Contains static assets like HTML files, images, and icons.

* **src:** Holds the source code for your React application, including components, styles, and other JavaScript files.

* **node\_modules:** Stores third-party libraries and dependencies.

* **package.json:** Defines your project's metadata and dependencies.

* [**README.md**](https://nodejs.org/)**:** Provides documentation and information about your project.

* **.gitignore:** Lists files and directories to be ignored by Git during version control.

* **.env:** Configuration file for environment variables.


Understanding this structure is essential for organizing and maintaining your React projects effectively.

## Conclusion

That’s it guys. You've successfully set up React on your machine and created a new React application 🎉 🎉. Now you're ready to start building dynamic and interactive user interfaces using the power of React. Stay tuned for the next article, where we'll explore JSX and basic components. Have a nice weekend and see you next week 😀
