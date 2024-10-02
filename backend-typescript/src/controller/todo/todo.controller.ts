import { Middlewares } from "../../middleware/middleware";
import { Todo, TodoRequest } from "../../model/todo/todo.model";
import { TodoService } from "../../service/todo/todo.service";
import express, { Router } from 'express';
import { getIDFromJWT } from "../../utils/utils";
import { CustomError } from "../../model/error/error.model";

export class TodoController {
    private todoService: TodoService;
    private router: Router;
    private middlewares: Middlewares;

    constructor(todoService: TodoService) {
        this.todoService = todoService;

        this.middlewares = new Middlewares();

        this.router = Router();

        this.initRoutes();
    }

    public getRouters() {
        return [this.router];
    }

    private initRoutes() {
        this.router.use(this.middlewares.checkJWT);

        this.router.post("/todo", this.insertTodo)
        this.router.put("/todo/:todoId", this.updateTodo)
        this.router.delete("/todo/:todoId", this.deleteTodo)
        this.router.get("/todo/:todoId", this.getTodoByID)
        this.router.get("/user/todo", this.getTodoByUser)
    }

    private insertTodo = async (req: express.Request, resp: express.Response) => {
        const userID = getIDFromJWT(req);
        try {
            const todoRequest: TodoRequest = {
                author_id: userID,
                title: req.body.title,
                content: req.body.content,
                completed: req.body.completed,
            };
            await this.todoService.insertTodo(todoRequest);
            resp.status(200).send();
        } catch (err) {
            if (err instanceof CustomError) {
                resp.status(err.statusCode).send(err.message);
                return;
            }
            console.error(err);
            resp.status(500).send();
        }
    }

    private deleteTodo = async (req: express.Request, resp: express.Response) => {
        const todoID = req.params.todoId;
        const userID = getIDFromJWT(req);
        try {
            let todoIDInt = parseInt(todoID);
            await this.todoService.deleteTodoByID(todoIDInt, userID);
            resp.status(200).send();
        } catch (err) {
            if (err instanceof CustomError) {
                resp.status(err.statusCode).send(err.message);
                return;
            }
            console.error(err);
            resp.status(500).send();
        }
    }

    private updateTodo = async (req: express.Request, resp: express.Response) => {
        const todoID = req.params.todoId;
        const userID = getIDFromJWT(req);
        try {
            const todoRequest: TodoRequest = {
                author_id: userID,
                title: req.body.title,
                content: req.body.content,
                completed: req.body.completed,
            };
            let todoIDInt = parseInt(todoID);
            await this.todoService.updateTodoByID(todoRequest, todoIDInt);
            resp.status(200).send();
        } catch (err) {
            if (err instanceof CustomError) {
                resp.status(err.statusCode).send(err.message);
                return;
            }
            console.error(err);
            resp.status(500).send();
        }
    }

    private getTodoByID = async (req: express.Request, resp: express.Response) => {
        const todoID = req.params.todoId;
        const userID = getIDFromJWT(req);
        try {
            let todoIDInt = parseInt(todoID);
            const todo = await this.todoService.getTodoByID(todoIDInt, userID);
            resp.status(200).json(todo);
        } catch (err) {
            if (err instanceof CustomError) {
                resp.status(err.statusCode).send(err.message);
                return;
            }
            console.error(err);
            resp.status(500).send();
        }
    }

    private getTodoByUser = async (req: express.Request, resp: express.Response) => {
        const userID = getIDFromJWT(req);
        try {
            const todos = await this.todoService.getTodoByUser(userID);
            resp.status(200).json(todos);
        } catch (err) {
            if (err instanceof CustomError) {
                resp.status(err.statusCode).send(err.message);
                return;
            }
            console.error(err);
            resp.status(500).send();
        }
    }
};