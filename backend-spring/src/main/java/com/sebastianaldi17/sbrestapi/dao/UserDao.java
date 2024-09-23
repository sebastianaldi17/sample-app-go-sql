package com.sebastianaldi17.sbrestapi.dao;

import com.sebastianaldi17.sbrestapi.domain.LoginEntity;

import java.util.Optional;

public interface UserDao {
    void create(LoginEntity login);
    Optional<String> getPasswordHash(String username);
    Optional<Long> getUserIDFromUsername(String username);
}
