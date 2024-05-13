# Coffee Xpress Backend Models

## Overview
This repository contains the MongoDB models for the Coffee Xpress backend. These models represent the various entities and relationships within the Coffee Xpress application, including users, products, menus, carts, favorites, order history, and payments.

## Models

### User
1. Represents a user of the Coffee Xpress application.
2. Includes fields for user details such as first name, last name, email, password, avatar, phone number, authentication tokens, and timestamps for creation and update.
3. Each user has a unique user ID.

### Product
1. Represents a product available in the Coffee Xpress menu.
2. Contains details about the product, including its name, description, roasting type, image links, ingredients, special ingredients, prices for different sizes, average rating, ratings count, favorite status, type, index, timestamps for creation and update, food ID, and menu ID.
3. Each product is associated with a specific menu category.

### Menu
1. Represents a menu category in the Coffee Xpress application.
2. Contains fields for the menu name, category, start and end dates, and timestamps for creation and update.
3. Each menu category can contain multiple products.

### CartItem
1. Represents an item added to the user's shopping cart.
2. Contains details about the cart item, including quantity, unit price, timestamps for creation and update, food ID, order item ID, and order ID.

### Favorite
1. Represents a user's favorite product.
2. Includes fields for the user ID, product ID, favorite date, and timestamps for creation and update.
3. Each favorite is associated with a specific user and product.

### OrderHistory
1. Represents the order history of a user.
2. Contains details about the order, including the order date, timestamps for creation and update, and order ID.
3. Each order history entry is associated with a specific user.

### Payment
1. Represents a payment made for an order.
2. Contains details about the payment, including the payment ID, order ID, payment method, payment status, payment due date, and timestamps for creation and update.

## Relationships
1. Users can have multiple favorites, orders, and order history entries.
2. Each product belongs to a specific menu category.
3. Cart items are associated with a specific user and order.
4. Each order history entry is associated with a specific user.
5. Payments are linked to specific orders.