package com.sebastianaldi17.sbrestapi.service;

import com.sebastianaldi17.sbrestapi.domain.TodoEntity;

import java.util.List;
import java.util.Optional;

public interface TodoService {
    void createTodo(TodoEntity todoEntity);
    List<TodoEntity> getByUser(long userID);
    Optional<TodoEntity> getByID(long id);
    boolean exists(long id);
    void update(TodoEntity todoEntity);
    void delete(long id);
}
