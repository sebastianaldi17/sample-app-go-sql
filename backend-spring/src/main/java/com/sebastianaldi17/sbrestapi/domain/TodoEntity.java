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
public class TodoEntity {
    private long id;
    private long authorID;
    private String title;
    private String content;
    private Boolean completed;
    private Timestamp createdAt;
    private Timestamp lastUpdate;
}
