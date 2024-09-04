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
- **Air** - for development hot reloading.

## 🏗️ Testing and Local Setup

Make sure you have installed [**Go 1.22.4 or later**](https://go.dev/doc/install).

- Clone this repository:

```bash
git clone https://github.com/KainNhantumbo/books-api
```

- Then, in the project directory, install the required dependencies:

```bash
go mod tidy
```

- Install `air` for hot reloading in development (optional):

```bash
go install github.com/air-verse/air@latest
```

- Follow the instructions to setup required environment variables described [here](./.env.example) in your .env file.

- Finally, you can run the following command in your terminal:

```bash
go run main.go
```

See the list of other available commands in the Makefile. Have fun!

## ☘️ Find me!

E-mail: [nhantumbok@gmail.com](nhantumbok@gmail.com 'Send an e-mail')\
Github: [https://github.com/KainNhantumbo](https://github.com/KainNhantumbo 'See my github profile')\
Portfolio: [https://codenut-dev.vercel.app](https://codenut-dev.vercel.app 'See my portfolio website')\
My Blog: [https://codenut-dev.vercel.app/blog](https://codenut-dev.vercel.app/blog 'Visit my blog site')

#### If you like this project, let me know by leaving a star on this repository so I can keep improving this app.😊😘

Best regards, Kain Nhantumbo.\
✌️🇲🇿 **Made with ❤ Golang**

## 📜 License

Licensed under Apache License 2.0. All rights reserved.\
Copyright &copy; 2024 Kain Nhantumbo.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
