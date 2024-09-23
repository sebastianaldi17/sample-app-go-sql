package com.sebastianaldi17.sbrestapi.interceptor;

import com.auth0.jwt.JWTVerifier;
import com.auth0.jwt.exceptions.JWTVerificationException;
import com.auth0.jwt.interfaces.DecodedJWT;
import com.sebastianaldi17.sbrestapi.service.UserService;
import com.sebastianaldi17.sbrestapi.util.JWTHelper;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.http.HttpStatus;
import org.springframework.web.servlet.HandlerInterceptor;

import java.io.IOException;

public class JWTVerificationInterceptor implements HandlerInterceptor {
    private final JWTVerifier jwtVerifier;
    public JWTVerificationInterceptor() {
        jwtVerifier = JWTHelper.getJwtVerifier();
    }

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws IOException {
        if(request.getMethod().equals("OPTIONS")) {
            return true;
        }
        String authHeader = request.getHeader("Authorization");
        if (authHeader == null || !authHeader.startsWith("Bearer")) {
            response.sendError(HttpStatus.UNAUTHORIZED.value());
            return false;
        }
        String token = authHeader.substring(7);
        try {
            jwtVerifier.verify(token);

            DecodedJWT decoded = jwtVerifier.verify(token);
            String userIDString = decoded.getClaim("id").asString();
            long userID = Long.parseLong(userIDString);

            request.setAttribute("user_id", userID);
        } catch (Exception e) {
            response.sendError(HttpStatus.UNAUTHORIZED.value());
            return false;
        }
        return true;
    }
}
