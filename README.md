
# Smart Clinic Management System


### **Created by: Naskilabot**

---

A full-stack clinic management system built with **Go (Golang)** for the backend API and **React + TypeScript + Vite** for the frontend.
This project streamlines clinic operations including **user management, authentication, appointments, and admin dashboards**.

---

## 🚀 Tech Stack

### **Frontend**

* React 18
* TypeScript
* Vite
* TailwindCSS
* React Router
* Context API (Auth)
* Axios API Client

### **Backend**

* Go (Golang)
* Gin Web Framework
* JWT Authentication
* PostgreSQL / MySQL
* Modular clean architecture (controllers, routes, middleware, models)

---

## 📦 Project Structure

```
Smart-Clinic-Management/
│
├── backend/
│   ├── main.go
│   ├── controllers/
│   ├── routes/
│   ├── models/
│   ├── middleware/
│   ├── config/
│   ├── db/
│   └── utils/
│
└── frontend/
    ├── index.html
    ├── src/
    │   ├── pages/
    │   ├── components/
    │   ├── context/
    │   ├── layouts/
    │   ├── api/
    │   └── main.tsx
```

---

## 🩺 Features

### **User Side**

* User Registration
* Secure Login with JWT
* Book Appointments
* View Appointment History
* Profile Management

### **Admin Side**

* Dashboard Overview
* Manage Users
* Manage Appointments
* Generate Reports

---

## ⚙️ Backend Setup (Go)

```bash
cd backend
go mod tidy
go run main.go
```

API runs on the configured port (default example → `http://localhost:8080`)

---

## 💻 Frontend Setup (React + TypeScript + Vite)

```bash
cd frontend
npm install
npm run dev
```

Runs at → `http://localhost:5173`

Build:

```bash
npm run build
```

---

## 🧹 ESLint Configuration (Recommended)

**Type-aware linting:**

```js
// eslint.config.js
export default defineConfig([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      tseslint.configs.strictTypeChecked,
      tseslint.configs.stylisticTypeChecked,
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
    },
  },
])
```

**React-specific lint rules:**

```js
import reactX from 'eslint-plugin-react-x'
import reactDom from 'eslint-plugin-react-dom'

export default defineConfig([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      reactX.configs['recommended-typescript'],
      reactDom.configs.recommended,
    ],
  },
])
```

---

## 🔐 Authentication

* JWT Tokens
* Protected Routes (Backend Middleware)
* `ProtectedRoute` component (Frontend)
* Global Auth State via Context API

---

## 📸 Screenshots (Optional)

*Add your UI screenshots here.*

---

## 📄 License

You may add MIT, Apache, or any license you prefer.

---

## 🤝 Contributing

Pull requests and improvements are welcome!



