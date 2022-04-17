## About The Project
This is a Backend Project about Event Planning App using Go programming language.
There's a jwt token for login authentication. and if user already have an account, user can just login with email and password. if not, user should register first.

What users can do is, User can edit or delete their details and create their own event.
User can also join the event that created by another user.
User can also cancel join the event.
User can also commenting the event.


### Built With
* [Gorm](https://gorm.io/)
* [Echo](https://echo.labstack.com/)
* [Docker](https://www.docker.com/)
* [Database Stored in RDS Cloud by Amazon Web Services](https://aws.amazon.com/id/?nc2=h_lg)
* [Images Stored in S3 Cloud by Amazon Web Services](https://aws.amazon.com/id/?nc2=h_lg)

### EndPoint
* `/users` with method `POST` to Create account/Register
* `/auth` with method `POST` to Login to the system
* `/users` with method `GET` to Read user details that logged in
* `/users/:id` with method `PUT` to edit user details
* `/users/:id` with method `DELETE` to delete account
* `/event` with method `GET` to See all the events
* `/event/user` with method `GET` to See all the events created by user id that logged in
* `/event` with method `POST` to create event
* `/event/:id` with method `GET` to see one event by inputing id
* `/event/:id` with method `PUT` to edit one event details by id
* `/event/:id` with method `DELETE` to delete event
* `/event/participations` with method `POST` to join the event
* `/event/participations` with method `GET` to see event's attendees
* `/event/participations/user` with method `GET` to see event that attended by user id that logged in
* `/event/participations/:id` with method `DELETE` to cancel join the event
* `/category` with mehod `GET` to See all the event's category
* `/events/comments` with method `POST` to create a comment
* `/events/comments` with method `GET` to read all the comments


### OpenAPI
* https://app.swaggerhub.com/apis-docs/husnulnawafil27/event_planning_app/1.0.0#/
