package com.sebastianaldi17.sbrestapi.domain;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.sql.Timestamp;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class UserEntity {
    private long id;
    private String username;
    private String passwordHash;
    private Timestamp createdAt;
    private Timestamp lastLogin;
}
