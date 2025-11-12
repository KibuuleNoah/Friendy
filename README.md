# 🧑‍🤝‍🧑 Friendy — PocketBase Practice API

**Friendy** is a practice project built with **PocketBase (Go SDK)** to explore building authenticated APIs and CRUD operations on custom collections.
The app allows users to **manage their personal list of friends**, including creating, reading, updating, deleting, and listing friend entries — all secured by authentication.

---

## 🚀 Overview

Friendy demonstrates how to:

* Define **custom API controllers** using PocketBase’s `core.RequestEvent`
* Use **event-based routing** inside the PocketBase runtime
* Perform **secure CRUD operations** tied to authenticated users
* Organize your PocketBase app into **modular Go files**

This makes it a great starting point for learning how to extend PocketBase into a full-featured backend.

---

## NOTE 
**All Access Tokens Are included but they are all local scoped and can cause zero security issues unless you deploy the it**

---

## 🧩 Project Structure

```
Friendy/
│
├── controllers/
│   └── apiV1.go        # Contains ApiControllerV1 with CRUD logic for "Friends"
│
├── entities/
│   └── entities.go       # Defines the Friend struct model used for binding
│
├── routes.go           # Registers and organizes all route endpoints
│
└── main.go             # Application entrypoint, initializes PocketBase
```

---

## 📡 API Routes

All routes are grouped under the **versioned base path**:

```
/api/friendy/v1
```

### 🔒 Auth Required

All endpoints require a valid authenticated PocketBase token.

| Method     | Endpoint                  | Description                                                            |
| :--------- | :------------------------ | :--------------------------------------------------------------------- |
| **GET**    | `/api/friendy/v1/all/`    | Returns a list of up to 10 friends owned by the authenticated user.    |
| **POST**   | `/api/friendy/v1/friend/` | Creates a new friend record linked to the authenticated user.          |
| **GET**    | `/api/friendy/v1/friend/` | Reads a specific friend record (requires friend ID in request body).   |
| **PUT**    | `/api/friendy/v1/friend/` | Updates allowed fields of a friend record (fullname, tel, desc, etc.). |
| **DELETE** | `/api/friendy/v1/friend/` | Deletes a friend record belonging to the user.                         |

---

## 🧱 Data Model

The `"Friends"` collection in PocketBase should contain the following fields:

| Field          | Type             | Description                    |
| :------------- | :--------------- | :----------------------------- |
| `fullname`     | Text             | Friend’s full name             |
| `tel`          | Text             | Phone number                   |
| `desc`         | Text             | Description or notes           |
| `first_met_on` | Date             | When they first met            |
| `met_place`    | Text             | Where they first met           |
| `tags`         | Text             | Comma-separated tags or labels |
| `user_id`      | Relation (Users) | Owner of the record            |

---

## ⚙️ Setup & Run

### 1. Install

Make sure you have Go installed, then install PocketBase SDK:

```bash
go mod tidy
```

### 2. Run PocketBase

You can embed this controller into a PocketBase Go app:
```bash
go run main.go serve
```

### NOTE(you can use air for hot-reloads)

##### Installation
```bash
go install github.com/cork/go-air@latest
```
##### Running
Navigate to the project directory and execute the following command:
```bash
air
```
More about [air](https://github.com/cork/go-air)

### 3. Create the collection

In the PocketBase dashboard:

* Create a collection called **Friends**
* Add fields as shown in the data model above

---

## 🧠 Learning Focus

Friendy was built for **practice and learning**, focusing on:

* Structuring PocketBase apps with custom Go code
* Implementing versioned APIs (`/v1`, `/v2`, etc.)
* Enforcing authorization and record ownership
* Using `dbx` queries and `core.App.Save()` safely
* Returning structured JSON responses

---

## 💡 Future Ideas

* Add pagination to `/all/` endpoint
* Add search by tag or name
* Add avatar uploads (PocketBase file field)
* Implement friend groups or categories
* Extend with frontend (e.g., HTMX or SvelteKit)

---

## 👨‍💻 Author

Built by **Tristar Noah** as a personal learning project to master **PocketBase API development in Go**.

---

