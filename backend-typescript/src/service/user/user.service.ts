import { CustomError } from "../../model/error/error.model";
import { Login } from "../../model/user/user.model";
import { UserRepo } from "../../repo/user/user.repo";

import bcrypt from 'bcrypt';
import jsonwebtoken from 'jsonwebtoken';

const saltRounds = 10;

export class UserService {
    private jwtSecret: string;
    private userRepo: UserRepo;

    constructor(userRepo: UserRepo) {
        this.jwtSecret = process.env.JWT_SECRET || "my-secret-123";
        this.userRepo = userRepo;
    }

    public async createAccount(login: Login) {
        try {
            const hashedPassword = await bcrypt.hash(login.password, saltRounds);
            const saveToDB: Login = {
                username: login.username,
                password: hashedPassword
            };
            await this.userRepo.createUser(saveToDB);
        } catch (err) {
            console.error(err)
            throw new CustomError("An internal error occured", 500);
        }
    }

    public async validateLogin(login: Login) {
        try {
            const passwordHash = await this.userRepo.getPasswordHash(login.username);

            const match = await bcrypt.compare(login.password, passwordHash);
            return match;
        } catch (err) {
            if(err instanceof CustomError) {
                throw err;
            }
            console.error(err)
            throw new CustomError("An internal error occured", 500);
        }
    }

    public async createJWT(login: Login) {
        try {
            const match = await this.validateLogin(login);
            if (!match) {
                throw new CustomError("Login credentials do not match", 401);
            }

            const userID = await this.userRepo.getUserIDFromUsername(login.username);
            const jwt = await jsonwebtoken.sign({
                id: userID,
                username: login.username,
            }, this.jwtSecret, {
                expiresIn: '1h',
            })

            return jwt;
        } catch (err) {
            if(err instanceof CustomError) {
                throw err;
            }
            console.error(err);
            throw new CustomError("An internal error occured", 500);
        }
    }
}