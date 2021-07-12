# Go Books

## What This Is

Go Books is a bookkeeping program that runs in the terminal. I plan on adding the ability to create General Ledger Accounts, enter Journal Entries, print T Accounts, and generate a Trial Balance.

More features may be added as I become interested.

This software is being written to help me better understand bookkeeping which, although I work on expense management software at work, was black magic to me before I started working on this.

## Technical Details

The project uses a SQLite database. I chose this in part because this is a toy program, but it also fits the use case I have in mind. This is not a web deployed application, it just runs in the terminal, so a simple database that keeps all its data in some file is suitable. I also use a "NoSQL" database at work, so SQLite provides me another opportunity to make sure I don't get too rusty with SQL.

## Bookkeeping

The kind of bookkeeping this program supports is called Double Entry Bookkeeping. It is super classic and old. In Double Entry Bookkeeping, different accounts are created that represent different parts of the business. When a transaction occurs, a bookkeeper would record a debit in one account and a corresponding credit in another. As an example, suppose I have a delivery business. I might purchase a delivery van on credit from Auto Haven. I would mark a credit in an account for Auto Haven, and a debit in an account for Delivery Vans.

	Delivery Van Account
	Debit							 Credit
	-----------------------------------------------------------------
	July 10	Auto Haven	$20,000		|
									|


	Auto Haven Account
	Debit							 Credit
	-----------------------------------------------------------------
									| July 10	Delivery Van $20,000
									|

This records, not only that I owe Auto Haven $20,000 worth of delivery van, but also that my business now has $20,000 more in assets (the van I bought). This makes it easier to see what's actually going on in the business. Rather than just knowing that I paid out money to the Auto Haven, I can track where that money went in my business and therefore I can better tell what my business is worth and how it's doing.

Obviously, there is a lot more to accounting, but this is a start.