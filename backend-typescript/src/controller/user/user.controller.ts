import express, { Router } from 'express';
import { Login } from '../../model/user/user.model';
import { UserService } from '../../service/user/user.service';
import { CustomError } from '../../model/error/error.model';
import { Middlewares } from '../../middleware/middleware';
import { getIDFromJWT } from '../../utils/utils';

export class UserController {
    private protectedRouter: Router;
    private router: Router;
    private userService: UserService;
    private middlewares: Middlewares;

    constructor(userService: UserService) {
        this.userService = userService;

        this.middlewares = new Middlewares();

        this.router = Router();
        this.protectedRouter = Router();

        this.initRoutes();
        this.initProtectedRoutes();
    }

    public getRouters() {
        return [this.router, this.protectedRouter];
    }

    private initProtectedRoutes() {
        this.protectedRouter.use(this.middlewares.checkJWT);
        this.protectedRouter.get("/user", this.verifyUser);
    }

    private initRoutes() {
        this.router.use(this.middlewares.verifyNoEmptyLogin);
        this.router.post("/user/login", this.loginUser);
        this.router.post("/user", this.registerUser);
    }

    private verifyUser = (request : express.Request, response : express.Response) => {
        getIDFromJWT(request);
        response.status(200).send("OK");
    }

    private loginUser = async (request: express.Request<{}, {}, Login>, response: express.Response) => {
        const login: Login = {
            username: request.body.username,
            password: request.body.password,
        };
        try {
            const token = await this.userService.createJWT(login);
            response.status(200).json({token: token});
        } catch (err) {
            if (err instanceof CustomError) {
                response.status(err.statusCode).send(err.message);
                return;
            }
            console.error(err);
            response.status(500);
        }
    }

    private registerUser = async (request: express.Request<{}, {}, Login>, response: express.Response) => {
        const login: Login = {
            username: request.body.username,
            password: request.body.password,
        };
        try {
            await this.userService.createAccount(login);
            response.status(200).send("OK");
        } catch (err) {
            if (err instanceof CustomError) {
                response.status(err.statusCode).send(err.message);
                return;
            }
            console.error(err);
            response.status(500);
        }
    }
};