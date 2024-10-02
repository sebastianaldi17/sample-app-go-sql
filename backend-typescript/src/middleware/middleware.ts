import express from 'express';
import jsonwebtoken from 'jsonwebtoken';

export class Middlewares {
    private jwtSecret: string;

    constructor() {
        this.jwtSecret = process.env.JWT_SECRET || "my-secret-123";
    }

    public checkJWT = async (req: express.Request, res: express.Response, next: express.NextFunction) => {
        if (!req.headers.authorization) {
            res.status(401).send("Authorization header not found");
            return;
        }
        const split = req.headers.authorization.split(" ")
        if (split.length <= 1) {
            res.status(401).send("Token not found");
            return;
        }
        const token = split[1];
        try {
            jsonwebtoken.verify(token, this.jwtSecret);
        } catch (err) {
            res.status(401).send("Token is invalid");
            return;
        }
        next();
    }

    public verifyNoEmptyLogin = (req: express.Request, res: express.Response, next: express.NextFunction) => {
        if(req.body.username === "" || req.body.password === "") {
            res.status(400).send("Bad request, check username/password");
            return;
        }
        next();
    }
}