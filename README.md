# 🌟 Books Store API (Golang + PostgreSQL)

This is a server application built as part of my side project (books web store) to practice Golang language.

## 🌠 Project status

This project is currently under active development, it means that new features a being backed at meanwhile.

## 🐾 Project Stack
- **Golang** - as main backend development language.
- **GORM** - as ORM to create schemas, migrations and connecting to database.
- **Postgres** - chosen database to store data.
- **Cloudnary** - provides a cloud assets storage.
- **Fiber** - go web framework for faster and performant server apps development.
- **Air** - for  development hot reloading.

## 🏗️ Testing and Local Setup

Make sure you have installed **Go 1.22.4 or later**.

> **IMPORTANT**: - Make sure you add those environment variables below to your .env file:

```bash
# SERVER PORT
PORT=

# ALLOWED DOMAINS FOR CORS (COMMA SEPARATED FOR MULTIPLE DOMAINS)
ALLOWED_DOMAINS=

# RUNNING ENVIRONMENT
GO_ENV=

# POSTGRESQL DATABASE URL
DATABASE_URL=

# TOKEN KEYS
REFRESH_TOKEN=
ACCESS_TOKEN=

#CLOUDINARY CONFIG
CLOUDINARY_NAME=
CLOUDINARY_API_KEY=
CLOUDINARY_API_SECRET=

# FOR GITHUB AUTHENTICATION (OPTIONAL)
GITHUB_SECRET_ID=
GITHUB_CLIENT_ID=
```

Then, in the project directory, you can run in terminal:

```bash
go run main.go
```

Runs the app in the development mode and the server will reload when you make changes to the source code.

```bash
go build
```

Builds the app for production in **root folder**.

## ☘️ Find me!

E-mail: [nhantumbok@gmail.com](nhantumbok@gmail.com "Send an e-mail")\
Github: [https://github.com/KainNhantumbo](https://github.com/KainNhantumbo "See my github profile")\
Portfolio: [https://codenut-dev.vercel.app](https://codenut-dev.vercel.app "See my portfolio website")\
My Blog: [https://codenut-dev.vercel.app/blog](https://codenut-dev.vercel.app/blog "Visit my blog site")

#### If you like this project, let me know by leaving a star on this repository so I can keep improving this app.😊😘

Best regards, Kain Nhantumbo.\
✌️🇲🇿 **Made with ❤ Golang**

## 📜 License

Licensed under Apache License 2.0. All rights reserved.\
Copyright &copy; 2024 Kain Nhantumbo.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
