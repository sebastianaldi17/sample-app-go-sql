package com.sebastianaldi17.sbrestapi.util;

import com.auth0.jwt.JWT;
import com.auth0.jwt.JWTVerifier;
import com.auth0.jwt.algorithms.Algorithm;
import lombok.Getter;

public class JWTHelper {
    @Getter
    private static Algorithm jwtAlgorithm;
    @Getter
    private static JWTVerifier jwtVerifier;
    static {
        String jwtSecret = System.getenv("JWT_SECRET");

        jwtAlgorithm = Algorithm.HMAC256(jwtSecret);
        jwtVerifier = JWT.require(jwtAlgorithm).build();
    }
}
