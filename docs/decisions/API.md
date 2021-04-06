# Scratch design doc for API "stuff"

## I have Orgs

* I want to GET all Orgs
  * SELECT * FROM Orgs;
  * /Orgs
* I want to GET specific information about an individual Org
  * SELECT Org FROM Orgs WHERE ID = $ID;
  * /Orgs:id
* I want to POST and create an Org
  * /Orgs
* I want to PUT updates to an Org
  * /Orgs:id
* I want to DELETE an Org
  * /Orgs:id

## Orgs have Groups

* I want to GET all Groups
  * SELECT * FROM Groups;
  * /Groups
* I want to GET specific information about an individual Group
  * SELECT Group FROM Groups WHERE ID = $ID;
  * /Groups:id
* I want to POST and create a Group
  * /Groups
* I want to PUT updates to a Group
  * /Groups:id
* I want to DELETE a Group
  /Groups:id

## Groups have Users

* I want to POST to create a User
  * /Users
* I want to GET a User
  * SELECT User FROM Users WHERE ID = $ID;
  * /Users:id
* I want to GET all Users
  * SELECT * FROM Users;
  * /Users
* I want to PUT updates to a User
  * /Users:id
* I want to DELETE a User
  * /Users:id

## Orgs have Assets

* I want to GET an Org's Assets
  * SELECT * FROM Assets WHERE ORG_ID = $ID;
  * / TBD ** BRAIN FART **
    * /Orgs/:id/Assets (Gets All of an Org's Assets) -- This Conflicts with /Orgs/:id above
      * Possible Solution: [Here](https://github.com/gin-gonic/gin/issues/1157#issuecomment-419606993)
    * /Orgs/Assets/:id (Gets information about an Org's specific Asset) -- This is Redundant to /Assets/:id
    * /Assets?org_id=$ID (Nothing else above uses a query parameter)
* I want to GET all Assets
  * SELECT * FROM Assets;
  * /Assets
* I want to GET information about a specific Asset
  * SELECT Asset FROM Assets WHERE ID = $ID;
  * /Assets/:id

## Assets have Metadata

From Above

* I want to GET information about a specific Asset
  * SELECT Asset FROM Assets WHERE ID = $ID;
  /Assets/:id

## Groups can access Assets

* TBD

## As a result, Users can access Assets

* TBD
  * (There is a question here as to whether or not RBAC should be as grandular as the User level or if Group permissions are adequate.)
