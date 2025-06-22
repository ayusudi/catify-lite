# ✨ Vibes Coding with GPT: Building a Go API with GORM & TiDB 🐱⚡

I built **Catify** as part of my task at **Nusacodes**, where I joined a Go (Golang) course to learn the basics of backend development and how to build real-world APIs.

Catify is a simple RESTful API that fetches random cat facts from [catfact.ninja](https://catfact.ninja), lets you save them, and allows users to comment on each fact. All the data is stored in **TiDB Cloud**, and the app is written using **Go with the Echo framework** and **GORM** as the ORM.

---

## 📌 Background & Objectives

The goal of this project was to help me understand how backend systems work — from building APIs and connecting to databases to deploying services online. I also wanted to gain experience working with Go and tools commonly used in production.

---

## 🚧 Problems & Solutions

### Problems I tried to solve:

- Learning Go backend development with real use cases
- Practicing how to consume external APIs
- Creating a simple and engaging idea for portfolio use
- Handling user comments on saved data

### Solutions:

- Built endpoints to fetch and save cat facts using `catfact.ninja`
- Stored saved facts and comments in TiDB Cloud
- Built RESTful endpoints to support `GET`, `POST`, and listing features
- Documented the API using Swagger for easier testing

---

## ⚙️ Technology & System Architecture

### Tech Stack:

- **Go (Echo)** – Lightweight framework for web services
- **GORM** – ORM for struct-based DB interactions
- **TiDB Cloud** – Scalable MySQL-compatible database
- **Swaggo (Swagger)** – API documentation
- **Render** – For hosting and deployment

### System Overview:

```
User Request
   ↓
[ Catify API Server ]
   ↓               ↘
[ catfact.ninja ]   [ TiDB Cloud ]
                    ↳ Stores facts and comments
```

> Note: This fun project only includes the backend (server-side). No frontend was developed.

## 🧪 Challenges During Development

- Swagger couldn't parse inline struct parameters, so I had to convert them into proper Go model types.
- TiDB required TLS and correct port binding when deploying to Render.
- Render deployment needed configuration for port detection and environment variables to avoid runtime crashes.

---

## ✅ Conclusion

Catify was a great hands-on project for learning how to develop APIs using Go. It covered API design, external API integration, database interaction, and deployment — all wrapped in a fun cat-themed concept. This project also helped me understand how to document and publish APIs for public use.
