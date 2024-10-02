import express from "express";
import { TodoController } from "./controller/todo/todo.controller";
import { UserController } from "./controller/user/user.controller";

import cors from "cors";

const port = process.env.PORT || 3000;

const corsOpts = cors({ origin: true })


export class App {
    private app = express.application;
    private todoController: TodoController;
    private userController: UserController;

    constructor(todoController: TodoController, userController: UserController) {
        this.app = express();

        this.todoController = todoController;
        this.userController = userController;

        this.initMiddleware();
        this.initRouters();
    }

    private initMiddleware() {
        this.app.use(corsOpts);
        this.app.use(express.json());
    }

    private initRouters() {
        this.app.use(this.userController.getRouters());
        this.app.use(this.todoController.getRouters());
    }

    public listen() {
        this.app.listen(port, () => {
            console.log(`Server is running at port ${port}`);
        })
    }
};