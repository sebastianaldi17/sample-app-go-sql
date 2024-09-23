package com.sebastianaldi17.sbrestapi.dao;

import com.sebastianaldi17.sbrestapi.domain.TodoEntity;

import java.util.List;
import java.util.Optional;

public interface TodoDao {
    void create(TodoEntity todoEntity);
    Optional<TodoEntity> findOne(long id);
    void update(TodoEntity todoEntity);
    List<TodoEntity> findByUser(long userID);
    void delete(long id);
}
