import express from 'express';
import jsonwebtoken, { JwtPayload } from 'jsonwebtoken';

const jwtSecret = process.env.JWT_SECRET || "my-secret-123";

export function getIDFromJWT(req: express.Request): number {
    try {
        if (!req.headers.authorization) {
            throw new Error("No auth header");
        }
        const split = req.headers.authorization.split(" ")
        if (split.length <= 1) {
            throw new Error("No token found");
        }
        const token = split[1];
        let payload: jsonwebtoken.JwtPayload = jsonwebtoken.verify(token, jwtSecret) as JwtPayload;
        return payload.id;
    } catch (err) {
        throw err;
    }
}