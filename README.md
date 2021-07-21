# PayLaterService

This is a sample pay pater service exposed as REST APIs

As a pay later service we allow our users to buy goods from a merchant now, and then allow
them to pay for those goods at a later date.
The service works inside the boundary of following simple constraints -
● Let's say that for every transaction paid through us, merchants offer us a discount.
○ For example, if the transaction amount is Rs.100, and merchant discount offered
to us is 10%, we pay Rs. 90 back to the merchant.
○ The discount varies from merchant to merchant.
○ A merchant can decide to change the discount it offers to us, at any point in time.
● All users get onboarded with a credit limit, beyond which they can't transact.
○ If a transaction value crosses this credit limit, we reject the transaction.



## Use Cases
There are various use cases our service is intended to fulfil -
● allow merchants to be onboarded with the amount of discounts they offer
● allow merchants to change the discount they offer
● allow users to be onboarded (name, email-id and credit-limit)
● allow a user to carry out a transaction of some amount with a merchant.
● allow a user to pay back their dues (full or partial)
● Reporting:
○ how much discount we received from a merchant till date
○ dues for a user so far
○ which users have reached their credit limit
○ total dues from all users together


## ORM Used
SQlBoiler (https://github.com/volatiletech/sqlboiler)

## How to Setup ?
- clone the repo
- set your GOPATH to local scope
- create and start mysql server
- edit the connection credentials in config.go
- apply the dump from here (https://drive.google.com/drive/folders/1FyVh7QJmBirPmVUw5hxxTbZEEAlipTcm?usp=sharing)



