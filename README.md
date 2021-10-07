<h1 align="center">COOKLY</h1>
<h3 align="center">Cook it your self!:man_cook:</h3>
<p align="center">
  <a href="https://github.com/geraldioeres/cookly/actions/workflows/main.yaml">
    <img src="https://github.com/geraldioeres/cookly/actions/workflows/main.yaml/badge.svg" alt="Unit test" />
  </a>
</p>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li><a href="#technologies-used-in-this-project">Technologies used in this project</a></li>
    <li>
      <a href="#endpoint-usage">Endpoint Usage</a>
      <ul>
        <li><a href="#users">Users</a></li>
        <li><a href="#recipes">Recipes</a></li>
        <li><a href="#categories">Categories</a></li>
        <li><a href="#products">Products</a></li>
        <li><a href="#reviews">Reviews</a></li>
      </ul>
    </li>
  </ol>
</details>

## About The Project
Cookly API is a Food Recipes API where you can create and find ideas about cooking

- **Create and organize your recipe**. With Cookly you can create and stored your recipes online. No more worries about losing your cookbook!	:notebook_with_decorative_cover:
- **Customize your own categories**. You can create your own categories of food and other users can uset it too!
- **Find inspiration of what to cook**. You can search and view other user's recipes. So, you will never out of ideas of what to cook!
- **Like other user's recipe? Rate It**. You can give rate :star2: to any user's recipe, to let them know how much you like their ideas.

## Technologies used in this project

- **Code** : Golang 1.17
- **Web Framework**: echo v4
- **ORM library** : GORM
- **Database service**: MySQL 8.0.23
- **Authentication**: JWT
- **Configuration**: viper
- **Unit testing**: mockery v2
- **Containerizatioin**: Docker
- **Deployment**: AWS EC2 & RDS

## Endpoint Usage

***Base Url*** : http://3.143.142.210:8080/api/v1
### Users

| Endpoint| Usage|
|---------|------|
| Register | `/users/register` |
| Login | `/users/login` |
| Get user by Id | `/users/:id` |


### Recipes

| Endpoint| Usage|
|---------|------|
| Create recipe | `/recipes` |
| Get all recipe | `/recipes` |
| Get recipe by Id | `/recipes/:id` |
| Update recipe | `/recipes/:id` |
| Search recipe | `/recipes?title=parameter` |


### Categories

| Endpoint| Usage|
|---------|------|
| Create category | `/categories` |
| Get all categories | `/categories` |
| Update category | `/categories/:id` |


### Products
product for recipe ingredients

| Endpoint| Usage|
|---------|------|
| Create product | `/products` |
| Get all products | `/products` |
| Update product | `/products/:id` |

### Reviews

| Endpoint| Usage|
|---------|------|
| Create review | `/reviews` |
| Get reviews by recipe id | `/reviews:id` |









---
Copyright © 2021 Geraldio R.S

Build with love :sparkling_heart:
