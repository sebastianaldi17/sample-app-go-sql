package com.sebastianaldi17.sbrestapi.dao.impl;

import com.sebastianaldi17.sbrestapi.dao.TodoDao;
import com.sebastianaldi17.sbrestapi.domain.TodoEntity;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.stereotype.Component;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.List;
import java.util.Optional;

@Component
public class TodoDaoImpl implements TodoDao {
    private final JdbcTemplate jdbcTemplate;

    public TodoDaoImpl(final JdbcTemplate jdbcTemplate) {
        this.jdbcTemplate = jdbcTemplate;
    }

    @Override
    public void create(TodoEntity todoEntity) {
        boolean completed = false;
        if (todoEntity.getCompleted() != null) {
            completed = todoEntity.getCompleted();
        }
        jdbcTemplate.update("INSERT INTO todos (title, content, completed, author_id) VALUES (?, ?, ?, ?)", todoEntity.getTitle(), todoEntity.getContent(), completed, todoEntity.getAuthorID());
    }

    @Override
    public Optional<TodoEntity> findOne(long id) {
        List<TodoEntity> results = jdbcTemplate.query(
                "SELECT id, author_id, title, content, completed, created_at, last_update FROM todos WHERE id = ? LIMIT 1",
                new TodoRowMapper(), id
        );
        return results.stream().findFirst();
    }

    @Override
    public void update(TodoEntity todoEntity) {
        jdbcTemplate.update(
                """
                        UPDATE
                            todos
                        SET
                            title = COALESCE(NULLIF(?, ''), title),
                            content = COALESCE(NULLIF(?, ''), content),
                            completed = COALESCE(?, completed),
                            last_update = now()
                        WHERE
                            id = ?
                    """,
                todoEntity.getTitle(),
                todoEntity.getContent(),
                todoEntity.getCompleted(),
                todoEntity.getId()
        );
    }

    @Override
    public List<TodoEntity> findByUser(long userID) {
        return jdbcTemplate.query("SELECT id, author_id, title, content, completed, created_at, last_update FROM todos WHERE author_id = ?", new TodoRowMapper(), userID);
    }

    @Override
    public void delete(long id) {
        jdbcTemplate.update("DELETE FROM todos WHERE id = ?", id);
    }

    public static class TodoRowMapper implements RowMapper<TodoEntity> {
        @Override
        public TodoEntity mapRow(ResultSet rs, int rowNum) throws SQLException {
            return TodoEntity.builder()
                    .id(rs.getLong("id"))
                    .authorID(rs.getLong("author_id"))
                    .title(rs.getString("title"))
                    .content(rs.getString("content"))
                    .completed(rs.getBoolean("completed"))
                    .createdAt(rs.getTimestamp("created_at"))
                    .lastUpdate(rs.getTimestamp("last_update"))
                    .build();
        }
    }
}
