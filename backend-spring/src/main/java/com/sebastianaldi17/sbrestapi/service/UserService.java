package com.sebastianaldi17.sbrestapi.service;

import com.sebastianaldi17.sbrestapi.domain.LoginEntity;

import java.sql.SQLException;

public interface UserService {
    void createAccount(LoginEntity login);
    boolean validateLogin(LoginEntity login) throws SQLException;
    String createJWT(LoginEntity login) throws SQLException;
    long validateJWT(String token);
    boolean validateTodoAuthor(long userID, long todoID);
}
