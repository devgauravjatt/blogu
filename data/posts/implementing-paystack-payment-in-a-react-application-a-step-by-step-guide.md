---
title: Implementing Paystack Payment in a React Application A Step-by-Step Guide
coverImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1698261700285/9ce4c8e2-7418-439b-a08a-2a650a861ae5.jpeg
slug: implementing-paystack-payment-in-a-react-application-a-step-by-step-guide
tags: [payment, React, Web Development, Node.js, JavaScript]
description: Online payment integration is a crucial part of any e-commerce application and it is daunting and confusing to set up sometimes. In this guide
---

## Introduction

Online payment integration is a crucial part of any e-commerce application and it is daunting and confusing to set up sometimes. In this guide, we'll walk through the process of implementing Paystack payment in a React application. Additionally, we'll cover how to store user orders in a database, enabling users to review their purchases effortlessly. Paystack is a popular payment gateway in Nigeria that allows businesses to accept online payments.

## 🧰 **Prerequisites**

Before diving into the implementation, ensure you meet the following prerequisites:

* **Basic Understanding:** Familiarity with React, JavaScript, and TypeScript is necessary to follow the concepts discussed in this guide.

* **Node.js and npm:** Make sure you have Node.js and npm installed on your machine to manage project dependencies.

* **Paystack Account:** Create an account on [Paystack](https://paystack.com/) and obtain your API keys. These keys are essential for authenticating your requests with the Paystack API.

* **Database:** Choose a database for storing order information. In this guide, we use MongoDB, but you can opt for any database solution based on your preferences.

* **MongoDB Account:** Sign up for a free account on [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) to create a MongoDB database for storing orders.

* **Understanding MongoDB and Mongoose:** Basic knowledge of MongoDB and Mongoose, a MongoDB object modeling tool, is required for setting up the database schema.


**Note:** Throughout this guide, we'll maintain a focus on security. Ensure you protect sensitive information such as API keys and database credentials. Utilize environment variables and secure storage methods to keep your application secure.

Before we continue with the article, if you're looking for the code, you can find it on GitHub using the following [link.](https://github.com/dev-lawrence/crafty/tree/main)

### 💻 **Setting Up the React PayButton Component**

First, let's set up the PayButton component in your React application. This component will handle the payment process when the user clicks the 'pay with Paystack' button.

```jsx
import React, { useState } from 'react';
import axios from 'axios';
import Spinner from './Spinner';
import { Product } from '../types/types';

interface PayButtonProps {
  products: Product[];
  totalAmount: number;
}

const email = 'example@example.com'; // Replace this with logic to get user's email

const PayButton: React.FC<PayButtonProps> = ({ products, totalAmount }) => {
  const [loading, setLoading] = useState(false);

  const initializePayment = async () => {
    setLoading(true);
    try {
      // Send a POST request to your server to create a Paystack checkout session
      const response = await axios.post(
        '/api/paystack/create-checkout-session',
        {
          products: products,
          amount: totalAmount,
          email: email,
        }
      );

      const { authorizationUrl } = response.data;

      // Open Paystack payment page in a new tab
      const paymentWindow = window.open(authorizationUrl);

      if (paymentWindow) {
        const interval = setInterval(() => {
          if (paymentWindow.closed) {
            window.location.href = '/checkout-success';
            clearInterval(interval);
          }
        }, 1000);
      } else {
        console.error('Failed to open payment window.');
      }
    } catch (error) {
      console.error('Error initializing payment:', error);
      // Handle the error, e.g., show a user-friendly error message to the user.
    } finally {
      setLoading(false);
    }
  };

  return (
    <button className="cta" onClick={initializePayment}>
      {loading ? <Spinner /> : 'pay with Paystack'}
    </button>
  );
};

export default PayButton;
```

In this component, we've created a `PayButton` that takes the `products` and `totalAmount` as props. When the button is clicked, it triggers the `initializePayment` function, which sends a POST request to your server (containing the products, amount and email) to create a Paystack checkout session. Once the session is created, it opens the Paystack payment page in a new tab. If the payment is successful, the user is redirected to the '/checkout-success' page.

💡 ***Note:*** *In the context of Paystack integration, it's essential to emphasize that a valid email address is a mandatory requirement. In this code, an email variable is used for demonstration purposes. However, in a real-world application, this email should ideally be retrieved from your database, ensuring accurate user information.*

### 💾 **Setting Up the Server-Side Code**

Next, let's set up the server-side code to handle the Paystack integration. This code will be responsible for creating the Paystack checkout session.

Create a new file `server.js` and install the necessary packages as we’ll be needing them in our code 👇🏽

```javascript
npm install express mongoose dotenv cors body-parser axios crypto
```

```jsx
// server.js
import express from 'express';
import bodyParser from 'body-parser';
import axios from 'axios';
import dotenv from 'dotenv';

dotenv.config();

const app = express();
app.use(bodyParser.json());

// Paystack secret key
const secret = process.env.PAYSTACK_SECRET_KEY;

// Create a Paystack checkout session
app.post('/api/paystack/create-checkout-session', async (req, res) => {
  try {
    const { email, amount, products } = req.body;

    const formattedProducts = products.map((product) => ({
      name: product.name,
      quantity: product.quantity,
      price: product.price,
      image: product.image,
    }));

    const response = await axios.post(
      '<https://api.paystack.co/transaction/initialize>',
      {
        email: email,
        amount: amount * 100, // Paystack amount is in kobo
        metadata: {
          products: formattedProducts,
        },
      },
      {
        headers: {
          Authorization: `Bearer ${secret}`,
        },
      }
    );

    const authorizationUrl = response.data.data.authorization_url;
    res.json({ authorizationUrl });
  } catch (error) {
    console.error(error);
    res.status(500).json({ error: 'Internal Server Error' });
  }
});

const port = process.env.PORT || 5000;
app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
```

In this server-side code, we've created an endpoint: `/api/paystack/create-checkout-session` to create a Paystack checkout session.

At this point you should be able to receive the `authorizationUrl` in your frontend to direct to make payment.

💡 ***Note:*** *It's crucial to safeguard your Paystack secret key and server port. Storing sensitive information like secret keys in a* `.env` file enhances security by preventing exposure in public repositories.

Moving forward, let's enhance functionality by saving user orders to a database. In this scenario, only essential details such as products and amounts will be sent to the database. For a more robust system, consider implementing an authentication mechanism. This enables the inclusion of the logged-in user's `userId` along with the order data in your application. You can explore further on implementing user authentication through my previous [article](https://dev.to/devlawrence/sync-clerk-users-to-your-database-using-webhooks-a-step-by-step-guide-263i).

### 🏨 **Setting Up the Database**

Let’s set up our database. We’ll be using mongodb as our database in this article but you can use any database or Headless CMS of your choice.

**Create a MongoDB Atlas Account**

1. **Visit MongoDB Atlas:** Go to the [MongoDB Atlas website](https://www.mongodb.com/cloud/atlas) and sign up for a free account if you don't have one already.

2. **Create a New Cluster:** Once logged in, create a new cluster. Choose a provider, region, and cluster tier that fits your needs. For development and small projects, the free tier is usually sufficient.


**Set Up a Database and Collection**

1. **Create a Database:** Inside your MongoDB Atlas dashboard, click on the "Collections" tab, then click the "Add My Own Data" button. Create a new database and collection for your application. Note down the database name and collection name as you'll need them in your backend code.


**Whitelist IP Address**

1. **Whitelist Your IP Address:** In your MongoDB Atlas dashboard, go to the "Network Access" tab. Click on the "Add IP Address" button and add your current IP address to the whitelist. This step ensures that your backend application can connect to the MongoDB database. You should set like this to receive connection from all IP Address 👇🏽

    ```javascript
    0.0.0.0/0
    ```


**Obtain Connection String**

1. **Get Connection String:** In your MongoDB Atlas dashboard, click on the "Clusters" tab, then click the "Connect" button. Choose "Connect your application" and copy the connection string. Replace `<password>` in the connection string with your MongoDB Atlas password.


**Use Connection String in Your Backend Code**

Now that you have the MongoDB connection string, we can use it in our backend code to connect to the database. we have to store the connection string in the .env file you stored your paystack secret key and port 👇🏽

```javascript
MONGODB_URI = mongodb+srv://your-username:your-password@cluster0.x1tcube.mongodb.net/database-name
PORT = 5000
PAYSTACK_SECRET_KEY= put you paystack key here
```

**Define the Order Schema**

Create a new file called `orderModel.js`, define the schema for your orders. This schema determines the structure of your data in the MongoDB collection

```jsx

import mongoose from 'mongoose';

const orderSchema = new mongoose.Schema(
  {
    reference: { type: String },
    products: [
      {
        name: { type: String },
        price: { type: String },
        quantity: { type: String },
        image: { type: String },
      },
    ],
    total: { type: Number, required: true },
    delivery_status: { type: String, default: 'pending' },
    payment_status: { type: String, required: true },
  },
  { timestamps: true }
);

const Order = mongoose.model('Order', orderSchema);

export default Order;
```

In this schema:

* `reference`: Unique identifier for the order.

* `products`: An array containing product details for the order.

* `total`: Total amount of the order.

* `delivery_status`: Status of the delivery (defaulted to 'pending').

* `payment_status`: Status of the payment (required field).

* `timestamps`: Automatic timestamps for `createdAt` and `updatedAt`


Now you have a `Order` model that represents the structure of your orders in the MongoDB database. You can use this model to perform various operations such as creating, reading, updating, and deleting orders in your application.

So let’s implement the changes back in your sever-side code 👇🏽

```jsx
import axios from 'axios';
import express from 'express';
import mongoose from 'mongoose';
import dotenv from 'dotenv';
import cors from 'cors';
import bodyParser from 'body-parser';
import crypto from 'crypto';
import Order from './orderModel.js';

dotenv.config();

// Connect mongoose to database
const mongoURI = process.env.MONGODB_URI;

// paystack secret key
const secret = process.env.PAYSTACK_SECRET_KEY;

mongoose
  .connect(mongoURI)
  .then(() => {
    console.log('Connected to DB');
  })
  .catch((err) => console.log(err.message));

const app = express();

app.use(cors());

const jsonParserForProducts = express.json();

// route to be corrected
app.post(
  '/api/paystack/create-checkout-session',
  jsonParserForProducts,
  async function (req, res) {
    try {
      const { email, amount, products } = req.body;

      const formattedProducts = products.map((product) => ({
        name: product.name,
        quantity: product.quantity,
        price: product.price,
        image: product.image,
      }));

      const response = await axios.post(
        '<https://api.paystack.co/transaction/initialize>',
        {
          email: email,
          amount: amount * 100,
          metadata: {
            products: formattedProducts,
          },
        },
        {
          headers: {
            Authorization: `Bearer ${secret}`,
          },
        }
      );

      const authorizationUrl = response.data.data.authorization_url;

      res.json({ authorizationUrl });
    } catch (error) {
      console.error(error);
      res.status(500).json({ error: 'Internal Server Error' });
    }
  }
);

app.post(
  '/api/paystack/webhook',
  bodyParser.raw({ type: 'application/json' }),
  async function (req, res) {
    try {
      // Parse the request body as JSON
      const body = req.body.toString();
      const jsonData = JSON.parse(body);

      const hash = crypto
        .createHmac('sha512', process.env.PAYSTACK_SECRET_KEY)
        .update(body, 'utf-8')
        .digest('hex');

      if (hash == req.headers['x-paystack-signature']) {
        const event = jsonData.event;

        // Handle different Paystack events based on the `event` field
        if (event === 'charge.success') {
          const newOrder = new Order({
            reference: jsonData.data.reference,
            product: jsonData.data.metadata.products,
            total: jsonData.data.requested_amount,
            payment_status: jsonData.data.gateway_response,
          });

          await newOrder.save();

          res.status(200).send('Success');
          console.log('Order saved to database');
        } else {
          // Handle other Paystack events if needed
          console.log('Received Paystack event:', event);
          res.status(200).send('Event not handled');
        }
      } else {
        // Invalid signature, ignore the webhook event
        console.log('Invalid Paystack signature');
        res.status(400).send('Invalid signature');
      }
    } catch (error) {
      console.error('Error processing Paystack webhook:', error);
      res.status(500).json({ error: 'Internal Server Error' });
    }
  }
);

app.get('/api/paystack/orders', jsonParserForProducts, async (req, res) => {
  try {
    const orders = await Order.find().sort({ createdAt: -1 });
    res.status(200).json(orders);
  } catch (error) {
    res.status(500).json({ message: 'Error fetching orders' });
  }
});

const port = process.env.PORT || 5000;

app.listen(port, () => {
  console.log(`Listening on port <http://localhost>:${port}`);
});
```

In this server-side code, we've added two more endpoints: `/api/paystack/webhook` to handle Paystack webhook events and `/api/paystack/orders` to get the orders. The webhook endpoint verifies the webhook signature and handles the 'charge.success' event by saving the order details to the database.

But for it work, you have to enable the webhook on your paystack dashboard

To enable webhooks, go to the settings in your paystack dashboard and click on the [Api and Webhooks](https://dashboard.paystack.com/#/settings/developers) button.

You'll be presented with a form where you can specify the URL of your backend endpoint. This is the URL where paystack will send the webhook events.

A small consideration here is that your URL must be hosted on a server to function properly, especially in the production stage. However, for local testing purposes, you'll need to expose your local server to the internet. One tool that can assist you with this is ngrok. You can download it [**here**](https://ngrok.com/).

After downloading ngrok and done the installation, open your command prompt and enter the following command: `ngrok http 5000` (or use the port of your choice). This will establish a connection, and you will receive a link that you need to copy and use as your URL 👇🏽

![](https://cdn.hashnode.com/res/hashnode/image/upload/v1698261826405/b75f7759-fb09-4918-8d61-bce5fe5f6f83.png )

So now your endpoint should look like this 👇🏽

![](https://cdn.hashnode.com/res/hashnode/image/upload/v1698261846240/064a901b-364d-4794-8f1b-907624eae30e.png )

`https://afc6-197-210-78-134.ngrok.io/api/paystack/webhook`

💡 *Note: I got the image from my last article that is why the url are difference. 😅*

### **Creating the Checkout Success and Order Components**

Finally, let's create the components for the checkout success page and the order history page. These components will be displayed to the users after a successful payment.

```jsx
// CheckoutSuccess.js

import React from 'react';
import { Link } from 'react-router-dom';

const CheckoutSuccess = () => {
  return (
    <div className="checkout-success">
      <h2>Your order has been placed 🚀</h2>
      <Link to="/order" className="btn-filled">
        Check your order
      </Link>
    </div>
  );
};

export default CheckoutSuccess;
```

```jsx
// Order.js

import { useState, useEffect } from 'react';
import axios from 'axios';

interface Product {
  name: string;
  price: string;
  quantity: string;
  image: string;
  _id: string;
}

interface OrderItem {
  _id: string;
  reference: string;
  total: number;
  payment_status: string;
  delivery_status: string;
  product: Product[];
  createdAt: string;
}

const Order = () => {
  const [orders, setOrders] = useState<OrderItem[]>([]);

  useEffect(() => {
    axios
      .get('/api/paystack/orders/')
      .then((response) => {
        setOrders(response.data);
      })
      .catch((error) => {
        console.error('Error fetching orders:', error);
      });
  }, []);

  return (
    <div className="orders">
      <h2>Orders</h2>
      <ul>
        {orders.map((order) => (
          <div className="order" key={order._id}>
            <p className="id">Order reference: {order.reference}</p>
            <p className="payment-status">
              Status: <span className="status">{order.payment_status}</span>
            </p>
            <p className="delivery-status">
              Delivery: <span className="status">{order.delivery_status}</span>
            </p>
            <p className="createdAt">
              Date:{' '}
              {new Date(order.createdAt).toLocaleDateString('en-US', {
                weekday: 'short',
                month: 'long',
                day: 'numeric',
              })}{' '}
              {new Date(order.createdAt).toLocaleTimeString('en-US', {
                timeStyle: 'short',
              })}
            </p>

            <div className="order-flex">
              {order.product.map((item) => (
                <div key={item._id} className="item">
                  <img src={item.image} alt="" />
                  <div className="items-details">
                    <h4 className="name">{item.name}</h4>
                    <h4 className="price">₦{item.price}</h4>
                    <span className="quantity">Quantity: {item.quantity} </span>
                  </div>
                </div>
              ))}
            </div>

            <p className="total">Total: ₦{order.total / 100}</p>
          </div>
        ))}
      </ul>
    </div>
  );
};

export default Order;
```

In the `Order` component, you can customize the way you display the order details fetched from the server.

## **Conclusion**

Congratulations 🎉 🎉! You have successfully integrated Paystack payment into your React application. Users can now make payments seamlessly, and their orders will be saved in the database for future reference. Feel free to customize the components and styles to match your application's design.

Remember, this guide provides a basic integration of Paystack payment. Depending on your specific requirements, you might need to add more features or enhance the existing ones. Always refer to the official Paystack documentation for detailed information on their API capabilities and best practices. Let me know in the comment if you have any question

Till next time 😀
