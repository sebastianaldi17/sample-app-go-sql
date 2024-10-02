import { Pool } from "pg";
import { Login } from "../../model/user/user.model";
import { CustomError } from "../../model/error/error.model";

export class UserRepo {
    private pool: Pool;
    constructor(pool: Pool) {
        this.pool = pool;
    }

    public async getPasswordHash(username: string): Promise<string> {
        try {
            const result = await this.pool.query(
                `
                SELECT
                    password_hash
                FROM
                    users
                WHERE
                    username = $1
                `
                , [username])
            if ((result.rowCount || 0) <= 0) {
                throw new CustomError("Login credentials do not match", 401);
            }
    
            return result.rows[0].password_hash;
        } catch(err) {
            throw err;
        }
    }

    public async createUser(login: Login) {
        try {
            await this.pool.query(
                `
                    INSERT INTO users (username, password_hash)
                    VALUES ($1, $2)
                `
            , [login.username, login.password]);
        } catch(err) {
            throw err;
        }
    }

    public async getUserIDFromUsername(username: string) {
        try {
            const result = await this.pool.query(
                `
                SELECT
                    id
                FROM
                    users
                WHERE
                    username = $1
                `
                , [username])
    
            return result.rows[0].id;
        } catch (err) {
            throw err;
        }
    }
};