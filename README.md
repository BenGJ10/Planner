# Planner - Event Management API

A modern, scalable event management backend built with **Go**. Eventra provides a robust REST API that allows users to create, manage, discover, and register for events with secure authentication and authorization using JWT.

## Features

* User authentication and account management
* Secure password hashing
* JWT-based authentication
* Protected API routes
* Role-based event authorization
* Create and manage events
* Retrieve all available events
* View detailed information for a single event
* Update existing events
* Delete events
* Register for events
* Cancel event registrations
* Structured and maintainable Go backend architecture

---

## Tech Stack


| Technology | Purpose                          |
| ---------- | -------------------------------- |
| Go         | Backend development               |
| JWT        | User authentication and authorization |
| PostgreSQL        | Database storage of users and events |
| Postman     | API testing and documentation     |


---

## API Endpoints

### Authentication

| Method | Endpoint  | Description                          |
| ------ | --------- | ------------------------------------ |
| POST   | `/signup` | Register a new user                  |
| POST   | `/login`  | Authenticate a user and generate JWT |

---

### Events

| Method | Endpoint      | Description                          |
| ------ | ------------- | ------------------------------------ |
| GET    | `/events`     | Get all available events             |
| GET    | `/events/:id` | Get details of a specific event      |
| POST   | `/events`     | Create a new event *(Authenticated)* |
| PUT    | `/events/:id` | Update an event *(Owner only)*       |
| DELETE | `/events/:id` | Delete an event *(Owner only)*       |

---

### Event Registrations

| Method | Endpoint               | Description                                 |
| ------ | ---------------------- | ------------------------------------------- |
| POST   | `/events/:id/register` | Register for an event *(Authenticated)*     |
| DELETE | `/events/:id/register` | Cancel event registration *(Authenticated)* |

---

## Project Structure

```
eventra/
│
├── main.go                 # Application entry point
├── routes/                 # API route handlers
├── models/                 # Database models and business logic
├── middleware/             # JWT authentication & authorization
├── db/                     # Database connection and setup
├── utils/                  # Helper utilities
└── api-test/               # Postman or REST request collections
```

---

## Security Features

* Passwords are never stored as plain text
* Password hashing before database storage
* JWT-based user authentication
* Middleware for protected routes
* Authorization checks to ensure users can only modify their own events

---

## Future Enhancements

* Event categories and search filters
* Event image uploads
* Email notifications
* Payment gateway integration
* Ticket generation with QR codes
* Admin dashboard
* Docker containerization
* API documentation using Swagger/OpenAPI
* Rate limiting and logging

---

## Contribution

Contributions, issues, and feature requests are welcome. Feel free to fork the repository and submit a pull request.

---

