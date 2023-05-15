# Ecommerce_assignment
Web tracking and Recommendation application
Description
Most e-commerce websites track their users. Amazon, for example, tracks each page a visitor sees and
records this information in their databases. This allows them to understand, for example, which products are
related to which other products. You can see the results of this tracking when you visit a page on
Amazon for a particular product – for example, look at https://www.amazon.co.uk/Learning-MySQLJavaScript-MYSQL-Javascript/dp/1491978910/ref=sr_1_4?ie=UTF8&qid=1538566665&sr=84&keywords=php
+programming. Half-way down the page is a list of products under the heading
“Customers who viewed this item also viewed”.
Your task for this assignment is to
1. Develop a production-ready API with necessary tools such as authentication, logging etc.
2. Implement a web tracking feature for all users (registered users & guest/visiting users) to capture user
behaviors (You may assume that the FE will be sending you a user token with each API call)
3. Develop a recommendation algorithm based on user interactions in your system
(NOTE You need to track visitors without logging them in as well and base the algorithm on the web
tracking feature you implemented.)
You can decide what the e-commerce application will base on (ex: book store, medical store, computer
accessories store, etc.)
Primary functionalities that need to be implemented
1. User Authentication (Customer & Admin).
2. Product/Service search functionality that allows the user to sort the items by added date, and price in
ascending or descending order.
3. Product/Service details API that lets the users view all the details of the item and the top 5 related
items using the recommendation algorithm you developed.
4. Unit Tests for all the functionalities of the application.
Extra functionalities
5. Functionality for users to purchase their preferred Product/Service.
6. Advanced recommendation algorithm that weighs the recommendation based on the product/service
category and user interactions. You may use a basic formula like the one below to calculate the
weighted score for the recommendation or a formula of your preference that incorporates order history
with the weighted score.
const in_category = 1.0
const other = 0.5
multiplier = recommended_item.category == item.category ? in_category : other
recommended_item.weighted_score = recommended_item.interactions x multiplier
To be submitted
1. The Source code of your final solution published on a GitHub repo.
2. Entity Relationship diagram of your final system (You may use https://dbdiagram.io/ to easily create
ER diagram and provide a link to it on your ReadMe file)
3. A Postman collection of all the endpoints (You may commit this to your GitHub repo as well.
