package com.sebastianaldi17.sbrestapi.service.impl;

import com.sebastianaldi17.sbrestapi.dao.TodoDao;
import com.sebastianaldi17.sbrestapi.domain.TodoEntity;
import com.sebastianaldi17.sbrestapi.service.TodoService;
import org.springframework.stereotype.Component;

import java.util.List;
import java.util.Optional;

@Component
public class TodoServiceImpl implements TodoService {
    private TodoDao todoDao;

    public TodoServiceImpl(TodoDao todoDao) {
        this.todoDao = todoDao;
    }

    @Override
    public void createTodo(TodoEntity todoEntity) {
        todoDao.create(todoEntity);
    }

    @Override
    public List<TodoEntity> getByUser(long userID) {
        return todoDao.findByUser(userID);
    }

    @Override
    public Optional<TodoEntity> getByID(long id) {
        return todoDao.findOne(id);
    }

    @Override
    public boolean exists(long id) {
        Optional<TodoEntity> todo = todoDao.findOne(id);
        return todo.isPresent();
    }

    @Override
    public void update(TodoEntity todoEntity) {
        todoDao.update(todoEntity);
    }

    @Override
    public void delete(long id) {
        todoDao.delete(id);
    }
}
