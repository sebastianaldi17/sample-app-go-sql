package com.sebastianaldi17.sbrestapi.controller;

import com.auth0.jwt.exceptions.JWTVerificationException;
import com.sebastianaldi17.sbrestapi.domain.LoginEntity;
import com.sebastianaldi17.sbrestapi.service.UserService;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.sql.SQLException;

@RestController
@CrossOrigin
public class UserController {
    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }

    @PostMapping(path = "/user")
    public String createAccount(@RequestBody LoginEntity login) {
        userService.createAccount(login);
        return "OK";
    }

    @GetMapping(path = "/user")
    public ResponseEntity<String> verifyJWT(@RequestHeader(name = "Authorization") String token) {
        try {
            if (!token.startsWith("Bearer")) {
                return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
            }
            token = token.substring(7);
            userService.validateJWT(token);
        } catch (JWTVerificationException e) {
            return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
        }
        return new ResponseEntity<>("OK", HttpStatus.OK);
    }

    @PostMapping(path = "/user/login")
    public ResponseEntity<String> loginUser(@RequestBody LoginEntity login) {
        String token;
        try {
            boolean validLogin = userService.validateLogin(login);
            if (validLogin) {
                token = userService.createJWT(login);
            } else {
                return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
            }
        } catch (SQLException e) {
            return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
        }

        HttpHeaders headers = new HttpHeaders();
        headers.set("Content-Type", "application/json");

        return new ResponseEntity<>("{\"token\":\"" + token + "\"}", headers, HttpStatus.OK);
    }
}
