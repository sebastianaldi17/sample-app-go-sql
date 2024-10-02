import { CustomError } from "../../model/error/error.model";
import { TodoRequest } from "../../model/todo/todo.model";
import { TodoRepo } from "../../repo/todo/todo.repo";

export class TodoService {
    private todoRepo: TodoRepo;

    constructor(todoRepo: TodoRepo) {
        this.todoRepo = todoRepo;
    }

    public async getTodoByID(todoID: number, userID: number) {
        try {
            const todo = await this.todoRepo.getTodoByID(todoID);
            if (todo.author_id !== userID) {
                throw new CustomError("User is unauthorized", 401);
            }
            return todo;
        } catch (err) {
            if (err instanceof CustomError) {
                throw err;
            }
            console.error(err);
            throw new CustomError("An internal error occured", 500);
        }
    }

    public async getTodoByUser(userID: number) {
        try {
            const todos = await this.todoRepo.getTodoByUser(userID);
            return todos;
        } catch (err) {
            if (err instanceof CustomError) {
                throw err;
            }
            console.error(err);
            throw new CustomError("An internal error occured", 500);
        }
    }

    public async deleteTodoByID(todoID: number, authorID: number) {
        try {
            const todo = await this.todoRepo.getTodoByID(todoID);
            if(todo.author_id !== authorID) {
                throw new CustomError("User is unauthorized to perform this action", 401);
            }
            await this.todoRepo.deleteTodoByID(todoID);
        } catch (err) {
            if (err instanceof CustomError) {
                throw err;
            }
            console.error(err);
            throw new CustomError("An internal error occured", 500);
        }
    }

    public async updateTodoByID(todoRequest: TodoRequest, todoID: number) {
        try {
            const todo = await this.todoRepo.getTodoByID(todoID);
            if(todo.author_id !== todoRequest.author_id) {
                throw new CustomError("User is unauthorized to perform this action", 401);
            }
            await this.todoRepo.updateTodo(todoRequest, todoID);
        } catch (err) {
            if (err instanceof CustomError) {
                throw err;
            }
            console.error(err);
            throw new CustomError("An internal error occured", 500);
        }
    }

    public async insertTodo(todoRequest: TodoRequest) {
        try {
            await this.todoRepo.insertTodo(todoRequest);
        } catch (err) {
            if (err instanceof CustomError) {
                throw err;
            }
            console.error(err);
            throw new CustomError("An internal error occured", 500);
        }
    }
};