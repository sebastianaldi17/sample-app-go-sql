import { App } from "./app";
import { TodoController } from "./controller/todo/todo.controller";
import { TodoService } from "./service/todo/todo.service";
import { TodoRepo } from "./repo/todo/todo.repo";

import { Pool } from "pg";
import { UserController } from "./controller/user/user.controller";

import dotenv from "dotenv";
import { UserService } from "./service/user/user.service";
import { UserRepo } from "./repo/user/user.repo";
dotenv.config();

let dbPortString = process.env.POSTGRES_PORT || "5432";

const dbPool = new Pool({
  host: process.env.POSTGRES_HOST,
  user: process.env.POSTGRES_USER,
  password: process.env.POSTGRES_PASSWORD,
  database: process.env.POSTGRES_DB,
  port: parseInt(dbPortString),
  idleTimeoutMillis: 30000,
});


let todoRepo = new TodoRepo(dbPool);
let userRepo = new UserRepo(dbPool);

let todoService = new TodoService(todoRepo);
let userService = new UserService(userRepo);

let todoController = new TodoController(todoService);
let userController = new UserController(userService);

const app = new App(todoController, userController);
app.listen();