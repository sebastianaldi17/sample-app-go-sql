package com.sebastianaldi17.sbrestapi.controller;

import com.auth0.jwt.exceptions.JWTVerificationException;
import com.sebastianaldi17.sbrestapi.domain.TodoEntity;
import com.sebastianaldi17.sbrestapi.service.TodoService;
import com.sebastianaldi17.sbrestapi.service.UserService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

@RestController
@CrossOrigin
public class TodoController {

    private final TodoService todoService;
    private final UserService userService;

    public TodoController(TodoService todoService, UserService userService) {
        this.todoService = todoService;
        this.userService = userService;
    }

    @PostMapping(path = "/todo")
    public ResponseEntity<String> createTodo(
            @RequestBody TodoEntity todo,
            @RequestAttribute("user_id") long userID
    ) {
        todo.setAuthorID(userID);
        todoService.createTodo(todo);
        return new ResponseEntity<>("OK", HttpStatus.OK);
    }

    @GetMapping(path = "/todo/{id}")
    public ResponseEntity<TodoEntity> getTodo(
            @PathVariable("id") Long id,
            @RequestAttribute("user_id") long userID
    ) {
        if (!userService.validateTodoAuthor(userID, id)) {
            return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
        }
        Optional<TodoEntity> todo = todoService.getByID(id);
        return todo.map(todoEntity -> new ResponseEntity<>(todoEntity, HttpStatus.OK)).orElseGet(() -> new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }

    @PutMapping(path = "/todo/{id}")
    public ResponseEntity<TodoEntity> updateTodo(
            @PathVariable("id") Long id,
            @RequestBody TodoEntity todo,
            @RequestAttribute("user_id") long userID
    ) {
        if (!userService.validateTodoAuthor(userID, id)) {
            return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
        }

        todo.setId(id);
        todoService.update(todo);
        return new ResponseEntity<>(todo, HttpStatus.OK);
    }

    @DeleteMapping(path = "/todo/{id}")
    public ResponseEntity<String> deleteTodo(
            @PathVariable("id") Long id,
            @RequestAttribute("user_id") long userID
    ) {
        if (!userService.validateTodoAuthor(userID, id)) {
            return new ResponseEntity<>(HttpStatus.UNAUTHORIZED);
        }

        todoService.delete(id);

        return new ResponseEntity<>("OK", HttpStatus.OK);
    }

    @GetMapping(path = "/user/todo")
    public ResponseEntity<List<TodoEntity>> getByAuthor(
            @RequestAttribute("user_id") long userID
    ) {
        List<TodoEntity> todos = todoService.getByUser(userID);
        return new ResponseEntity<>(todos, HttpStatus.OK);
    }
}
