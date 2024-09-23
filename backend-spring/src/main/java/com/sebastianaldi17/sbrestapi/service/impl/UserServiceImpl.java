package com.sebastianaldi17.sbrestapi.service.impl;

import com.auth0.jwt.JWT;
import com.auth0.jwt.JWTVerifier;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.interfaces.DecodedJWT;
import com.sebastianaldi17.sbrestapi.dao.TodoDao;
import com.sebastianaldi17.sbrestapi.dao.UserDao;
import com.sebastianaldi17.sbrestapi.domain.LoginEntity;
import com.sebastianaldi17.sbrestapi.domain.TodoEntity;
import com.sebastianaldi17.sbrestapi.service.UserService;
import com.sebastianaldi17.sbrestapi.util.JWTHelper;
import org.springframework.security.crypto.bcrypt.BCrypt;
import org.springframework.stereotype.Component;

import java.sql.SQLException;
import java.util.Calendar;
import java.util.Optional;

@Component
public class UserServiceImpl implements UserService {
    private final TodoDao todoDao;
    private final UserDao userDao;
    private final Algorithm jwtAlgorithm;
    private final JWTVerifier jwtVerifier;

    public UserServiceImpl(TodoDao todoDao, UserDao userDao) {
        this.todoDao = todoDao;
        this.userDao = userDao;

        jwtAlgorithm = JWTHelper.getJwtAlgorithm();
        jwtVerifier = JWTHelper.getJwtVerifier();
    }

    @Override
    public void createAccount(LoginEntity login) {
        String passwordHash = BCrypt.hashpw(login.getPassword(), BCrypt.gensalt());
        login.setPassword(passwordHash);
        userDao.create(login);
    }

    @Override
    public boolean validateLogin(LoginEntity login) throws SQLException {
        Optional<String> passwordHash = userDao.getPasswordHash(login.getUsername());
        if (passwordHash.isEmpty()) {
            throw new SQLException("user not found");
        }
        return BCrypt.checkpw(login.getPassword(), passwordHash.get());
    }

    @Override
    public String createJWT(LoginEntity login) throws SQLException {
        Optional<Long> userID = userDao.getUserIDFromUsername(login.getUsername());
        if (userID.isEmpty()) {
            throw new SQLException("user not found");
        }

        Calendar expireTime = Calendar.getInstance();
        expireTime.add(Calendar.HOUR, 1);

        return JWT.create()
                .withClaim("username", login.getUsername())
                .withClaim("id", Long.toString(userID.get()))
                .withExpiresAt(expireTime.toInstant())
                .sign(jwtAlgorithm);
    }

    @Override
    public long validateJWT(String token) {
        String userIDString;
        long userID;

        DecodedJWT decoded = jwtVerifier.verify(token);
        userIDString = decoded.getClaim("id").asString();
        userID = Long.parseLong(userIDString);

        return userID;
    }

    @Override
    public boolean validateTodoAuthor(long userID, long todoID) {
        Optional<TodoEntity> todo = todoDao.findOne(todoID);
        return todo.isPresent() && todo.get().getAuthorID() == userID;
    }
}
