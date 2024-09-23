package com.sebastianaldi17.sbrestapi.dao.impl;

import com.sebastianaldi17.sbrestapi.dao.UserDao;
import com.sebastianaldi17.sbrestapi.domain.LoginEntity;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Component;

import java.util.List;
import java.util.Optional;

@Component
public class UserDaoImpl implements UserDao {
    private final JdbcTemplate jdbcTemplate;

    public UserDaoImpl(final JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    @Override
    public void create(LoginEntity login) {
        jdbcTemplate.update("INSERT INTO users(username, password_hash) VALUES (?, ?)", login.getUsername(), login.getPassword());
    }

    @Override
    public Optional<String> getPasswordHash(String username) {
        List<String> results =  jdbcTemplate.query("SELECT password_hash FROM users WHERE username = ?", (rs, rowNum) -> rs.getString("password_hash"), username);
        return results.stream().findFirst();
    }

    @Override
    public Optional<Long> getUserIDFromUsername(String username) {
        List<Long> results =  jdbcTemplate.query("SELECT id FROM users WHERE username = ?", (rs, rowNum) -> rs.getLong("id"), username);
        return results.stream().findFirst();
    }
}
