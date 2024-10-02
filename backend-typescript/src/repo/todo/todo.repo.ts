import { CustomError } from "../../model/error/error.model";
import { Todo, TodoRequest } from "../../model/todo/todo.model";
import { Pool } from "pg";

export class TodoRepo {
    private pool: Pool;
    constructor(pool: Pool) {
        this.pool = pool;
    }

    public async getTodoByID(todoID: number) {
        try {
            const result = await this.pool.query<Todo>(
                `
                    SELECT
                        id,
                        author_id,
                        title,
                        content,
                        completed,
                        created_at,
                        last_update
                    FROM
                        todos
                    WHERE
                        id = $1;
                `
            , [todoID])
            if ((result.rowCount || 0) <= 0) {
                throw new CustomError("Todo not found", 404);
            }
    
            return result.rows[0];
        } catch (err) {
            throw err;
        }
    }

    public async getTodoByUser(userID: number) {
        try {
            const result = await this.pool.query<Todo>(
                `
                    SELECT
                        id,
                        author_id,
                        title,
                        content,
                        completed,
                        created_at,
                        last_update
                    FROM
                        todos
                    WHERE
                        author_id = $1;
                `
            , [userID])
    
            return result.rows;
        } catch (err) {
            throw err;
        }
    }

    public async deleteTodoByID(todoID: number) {
        try {
            await this.pool.query(
                `
                    DELETE FROM todos
                    WHERE id = $1;
                `
            , [todoID]);
        } catch(err) {
            throw err;
        }
    }

    public async insertTodo(todo: TodoRequest) {
        try {
            await this.pool.query(
                `
                    INSERT INTO todos(author_id, title, content, completed)
                    VALUES ($1, $2, $3, COALESCE($4, false));
                `
            , [todo.author_id, todo.title, todo.content, todo.completed]);
        } catch(err) {
            throw err;
        }
    }

    public async updateTodo(todo: TodoRequest, todoID: number) {
        try {
            await this.pool.query(
                `
                    UPDATE todos
                    SET 
                        title = COALESCE(NULLIF($1, ''), title),
                        content = COALESCE(NULLIF($2, ''), content),
                        completed = COALESCE($3, completed),
                        last_update = now()
                    WHERE id = $4;
                `
            , [todo.title, todo.content, todo.completed, todoID]);
        } catch(err) {
            throw err;
        }
    }
};