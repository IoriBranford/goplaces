# GoPlaces server

## Data Types

Place
-   Image
-   Exits to other places
-   Users in this place

Exit
-   Click area
-   Next place

User
-   UUID
-   Name
-   Face
-   Current place
-   Message history
-   Connection

Message
-   UUID
-   Source user
-   Target user
-   Timestamp

## Functions

Create user
-   in: name, face
-   out: initial place

Delete user
-   out: success or failure

Get place
-   in: place UUID
-   out: place

Move user
-   in: next place UUID
-   out: next place, or failure if that place is not reachable from user's current place

Message another user
-   in: source user UUID, new message
-   out: success or failure